package auth

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("clave_secreta_super_segura")

func GenerateToken(email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": email,
        "exp":   time.Now().Add(time.Hour * 24).Unix(),
    })

    return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
}
