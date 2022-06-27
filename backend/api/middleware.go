package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthErrorResponse struct {
	Error string `json:"error"`
}

//untuk allow origin front end
func (api *API) alloworigin(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	if c.Request.Method == "OPTIONS" {
		c.Writer.WriteHeader(200)
	}
}

//Jtw middleware
func (api *API) AuthMiddleWare(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		api.alloworigin(c)
		token, err := c.Request.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				c.Writer.WriteHeader(http.StatusUnauthorized)
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "anda belum login",
				})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		tknStr := token.Value

		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Writer.WriteHeader(http.StatusUnauthorized)
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": err.Error(),
				})
				return
			}
			c.Writer.WriteHeader(http.StatusBadRequest)
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": err.Error(),
			})
			return
		}

		if !tkn.Valid {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "token invalid!",
			})
			return
		}

		ctx := context.WithValue(c, "username", claims.Username)
		ctx = context.WithValue(ctx, "role", claims.Role)
		next(c)
	}
}

func (api *API) AdminMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		api.alloworigin(c)

		token, _ := c.Request.Cookie("token")

		tknStr := token.Value

		claims := &Claims{}

		_, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Writer.WriteHeader(http.StatusUnauthorized)
				c.JSON(http.StatusUnauthorized, gin.H{
					"status":  "false",
					"code":    http.StatusUnauthorized,
					"message": err.Error(),
				})
				return
			}
			c.Writer.WriteHeader(http.StatusBadRequest)
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "false",
				"code":    http.StatusUnauthorized,
				"message": err.Error(),
			})
			return
		}

		if claims.Role != "admin" {
			c.Writer.WriteHeader(http.StatusForbidden)
			c.JSON(http.StatusForbidden, gin.H{
				"status":  "false",
				"code":    http.StatusForbidden,
				"message": "forbidden access",
			})
			return
		}
		next(c)
	}
}
