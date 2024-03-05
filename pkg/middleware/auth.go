// package middleware

// import (
// 	"net/http"
// 	helper "real_nimi_project/utils"

// 	"github.com/gin-gonic/gin"
// )

// func Auntentifikation() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		vertifyToken, err := helper.VerifyToke(c)
// 		_ = vertifyToken

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error":   "Unauthorization",
// 				"message": err.Error(),
// 			})

// 		}

// 		c.Set("useData", vertifyToken)
// 		c.Next()
// 	}

// }

// package middleware

// import (
// 	"net/http"

// 	"net/http"
// 	"strings"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt"
// )

// func Authorization() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         verifyToken, err := helper.VerifyToke(c)
//         if err != nil {
//             c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
//                 "error":   "Unauthorized",
//                 "message": err.Error(),
//             })
//             return
//         }

//         claims, ok := verifyToken.(jwt.MapClaims)
//         if !ok || !verifyToken.Valid {
//             c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
//                 "error":   "Unauthorized",
//                 "message": "Invalid token",
//             })
//             return
//         }

//         role, ok := claims["role"].(string)
//         if !ok || role != "admin" {
//             c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
//                 "error":   "Forbidden",
//                 "message": "You are not authorized to access this page",
//             })
//             return
//         }

//         // Jika pengguna adalah admin, lanjutkan ke handler berikutnya
//         c.Next()
//     }
// }

// package middleware
// package middleware

// import (
// 	"errors"
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt"
// )

// const secretKey = "access-login"

// func JWTMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		errResponse := errors.New("token invalid")
// 		headerToken := c.Request.Header.Get("Authorization")
// 		bearer := strings.HasPrefix(headerToken, "Bearer")

// 		if !bearer {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 			return
// 		}

// 		stringToken := strings.Split(headerToken, " ")[1]

// 		token, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, errResponse
// 			}

// 			return []byte(secretKey), nil

// 		})

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 			return
// 		}

// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			c.Set("claims", claims)
// 			c.Next()
// 		} else {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 		}
// 	}
// }

// package middleware

// import (
// 	"net/http"
// 	"strings"

// 	helper "real_nimi_project/utils"
// 	// "real_nimi_project/utils/helper"

// 	"github.com/gin-gonic/gin"
// )

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		headerToken := c.Request.Header.Get("Authorization")
// 		if headerToken == "" {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
// 			return
// 		}

// 		bearer := strings.HasPrefix(headerToken, "Bearer")
// 		if !bearer {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
// 			return
// 		}

// 		stringToken := strings.Split(headerToken, " ")[1]
// 		claims, err := helper.VerifyToken(stringToken)
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			return
// 		}

// 		c.Set("claims", claims)
// 		c.Next()
// 	}
// }

package middleware

import (
	"net/http"

	helper "real_nimi_project/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieToken, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token cookie provided"})
			return
		}

		claims, err := helper.VerifyToken(cookieToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}

}

func AuthMiddlewareadmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieToken, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token cookie provided"})
			return
		}

		claims, err := helper.VerifyToken(cookieToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Check if the user's role is "admin"
		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this resource"})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}

}
