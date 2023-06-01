package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"os"
)

// apply(create or update) connector
func applyConnector(c *cli.Context) error {
	fmt.Println("File path: ", c.String("file"))
	fmt.Println("Trying to create or update a connector using the connector yaml")

	// Getting the account details
	/*reqUrl := GetUrlWithQueryParams("", "", "connectors", map[string]string{
		"accountIdentifier": cliCdRequestData.Account,
	})*/
	reqUrl := GetUrlWithQueryParams("", "", "connectors", map[string]string{
		"accountIdentifier": "vpCkHKsDSxK9_KYfjCTMKA",
	})
	var body = readFromFile(cliCdRequestData.File)
	fmt.Println("reqUrl: ", reqUrl)
	//resp, err := Post(reqUrl, cliCdRequestData.AuthToken, body, "text/yaml")
	resp, err := Post(reqUrl, "pat.vpCkHKsDSxK9_KYfjCTMKA.64769c0d2b0261625f875f13.8IhErJ1DSmu7KKHPLya4", body, "text/yaml")

	fmt.Println("Response Headers: ", resp)
	if err == nil {
		fmt.Printf("Response status: %s \n", resp.Status)
		fmt.Printf("Response code: %s \n", resp.Code)
		fmt.Printf("Response resource: %s \n", resp.Resource)
		fmt.Printf("Response messages: %s \n", resp.Messages)
		printJson(resp.Data)
		return nil
	}

	//fmt.Println("Connector yaml details", connectorDetails)
	return nil
}

// Delete an existing connector
func deleteConnector(*cli.Context) error {
	return nil
}

// Delete an existing connector
func listConnector(*cli.Context) error {
	return nil
}

func readFromFile(filepath string) (s string) {
	bufferedContent := make([]byte, 1024)
	f, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	fileContent := ""
	for {
		numBytes, err1 := f.Read(bufferedContent)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("\n Reached end of the file")
				break
			} else {
				fmt.Println("Error reading from file:", err.Error())
				break
			}
		}

		fileContent = string(bufferedContent[:numBytes])
		//fmt.Printf("File content is %s\n  ", fileContent)
	}
	return fileContent
}
