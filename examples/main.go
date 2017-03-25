package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/eguevara/go-books"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

const (
	clientEmail      = "eguevara-books-p12-new@eguevara-books.iam.gserviceaccount.com"
	impersonateEmail = "erick.guevara@gmail.com"
)

func exampleShelvesList(c *books.Client) {
	opts := &books.ShelvesListOptions{}

	shelves, _, err := c.Shelves.List(opts)
	if err != nil {
		log.Fatalf("error in List(): %v", err)
	}

	for _, v := range shelves {
		fmt.Printf("Id: %d, Title: %v, VolumeCount: %d\n", v.ID, v.Title, v.VolumeCount)
	}
}

func exampleVolumesList(c *books.Client) {
	opts := &books.VolumesListOptions{
		Fields:     "items(id,volumeInfo(contentVersion,title)),totalItems",
		MaxResults: 1,
	}

	volumes, _, err := c.Volumes.List("1", opts)
	if err != nil {
		log.Fatalf("error in List(): %v", err)
	}

	for _, v := range volumes {
		fmt.Printf("VolumeId: %s, Title: %v, ContentVersion: %v\n", v.ID, v.Info.Title, v.Info.ContentVersion)
	}
}

func exampleAnnotationsList(c *books.Client) {
	opts := &books.AnnotationsListOptions{
		VolumeID:       "VN2jCgAAAEAJ",
		ContentVersion: "full-1.0.0",
		LayerID:        "notes",
		MaxResults:     1,
		Source:         "ge-web-app1",
		Fields:         "items(layerId,selectedText,volumeId),totalItems",
	}

	list, _, err := c.Annotations.List(opts)
	if err != nil {
		log.Fatalf("error in list(): %v ", err)
	}

	for idx, note := range list {
		fmt.Printf("%d - %s\n\n", idx, note.SelectedText)
	}
}

func main() {

	oauthClient := getOAuthClient()
	client, err := books.New(oauthClient)
	if err != nil {
		log.Fatalf("http: error %v", err)
	}

	exampleAnnotationsList(client)
	exampleVolumesList(client)
	exampleShelvesList(client)

}

func getOAuthClient() *http.Client {
	data, err := ioutil.ReadFile("key.pem")
	if err != nil {
		log.Fatal(err)
	}

	conf := &jwt.Config{
		Email:      clientEmail,
		PrivateKey: data,
		Scopes: []string{
			"https://www.googleapis.com/auth/books",
		},
		TokenURL: google.JWTTokenURL,
		Subject:  impersonateEmail,
	}
	client := conf.Client(oauth2.NoContext)

	return client
}
