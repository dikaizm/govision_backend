package response

type (
	Register struct {
		UserID      string `json:"user_id"`
		Name        string `json:"name"`
		Role        string `json:"role"`
		Email       string `json:"email"`
		AccessToken string `json:"access_token"`
	}
)

type (
	Login struct {
		UserID      string `json:"user_id"`
		Name        string `json:"name"`
		Role        string `json:"role"`
		Email       string `json:"email"`
		AccessToken string `json:"access_token"`
	}
)
