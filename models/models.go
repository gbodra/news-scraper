package models

// ScraperProperties the object
type ScraperProperties struct {
	Domain     string
	Message    string
	BaseURL    string
	SearchPath string
	Source     string
}

// Source types
const (
	Enterprisers = "Enterprisers"
	TechCrunch   = "TechCrunch"
	TheVerge     = "TheVerge"
	HBR          = "HBR"
)
