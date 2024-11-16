package middleware_intf

type contextKey string

type sContextKey struct {
	UserID   contextKey
	UserRole contextKey
}

var ContextKey = sContextKey{
	UserID:   "userID",
	UserRole: "userRole",
}
