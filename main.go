package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Course struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
}

var courses = []Course{
	{ID: "1", Name: "armboy", Description: "test d bug"},
	{ID: "2", Name: "offboy", Description: "test da bug"},
	{ID: "3", Name: "wayboy", Description: "test dd bug"},
}

func main() {
	r := gin.Default()

	r.GET("/courses", listCourses)
	r.GET("/courses/:id", getCourses)

	r.Run(":8080")
}

func listCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func getCourses(c *gin.Context) {
	id := c.Param("id")
	for _, course := range courses {
		if course.ID == id {
			c.IndentedJSON(http.StatusOK, course)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "course not found",
	})
}
