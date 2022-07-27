package sonarcloud

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ErrorResponse struct {
	StatusCode int
	Errors     []struct {
		Msg string `json:"msg"`
	} `json:"errors"`
}

func (e *ErrorResponse) Error() string {
	messages := make([]string, len(e.Errors))
	for i := 0; i < len(e.Errors); i++ {
		messages[i] = e.Errors[i].Msg
	}
	messagesString := strings.Join(messages, ",")

	return fmt.Sprintf("received non 2xx status code (%d): %s", e.StatusCode, messagesString)
}

func ErrorResponseFrom(res *http.Response) (*ErrorResponse, error) {
	errorResponse := &ErrorResponse{}
	err := json.NewDecoder(res.Body).Decode(&errorResponse)
	if err != nil {
		return nil, fmt.Errorf("could not decode response into ErrorResponse: %+v", err)
	}
	errorResponse.StatusCode = res.StatusCode
	return errorResponse, nil
}
