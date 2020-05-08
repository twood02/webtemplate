package main

import (
	"encoding/json"
	"log"
	"time"
)

//LogFile represents a logfile with multiple lines
type LogFile struct {
	Logs []LogLine
}

//LogLine represents fields in a given log line
type LogLine struct {
	Name          string
	RawLog        string
	RemoteAddr    string
	TimeLocal     string
	RequestType   string
	RequestPath   string
	Status        int
	BodyBytesSent int
	HTTPReferer   string
	HTTPUserAgent string
	Created       time.Time
}

//ResponseInt HTTP RESPONSE for messages
type ResponseInt struct {
	StatusCode int            `json:"statusCode"`
	Message    string         `json:"message"`
	Data       map[string]int `json:"data"`
}

//ResponseString HTTP RESPONSE for messages
type ResponseString struct {
	StatusCode int               `json:"statusCode"`
	Message    string            `json:"message"`
	Data       map[string]string `json:"data"`
}

//Response HTTP RESPONSE for messages
type Response struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

//ErrorResponse, response to send when erroring out
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	json       string
}

//NewError generates a new http error
func NewError(code int, e string) ErrorResponse {
	err := ErrorResponse{code, e, ""}
	temp, _ := err.JSON()
	err.json = temp
	return err
}

//ConvertMapInterfaceToMapString converts generic interface to string/string
func ConvertMapInterfaceToMapString(m interface{}) map[string]string {

	result_map := map[string]string{}
	for key, value := range m.(map[string]interface{}) {
		result_map[key] = value.(string)
	}

	return result_map
}

//ConvertMapInterfaceToMapInt converts generic interface to string/int
func ConvertMapInterfaceToMapInt(m interface{}) map[string]int {

	result_map := map[string]int{}
	for key, value := range m.(map[string]interface{}) {
		result_map[key] = int(value.(float64))
	}

	return result_map
}

//JSON returns json version of type
func (r *ResponseInt) JSON() (string, error) {
	jOut, err := json.Marshal(r)
	if err != nil {
		log.Println("Error Unmarshalling data", err)
		return "", err
	}

	return string(jOut), nil
}

//JSON returns json version of type
func (r *ResponseString) JSON() (string, error) {
	jOut, err := json.Marshal(r)
	if err != nil {
		log.Println("Error Unmarshalling data", err)
		return "", err
	}

	return string(jOut), nil
}

//JSON returns json version of type
func (r *Response) JSON() (string, error) {
	jOut, err := json.Marshal(r)
	if err != nil {
		log.Println("Error Unmarshalling data", err)
		return "", err
	}

	return string(jOut), nil
}

//JSON returns json version of type
func (r *ErrorResponse) JSON() (string, error) {
	jOut, err := json.Marshal(r)
	if err != nil {
		log.Println("Error Unmarshalling data", err)
		return "", err
	}

	return string(jOut), nil
}
