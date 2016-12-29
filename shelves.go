package gobooks

// ShelvesService defines the behavior required by types that want to implement a new Shelf type.
type ShelvesService interface {
	List(*ShelvesListOptions) ([]Shelf, *Response, error)
}

// GoogleShelvesService implements the VolumesService interface.
type GoogleShelvesService struct {
	client *Client
}

// Shelf represents a Google Book Volume resource.
// https://developers.google.com/books/docs/v1/reference/bookshelves#resource
type Shelf struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	VolumeCount int    `json:"volumeCount"`
	Description string `json:"description"`
	Updated     string `json:"updated"`
}

// shelvesRoot represents a response from Google Books API.
// https://developers.google.com/books/docs/v1/reference/mylibrary/bookshelves/list#response
type shelvesRoot struct {
	Shelves []Shelf `json:"items"`
}

// ShelvesListOptions specifies the optional parameters needed to make API request.
// books.mylibrary.bookshelves.list
type ShelvesListOptions struct {
	Source string `url:"source,omitempty"`
	Fields string `url:"fields, omitempty,omitempty"`
}

// List will call the books.mylibrary.bookshelves.list API.
// https://www.googleapis.com/books/v1/mylibrary/bookshelves
func (v *GoogleShelvesService) List(opt *ShelvesListOptions) ([]Shelf, *Response, error) {
	url := "mylibrary/bookshelves"
	url, err := addOptions(url, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := v.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(shelvesRoot)
	resp, err := v.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Shelves, resp, err
}
