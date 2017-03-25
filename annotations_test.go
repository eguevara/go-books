package books

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAnnotations_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/mylibrary/annotations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"totalItems":1,"items":[{"volumeId":"VN2jCgAAAEAJ","layerId":"notes","selectedText":"Go"}]}`)
	})

	opts := &AnnotationsListOptions{
		VolumeID:       "SJHvCgAAQBAJ",
		ContentVersion: "test11",
		LayerID:        "notes",
		Source:         "ge-web-app1",
	}
	list, _, err := client.Annotations.List(opts)
	if err != nil {
		t.Errorf("List() returned an error: %v", err)
	}

	expected := []Annotation{{VolumeID: "VN2jCgAAAEAJ", LayerID: "notes", SelectedText: "Go"}}

	if !reflect.DeepEqual(list, expected) {
		t.Errorf("List() returned %+v, expected %+v", list, expected)
	}
}

func TestAnnotationsList_badBody(t *testing.T) {
	setup()
	defer teardown()

	opts := &AnnotationsListOptions{}
	_, resp, err := client.Annotations.List(opts)

	// Check that response is error on nil request body
	if err == nil {
		t.Error("List() Expected Request body error.")
	}

	// Check that response status code is http.StatusNotFound.
	if got, want := resp.StatusCode, http.StatusNotFound; got != want {
		t.Errorf("List() Expected Status code got %v, want %v", got, want)
	}
}
