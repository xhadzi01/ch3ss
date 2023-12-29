package main

type IGameSessionManagement interface {
	StartNewGame(player1ID PlayerID) (Session, error)
	JoinGame(sessionID SessionID, player2ID PlayerID) (Session, error)
	IsReadyToProceed(sessionID SessionID, sessionToken SessionToken, playerID PlayerID) (bool, error)
}
