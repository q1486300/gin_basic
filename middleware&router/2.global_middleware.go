package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.Use(m10, m11)

	router.GET("/index1", func(c *gin.Context) {
		fmt.Println("index1 ...in")

		name, _ := c.Get("name")
		fmt.Println(name)

		_user, _ := c.Get("user")
		user := _user.(User)
		fmt.Println(user.Name, user.Age)

		c.JSON(http.StatusOK, gin.H{"msg": user})

		c.Next()
		fmt.Println("index1 ...out")
	})

	router.GET("/index2", func(c *gin.Context) {
		fmt.Println("index2 ...in")

		c.JSON(http.StatusOK, gin.H{"msg": "index2"})

		c.Next()
		fmt.Println("index2 ...out")
	})

	router.Run(":8080")
}

type User struct {
	Name string
	Age  int
}

func m10(c *gin.Context) {
	fmt.Println("m10 ...in")

	//c.JSON(http.StatusOK, gin.H{"msg": "響應被 m10 中間件給中斷了"})
	//c.Abort()

	c.Set("name", "名稱")
	c.Set("user", User{
		Name: "名稱",
		Age:  26,
	})

	c.Next()
	fmt.Println("m10 ...out")
}

func m11(c *gin.Context) {
	fmt.Println("m11 ...in")
	c.Next()
	fmt.Println("m11 ...out")
}
