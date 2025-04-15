package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"campaign-engine/engine"
	"campaign-engine/models"
)

var platforms []models.Platform

func main() {
	rand.Seed(time.Now().UnixNano())

	platforms = []models.Platform{
		{ID: "1", Name: "Google", CPC: 1.0, CVR: 0.05},
		{ID: "2", Name: "Meta", CPC: 0.8, CVR: 0.07},
		{ID: "3", Name: "TikTok", CPC: 0.5, CVR: 0.03},
	}

	go simulatePlatformFluctuations()

	campaigns := []models.Campaign{
		{ID: "C1", Name: "Campaign 1", Budget: 10, TargetConversions: 5, PreferredPlatforms: []string{"Google", "Meta"}},
		{ID: "C2", Name: "Campaign 2", Budget: 12, TargetConversions: 8, PreferredPlatforms: []string{"Meta", "TikTok"}},
	}

	for _, campaign := range campaigns {
		engine := engine.CampaignEngine{
			Campaign:  &campaign,
			Platforms: &platforms,
		}
		go engine.Run()
	}

	select {}
}

func simulatePlatformFluctuations() {
	ticker := time.NewTicker(3 * time.Second)
	for range ticker.C {
		for i := range platforms {
			platforms[i].CPC = round(rand.Float64()*2, 2)
			platforms[i].CVR = round(rand.Float64()*0.1, 4)
		}
		printPlatformStats()
	}
}

func printPlatformStats() {
	fmt.Println("------ Real-Time Platform Stats ------")
	for _, p := range platforms {
		fmt.Printf("%s: CPC = $%.2f | CVR = %.4f", p.Name, p.CPC, p.CVR)
	}
	fmt.Println("--------------------------------------")
}

func round(val float64, precision int) float64 {
	pow := math.Pow(10, float64(precision))
	return float64(int(val*pow)) / pow
}
