package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alecthomas/chroma/lexers"
	god "github.com/alecthomas/chroma/lexers/go"
	"github.com/russross/blackfriday/v2"
)

func main() {
	http.HandleFunc("/", markdownHandler)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func markdownHandler(w http.ResponseWriter, r *http.Request) {
	// Read the Markdown file
	mdFile, err := ioutil.ReadFile("article.md")
	if err != nil {
		http.Error(w, "Failed to read Markdown file", http.StatusInternalServerError)
		return
	}

	// Convert Markdown to HTML
	html := blackfriday.Run(mdFile)

	// Apply syntax highlighting to Go code blocks
	highlightedHTML := highlightGoCode(html)

	// Set the Content-Type header to indicate HTML response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(highlightedHTML)
}

func highlightGoCode(html []byte) []byte {
	// Create a new Chroma HTML formatter
	formatter := html.New(html.WithClasses())

	// Retrieve the Go lexer from Chroma
	lexer := lexers.Get("go")
	if lexer == nil {
		 lexer = god.New()
	}

	// Create a new Chroma iterator for syntax highlighting
	iterator, err := lexer.Tokenise(nil, string(html))
	if err != nil {
		// Handle lexer error
		fmt.Println("Lexer error:", err)
	}

	// Create a new buffer for storing the highlighted code
	buffer := &bytes.Buffer{}

	// Apply syntax highlighting using the Chroma formatter
	err = formatter.Format(buffer, chromaStyle, iterator)
	if err != nil {
		// Handle formatter error
		fmt.Println("Formatter error:", err)
	}

	return buffer.Bytes()
}
