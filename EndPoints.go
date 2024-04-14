package main

import (
	"SplitKaro/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

// connect to db
func main() {
	r := gin.Default()

	r.GET("/Users/GetAllUsers", func(c *gin.Context) {
		// Create a Response struct instance
		response := Models.User{}.ReadAllUsers()
		// Send the response as JSON
		if response == nil {
			c.JSON(http.StatusOK, "No records found")
		}
		c.JSON(http.StatusOK, response)
	})

	r.POST("Users/CreateUser", func(c *gin.Context) {
		var user Models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		Models.User{}.InsertUser(user)
		c.JSON(http.StatusOK, "record created")
	})

	r.GET("Users/GetUserByEmail", func(c *gin.Context) {
		var email string
		if err := c.ShouldBindJSON(&email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(email)
		response, err := Models.User{}.ReadUserByEmail(email)
		if err != nil {
			c.JSON(http.StatusOK, "no user found")
		}

		c.JSON(http.StatusOK, response)
	})

	r.POST("Users/CreateGroup", func(c *gin.Context) {
		var group Models.Group
		if err := c.ShouldBindJSON(&group); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i := 0; i < len(group.Email); i++ {
			user, err := Models.User{}.ReadUserByEmail(group.Email[i])
			if err == nil {
				group.User = append(group.User, user)

			}
		}

		err := Models.Group{}.CreateGroup(group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, "group created")
	})
	r.POST("Balance/AddBalance", func(c *gin.Context) {
		var balacne Models.Balance
		if err := c.ShouldBindJSON(&balacne); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		Models.Balance{}.AddBalance(balacne)
		c.JSON(http.StatusOK, "record created")
	})

	// Run the server
	err := r.Run(":8080")
	if err != nil {
		return
	}

}
