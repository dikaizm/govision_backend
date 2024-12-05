package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
)

// Initialize the logger at the package level
var logger *log.Logger

func init() {
	// Open the log file in append mode, create it if it doesn't exist
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Create a new logger that writes to the file
	logger = log.New(file, "", log.LstdFlags|log.Lshortfile)
}

// LoggingMiddleware logs the request and response
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log Request
		logRequest(r)

		// Create a ResponseRecorder to capture the response
		recorder := &ResponseRecorder{
			ResponseWriter: w,
			StatusCode:     http.StatusOK, // Default status code
			Body:           bytes.NewBuffer(nil),
		}

		// Call the next handler
		next.ServeHTTP(recorder, r)

		// Log Response
		logResponse(recorder)
	})
}

// logRequest logs details of the incoming request
func logRequest(r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Printf("Error reading request body: %v\n", err)
		log.Printf("Error reading request body: %v\n", err)
		return
	}
	// Reset request body for further processing
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	logger.Printf("Request: Method=%s, URL=%s, Headers=%v, Body=%s\n",
		r.Method, r.URL.String(), r.Header, string(body))
	log.Printf("Request: Method=%s, URL=%s, Headers=%v, Body=%s\n",
		r.Method, r.URL.String(), r.Header, string(body))
}

// logResponse logs details of the outgoing response
func logResponse(rec *ResponseRecorder) {
	logger.Printf("Response: Status=%d, Headers=%v, Body=%s\n",
		rec.StatusCode, rec.Header(), rec.Body.String())
	log.Printf("Response: Status=%d, Headers=%v, Body=%s\n",
		rec.StatusCode, rec.Header(), rec.Body.String())
}

// ResponseRecorder wraps http.ResponseWriter to capture response details
type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body       *bytes.Buffer
}

func (rec *ResponseRecorder) WriteHeader(code int) {
	rec.StatusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *ResponseRecorder) Write(data []byte) (int, error) {
	rec.Body.Write(data) // Capture the response body
	return rec.ResponseWriter.Write(data)
}
