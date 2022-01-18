package shortlink

import "example.com/shortlink-kata/db"

var Insert = db.InsertNewLink

func createNewShortlink(req *NewShortRequest) (NewShortResponse, error) {
	newSlug, err := generateHash(req.Link)
	if err != nil {
		return NewShortResponse{}, err
	}
	newLink := req.Link

	err = Insert(newSlug, newLink)
	if err != nil {
		return NewShortResponse{}, err
	}

	return NewShortResponse{Slug: newSlug, Link: newLink}, nil
}
