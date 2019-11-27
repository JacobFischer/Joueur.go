package client

import (
	"errors"
	"joueur/base/errorhandler"
)

func autoHandleEventFatal(eventBytes []byte) {
	errorhandler.HandleError(
		errorhandler.FatalEvent,
		errors.New("Unexpected fatal event from server"),
		"Got a fatal event from the server: "+string(eventBytes),
	)
}
