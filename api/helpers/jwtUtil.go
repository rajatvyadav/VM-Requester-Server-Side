package helpers

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"

	models "VMCreationWorkflow/api/model"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

var GlobalJWTKey string

func init() {
	GlobalJWTKey = "TODO"
}

type jwtCustomClaim struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	UserRole string `json:"userrole"`
	Token    string `json:"token"`
	jwt.StandardClaims
}

func GenerateToken(username string, password string, userrole string, expirationTime time.Duration) (string, error) {
	claims := jwtCustomClaim{
		UserName: username,
		Password: password,
		UserRole: userrole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(GlobalJWTKey))
	if err != nil {
		fmt.Println("Error during token generation", err)
	}
	return t, nil
}

// GetLoginFromToken login object from JWT Token
func GetLoginFromToken(c *gin.Context) (models.User, error) {

	login := models.User{}
	decodedToken, err := DecodeToken(c.GetHeader("Authorization"), GlobalJWTKey)
	fmt.Println("GetLoginFromToken -- decodedToken", decodedToken)
	if err != nil {
		return login, errors.New("GetLoginFromToken - unable to decode token")
	}
	// login ID is the compulsary field, so haven't added check for nil
	if decodedToken["username"] == nil || decodedToken["username"] == "" {
		return login, errors.New("GetLoginFromToken - login id not found")
	}
	login.UserName = decodedToken["username"].(string)
	login.Password = decodedToken["password"].(string)
	login.UserRole = decodedToken["userrole"].(string)
	return login, nil
}

func DecodeToken(tokenFromRequest, jwtKey string) (jwt.MapClaims, error) {

	// get data i.e.Claims from token
	token, err := jwt.Parse(tokenFromRequest, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})

	if err != nil {
		fmt.Println("Error while parsing JWT Token: ", err)
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Error while getting claims")
	}
	return claims, nil
}
