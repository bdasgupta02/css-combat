package main

import "time"

const (
	writeWait      = 10 * time.Second
	pongWait       = 180 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	gamePeriod     = 30 * time.Minute
	gameFreq       = 1 * time.Second
	gameFreqDouble = gameFreq / 2
	maxMessageSize = 512
)
