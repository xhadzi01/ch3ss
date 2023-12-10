package main

type startNewGameResponseModel struct {
	SessionID    uint64 `json:"sessionID"`
	SessionToken string `json:"sessionToken"`
	PlayerID     uint64 `json:"playerID"`
}

type joinGameRequestModel struct {
	SessionID uint64 `json:"sessionID"`
}

type joinGameResponseModel struct {
	PlayerID     uint64 `json:"playerID,omitempty"`
	SessionToken string `json:"sessionToken,omitempty"`
	FailedReason string `json:"failedReason,omitempty"`
}

type proceedToGameGameRequestModel struct {
	SessionID    uint64 `json:"sessionID"`
	SessionToken string `json:"sessionToken"`
	PlayerID     uint64 `json:"playerID"`
}

type proceedToGameGameResponseModel struct {
	Proceed bool `json:"proceed"`
}
