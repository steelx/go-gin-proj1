package app

import (
	"encoding/json"
	"time"
)

var employees = map[string]Employee{
	"962134": Employee{
		ID:        962134,
		FirstName: "Jennifer",
		LastName:  "Watson",
		Position:  "CEO",
		StartDate: time.Now().Add(-13 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  30,
	},
	"176158": Employee{
		ID:        176158,
		FirstName: "Allison",
		LastName:  "Jane",
		Position:  "COO",
		StartDate: time.Now().Add(-4 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  20,
	},
	"160898": Employee{
		ID:        160898,
		FirstName: "Aakar",
		LastName:  "Uppal1",
		Position:  "CTO",
		StartDate: time.Now().Add(-6 * time.Hour * 24 * 365),
		TotalPTO:  20,
	},
	"297365": Employee{
		ID:        297365,
		FirstName: "Jonathon",
		LastName:  "Anderson",
		Position:  "Worker Bee",
		StartDate: time.Now().Add(-12 * time.Hour * 24 * 365),
		TotalPTO:  30,
	},
}

var TimesOff = map[string][]TimeOff{
	"962134": []TimeOff{
		{
			Type:      "Holiday",
			Amount:    "8",
			StartDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
			Status:    "Taken",
		}, {
			Type:      "PTO",
			Amount:    "16",
			StartDate: time.Date(2016, 8, 16, 0, 0, 0, 0, time.UTC),
			Status:    "Scheduled",
		}, {
			Type:      "PTO",
			Amount:    "16",
			StartDate: time.Date(2016, 12, 8, 0, 0, 0, 0, time.UTC),
			Status:    "Requested",
		},
	},
}

type Employee struct {
	ID        uint
	FirstName string
	LastName  string
	StartDate time.Time
	Position  string
	TotalPTO  float32
	Status    string
	TimesOff  []TimeOff
}

type TimeOff struct {
	Type      string      `json:"reason"`
	Amount    json.Number `json:"hours"`
	StartDate time.Time   `json:"startDate"`
	Status    string      `json:"status"`
}
