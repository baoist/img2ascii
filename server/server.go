package server

import (
	"encoding/json"
	"image"
	"log"
	"net/http"

	_ "image/jpeg"
	_ "image/png"

	processor "github.com/baoist/img2ascii/image_processor"
	units "github.com/docker/go-units"
)

type HttpResponseError struct {
	Message    string `json:"error"`
	statusCode int
}

func handleError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	requestError := &HttpResponseError{
		Message:    message,
		statusCode: statusCode,
	}

	json.NewEncoder(w).Encode(requestError)
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	maxFileSize, err := units.FromHumanSize("2MB")
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxFileSize)
	err = r.ParseMultipartForm(1024)
	if err != nil {
		handleError(w, err.Error(), http.StatusRequestEntityTooLarge)
		return
	}

	multipart := r.MultipartForm

	files := multipart.File["image"]
	file, err := files[0].Open()
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	img, _, err := image.Decode(file)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ascii_image := processor.Convert(img)

	json.NewEncoder(w).Encode(ascii_image)
}

func StartServer() {
	http.HandleFunc("/process", processHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
