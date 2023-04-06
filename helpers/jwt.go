package helpers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "rahasia"

func GenerateToken(id uint, email string, role string) (string, string) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken, role
}

func VerifyToken(c *gin.Context) (jwt.MapClaims, error) {
	errResponse := errors.New("unauthorized access")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errResponse
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errResponse
	}

	return claims, nil
}
