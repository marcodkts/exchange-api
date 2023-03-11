package middleware

import (
	"exchange-api/initializers"
	"exchange-api/models"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		fmt.Println("Cookie error:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		fmt.Println("Token error:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if !token.Valid {
		fmt.Println("Token is invalid")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Claims are invalid")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		fmt.Println("Token is expired")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User
	initializers.DB.First(&user, claims["sub"])

	if time.Now().Day() > user.LastResetTime.Day() {
		initializers.DB.Model(&user).Updates(models.User{
			UserRequestCount: 0,
			LastResetTime:    time.Now(),
		})
	}

	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	path := c.Request.URL.Path

	if strings.Contains(path, "/exchange") {
		today := time.Now().Weekday()

		if today == time.Saturday || today == time.Sunday {
			if user.UserRequestCount >= 200 {
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Request quota exceeded"})
				return
			}
			initializers.DB.Model(&user).Updates(models.User{
				UserRequestCount: user.UserRequestCount + 1,
			})

		} else {
			if user.UserRequestCount >= 100 {
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Request quota exceeded"})
				return
			}
			initializers.DB.Model(&user).Updates(models.User{
				UserRequestCount: user.UserRequestCount + 1,
			})
		}
	}

	c.Set("user", user)
	c.Next()
}