package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)


func LogMessage(message any){
	log.Println(message)
}

func ErrorNotNil(err error) bool {
	return err != nil
}

func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}


func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if ErrorNotNil(err) {
		return "", err
	}
	return string(hashedPassword), nil
}

type CustomClaims struct {
	UUID string `json:"uuid"`
	jwt.RegisteredClaims
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		validateAuthHeader(authHeader, ctx)

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := getAndParseClaims(tokenString)

		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			ctx.Set("uuid", claims.UUID)
		}

		// Check if the token is valid
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func validateAuthHeader(authHeader string, ctx *gin.Context) {
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		ctx.Abort()
		return
	}
}

func GenerateUuid() string {
	return uuid.NewString()
}

func getAndParseClaims(tokenString string) (*jwt.Token, error){
	return jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
}