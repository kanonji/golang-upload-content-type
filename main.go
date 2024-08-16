package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// multipart/form-dataの処理を行う関数
func handleMultipartFormData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Multipart form data")
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Error parsing multipart form", http.StatusInternalServerError)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileMIMEType := fileHeader.Header.Get("Content-Type")
	contentLength := fileHeader.Size
	fileName := fileHeader.Filename

	output := fmt.Sprintf("Content-Type: %s, Content-Length: %d, File Name: %s\n", fileMIMEType, contentLength, fileName)
	fmt.Println(output)
	fmt.Fprintln(w, output)
}

// 直接送信されたファイルの処理を行う関数
func handleDirectFileUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Direct file upload")
	file, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading the file", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	contentType := r.Header.Get("Content-Type")
	contentLength := len(file)
	fileName := "unknown"

	output := fmt.Sprintf("Content-Type: %s, Content-Length: %d, File Name: %s\n", contentType, contentLength, fileName)
	fmt.Println(output)
	fmt.Fprintln(w, output)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "multipart/form-data") {
		handleMultipartFormData(w, r)
	} else {
		handleDirectFileUpload(w, r)
	}
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	port := flag.String("port", "8080", "Port to run the server on")
	flag.Parse()

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", htmlHandler)
	fmt.Printf("Server started at :%s\n", *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		panic(err)
	}
}
