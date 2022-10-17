package main

import (
	"context"
	"math/rand"
	"game-service/proto/problem"
	"log"
	"time"
)

func getRandomType() string {
	// TODO: add lists code (303) when done
	codes := []string{
		// "301",
		"302",
		// "303",
		// "304",
		// "305",
		// "306",
	}
	idx := rand.Int() % len(codes)
	return codes[idx]
}

func (app *serverConfig) GenerateProblem() (string, string) {
	var cancel context.CancelFunc
	probC.ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	typeObj := problem.TypeMessage{Type: getRandomType()}

	probResp, err := probC.client.GenerateProblem(probC.ctx, &typeObj)
	if err != nil {
		log.Printf("Cannot login via User Service: %v", err)
		return "", ""
	}

	return probResp.Type, probResp.Img
}
