package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	NAME string `json:"username"`
	AGE  uint8  `json:"age"`
}
type Cart struct {
	ID    string  `json:"id"`
	TITLE string  `json:"title"`
	PRICE float32 `json:"price"`
	USER  User    `json:"user"`
}

var users = []User{
	{ID: "1", NAME: "Valentin", AGE: 16},
	{ID: "2", NAME: "Arkado$", AGE: 16},
	{ID: "3", NAME: "Ilya", AGE: 17},
}

var items = []Cart{
	{ID: "1", TITLE: "Watch", PRICE: 32.90, USER: users[0]},
	{ID: "2", TITLE: "Glasses", PRICE: 12.20, USER: users[1]},
	{ID: "3", TITLE: "T-Shirt", PRICE: 52.30, USER: users[2]},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/about", funcAbout)
	router.GET("/", funcHome)
	router.GET("/users/:id", getUser)
	router.POST("/add", pushUser)
  router.POST("/createtask", createT)
  router.GET("/task", getT)
  router.GET("/task/:id", getDetailT)
	router.GET("/items/:id", getDetailItems)
	router.GET("/items/", getItems)

	router.Run("localhost:8000")

}
func createT(c *gin.Context){
  desc := c.PostForm("description")
  title := c.PostForm("title")

  a := createTask(title, desc)
  c.IndentedJSON(http.StatusOK, a)
}
func getT(c *gin.Context){
  a := getTasks()
  c.IndentedJSON(http.StatusOK, a)
}
func getDetailT(c *gin.Context){
  id := c.Param("id")
  taskId, _:= strconv.Atoi(id)
  result := getDetailTask(taskId)
  fmt.Println(result)
  c.IndentedJSON(http.StatusOK, result)

}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func getDetailItems(c *gin.Context) {
	itemId := c.Param("id")
	for _, val := range items {
		if val.ID == itemId {
			c.IndentedJSON(http.StatusOK, val)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Item not found("})
}
func funcAbout(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"about.html",
		gin.H{
			"title": "About",
		},
	)
}

func funcHome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"/": "home"})
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func pushUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
