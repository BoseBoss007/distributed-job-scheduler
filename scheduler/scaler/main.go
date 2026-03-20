package main

import (
	"fmt"
	"os/exec"
	"time"

	"scheduler/shared"
)

func main() {

	// 🔥 Initialize Redis (VERY IMPORTANT)
	shared.InitRedis()

	fmt.Println("🚀 Auto-scaler started")

	for {

		// 🔥 Get queue sizes
		high, _ := shared.Rdb.LLen(shared.Ctx, "high_priority_queue").Result()
		medium, _ := shared.Rdb.LLen(shared.Ctx, "medium_priority_queue").Result()
		low, _ := shared.Rdb.LLen(shared.Ctx, "low_priority_queue").Result()

		totalJobs := high + medium + low

		// 🔥 Scaling logic
		var workers int

		if totalJobs == 0 {
			workers = 1
		} else if totalJobs <= 5 {
			workers = 2
		} else if totalJobs <= 10 {
			workers = 3
		} else {
			workers = 5
		}

		fmt.Println("📊 Jobs:", totalJobs, "| Scaling workers to:", workers)

		// 🔥 CRITICAL FIX: same project + correct directory
		cmd := exec.Command(
			"docker", "compose",
			"-p", "distributed-job-scheduler", // ✅ SAME PROJECT NAME
			"up",
			"--scale", fmt.Sprintf("worker=%d", workers),
			"-d",
		)

		// 🔥 CRITICAL FIX: run from project root
		cmd.Dir = "../../"

		err := cmd.Run()
		if err != nil {
			fmt.Println("❌ Scaling error:", err)
		}

		// 🔥 Faster polling (important for scaling)
		time.Sleep(2 * time.Second)
	}
}