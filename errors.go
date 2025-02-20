package jsonres

import (
	"fmt"
)

func InternalServerError() ([]byte, error) {
	res, err := JsonStatusResponse{
		Status:  "error",
		Message: "internal server error",
	}.ToJson()
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON response")
	}
	return res, nil
}

func NotFound() ([]byte, error) {
	res, err := JsonStatusResponse{
		Status:  "error",
		Message: "not found",
	}.ToJson()
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON response")
	}
	return res, nil
}

func BadRequest() ([]byte, error) {
	res, err := JsonStatusResponse{
		Status:  "error",
		Message: "bad request",
	}.ToJson()
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON response")
	}
	return res, nil
}

func Unauthorized() ([]byte, error) {
	res, err := JsonStatusResponse{
		Status:  "error",
		Message: "unauthorized",
	}.ToJson()
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON response")
	}
	return res, nil
}

func Forbidden() ([]byte, error) {
	res, err := JsonStatusResponse{
		Status:  "error",
		Message: "forbidden",
	}.ToJson()
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON response")
	}
	return res, nil
}
