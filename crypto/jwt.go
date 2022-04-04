package crypto

import (
	"errors"
	"fmt"
	"go_sample_login_register/enums"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type JSONWebKeys struct {
	Keys []JSONWebKey `json:"keys"`
}

type JSONWebKey struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type Payload struct {
	UserID string
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func extractHttpRequestAuthToken(request *http.Request) (string, error) {
	token := request.Header.Get("Authorization")
	fields := strings.Fields(token)
	if len(fields) != 2 || fields[0] != "Bearer" {
		return "", errors.New("Invalid authorization header ")
	}
	return fields[1], nil
}

func GetHttpRequestAuthorizationClaim(request *http.Request) (*Payload, enums.ResultCode, error) {
	authToken, err := extractHttpRequestAuthToken(request)
	if err != nil {
		return nil, enums.INVALID_AUTH_TOKEN, err
	}

	JWTKey := []byte(os.Getenv("JWT_KEY"))

	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})
	if err != nil {
		return nil, enums.INVALID_AUTH_TOKEN, err
	}
	if !token.Valid {
		return nil, enums.INVALID_AUTH_TOKEN, errors.New("Auth token is invalid ")
	}

	claims := token.Claims.(jwt.MapClaims)

	payload := &Payload{
		UserID: fmt.Sprintf("%v", claims["user_id"]),
	}
	return payload, enums.SUCCESS, nil
}

func GenerateJWT(userID string) (tokenString string, err error) {
	JWTKey := []byte(os.Getenv("JWT_KEY"))

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewTime(float64(expirationTime.Unix())),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(JWTKey)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func convertInterfaceArray(list interface{}) []string {
	tempList := list.([]interface{})
	resultList := make([]string, len(tempList))
	for _, data := range tempList {
		resultList = append(resultList, data.(string))
	}
	return resultList
}
