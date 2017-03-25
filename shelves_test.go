package books

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestShelvesList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/mylibrary/bookshelves", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"items":[{"id":7,"title":"My Google eBooks","volumeCount":13},{"id":1,"title":"Purchased","volumeCount":11}]}`)
	})

	opts := &ShelvesListOptions{Fields: "items(description,id,title,volumeCount)"}

	list, _, err := client.Shelves.List(opts)
	if err != nil {
		t.Errorf("List() returned an error: %v", err)
	}

	expected := []Shelf{{ID: 7, Title: "My Google eBooks", VolumeCount: 13}, {ID: 1, Title: "Purchased", VolumeCount: 11}}

	if !reflect.DeepEqual(list, expected) {
		t.Errorf("List() returned %+v, expected %+v", list, expected)
	}
}

func TestShelvesList_badBody(t *testing.T) {
	setup()
	defer teardown()

	opts := &ShelvesListOptions{}
	_, resp, err := client.Shelves.List(opts)

	// Check that response is error on nil request body
	if err == nil {
		t.Error("List() Expected Request body error.")
	}

	// Check that response status code is http.StatusNotFound.
	if got, want := resp.StatusCode, http.StatusNotFound; got != want {
		t.Errorf("List() Expected Status code got %v, want %v", got, want)
	}
}
