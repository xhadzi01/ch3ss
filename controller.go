package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Controller struct {
	behavior   IBehaviorModel
	management IGameSessionManagement
}

func NewController() IController {
	return &Controller{
		behavior:   NewBehaviorModel(),
		management: NewGameSessionManagement(),
	}
}

func (controller *Controller) ShowMainScreen(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	// delete all cookies except for Player ID if any was set beforehand
	ResetCookies(writter)

	playerIDText := PlayerID("")
	// check whether Player ID was set from the last time
	if playerIDTmp, playerIDErr := GetPlayerIDCookie(request); playerIDErr == nil {
		playerIDText = playerIDTmp
	}

	// load teplates
	if templ, err := LoadTemplates([]string{"header.html", "footer.html", "index.html"}); err != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, err)
	} else {
		data := map[string]string{
			"Title":    "Ch3ss",
			"PlayerID": string(playerIDText),
		}
		templ.ExecuteTemplate(writter, "index", data)
	}
}

func (controller *Controller) StartNewGame(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	// parse form
	if parseErr := request.ParseForm(); parseErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, parseErr)
		return
	}

	// retrieve new player ID if any was specified
	playerID := request.Form.Get("player-id-text-box-value")
	if playerID == "" {
		encodeResponseAsText(writter, http.StatusBadRequest, "playerID is invalid (empty)")
		return
	}
	// set player ID cookie
	SetPlayerIDCookie(writter, PlayerID(playerID))

	// try to start the game
	session, startGameErr := controller.management.StartNewGame(PlayerID(playerID))
	if startGameErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, startGameErr)
		return
	}

	// verify that data matches
	if session.Player1Info.PlayerID != PlayerID(playerID) {
		encodeResponseAsText(writter, http.StatusBadRequest, "playerID in started session does not match")
		return
	}

	// set session cookies before redirection
	SetSessionIDCookie(writter, session.SessionID)
	SetSessionTokenCookie(writter, session.SessionToken)

	// redirect to waiting screen
	http.Redirect(writter, request, "/wait-for-oponent", http.StatusSeeOther)
}

func (controller *Controller) JoinGame(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	// parse form
	if parseErr := request.ParseForm(); parseErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, parseErr)
		return
	}

	// retrieve new player ID if any was specified
	playerID := request.Form.Get("player-id-text-box-value")
	if playerID == "" {
		encodeResponseAsText(writter, http.StatusBadRequest, "playerID is invalid (empty)")
		return
	}
	// set player ID cookie
	SetPlayerIDCookie(writter, PlayerID(playerID))

	// retrieve session ID if any was specified
	sessionIDText := request.Form.Get("join-game-text-value")
	var sessionID SessionID
	if sessionIDText == "" {
		encodeResponseAsText(writter, http.StatusBadRequest, "sessionID is invalid (empty)")
		return
	} else if sessionIDTmp, sessionIDErr := strconv.ParseUint(sessionIDText, 10, 64); sessionIDErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, fmt.Sprintf("sessionID could not be parsed, reason: %v", sessionIDErr))
		return
	} else {
		sessionID = SessionID(sessionIDTmp)
	}

	// try to join the game
	session, joinGameErr := controller.management.JoinGame(sessionID, PlayerID(playerID))
	if joinGameErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, joinGameErr)
		return
	}

	// verify that data matches
	if session.Player2Info == nil || session.Player2Info.PlayerID != PlayerID(playerID) {
		encodeResponseAsText(writter, http.StatusBadRequest, "playerID in started session does not match")
		return
	} else if sessionID != session.SessionID {
		encodeResponseAsText(writter, http.StatusBadRequest, "sessionID in started session does not match")
		return
	}

	// set session cookies before redirection
	SetSessionIDCookie(writter, session.SessionID)
	SetSessionTokenCookie(writter, session.SessionToken)

	// redirect to waiting screen
	http.Redirect(writter, request, "/wait-for-oponent", http.StatusSeeOther)
}

func (controller *Controller) WaitForOponent(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	// check whether cookies are set and contain correct values
	var playerID PlayerID
	var sessionID SessionID
	var sessionToken SessionToken
	_ = sessionToken // silence "declared but not used"
	if playerIDTmp, playerIDErr := GetPlayerIDCookie(request); playerIDErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, "playerID is not set, try again")
		return
	} else if sessionIDTmp, sessionIDErr := GetSessionIDCookie(request); sessionIDErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, "sessionID is not set, try again")
		return
	} else if sessionTokenTmp, sessionTokenErr := GetSessionTokenCookie(request); sessionTokenErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, "sessionToken is not set, try again")
		return
	} else if _, isReadyErr := controller.management.IsReadyToProceed(sessionIDTmp, sessionTokenTmp, playerIDTmp); isReadyErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, isReadyErr)
		return
	} else {
		playerID = playerIDTmp
		sessionID = sessionIDTmp
		sessionToken = sessionTokenTmp
	}

	// load teplates
	if templ, err := LoadTemplates([]string{"header.html", "footer.html", "waiting-for-game-ready.html"}); err != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, err)
	} else {
		data := map[string]string{
			"Title":     "Ch3ss",
			"PlayerID":  string(playerID),
			"SessionID": fmt.Sprint(sessionID),
		}
		templ.ExecuteTemplate(writter, "waiting-for-game-ready", data)
	}
}

func (controller *Controller) IsReadyToProceed(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	// check whether cookies are set and contain correct values
	var playerID PlayerID
	var sessionID SessionID
	var sessionToken SessionToken
	_ = sessionToken // silence "declared but not used"
	if playerIDTmp, playerIDErr := GetPlayerIDCookie(request); playerIDErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, "playerID is not set, try again")
		return
	} else if sessionIDTmp, sessionIDErr := GetSessionIDCookie(request); sessionIDErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, "sessionID is not set, try again")
		return
	} else if sessionTokenTmp, sessionTokenErr := GetSessionTokenCookie(request); sessionTokenErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, "sessionToken is not set, try again")
		return
	} else if _, isReadyErr := controller.management.IsReadyToProceed(sessionIDTmp, sessionTokenTmp, playerIDTmp); isReadyErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, isReadyErr)
		return
	} else {
		playerID = playerIDTmp
		sessionID = sessionIDTmp
		sessionToken = sessionTokenTmp
	}

	// load teplates
	if templ, err := LoadTemplates([]string{"header.html", "footer.html", "waiting-for-game-ready.html"}); err != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, err)
	} else {
		data := map[string]string{
			"Title":     "Ch3ss",
			"PlayerID":  string(playerID),
			"SessionID": fmt.Sprint(sessionID),
		}
		templ.ExecuteTemplate(writter, "waiting-for-game-ready", data)
	}
}

func (controller *Controller) ProceedToGame(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	// check whether cookies are set and contain correct values
	var playerID PlayerID
	var sessionID SessionID
	var sessionToken SessionToken
	_ = sessionToken // silence "declared but not used"
	if playerIDTmp, playerIDErr := GetPlayerIDCookie(request); playerIDErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, "playerID is not set, try again")
		return
	} else if sessionIDTmp, sessionIDErr := GetSessionIDCookie(request); sessionIDErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, "sessionID is not set, try again")
		return
	} else if sessionTokenTmp, sessionTokenErr := GetSessionTokenCookie(request); sessionTokenErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, "sessionToken is not set, try again")
		return
	} else if _, isReadyErr := controller.management.IsReadyToProceed(sessionIDTmp, sessionTokenTmp, playerIDTmp); isReadyErr != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, isReadyErr)
		return
	} else {
		playerID = playerIDTmp
		sessionID = sessionIDTmp
		sessionToken = sessionTokenTmp
	}

	// load teplates
	if templ, err := LoadTemplates([]string{"header.html", "footer.html", "waiting-for-game-ready.html"}); err != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, err)
	} else {
		data := map[string]string{
			"Title":     "Ch3ss",
			"PlayerID":  string(playerID),
			"SessionID": fmt.Sprint(sessionID),
		}
		templ.ExecuteTemplate(writter, "waiting-for-game-ready", data)
	}
}

/*
type ChessPieceInfo struct{

}

{
	activePlayer: "Player1",
	chessPieces: []ChessPieceInfo{
		ChessPieceInfo{},
		ChessPieceInfo{},
		ChessPieceInfo{},
		ChessPieceInfo{},
		ChessPieceInfo{},
		ChessPieceInfo{},
		ChessPieceInfo{},
		ChessPieceInfo{},
		ChessPieceInfo{},
	},






}
*/

func (controller *Controller) GetGameInfo(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	fmt.Println("GetGameInfo")
}

func (controller *Controller) GetPlayerInfo(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}
	fmt.Println("GetPlayerInfo")
}

func (controller *Controller) GetCurrentScore(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	templ, err := LoadTemplates([]string{"header.html", "footer.html", "my-score.html"})
	if err != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, err)
		return
	}

	//proceed, poceedToGameErr := controller.behavior.ProceedToGame(SessionID(proceedToGameRequest.SessionID), SessionToken(proceedToGameRequest.SessionToken), PlayerID(proceedToGameRequest.PlayerID))

	data := map[string]string{
		"Title":        "Ch3ss",
		"CurrentScore": "275",
	}
	templ.ExecuteTemplate(writter, "my_score", data)
}

func (controller *Controller) ResetScore(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	templ, err := LoadTemplates([]string{"header.html", "footer.html", "my-score.html"})
	if err != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, err)
		return
	}

	//proceed, poceedToGameErr := controller.behavior.ProceedToGame(SessionID(proceedToGameRequest.SessionID), SessionToken(proceedToGameRequest.SessionToken), PlayerID(proceedToGameRequest.PlayerID))

	data := map[string]string{
		"Title":        "Ch3ss",
		"CurrentScore": "275",
	}
	templ.ExecuteTemplate(writter, "my_score", data)
}

func (controller *Controller) GetLeaderboard(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	if templ, err := LoadTemplates([]string{"header.html", "footer.html", "leaderboard.html"}); err != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, err)
	} else {
		templ.ExecuteTemplate(writter, "leaderboard", map[string]string{"Title": "Ch3ss"})
	}
}

func (controller *Controller) GetActiveSessions(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}
	fmt.Println("GetPlayerInfo")
}

func (controller *Controller) GetStatic(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	fileName := request.URL.Path[strings.LastIndex(request.URL.Path, "/")+1:]
	suffixString := fileName[strings.LastIndex(fileName, ".")+1:]

	if IsImageResource(suffixString) {
		serveStaticResource(writter, request, "resources", fileName)
	} else if IsScriptResource(suffixString) {
		serveStaticResource(writter, request, "scripts", fileName)
	}
}

func IsImageResource(suffix string) bool {
	suffix = strings.ToLower(suffix)
	for _, imgSuffix := range []string{"jpg", "jpeg", "png", "tiff", "bmp", "svg"} {
		if imgSuffix == suffix {
			return true
		}
	}
	return false
}

func IsScriptResource(suffix string) bool {
	return strings.ToLower(suffix) == "js"
}

func serveStaticResource(writter http.ResponseWriter, request *http.Request, subpaths ...string) error {
	if fullResourcePath, resourceErr := getResourcePath(subpaths...); resourceErr != nil {
		return encodeResponseAsText(writter, http.StatusBadRequest, resourceErr)
	} else {
		http.ServeFile(writter, request, fullResourcePath)
		return nil
	}
}
