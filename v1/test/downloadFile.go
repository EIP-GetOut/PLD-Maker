package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fileName := "downloaded.png"
	URL := "https://v5.airtableusercontent.com/v1/20/20/1694880000000/QjBDs8Bx0ay6csYpQ6ebMQ/qbDEupysDgnoCx4LyhtPYFpkGQquq2NvfTyeVckAqB7peWluCvc4MM5Ohsv1PRrvfjHBuxTwQpNqcvIaB_5eZP_1E-hY_WKmeGSx1g_ZaaA_VtQZ9H1o1mjByZqe9FQX/eqmmZ3V-eCj_z9-neT-LaQykFcEqzqtS0NvBNlzSB4s"
	err := downloadFile(URL, fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File %s downlaod in current working directory", fileName)
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
