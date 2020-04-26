package books

// AnnotationsService defines the behavior required by types that want to implement a new Annotation type.
type AnnotationsService interface {
	List(*AnnotationsListOptions) ([]Annotation, *Response, error)
}

// GoogleAnnotationsService implements the AnnotationService interface.
type GoogleAnnotationsService struct {
	client *Client
}

// Annotation represents a Google Book Annotation resource.
type Annotation struct {
	SelectedText *string  `json:"selectedText,omitempty"`
	VolumeID     *string  `json:"VolumeID,omitempty"`
	ID           *string  `json:"id,omitempty"`
	LayerID      *string  `json:"layerId,omitempty"`
	PageIds      []string `json:"pageIds,omitempty"`
}

// annotationRoot represents a response from Google Books API.
type annotationRoot struct {
	TotalItems    *int         `json:"totalItems,omitempty"`
	NextPageToken *string      `json:"nextPageToken,omitempty"`
	Annotations   []Annotation `json:"items,omitempty"`
}

// AnnotationsListOptions specifies the optional parameters needed for AnnotationsListOptions.
type AnnotationsListOptions struct {
	ContentVersion string `url:"contentVersion,omitempty"`
	LayerID        string `url:"layerId,omitempty"`
	MaxResults     int    `url:"maxResults,omitempty"`
	Source         string `url:"source,omitempty"`
	VolumeID       string `url:"volumeId,omitempty"`
	Fields         string `url:"fields,omitempty"`
	PageToken      string `url:"pageToken,omitempty"`
}

// List will call Annotation service with opts param.
// books.mylibrary.annotations.list
func (u *GoogleAnnotationsService) List(opt *AnnotationsListOptions) ([]Annotation, *Response, error) {
	url := "mylibrary/annotations"
	url, err := addOptions(url, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := u.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(annotationRoot)
	resp, err := u.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	if n := root.NextPageToken; n != nil {
		resp.NextPageToken = *n
	}

	return root.Annotations, resp, err
}
