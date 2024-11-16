package response

type (
	Register struct {
		Name        string `json:"name"`
		Role        string `json:"role"`
		Email       string `json:"email"`
		AccessToken string `json:"access_token"`
	}
)

type (
	Login struct {
		Name        string `json:"name"`
		Role        string `json:"role"`
		Email       string `json:"email"`
		AccessToken string `json:"access_token"`
	}
)
