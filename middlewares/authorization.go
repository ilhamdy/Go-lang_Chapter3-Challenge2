package middlewares

import (
	"net/http"
	"strconv"

	"jwt/database"
	"jwt/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": "invalid parameter",
			})
			return
		}

		// Ambil data pengguna dari token JWT
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		userRole := userData["role"].(string)

		// Jika pengguna memiliki peran "admin", maka pengguna dapat mengakses semua produk
		if userRole == "admin" {
			c.Next()
			return
		}

		// Jika pengguna memiliki peran "user", maka pengguna hanya dapat mengakses produk dengan "Create" dan "Get Product by Id"
		if c.Request.Method != "POST" && c.Request.Method != "GET" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		Product := models.Product{}
		err = db.Select("user_id").First(&Product, uint(productId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data doesn't exist",
			})
			return
		}

		if Product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}
