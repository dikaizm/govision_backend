package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/dikaizm/govision_backend/internal/dto/request"
	middleware_intf "github.com/dikaizm/govision_backend/internal/http/middleware/interfaces"
	"github.com/gorilla/mux"
)

func JsonBodyDecoder(body io.ReadCloser, req any) error {
	if err := json.NewDecoder(body).Decode(&req); err != nil {
		log.Printf("Failed to decode request body: %v", err)
		return err
	}
	return nil
}

func QueryDecoder(r *http.Request, req *map[string]string) error {
	queryParams := r.URL.Query()

	for key, val := range queryParams {
		if len(val) > 0 {
			(*req)[key] = val[0]
		}
	}
	return nil
}

func SendResponse(w http.ResponseWriter, response interface{}, status int) {
	res, err := json.Marshal(response)
	if err != nil {
		log.Printf("Failed to parse response: %v", err)
		http.Error(w, "Failed to parse response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(res)
}

func GetCurrentUser(r *http.Request) (*request.CurrentUser, error) {
	// Extract values from context
	userID, ok := r.Context().Value(middleware_intf.ContextKey.UserID).(string)
	if !ok {
		return nil, errors.New("user id is required or invalid")
	}

	userRole, ok := r.Context().Value(middleware_intf.ContextKey.UserRole).(string)
	if !ok {
		return nil, errors.New("user role is required or invalid")
	}

	return &request.CurrentUser{
		ID:   userID,
		Role: userRole,
	}, nil
}

func UrlVars(r *http.Request, key string) string {
	return mux.Vars(r)[key]
}
