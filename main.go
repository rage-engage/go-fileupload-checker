package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	defaultPort := 8080

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // All origins
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		// 	AllowedMethods: []string{"GET"}, // Allowing only get, just an example
	})
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/upload", uploadPath).Methods("PUT")

	// Optional port flag
	flagPort := flag.Int("p", defaultPort, "port number, default 8080")
	flag.Parse()
	port := fmt.Sprintf(":%v", *flagPort)

	fmt.Println("Listening on: ", port)
	log.Fatal(http.ListenAndServe(port, c.Handler(router)))
}
