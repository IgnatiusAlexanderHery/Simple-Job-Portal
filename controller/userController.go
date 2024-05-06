package controller

import (
	"Simple-Job-Portal/auth"
	"Simple-Job-Portal/database"
	"Simple-Job-Portal/model"
	"Simple-Job-Portal/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	csrf "github.com/srbry/gin-csrf"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		utils.BadRequestError(c, err.Error())
		return
	}

	if user.Username == "" || user.Password == "" || user.Role == "" {
		utils.BadRequestError(c, "All fields (username, password, role) are required")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.BadRequestError(c, err.Error())
		return
	}
	user.Password = string(hashedPassword)

	if user.Role != "talent" && user.Role != "employer" {
		utils.BadRequestError(c, "Invalid role. Must be either 'talent' or 'employer'")
		return
	}

	user.UUID = uuid.New().String()

	err = database.AddUser(&user)
	if err != nil {
		utils.BadRequestError(c, err.Error())
		return
	}

	user.Password = ""
	utils.SuccessfulResponse(c, user)
}

func Login(c *gin.Context, csrfManager *csrf.DefaultCSRFManager) {
	_, tokenErr := c.Cookie("token")
	_, csrfErr := c.Cookie("X-CSRF-Token")
	if tokenErr == nil && csrfErr == nil {
		utils.UnauthorizedError(c, "Already logged in")
		return
	}
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		utils.BadRequestError(c, err.Error())
		return
	}
	if user.Username == "" || user.Password == "" {
		utils.BadRequestError(c, "All fields (username, password) are required")
		return
	}

	inputPassword := user.Password

	foundUser, err := database.GetUserByUsername(user.Username)

	if err != nil {
		utils.UnauthorizedError(c, "Invalid username or password")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(inputPassword))
	if err != nil {
		utils.UnauthorizedError(c, "Invalid username or password")
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &model.Claims{
		Username: foundUser.Username,
		Role:     foundUser.Role,
		UUID:     foundUser.UUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(auth.GetJWTKey())
	if err != nil {
		utils.UnauthorizedError(c, "Error generating token")
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
		Secure:   true,
	})
	utils.SuccessfulResponse(c, gin.H{"message": "Login successful", "X-CSRF-Token": csrfManager.GetToken(c)})
}

func Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(),
	})
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "X-CSRF-Token",
		Value:   "",
		Expires: time.Now(),
	})
	utils.SuccessfulResponse(c, gin.H{"message": "Logout successful"})
}

func GetUserData(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		utils.UnauthorizedError(c, "Unauthorized")
		return
	}

	tokenString := cookie
	claims := &model.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return auth.GetJWTKey(), nil
	})
	if err != nil && !token.Valid {
		utils.UnauthorizedError(c, "Unauthorized")
		return
	}

	c.Set("user", claims)
	user, _ := c.Get("user")
	utils.SuccessfulResponse(c, user)
}
