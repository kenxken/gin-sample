package main

import (
	"strconv"

	_ "net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type person struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to database¥n")
	}
	defer db.Close()
	db.AutoMigrate(&person{})

	router.GET("/", func(c *gin.Context) {
		var people []person
		db.Find(&people)
		c.HTML(200, "index.tmpl", gin.H{
			"people": people,
		})
	})
	router.POST("/new", func(c *gin.Context) {
		name := c.PostForm("name")
		age, _ := strconv.Atoi(c.PostForm("age"))
		db.Create(&person{Name: name, Age: age})
		c.Redirect(302, "/")
	})

	router.Run(":8080")
}
