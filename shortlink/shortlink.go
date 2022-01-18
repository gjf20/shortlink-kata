package shortlink

import "example.com/shortlink-kata/db"

var insert = db.InsertNewLink
var getHash = generateHash

func createNewShortlink(req *NewShortRequest) (NewShortResponse, error) {
	var newSlug string
	if req.CustomSlug == "" {
		var err error
		newSlug, err = getHash(req.Link)
		if err != nil {
			return NewShortResponse{}, err
		}
	} else {
		newSlug = req.CustomSlug
	}

	newLink := req.Link //todo check that the URL is a valid URL -- trust user for now

	err := insert(newSlug, newLink)
	if err != nil {
		return NewShortResponse{}, err
	}

	return NewShortResponse{Slug: newSlug, Link: newLink}, nil
}
