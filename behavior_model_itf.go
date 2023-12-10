package main

type IBehaviorModel interface {
	StartNewGame() (Session, error)
	JoinGame(SessionID) (Session, error)
	ProceedToGame(SessionID, SessionToken, PlayerID) (bool, error)
}
