package main

import (
	"errors"
	"math/rand"
	"time"
)

type SessionID uint64
type SessionToken string
type SessionIDs []SessionID

type Session struct {
	SessionID
	SessionToken
	Player1Info PlayerInfo
	Player2Info *PlayerInfo
}

type Sessions []*Session

func generateNewSessionID() SessionID {
	return SessionID(rand.Uint64())
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

func NewSession(player1ID PlayerID) *Session {
	return &Session{
		SessionID:    generateNewSessionID(),
		SessionToken: generateNewSessionToken(),
		Player1Info:  NewPlayerInfo(player1ID, Player1),
		Player2Info:  nil,
	}
}

func (s *Session) JoinSession(player2ID PlayerID) (err error) {
	if s == nil {
		err = errors.New("session is invalid (nil)")
		return
	} else if s.Player2Info != nil {
		err = errors.New("session is already full")
		return
	} else {
		player2Info := NewPlayerInfo(player2ID, Player2)
		s.Player2Info = &player2Info
		return
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
