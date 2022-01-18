package shortlink

import "example.com/shortlink-kata/db"

var insert = db.InsertNewLink
var getHash = generateHash
var getLink = db.GetLink

func createNewShortlink(req *NewShortRequest) (NewShortResponse, error) {
	var newSlug string
	if req.CustomSlug == "" {
		var err error
		newSlug, err = getHash(req.Link)
		if err != nil {
			return NewShortResponse{}, err
		}
	} else {

		newSlug = req.CustomSlug //todo check that the URL is valid -- trust user for now
	}

	//todo check that newSlug does not collide with server paths

	newLink := req.Link //todo check that the URL is a valid URL -- trust user for now

	err := insert(newSlug, newLink)
	if err != nil {
		return NewShortResponse{}, err
	}

	return NewShortResponse{Slug: newSlug, Link: newLink}, nil
}

func getRedirectLink(slug string) (string, error) {
	return getLink(slug)
}
