package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dikaizm/govision_backend/internal/dto/response"
	middleware_intf "github.com/dikaizm/govision_backend/internal/http/middleware/interfaces"
	"github.com/dikaizm/govision_backend/pkg/helpers"
)

func Authentication(secretKey string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helpers.SendResponse(w, response.Response{
				Status:  "error",
				Message: "Authorization header is required",
				Error:   "Unauthorized",
			}, http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate JWT token
		claims, err := helpers.ValidateJWT(&helpers.ParamsValidateJWT{
			Token:     token,
			SecretKey: secretKey,
		})

		if err != nil {
			helpers.SendResponse(w, response.Response{
				Status:  "error",
				Message: "Invalid token",
			}, http.StatusUnauthorized)
			return
		}

		// Store claims in context
		ctx := r.Context()
		ctx = context.WithValue(ctx, middleware_intf.ContextKey.UserID, claims.UserID)
		ctx = context.WithValue(ctx, middleware_intf.ContextKey.UserRole, claims.UserRole)

		// Pass the context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
