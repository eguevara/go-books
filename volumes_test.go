package gobooks

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestVolumesList(t *testing.T) {
	setup()
	defer teardown()

	volumeID := "1"
	url := fmt.Sprintf("/mylibrary/bookshelves/%s/volumes", volumeID)
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"totalItems":1,"items":[{"id":"VN2jCgAAAEAJ","volumeInfo":{"title":"Go in Action","contentVersion":"full-1.0.0"}}]}`)
	})

	opts := &VolumesListOptions{
		Fields:     "items(id,volumeInfo(contentVersion,title)),totalItems",
		MaxResults: 1,
	}

	list, _, err := client.Volumes.List(volumeID, opts)
	if err != nil {
		t.Errorf("List() returned an error: %v", err)
	}

	expected := []Volume{{ID: "VN2jCgAAAEAJ", Info: VolumeInfo{Title: "Go in Action", ContentVersion: "full-1.0.0"}}}

	if !reflect.DeepEqual(list, expected) {
		t.Errorf("List() returned %+v, expected %+v", list, expected)
	}
}

func TestVolumesList_badBody(t *testing.T) {
	setup()
	defer teardown()

	opts := &VolumesListOptions{}
	_, resp, err := client.Volumes.List("1", opts)

	// Check that response is error on nil request body
	if err == nil {
		t.Error("List() Expected Request body error.")
	}

	// Check that response status code is http.StatusNotFound.
	if got, want := resp.StatusCode, http.StatusNotFound; got != want {
		t.Errorf("List() Expected Status code got %v, want %v", got, want)
	}
}

func TestVolumesList_emptyVolume(t *testing.T) {
	setup()
	defer teardown()

	opts := &VolumesListOptions{}
	_, _, err := client.Volumes.List("", opts)

	// Check that response is error on nil request body
	if err == nil {
		t.Error("List() Expected Request body error.")
	}

}
