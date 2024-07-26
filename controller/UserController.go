package controller

import (
	"github.com/chuxin-cs/go-crud/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func InitUser(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Register Swagger routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/addUser", addUser(db))
		v1.GET("/deleteUser", deleteUser(db))
		v1.GET("/getUserList", getUserList(db))
		v1.GET("/updateUser", updateUser(db))
	}

	return r
}

func getUserList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []entity.User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"list": users,
		})
	}
}

// 删除用户
func deleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
		})
	}
}

// 创建用户
func addUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entity.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}

// 更新用户
func updateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
		})
	}
}
