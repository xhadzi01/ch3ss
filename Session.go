package main

import (
	"math/rand"
	"time"
)

type SessionID uint64
type SessionToken string
type SessionIDs []SessionID
type PlayerID uint64

type Session struct {
	SessionID
	SessionToken
	Player1ID   PlayerID
	Player2ID   *PlayerID
	GameSession IGameSession
}

type Sessions []*Session

func generateNewSessionID() SessionID {
	return SessionID(rand.Uint64())
}

func generateNewPlayerID() PlayerID {
	return PlayerID(rand.Uint64())
}

const (
	allowedSessionTokenLength  = 30
	allowedSessionTokenLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-*@#!$"
)

func generateNewSessionToken() SessionToken {
	var letters [allowedSessionTokenLength]byte

	for i := 0; i < allowedSessionTokenLength; i++ {
		randByteIdx := rand.Intn(len(allowedSessionTokenLetters))
		letters[i] = allowedSessionTokenLetters[randByteIdx]
	}

	return SessionToken(string(letters[:]))
}

func NewSession() *Session {
	return &Session{
		SessionID:    generateNewSessionID(),
		SessionToken: generateNewSessionToken(),
		Player1ID:    generateNewPlayerID(),
		Player2ID:    nil,
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
