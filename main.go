package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Set content type explicitly
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
	// Set content type explicitly
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Render the course html page
	http.ServeFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Set content type explicitly
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Render the about html page
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Set content type explicitly
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Render the contact html page
	http.ServeFile(w, r, "static/contact.html")
}

func main() {
	// Add a root route
	http.HandleFunc("/", homePage) // Serve the home page for the root

	// Handle other pages
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/courses", coursePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	// Start the server
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
