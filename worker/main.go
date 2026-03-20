package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os/exec"
	"time"

	"scheduler/shared"
)

func generateWorkerID() string {
	return fmt.Sprintf("worker-%x", rand.Intn(1000000))
}

// 🔥 PANIC RECOVERY (VERY IMPORTANT)
func safeExecute(workerID string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("❌ Worker crashed, recovered:", r)
		}
	}()
	startWorker(workerID)
}

// 🔥 HEARTBEAT
func sendHeartbeat(workerID string) {
	for {
		if shared.Rdb != nil {
			err := shared.Rdb.Set(shared.Ctx, "worker_heartbeat:"+workerID, "alive", 10*time.Second).Err()
			if err != nil {
				fmt.Println("❌ Heartbeat error:", err)
			}
		}
		time.Sleep(5 * time.Second)
	}
}

// 🔥 WORKER LOOP
func startWorker(workerID string) {

	if shared.Rdb == nil {
		fmt.Println("❌ Redis not initialized")
		return
	}

	queues := []struct {
		main       string
		processing string
	}{
		{"high_priority_queue", "processing_high"},
		{"medium_priority_queue", "processing_medium"},
		{"low_priority_queue", "processing_low"},
	}

	for {

		var jobData string
		var err error
		var currentQueue struct {
			main       string
			processing string
		}

		// 🔥 PRIORITY ORDER
		for _, q := range queues {
			jobData, err = shared.Rdb.BRPopLPush(shared.Ctx, q.main, q.processing, 2*time.Second).Result()
			if err == nil {
				currentQueue = q
				break
			}
		}

		if err != nil {
			continue
		}

		var job shared.Job

		// 🔥 FIX: JSON error handling
		err = json.Unmarshal([]byte(jobData), &job)
		if err != nil {
			fmt.Println("❌ JSON parse error:", err)
			continue
		}

		fmt.Println("📌", workerID, "picked job:", job.ID, "| Priority:", job.Priority)

		job.Status = "RUNNING"

		// 🔥 Default retries
		if job.MaxRetries == 0 {
			job.MaxRetries = 3
		}

		// 🔥 Execute command
		cmd := exec.Command("sh", "-c", job.Command)
		output, err := cmd.CombinedOutput()

		if err != nil {

			job.RetryCount++

			fmt.Println("❌ Job failed:", job.ID, "| Attempt:", job.RetryCount)

			// 🔥 FIXED retry condition
			if job.RetryCount < job.MaxRetries {

				fmt.Println("🔁 Retrying job:", job.ID)

				jobDataUpdated, _ := json.Marshal(job)

				err = shared.Rdb.LPush(shared.Ctx, currentQueue.main, jobDataUpdated).Err()
				if err != nil {
					fmt.Println("❌ Redis push error:", err)
				}

			} else {

				fmt.Println("💀 Moving job to DEAD QUEUE:", job.ID)

				job.Status = "FAILED"
				job.Output = err.Error()

				jobDataUpdated, _ := json.Marshal(job)

				err = shared.Rdb.LPush(shared.Ctx, "dead_letter_queue", jobDataUpdated).Err()
				if err != nil {
					fmt.Println("❌ Dead queue push error:", err)
				}
			}

		} else {

			job.Status = "COMPLETED"
			job.Output = string(output)

			fmt.Println("✅ Job completed:", job.ID)
			fmt.Println("Output:", job.Output)
		}

		// 🔥 Remove from processing queue (safe)
		err = shared.Rdb.LRem(shared.Ctx, currentQueue.processing, 1, jobData).Err()
		if err != nil {
			fmt.Println("❌ Failed to remove from processing queue:", err)
		}
	}
}

func main() {

	// 🔥 INIT REDIS FIRST
	shared.InitRedis()

	// 🔥 Prevent startup race condition
	time.Sleep(2 * time.Second)

	rand.Seed(time.Now().UnixNano())

	workerID := generateWorkerID()

	fmt.Println("🚀 Worker started:", workerID)

	// 🔥 Start heartbeat
	go sendHeartbeat(workerID)

	// 🔥 Safe worker execution (prevents crash)
	safeExecute(workerID)
}