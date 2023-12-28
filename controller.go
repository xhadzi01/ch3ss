package main

import (
	"encoding/json"
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

	http.SetCookie(writter, &http.Cookie{
		Name:     "sessionID",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1,
	})
	http.SetCookie(writter, &http.Cookie{
		Name:     "sessionToken",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1,
	})

	if templ, err := LoadTemplates([]string{"header.html", "footer.html", "index.html"}); err != nil {
		encodeResponseAsText(writter, http.StatusBadRequest, err)
	} else {
		templ.ExecuteTemplate(writter, "index", map[string]string{"Title": "Ch3ss"})
	}
}

func (controller *Controller) StartNewGame(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	if parseErr := request.ParseForm(); parseErr != nil {
		encodeResponseAsJSON(writter, http.StatusBadRequest, parseErr)
		return
	}

	playerID := request.Form.Get("player-id-text-box-value")
	parsedPlayerID, parseErr := strconv.ParseUint(playerID, 10, 64)
	if parseErr != nil {
		encodeResponseAsJSON(writter, http.StatusBadRequest, parseErr)
		return
	}

	// update logic
	session, startGameErr := controller.management.StartNewGame(PlayerID(parsedPlayerID))
	if startGameErr != nil {
		encodeResponseAsJSON(writter, http.StatusBadRequest, startGameErr)
		return
	}

	http.SetCookie(writter, &http.Cookie{
		Name:     "sessionID",
		Value:    strconv.FormatUint(uint64(session.SessionID), 10),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   0,
	})
	http.SetCookie(writter, &http.Cookie{
		Name:     "sessionToken",
		Value:    string(session.SessionToken),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   0,
	})

	http.Redirect(writter, request, "/proceed-to-game", http.StatusSeeOther)
}

func (controller *Controller) JoinGame(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	// // parse input
	// var joinGameRequest joinGameRequestModel
	// if errParsing := parseJSONMessage(&joinGameRequest, request); errParsing != nil {
	// 	encodeResponseAsJSON(writter, http.StatusBadRequest, errParsing)
	// 	return
	// }

	// // update logic
	// session, joinGameErr := controller.behavior.JoinGame(SessionID(joinGameRequest.SessionID))

	// // respond
	// var responseModel joinGameResponseModel

	// if joinGameErr == nil {
	// 	responseModel = joinGameResponseModel{
	// 		PlayerID:     uint64(*session.Player2ID),
	// 		SessionToken: string(session.SessionToken),
	// 	}
	// } else {
	// 	responseModel = joinGameResponseModel{
	// 		FailedReason: joinGameErr.Error(),
	// 	}
	// }

	// encodeResponseAsJSON(writter, http.StatusOK, responseModel)
}

func (controller *Controller) ProceedToGame(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	cookie, err := request.Cookie("sessionID")
	if err != nil || cookie == nil {
		panic("Controller instance is nil")
	}

	// parse input
	var proceedToGameRequest proceedToGameGameRequestModel
	if errParsing := parseJSONMessage(&proceedToGameRequest, request); errParsing != nil {
		encodeResponseAsJSON(writter, http.StatusBadRequest, errParsing)
		return
	}

	// redirect to a waiting screen
	// encodeResponseAsJSON(writter, http.StatusOK, startNewGameResponseModel{
	// 	SessionID:    uint64(session.SessionID),
	// 	SessionToken: string(session.SessionToken),
	// 	PlayerID:     uint64(session.Player1ID),
	// })

	return
	// // update logic
	// proceed, poceedToGameErr := controller.behavior.ProceedToGame(SessionID(proceedToGameRequest.SessionID), SessionToken(proceedToGameRequest.SessionToken), PlayerID(proceedToGameRequest.PlayerID))

	// // respond
	// if poceedToGameErr == nil {
	// 	encodeResponseAsJSON(writter, http.StatusOK, proceedToGameGameResponseModel{
	// 		Proceed: proceed,
	// 	})
	// } else {
	// 	encodeResponseAsJSON(writter, http.StatusBadRequest, poceedToGameErr.Error())
	// }
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

func encodeResponseAsJSON(w http.ResponseWriter, statusCode int, itf interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(itf)
}

func encodeResponseAsText(w http.ResponseWriter, statusCode int, itf interface{}) error {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(statusCode)
	_, err := fmt.Fprint(w, itf)
	return err
}
