🚀 Distributed Job Scheduler (Go + Redis + Docker)

A production-grade Distributed Job Scheduling System built using Go, Redis, and Docker, supporting priority-based execution, auto-scaling workers, retry mechanisms, and dead letter queue handling.

🧠 Overview

This project simulates a real-world distributed system where jobs are:
	•	Submitted via an API server
	•	Stored in Redis priority queues
	•	Processed by multiple worker nodes
	•	Automatically scaled based on load

It mimics systems like Kubernetes Job Controllers and message queue systems (RabbitMQ/Kafka).

⚙️ Features
	•	✅ Priority-based job scheduling (High, Medium, Low)
	•	✅ Distributed worker execution
	•	✅ Auto-scaling workers based on queue size
	•	✅ Retry mechanism for failed jobs
	•	✅ Dead Letter Queue (DLQ) for failed jobs
	•	✅ Worker heartbeat monitoring
	•	✅ Fault-tolerant processing using Redis
	•	✅ Docker-based deployment

🏗️ Architecture

                    ┌─────────────────────┐
                    │       CLIENT        │
                    │  (curl / UI / app)  │
                    └─────────┬───────────┘
                              │
                              ▼
                    ┌─────────────────────┐
                    │     API SERVER      │
                    │  (Go + Gin/Fiber)  │
                    └─────────┬───────────┘
                              │
                              ▼
              ┌──────────────────────────────────┐
              │             REDIS                │
              │                                  │
              │  high_priority_queue             │
              │  medium_priority_queue           │
              │  low_priority_queue              │
              │                                  │
              │  processing_high                 │
              │  processing_medium               │
              │  processing_low                  │
              │                                  │
              │  dead_letter_queue               │
              │                                  │
              │  worker_heartbeat:<worker_id>    │
              └──────────────┬───────────────────┘
                             │
        ┌────────────────────┼────────────────────┐
        │                    │                    │
        ▼                    ▼                    ▼
┌──────────────┐    ┌──────────────┐    ┌──────────────┐
│   WORKER-1   │    │   WORKER-2   │    │   WORKER-N   │
│              │    │              │    │              │
│ Pull jobs    │    │ Pull jobs    │    │ Pull jobs    │
│ Execute cmd  │    │ Execute cmd  │    │ Execute cmd  │
│ Retry / DLQ  │    │ Retry / DLQ  │    │ Retry / DLQ  │
└──────┬───────┘    └──────┬───────┘    └──────┬───────┘
       │                   │                   │
       └──────────────┬────┴──────────────┬────┘
                      ▼                   ▼
                ┌────────────────────────────┐
                │      JOB EXECUTION         │
                │   (Shell / OS commands)    │
                └────────────────────────────┘


                    ┌─────────────────────┐
                    │     AUTO-SCALER     │
                    │  (Go service)       │
                    │                     │
                    │ Reads Redis queue   │
                    │ Decides worker cnt  │
                    │ docker scale up/down│
                    └─────────────────────┘

🧰 Tech Stack
	•	Backend: Go (Golang)
	•	Queue System: Redis
	•	Containerization: Docker & Docker Compose
	•	Concurrency: Goroutines
	•	Communication: REST API

