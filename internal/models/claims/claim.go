package claims

import "github.com/dgrijalva/jwt-go"

// UserClaims ...
type UserClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
	Role  int32  `json:"role"`
}
