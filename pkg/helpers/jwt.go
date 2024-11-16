package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

type (
	ParamsGenerateJWT struct {
		ExpiredInMinute int
		SecretKey       string
		UserID          string
		UserRole        string
	}

	ResultGenerateJWT struct {
		Token  string
		Expire int64
	}

	ParamsValidateJWT struct {
		Token     string
		SecretKey string
	}

	Claims struct {
		jwt.StandardClaims
		UserID   string `json:"user_id"`
		UserRole string `json:"user_role"`
	}

	ClaimsResult struct {
		UserID   string `json:"user_id"`
		UserRole string `json:"user_role"`
	}
)

func GenerateJWT(p *ParamsGenerateJWT) (ResultGenerateJWT, error) {
	expiredAt := time.Now().Add(time.Duration(p.ExpiredInMinute) * time.Minute).Unix()
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
		},
		UserID:   p.UserID,
		UserRole: p.UserRole,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString([]byte(p.SecretKey))

	return ResultGenerateJWT{
		signedToken,
		expiredAt,
	}, err
}

func ValidateJWT(p *ParamsValidateJWT) (*ClaimsResult, error) {
	token, err := jwt.Parse(p.Token, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || method != JWT_SIGNING_METHOD {
			return nil, errors.New("invalid token")
		}

		return []byte(p.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return &ClaimsResult{
		UserID:   claims["user_id"].(string),
		UserRole: claims["user_role"].(string),
	}, nil
}
