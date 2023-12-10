package main

import (
	"math"
	"math/rand"
	"time"
)

type SessionID uint64
type SessionIDs []SessionID
type PlayerID uint64

type Session struct {
	SessionID
	Player1ID PlayerID
	Player2ID PlayerID
}

type Sessions []*Session

func generateNewSessionID() SessionID {
	return SessionID(rand.Uint64())
}

func NewSession() *Session {
	return &Session{
		SessionID: generateNewSessionID(),
		Player1ID: math.MaxUint64,
		Player2ID: math.MaxUint64,
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
