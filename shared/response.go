package shared

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// Response - jsust a shortalnd
type Response = events.APIGatewayProxyResponse

var jsonHeaders = map[string]string{
	"Content-Type": "application/json",
}

// JSONError - error struct
type JSONError struct {
	IP      string
	Error   string
	Message string
}

func jsonError(ip string, err error) string {
	// TODO
	// detect error kind
	errData := JSONError{ip, "some_error", err.Error()}
	data, _ := json.Marshal(errData)
	return string(data)
}

// AsError - simple aws json error response
func AsError(ip string, err error) (Response, error) {
	return Response{
		StatusCode: 422,
		Headers:    jsonHeaders,
		Body:       jsonError(ip, err),
	}, nil
}

// AsData - simple aws proxy response
func AsData(ip string, data interface{}) (resp Response, err error) {
	raw, err := json.Marshal(data)
	if err != nil {
		return AsError(ip, err)
	}

	resp = Response{
		StatusCode: 200,
		Headers:    jsonHeaders,
		Body:       string(raw),
	}
	return
}
