package shortlink

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	maxURLSize     = 2048 // the maximum length of bytes in allowed in a URL for Google Chrome
	jsonFormatSize = 32
	maxBodySize    = 2 * maxURLSize // overestimate for the time being - the short url is strictly smaller than the source URL
)

type NewShortRequest struct {
	Link string `json:"link"`
}

type NewShortResponse struct {
	Slug string `json:"slug"`
	Link string `json:"link"`
}

func CreateHandler(w http.ResponseWriter, req *http.Request) {

	shortLink, err := unmarshalRequest(w, req)
	if err != nil {
		return
	}

	t := NewShortResponse{Slug: "foo", Link: shortLink.Link} //todo insert to db here
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(t); err != nil {
		log.Fatalf("Encountered error encoding the created shortlink: %v", err)
	}
}

func unmarshalRequest(w http.ResponseWriter, req *http.Request) (NewShortRequest, error) {
	var shortLink NewShortRequest

	if req.ContentLength > maxBodySize {
		encodeError(w, errors.New(fmt.Sprintf("Request body too large")), http.StatusRequestEntityTooLarge)
	}

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, maxBodySize))
	if err != nil {
		encodeError(w, errors.New(fmt.Sprintf("Encountered error reading request body: %v", err)), http.StatusBadRequest)
		return shortLink, err
	}
	if err := req.Body.Close(); err != nil {
		encodeError(w, errors.New(fmt.Sprintf("Encountered error closing request body: %v", err)), http.StatusInternalServerError)
		return shortLink, err
	}
	if err = json.Unmarshal(body, &shortLink); err != nil {
		encodeError(w, errors.New(fmt.Sprintf("Encountered error unmarshalling the request body: %v", err)), http.StatusUnprocessableEntity)
		return shortLink, err
	}
	return shortLink, nil
}

func encodeError(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(err); err != nil {
		log.Printf("Encountered error while encoding following error into response: %v", err)
	}
}
