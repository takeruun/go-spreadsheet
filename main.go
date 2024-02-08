package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var spreadsheetID = "1BmvcLURB9uRNZMwz-sAS8DpL785P4Zh6aQQKxgMYjd8"

func main() {
	credential := option.WithCredentialsFile("secret.json")
	srv, err := sheets.NewService(context.TODO(), credential)
	if err != nil {
		log.Fatal(err)
	}
	readRange := "短文!A2:B2"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalln(err)
	}
	if len(resp.Values) == 0 {
		log.Fatalln("data not found")
	}
	for _, row := range resp.Values {
		fmt.Printf("%s, %s\n", row[0], row[1])
	}
}
