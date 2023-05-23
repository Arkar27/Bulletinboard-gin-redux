package services

import (
	"os"
	"time"

	"github.com/Arkar27/gin-bulletinboard/backend/dao/authDao"
	"github.com/Arkar27/gin-bulletinboard/backend/models"
	"github.com/Arkar27/gin-bulletinboard/backend/types/request"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginService struct {
	LoginDaoInterface authDao.AuthDaoInterface
}

type Claims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

// Create implements PostServiceInterface.
func (service *LoginService) Authenticate(user request.LoginRequest, ctx *gin.Context) interface{} {
	email := user.Email
	password := user.Password

	userData := service.LoginDaoInterface.Login(email, password, ctx)
	if userData.ID != 0 {
		token, _ := GenerateToken(email, userData.Name)
		retData := models.AuthUser{
			User:  userData,
			Token: token,
		}
		return retData
	} else {
		return struct{}{}
	}
}

func GenerateToken(email string, name string) (string, error) {
	// Set the expiration time for the token (1 day)
	expirationTime := time.Now().Add(time.Hour * 24)

	// Create the claims containing the Email and expiration time
	claims := &Claims{
		Email: email,
		Name:  name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func NewLoginService(LoginDaoInterface authDao.AuthDaoInterface) AuthServiceInterface {
	return &LoginService{
		LoginDaoInterface: LoginDaoInterface,
	}
}
