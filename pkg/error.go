package spyse

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	// ErrReadBody is returned when response's body cannot be read.
	ErrReadBody = errors.New("could not read error response")
)

type ErrResponse struct {
	Err *Error `json:"error,omitempty"`
}

func (e *ErrResponse) Error() string {
	var err string
	if e.Err != nil {
		err = fmt.Sprintf("code:%s message:\"%s\"\n", e.Err.Code, e.Err.Message)
	}
	if e.Err.Errors != nil {
		for _, v := range e.Err.Errors {
			err += fmt.Sprintf("code:%s location:%s message:\"%s\"", v.Code, v.Location, v.Message)
		}
	}
	return err
}

func (e *ErrResponse) Is(target error) bool {
	var err *ErrResponse
	return errors.As(target, &err)
}

type FieldError struct {
	Code     string `json:"code"`
	Location string `json:"location"`
	Message  string `json:"message,omitempty"`
}

// Error message from Spyse
type Error struct {
	Status  int           `json:"-"`
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Errors  []*FieldError `json:"errors,omitempty"`
}

func (e *Error) Error() string {
	err := e.Code + ": " + e.Message
	if e.Errors != nil {
		for _, v := range e.Errors {
			err += fmt.Sprintf("\ncode:%s, location:%s, message:%s", v.Code, v.Location, v.Message)
		}
	}
	return err
}

// NewSpyseError creates a new spyse Error
func NewSpyseError(err error) error {
	if errors.Is(&ErrResponse{}, err) {
		return err
	}
	return &ErrResponse{Err: &Error{Message: err.Error()}}
}

func getErrorFromResponse(r *http.Response) error {
	errorResponse := new(ErrResponse)
	message, err := ioutil.ReadAll(r.Body)
	if err == nil {
		if err = json.Unmarshal(message, &errorResponse); err == nil {
			errorResponse.Err.Status = r.StatusCode
			return errorResponse
		}
		return errors.New(strings.TrimSpace(string(message)))
	}
	return ErrReadBody
}
