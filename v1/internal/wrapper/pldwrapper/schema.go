package pldwrapper

import (
	"errors"
	"io"
	"net/http"
	"os"
	"pld-maker/v1/internal/interface/db"
	"pld-maker/v1/internal/interface/pdf"
	"pld-maker/v1/internal/tools"
	"strings"
)

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

// Display Schema Images
func (cli *Client) Schemas(schemas []db.Schema) {
	if err := os.Mkdir("bin", os.ModePerm); err != nil && err.Error() != "mkdir bin: file exists" {
		panic(err)
	}
	for i, schema := range schemas {
		(*cli.PdfClient).NewPage()
		if i == 0 {
			(*cli.PdfClient).Text(pdf.Text{Data: "1. Schéma Fonctionnel:", Params: &pdf.TextParams{Bold: true}})
			(*cli.PdfClient).NewLine()
		}
		parts := strings.Split(schema.Type, "/")

		if schema.Width > schema.Height {
			// Width is larger, so scale width to maxSize and adjust height accordingly
			schema.Height = (schema.Height * 180) / schema.Width
			schema.Width = 180
		} else {
			// Height is larger or equal, so scale height to maxSize and adjust width accordingly
			schema.Width = (schema.Width * 180) / schema.Height
			schema.Height = 180
		}

		if len(parts) == 2 && parts[0] == "image" {
			tools.Must("", downloadFile(schema.Url, "bin/"+schema.Title+"."+parts[1]))
			(*cli.PdfClient).Text(pdf.Text{Data: schema.Title + ":", Params: &pdf.TextParams{Underline: true}})
			(*cli.PdfClient).Image(pdf.Image{Filepath: "bin/" + schema.Title + "." + parts[1], Width: float64(schema.Width), Height: float64(schema.Height)})
			os.Remove("bin/" + schema.Title + "." + parts[1])
		}
	}
}
