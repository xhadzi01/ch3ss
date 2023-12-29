package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func ResetCookies(writter http.ResponseWriter) {
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
}

func GetPlayerIDCookie(request *http.Request) (playerID PlayerID, err error) {
	if playerIDCookie, playerIDCookieErr := request.Cookie("playerID"); playerIDCookieErr == nil && playerIDCookie != nil {
		playerID = PlayerID(playerIDCookie.Value)
	} else {
		err = errors.New("Could not retrieve Cookie with name: playerID")
	}
	return
}

func SetPlayerIDCookie(writter http.ResponseWriter, playerID PlayerID) {
	http.SetCookie(writter, &http.Cookie{
		Name:     "playerID",
		Value:    string(playerID),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   0,
	})
}

func GetSessionIDCookie(request *http.Request) (sessionID SessionID, err error) {
	if sessionIDCookie, sessionIDCookieErr := request.Cookie("sessionID"); sessionIDCookieErr == nil && sessionIDCookie != nil {
		if parsedSessionID, parseErr := strconv.ParseUint(sessionIDCookie.Value, 10, 64); parseErr == nil {
			sessionID = SessionID(parsedSessionID)
			return
		}
	}

	err = errors.New("Could not retrieve Cookie with name: sessionID")
	return
}

func SetSessionIDCookie(writter http.ResponseWriter, sessionID SessionID) {
	http.SetCookie(writter, &http.Cookie{
		Name:     "sessionID",
		Value:    fmt.Sprint(sessionID),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   0,
	})
}

func GetSessionTokenCookie(request *http.Request) (sessionToken SessionToken, err error) {
	if sessionTokenCookie, sessionTokenCookieErr := request.Cookie("sessionToken"); sessionTokenCookieErr == nil && sessionTokenCookie != nil {
		sessionToken = SessionToken(sessionTokenCookie.Value)
	} else {
		err = errors.New("Could not retrieve Cookie with name: sessionToken")
	}
	return
}

func SetSessionTokenCookie(writter http.ResponseWriter, sessionToken SessionToken) {
	http.SetCookie(writter, &http.Cookie{
		Name:     "sessionToken",
		Value:    string(sessionToken),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   0,
	})
}
