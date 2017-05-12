package main

import (
	"fmt"
	"net/http"
)

const listenPort = 3000

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received: ", r.URL.Path)

	if r.URL.Path == "/" {
		http.ServeFile(w, r, "views/index.html")
	} else {
		errorHandler(w, r, http.StatusNotFound)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Favicon request received: ", r.URL.Path)

	http.ServeFile(w, r, "favicon.ico")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)

	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 page not found")
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)

	// Serve static files directly (CSS, JS, images, etc.)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// Fire up the server
	fmt.Printf("Listening on port %d\n", listenPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)

	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
