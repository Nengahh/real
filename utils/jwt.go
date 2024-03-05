package helper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const secretKey = "access-login"

func GenerateToken(id uint, email, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// func VerifyToke(c *gin.Context) (interface{}, error) {
// 	errResponse := errors.New("token invalid")
// 	headerToken := c.Request.Header.Get("Authorization")
// 	bearer := strings.HasPrefix(headerToken, "Bearer")

// 	if !bearer {
// 		return nil, errResponse
// 	}

// 	stringToken := strings.Split(headerToken, " ")[1]

// 	token, _ := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errResponse
// 		}
// 		return []byte(secretKey), nil
// 	})

// 	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
// 		return nil, errResponse
// 	}
// 	return token.Claims.(jwt.MapClaims), nil

// }

//  ini code benar
// func VerifyToken(stringToken string) (interface{}, error) {
// 	errResponse := errors.New("token invalid")

// 	token, _ := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errResponse
// 		}
// 		return []byte(secretKey), nil
// 	})

// 	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
// 		return nil, errResponse
// 	}
// 	return token.Claims.(jwt.MapClaims), nil
// }

//sampai sini

func VerifyToken(stringToken string) (jwt.MapClaims, error) {
	errResponse := errors.New("token invalid")

	token, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
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

// package helper

// import (
// 	"net/http"

// 	"github.com/golang-jwt/jwt"
// 	"github.com/jeypc/go-jwt-mux/config"
// 	"github.com/jeypc/go-jwt-mux/helper"
// )

// func JWTMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		c, err := r.Cookie("token")
// 		if err != nil {
// 			if err == http.ErrNoCookie {
// 				response := map[string]string{"message": "Unauthorized"}
// 				helper.ResponseJSON(w, http.StatusUnauthorized, response)
// 				return
// 			}
// 		}
// 		// mengambil token value
// 		tokenString := c.Value

// 		claims := &config.JWTClaim{}
// 		// parsing token jwt
// 		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
// 			return config.JWT_KEY, nil
// 		})

// 		if err != nil {
// 			v, _ := err.(*jwt.ValidationError)
// 			switch v.Errors {
// 			case jwt.ValidationErrorSignatureInvalid:
// 				// token invalid
// 				response := map[string]string{"message": "Unauthorized"}
// 				helper.ResponseJSON(w, http.StatusUnauthorized, response)
// 				return
// 			case jwt.ValidationErrorExpired:
// 				// token expired
// 				response := map[string]string{"message": "Unauthorized, Token expired!"}
// 				helper.ResponseJSON(w, http.StatusUnauthorized, response)
// 				return
// 			default:
// 				response := map[string]string{"message": "Unauthorized"}
// 				helper.ResponseJSON(w, http.StatusUnauthorized, response)
// 				return
// 			}
// 		}

// 		if !token.Valid {
// 			response := map[string]string{"message": "Unauthorized"}
// 			helper.ResponseJSON(w, http.StatusUnauthorized, response)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }
