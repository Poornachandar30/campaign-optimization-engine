# Campaign Optimization Engine — 

## 🚀 Overview

This is a real-time multi-platform Campaign Optimization Engine built in **Golang**, designed to simulate ad campaigns that dynamically allocate budgets across fluctuating ad platforms (Google, Meta, TikTok). Each campaign operates concurrently, making bidding decisions based on platform performance data (CPC & CVR) in real time.

---

## 🧩 Features Implemented

### ✅ Campaign Simulation
- Each campaign has:
  - Unique ID and name
  - Initial budget
  - Conversion goals
  - Preferred platforms
- Runs as a separate goroutine (simulating concurrent processing)

### ✅ Platform Simulation
- Platforms (Google, Meta, TikTok) with:
  - Fluctuating **Cost Per Click (CPC)** and **Conversion Rate (CVR)**
  - Updated every 3 seconds using `math/rand`

### ✅ Bidding & Decision Engine
- Every 5 seconds, a campaign:
  - Selects the best-performing platform using a score `CVR / CPC`
  - Bids a random amount (up to $2) if within budget
  - Stops when the budget is exhausted

### ✅ Concurrency & Real-Time Logging
- Campaigns run in parallel using goroutines
- Live logs for:
  - Platform stats
  - Bidding decisions
  - Budget usage
  - Campaign termination

---

## 📁 Project Structure
campaign-engine/ ├── main.go // Entry point ├── models/ │ ├── campaign.go // Campaign struct │ └── platform.go // Platform struct ├── engine/ │ └── campaign_engine.go // Core engine logic


---

## ▶️ How to Run

Make sure you have Go installed.

```bash
git clone <your repo OR extract the zip>
cd campaign-engine
go mod init campaign-engine
go mod tidy
go run main.go

