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

          +-------------+
          |   Client    |
          +-------------+
                 |
                 v
          +-------------+
          | API Server  |
          +-------------+
                 |
                 v
        +-------------------+
        |   Redis Queues    |
        | (Priority Based)  |
        +-------------------+
          /       |        \
         v        v         v
   Worker-1  Worker-2  Worker-N
         \        |        /
          \       |       /
           v      v      v
        +-------------------+
        |   Job Execution   |
        +-------------------+
                 |
                 v
        +-------------------+
        | Dead Letter Queue |
        +-------------------+

🧰 Tech Stack
	•	Backend: Go (Golang)
	•	Queue System: Redis
	•	Containerization: Docker & Docker Compose
	•	Concurrency: Goroutines
	•	Communication: REST API

