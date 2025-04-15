package models

type Campaign struct {
    ID                 string
    Name               string
    Budget             float64
    TargetConversions  int
    PreferredPlatforms []string
    CurrentSpend       float64
}