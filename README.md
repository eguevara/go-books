# Gobooks 
Gobooks is a Go client library for accessing Google Books Annotations API.

[https://developers.google.com/apis-explorer/?hl=en_US#s/books/v1/](https://developers.google.com/apis-explorer/?hl=en_US#s/books/v1/)

## Usage
```go
 import 	"github.com/eguevara/gobooks"
```

## Authenication

OAuth2 http client is required to to access Google Books API on behalf of a user.

```go
data, err := ioutil.ReadFile("pemFile")
if err != nil {
    log.Fatal(err)
}

conf := &jwt.Config{
		Email:      "xxx@eguevara-books.iam.gserviceaccount.com",
		PrivateKey: []byte(data),
		Scopes: []string{
			"https://www.googleapis.com/auth/books",
		},
		TokenURL: google.JWTTokenURL,
		Subject:  "xxx@gmail.com",
	}
client := conf.Client(oauth2.NoContext)
```

## Examples

```go
client, err := gobooks.New(oauthClient)
if err != nil {
    log.Fatalf("http: error %v", err)
}

opts := &gobooks.AnnotationsListOptions{
    VolumeID:       gobooks.String("volumeId"),
    ContentVersion: gobooks.String("version"),
    LayerID:        gobooks.String("notes"),
    Source:         gobooks.String("app"),
}

list, _, err := client.Annotations.List(opts)
if err != nil {
    log.Fatalf("error in list(): %v ", err)
}
```

