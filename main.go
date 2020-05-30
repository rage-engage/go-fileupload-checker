package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func uploadPath(w http.ResponseWriter, r *http.Request) {

	// Max upload size of 20mb
	r.ParseMultipartForm(20 << 20)

	// File handlers
	file, handler, err := r.FormFile("fileName")

	// Error handling
	if err != nil {
		fmt.Fprintf(w, "Upload error occured %v", err)
		return
	}

	if file == nil {
		fmt.Fprintf(w, "Please upload a file")
		return
	}

	defer file.Close()
	fmt.Println("Saved file to memory")
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	fmt.Printf("Filename: %+v\n", handler.Filename)

	// Create our types
	type response struct {
		Path     string `json:"Path"`
		FileName string `json:"fileName"`
	}

	resp := response{
		Path:     "in-memory",
		FileName: "Uploaded: " + handler.Filename,
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	port := ":8080"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/upload", uploadPath).Methods("PUT")
	fmt.Println("Listening on: ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
