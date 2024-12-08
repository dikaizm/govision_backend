package response

type (
	Register struct {
		UserID           string `json:"user_id"`
		Name             string `json:"name"`
		Role             string `json:"role"`
		Email            string `json:"email"`
		Photo            string `json:"photo"`
		CompletedProfile bool   `json:"completed_profile"`
		AccessToken      string `json:"access_token"`
	}
)

type (
	Login struct {
		UserID           string `json:"user_id"`
		Name             string `json:"name"`
		Role             string `json:"role"`
		Email            string `json:"email"`
		Photo            string `json:"photo"`
		AccessToken      string `json:"access_token"`
		CompletedProfile bool   `json:"completed_profile"`
	}
)
