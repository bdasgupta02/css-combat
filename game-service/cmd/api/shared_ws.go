package main

import "time"

const (
	writeWait      = 10 * time.Second
	pongWait       = 180 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)
