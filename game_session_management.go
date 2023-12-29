package main

import "errors"

type GameSessionManagement struct {
	Sessions
}

func NewGameSessionManagement() IGameSessionManagement {
	return &GameSessionManagement{
		Sessions: make(Sessions, 0),
	}
}

func (s *GameSessionManagement) StartNewGame(player1ID PlayerID) (session Session, err error) {
	if s == nil {
		panic("Behvior instance is nil")
	}

	sessionTmp := NewSession(player1ID)
	s.Sessions = append(s.Sessions, sessionTmp)
	session = *sessionTmp
	return
}

func (s *GameSessionManagement) JoinGame(sessionID SessionID, player2ID PlayerID) (session Session, err error) {
	if s == nil {
		panic("Behvior instance is nil")
	}

	for _, sessionInst := range s.Sessions {
		if sessionInst.SessionID == sessionID {
			err = sessionInst.JoinSession(player2ID)
			session = *sessionInst
			return
		}
	}

	err = errors.New("session does not exist")
	return
}

func (s *GameSessionManagement) IsReadyToProceed(sessionID SessionID, sessionToken SessionToken, playerID PlayerID) (proceed bool, err error) {
	if s == nil {
		panic("Behvior instance is nil")
	}

	for _, sessionInst := range s.Sessions {
		if sessionInst.SessionID == sessionID {
			if sessionInst.SessionToken != sessionToken {
				err = errors.New("session token is invalid")
				return
			} else if sessionInst.Player1Info.PlayerID == playerID {
				if sessionInst.Player2Info == nil {
					proceed = false
				} else {
					proceed = true
				}
			} else if sessionInst.Player2Info != nil && sessionInst.Player2Info.PlayerID == playerID {
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
