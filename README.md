# 🚀 Distributed Job Scheduler (Go + Redis + Docker)

A production-grade **Distributed Job Scheduling System** built using **Go, Redis, and Docker**, supporting:

- Priority-based execution  
- Auto-scaling workers  
- Retry mechanisms  
- Dead Letter Queue (DLQ) handling  

---

## 🧠 Overview

This project simulates a real-world distributed system where jobs are:

- Submitted via an API server  
- Stored in Redis priority queues  
- Processed by multiple worker nodes  
- Automatically scaled based on system load  

It mimics systems like **Kubernetes Job Controllers** and **message queue systems (RabbitMQ / Kafka)**.

---

## ⚙️ Features

- ✅ Priority-based job scheduling (**High, Medium, Low**)  
- ✅ Distributed worker execution  
- ✅ Auto-scaling workers based on queue size  
- ✅ Retry mechanism for failed jobs  
- ✅ Dead Letter Queue (DLQ) for failed jobs  
- ✅ Worker heartbeat monitoring  
- ✅ Fault-tolerant processing using Redis  
- ✅ Docker-based deployment  

---

## 🏗️ Architecture

<p align="center">
  <img src="https://github.com/user-attachments/assets/8e735a61-b912-44d7-bf15-cbbca9631951" alt="Architecture Diagram" width="800"/>
</p>

---

## 🧰 Tech Stack

- **Backend:** Go (Golang)  
- **Queue System:** Redis  
- **Containerization:** Docker & Docker Compose  
- **Concurrency:** Goroutines  
- **Communication:** REST API  

--
