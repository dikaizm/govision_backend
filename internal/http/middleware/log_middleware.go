package middleware

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// LoggerManager manages the logger and its associated file
type LoggerManager struct {
	logger     *log.Logger
	logFile    *os.File
	currentDay string
	sync.Mutex
}

// NewLoggerManager initializes and returns a LoggerManager
func NewLoggerManager() *LoggerManager {
	lm := &LoggerManager{}
	if err := lm.RotateLogFile(); err != nil {
		log.Fatalf("Failed to initialize LoggerManager: %v", err)
	}
	return lm
}

// Log logs a message safely
func (lm *LoggerManager) Log(format string, v ...interface{}) {
	// Ensure log file is rotated
	if err := lm.RotateLogFile(); err != nil {
		log.Printf("Failed to rotate log file: %v\n", err)
		return
	}

	// Safeguard against nil logger
	if lm.logger == nil {
		log.Printf("Logger is nil; message: "+format, v...)
		return
	}

	lm.logger.Printf(format, v...)
}

// RotateLogFile rotates the log file when the date changes
func (lm *LoggerManager) RotateLogFile() error {
	lm.Lock()
	defer lm.Unlock()

	currentDay := time.Now().Format("2006-01-02")
	if lm.currentDay == currentDay {
		return nil // No rotation needed
	}

	// Close the current log file
	if lm.logFile != nil {
		lm.logFile.Close()
	}

	// Ensure the logs directory exists
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create logs directory: %w", err)
		}
	}

	// Open a new log file
	logFileName := logDir + currentDay + ".log"
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	lm.logger = log.New(file, "", log.LstdFlags|log.Lshortfile)
	lm.logFile = file
	lm.currentDay = currentDay

	return nil
}

// Initialize the global LoggerManager
var loggerManager = NewLoggerManager()

// LoggingMiddleware logs the request and response
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Recovered from panic: %v\n", rec)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		// Log the request
		logRequest(r)

		// Create a ResponseRecorder to capture the response
		recorder := &ResponseRecorder{
			ResponseWriter: w,
			StatusCode:     http.StatusOK, // Default status code
			Body:           bytes.NewBuffer(nil),
		}

		// Call the next handler
		next.ServeHTTP(recorder, r)

		// Log the response
		logResponse(recorder)
	})
}

// logRequest logs details of the incoming request
func logRequest(r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		loggerManager.Log("Error reading request body: %v\n", err)
		log.Printf("Error reading request body: %v\n", err)
		return
	}
	// Reset request body for further processing
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	loggerManager.Log("Request: Method=%s, URL=%s, Headers=%v, Body=%s\n",
		r.Method, r.URL.String(), r.Header, string(body))
	log.Printf("Request: Method=%s, URL=%s, Headers=%v, Body=%s\n", r.Method, r.URL.String(), r.Header, string(body))
}

// logResponse logs details of the outgoing response
func logResponse(rec *ResponseRecorder) {
	loggerManager.Log("Response: Status=%d, Headers=%v, Body=%s\n",
		rec.StatusCode, rec.Header(), rec.Body.String())
	log.Printf("Response: Status=%d, Headers=%v, Body=%s\n", rec.StatusCode, rec.Header(), rec.Body.String())
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
