package main

import (
	"fmt"
)

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

type MoveFigureRequestData struct {
	PlayerType
	FigureIndex
	TargetFigurePosition uint64
}

func translateMoveFigureRequestData(modelData MoveFigureRequestModel) (data MoveFigureRequestData, err error) {
	// figure template "figure_playerX_X_id"
	// target position template "place_X_id"
	var playerType PlayerType
	var figureIndex FigureIndex
	var targetFigurePosition uint64

	if _, scanErr := fmt.Sscanf(modelData.Figure, "figure_player%d_%d_id", &playerType, &figureIndex); scanErr != nil {
		err = scanErr
		return
	} else if _, scanErr := fmt.Sscanf(modelData.TargetPosition, "place_%d_id", &targetFigurePosition); scanErr != nil {
		err = scanErr
		return
	}

	data = MoveFigureRequestData{
		PlayerType:           playerType,
		FigureIndex:          figureIndex,
		TargetFigurePosition: targetFigurePosition,
	}

	return
}

type MoveFigureRequestModel struct {
	Figure         string `json:"figure"`
	TargetPosition string `json:"targetPosition"`
}

type MoveFigureResponseModel struct {
	Success bool
}
