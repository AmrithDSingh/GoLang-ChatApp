package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Question struct {
	Text      string   `json:"text"`
	Responses []string `json:"responses"`
}

var questions []Question

func main() {
	router := gin.Default()

	router.POST("/questions", func(c *gin.Context) {
		var question Question
		if err := c.BindJSON(&question); err != nil {
			c.AbortWithError(400, err)
			return
		}
		questions = append(questions, question)
		c.JSON(200, gin.H{"message": "Question added successfully"})
	})

	router.GET("/questions", func(c *gin.Context) {
		c.JSON(200, gin.H{"questions": questions})
	})

	router.POST("/questions/:id/responses", func(c *gin.Context) {
		id := c.Param("id")
		index, err := findQuestionIndex(id)
		if err != nil {
			c.AbortWithError(400, err)
			return
		}
		var response string
		if err := c.BindJSON(&response); err != nil {
			c.AbortWithError(400, err)
			return
		}
		questions[index].Responses = append(questions[index].Responses, response)
		c.JSON(200, gin.H{"message": "Response added successfully"})
	})

	router.Run()
}

func findQuestionIndex(id string) (int, error) {
	for i, q := range questions {
		if q.Text == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Question not found: %s", id)
}
