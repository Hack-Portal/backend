package controllers

import (
	"bytes"
	"net/http"
)

func CheckFile(r *http.Request) ([]byte, error) {
	file, _, err := r.FormFile("image")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	icon := bytes.NewBuffer(nil)
	if _, err := icon.ReadFrom(file); err != nil {
		return nil, err
	}
	return icon.Bytes(), nil
}
