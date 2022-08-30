package models

import "time"

type EndUser struct {
	ID            uint              `json:"id"`
	Email         string            `json:"email"`
	Username      string            `json:"username"`
	PassHash      string            `json:"-"`
	PassSalt      string            `json:"-"`
	FullName      string            `json:"fullName"`
	ResumeLink    string            `json:"resumeLink"`
	PortfolioLink string            `json:"portfolioLink"`
	IsBlocked     bool              `json:"isBlocked"`
	BlockedTill   time.Time         `json:"blockedTill"`
	BlockHistory  map[string]string `json:"blockHistory"`
	IsDeactivated bool              `json:"isDeactivated"`
	Preferences   map[string]string `json:"preferences"`
	Currency      int64             `json:"currency"`
}
