package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
    cacheFile := "./go-quickstart.json"
    tok, err := tokenFromFile(cacheFile)
    if err != nil {
        tok = getTokenFromWeb(config)
        saveToken(cacheFile, tok)
    }
    return config.Client(ctx, tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
    authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
    fmt.Printf("Go to the following link in your browser then type the "+
        "authorization code: \n%v\n", authURL)
    var code string
    if _, err := fmt.Scan(&code); err != nil {
        log.Fatalf("Unable to read authorization code %v", err)
    }
    tok, err := config.Exchange(oauth2.NoContext, code)
    if err != nil {
        log.Fatalf("Unable to retrieve token from web %v", err)
    }
    return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
    f, err := os.Open(file)
    if err != nil {
        return nil, err
    }
    t := &oauth2.Token{}
    err = json.NewDecoder(f).Decode(t)
    defer f.Close()
    return t, err
}

func saveToken(file string, token *oauth2.Token) {
    fmt.Printf("Saving credential file to: %s\n", file)
    f, err := os.Create(file)
    if err != nil {
        log.Fatalf("Unable to cache oauth token: %v", err)
    }
    defer f.Close()
    json.NewEncoder(f).Encode(token)
}

func main() {
    ctx := context.Background()
    b, err := ioutil.ReadFile("client_secret.json")
    if err != nil {
        log.Fatalf("Unable to read client secret file: %v", err)
    }
    config, err := google.ConfigFromJSON(b, drive.DriveScope)
    if err != nil {
        log.Fatalf("Unable to parse client secret file to config: %v", err)
    }
    client := getClient(ctx, config)
    srv, err := drive.New(client)
    if err != nil {
        log.Fatalf("Unable to retrieve drive Client %v", err)
    }


    // Upload CSV and convert to Spreadsheet
    filename := "search_results.csv"                                       // File you want to upload
    baseMimeType := "text/csv"                                     // mimeType of file you want to upload
    convertedMimeType := "application/vnd.google-apps.spreadsheet" // mimeType of file you want to convert on Google Drive

    file, err := os.Open(filename)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
    defer file.Close()
    f := &drive.File{
        Name:     filename,
        MimeType: convertedMimeType,
    }
    res, err := srv.Files.Create(f).Media(file, googleapi.ContentType(baseMimeType)).Do()
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
    fmt.Printf("%s, %s, %s\n", res.Name, res.Id, res.MimeType)


    // Modify permissions
    permissiondata := &drive.Permission{
        Type:               "domain",
        Role:               "writer",
        Domain:             "google.com",
        AllowFileDiscovery: true,
    }
    pres, err := srv.Permissions.Create(res.Id, permissiondata).Do()
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
    fmt.Printf("%s, %s\n", pres.Type, pres.Role)
}