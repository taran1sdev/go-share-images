package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my site</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>Contact <a href=\"mailto:thommsairddev@gmail.com\">thomasairddev@gmail.com</>.</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprint(w, `
		<h1>FAQ</h1>

		<ul>
			<li>Is there a free version? <strong>Yes, there is!</strong></li>
			<li>What are your support hours? 24/7!</li>
			<li>How do I contact support? Send an email to <a href="mailto:support@site.com">Support</a>.</li>
		</ul>
	`)
}

func faqIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	questionID := chi.URLParam(r, "QID")
	printStr := `<h1>FAQ</h1>`

	switch questionID {
	case "1":
		printStr += `
		<ul>
			<li>Is there a free version? Yes, there is!</li>
		</ul>`
	case "2":
		printStr += `
		<ul>
			<li>What are your support hours? <strong>24/7!</strong></li>
		</ul>`
	case "3":
		printStr += `
		<ul>
			<li>How do I contact support? <strong>Send an email to <a href="mailto:support@site.com">Support</a></strong></li>
		</ul>`
	default:
		printStr += `
		<ul>
			<li>Is there a free version? <strong>Yes, there is!</strong></li>
			<li>What are your support hours? 24/7!</li>
			<li>How do I contact support? Send an email to <a href="mailto:support@site.com">Support</a>.</li>
		</ul>`
	}

	fmt.Fprint(w, printStr)
}

func main() {
	r := chi.NewRouter()

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)

	r.Get("/faq", faqHandler)
	r.Get("/faq/{QID}", faqIdHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting http server on port 3000")
	http.ListenAndServe(":3000", r)
}
