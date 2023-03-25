package main

import (
	_ "embed"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/spec"
	"sigs.k8s.io/yaml"
)

// Product represents a product in the e-commerce store
// swagger:model
type Product struct {
    ID          int64   `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description,omitempty"`
    Price       float64 `json:"price"`
}

// ErrorResponse represents an error response from the server
// swagger:model
type ErrorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}


func getProducts(w http.ResponseWriter, r *http.Request) {
    // ...
}


func getProductByID(w http.ResponseWriter, r *http.Request) {
    // ...
}

func serveDocs(w http.ResponseWriter, r *http.Request) {
    // Load the Swagger spec from the YAML file
    specJson, err := yaml.YAMLToJSON(embeddedSwaggerSpec)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var specDoc spec.Swagger
    if err := json.Unmarshal(specJson, &specDoc); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Serve the Swagger UI
    uiHandler := middleware.SwaggerUI(middleware.SwaggerUIOpts{
        SpecURL:  "/swagger.json",
        BasePath: "/",
    }, nil)
    uiHandler.ServeHTTP(w, r)
}

//go:embed "docs/e-commerce.yaml"
var embeddedSwaggerSpec []byte

func main() {
    // Create a new HTTP server
    server := http.NewServeMux()

    // Register the API handlers
    server.HandleFunc("/products", getProducts)
    server.HandleFunc("/products/{id}", getProductByID)

    // Register the Swagger documentation handler
    server.HandleFunc("/docs", serveDocs)
    server.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write(embeddedSwaggerSpec)
    })

    // Start the server
    http.ListenAndServe(":8080", server)
}
