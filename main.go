package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Route for authentication
	http.HandleFunc("/auth", authHandler)

	// Route for API
	http.HandleFunc("/api", apiHandler)

	// Specify the path to your SSL certificate and key files
	certFile := "/Users/siva/sidecarauth/certs/server.crt"
	keyFile := "/Users/siva/sidecarauth/certs/server.key"

	// Start the HTTPS server
	port1 := 8445
	fmt.Printf("Listening on :%d...\n", port1)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", port1), certFile, keyFile, nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authentication handler")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API handler")
}

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getXMLFiles(directory string) []string {
	var xmlFiles []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".xml" {
			xmlFiles = append(xmlFiles, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

	return xmlFiles
}

func main() {
	datapowerDirectory := "/path/to/datapower/deployment"
	xmlFilesList := getXMLFiles(datapowerDirectory)

	fmt.Println("XML Files in Datapower Deployment:")
	for _, xmlFile := range xmlFilesList {
		fmt.Println(xmlFile)
	}
}


package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getXMLFiles(directory string) map[string][]string {
	xmlFilesByParent := make(map[string][]string)

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.Name() == "temp" {
			parentDir := filepath.Base(filepath.Dir(path))
			files, err := os.ReadDir(path)
			if err != nil {
				return err
			}
			for _, file := range files {
				if !file.IsDir() && filepath.Ext(file.Name()) == ".xml" {
					xmlFilesByParent[parentDir] = append(xmlFilesByParent[parentDir], filepath.Join(path, file.Name()))
				}
			}
			return filepath.SkipDir // Skip subdirectories of "temp"
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

	return xmlFilesByParent
}

func main() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	xmlFilesByParent := getXMLFiles(currentDirectory)

	fmt.Println("XML Files in 'temp' Subdirectory:")
	for parentDir, xmlFiles := range xmlFilesByParent {
		fmt.Printf("Parent Folder: %s\n", parentDir)
		for _, xmlFile := range xmlFiles {
			fmt.Println(xmlFile)
		}
		fmt.Println()
	}
}
