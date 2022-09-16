package middlewares

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type JwtClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

var (
	JWT_SECRET         = []byte(os.Getenv("JWT_SECRET"))
	JWT_EXP            = time.Duration(1) * time.Hour
	JWT_SIGNING_METHOD = jwt.SigningMethodHS256
)

func getTokenString(authHeader string) (*string, error) {
	var token string
	if strings.Contains(authHeader, "Bearer") {
		token = strings.Replace(authHeader, "Bearer ", "", -1)
		return &token, nil
	}
	return nil, fmt.Errorf("authorization not found")
}

func CreateToken(id uint) string {
	claims := &JwtClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(JWT_SECRET)
	if err != nil {
		log.Error(err)
	}

	return t
}

func ParseJWTToken(authHeader string) (*JwtClaims, error) {
	tokenString, err := getTokenString(authHeader)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("invalid signing method")
		}
		return JWT_SECRET, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimsStr, err := json.Marshal(claims)
		if err != nil {
			return nil, fmt.Errorf("error when marshalling token")
		}

		var customClaims JwtClaims
		if err := json.Unmarshal(claimsStr, &customClaims); err != nil {
			return nil, fmt.Errorf("error when unmarshalling token")
		}

		return &customClaims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func JWTMiddleware(claims JwtClaims, signingKey []byte) echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &JwtClaims{},
		SigningKey: signingKey,
	}
	return middleware.JWTWithConfig(config)
}