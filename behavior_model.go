package main

import (
	"errors"
)

type BehaviorModel struct {
	Sessions
}

func NewBehaviorModel() IBehaviorModel {
	return &BehaviorModel{
		Sessions: make(Sessions, 0),
	}
}

func (bhv *BehaviorModel) StartNewGame() (s Session, err error) {
	if bhv == nil {
		panic("Behvior instance is nil")
	}

	session := NewSession()
	bhv.Sessions = append(bhv.Sessions, session)

	s = *session
	return
}

func (bhv *BehaviorModel) JoinGame(sessionID SessionID) (session Session, err error) {
	if bhv == nil {
		panic("Behvior instance is nil")
	}

	for _, sessionInst := range bhv.Sessions {
		if sessionInst.SessionID == sessionID {
			if sessionInst.Player2ID == nil {
				newPlayerID := generateNewPlayerID()
				sessionInst.Player2ID = &newPlayerID
				sessionInst.GameSession = NewGameSession()
				session = *sessionInst
				return
			} else {
				err = errors.New("session is already full")
				return
			}
		}
	}

	err = errors.New("session does not exist")
	return
}

func (bhv *BehaviorModel) ProceedToGame(sessionID SessionID, sessionToken SessionToken, playerID PlayerID) (proceed bool, err error) {
	if bhv == nil {
		panic("Behvior instance is nil")
	}

	for _, sessionInst := range bhv.Sessions {
		if sessionInst.SessionID == sessionID {
			if sessionInst.SessionToken != sessionToken {
				err = errors.New("session token is invalid")
				return
			} else if sessionInst.Player1ID == playerID {
				if sessionInst.Player2ID == nil {
					proceed = false
				} else {
					proceed = true
				}
			} else if sessionInst.Player2ID != nil && *sessionInst.Player2ID == playerID {
				proceed = true
			} else {
				err = errors.New("invalid player ID")
			}
			return
		}
	}

	err = errors.New("session does not exist")
	return
}
