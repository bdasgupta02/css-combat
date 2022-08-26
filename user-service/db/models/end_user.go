package models

import "time"

type EndUser struct {
	ID            uint              `json:"id"`
	Email         string            `json:"email"`
	Username      string            `json:"username"`
	PassHash      string            `json:"-"`
	PassSalt      string            `json:"-"`
	FullName      string            `json:"full_name"`
	ResumeLink    string            `json:"resume_link"`
	PortfolioLink string            `json:"portfolio_link"`
	IsBlocked     bool              `json:"is_blocked"`
	BlockedTill   time.Time         `json:"blocked_till"`
	BlockHistory  map[string]string `json:"block_history"`
	IsDeactivated bool              `json:"is_deactivated"`
	Preferences   map[string]string `json:"preferences"`
}