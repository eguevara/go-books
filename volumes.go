package gobooks

import (
	"errors"
	"fmt"
)

// VolumesService defines the behavior required by types that want to implement a new Volumes type.
type VolumesService interface {
	List(string, *VolumesListOptions) ([]Volume, *Response, error)
}

// GoogleVolumesService implements the VolumesService interface.
type GoogleVolumesService struct {
	client *Client
}

// Volume represents a Google Book Volume resource.
type Volume struct {
	ID   string     `json:"id"`
	Info VolumeInfo `json:"volumeInfo"`
}

// VolumeInfo represents a google.book.volumes.volumeInfo
type VolumeInfo struct {
	Title          string `json:"title"`
	ContentVersion string `json:"contentVersion"`
}

// volumesRoot represents a response from Google Books API.
type volumesRoot struct {
	TotalItems int      `json:"totalItems"`
	Volumes    []Volume `json:"items"`
}

// VolumesListOptions specifies the optional parameters needed to make API request.
// books.mylibrary.bookshelves.volumes.list
type VolumesListOptions struct {
	Shelf      int    `url:"shelf,omitempty"`
	MaxResults int    `url:"maxResults,omitempty"`
	Quey       string `url:"q,omitempty"`
	Source     string `url:"source,omitempty"`
	Fields     string `url:"fields, omitempty,omitempty"`
}

// List will call the books.mylibrary.bookshelves.volumes.list API.
func (v *GoogleVolumesService) List(volumeID string, opt *VolumesListOptions) ([]Volume, *Response, error) {
	if volumeID == "" {
		return nil, nil, errors.New("volumeID is a required field")
	}

	url := fmt.Sprintf("mylibrary/bookshelves/%s/volumes", volumeID)
	url, err := addOptions(url, opt)

	req, err := v.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(volumesRoot)
	resp, err := v.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Volumes, resp, err
}
