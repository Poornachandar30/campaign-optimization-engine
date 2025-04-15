package engine

import (
    "fmt"
    "math/rand"
    "time"

    "campaign-engine/models"
)

type CampaignEngine struct {
    Campaign  *models.Campaign
    Platforms *[]models.Platform
}

func (ce *CampaignEngine) Run() {
    ticker := time.NewTicker(5 * time.Second)
    for range ticker.C {
        if ce.Campaign.Budget <= 0 {
            fmt.Printf("[%s] Budget exhausted. Stopping campaign.\n", ce.Campaign.Name)
            return
        }

        bestPlatform := ce.selectPlatform()
        if bestPlatform == nil {
            fmt.Printf("[%s] No valid platform found.\n", ce.Campaign.Name)
            continue
        }

        bid := rand.Float64() * 2
        ce.Campaign.Budget -= bid
        ce.Campaign.CurrentSpend += bid

        fmt.Printf("[%s] Bidding $%.2f on %s | Remaining Budget: $%.2f\n",
            ce.Campaign.Name, bid, bestPlatform.Name, ce.Campaign.Budget)
    }
}

func (ce *CampaignEngine) selectPlatform() *models.Platform {
    var best *models.Platform
    bestScore := 0.0

    for _, platform := range *ce.Platforms {
        if !contains(ce.Campaign.PreferredPlatforms, platform.Name) {
            continue
        }
        score := platform.CVR / platform.CPC
        if score > bestScore {
            bestScore = score
            best = &platform
        }
    }
    return best
}

func contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}