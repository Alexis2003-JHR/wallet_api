package middlewares

import (
	"crypto/rsa"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username  string `json:"username"`
	TokenType string `json:"token_type"`
	jwt.StandardClaims
}

type JWTAuth struct {
	publicKey *rsa.PublicKey
}

func NewJWTAuth(publicKeyPath string) (*JWTAuth, error) {
	keyData, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, errors.New("error leyendo la clave pública")
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return nil, errors.New("error parseando la clave pública")
	}

	return &JWTAuth{publicKey: publicKey}, nil
}

func (auth *JWTAuth) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return auth.publicKey, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error al autenticar el token"})
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !token.Valid || !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		}

		if claims.TokenType != "access_token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido por tipo de token"})
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
