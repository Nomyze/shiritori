package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"maj9.org/shiritori/shiritori"
)

type Word struct {
	word string `uri:"word" binding:"required"`
}

func main() {
	shiritori.Read_dict("./dicts/JMdict_e")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/dict/:word", func(c *gin.Context) {
		var word = c.Param("word")
		fmt.Println(word)
		var term = shiritori.Get_entry(word)
		if reflect.ValueOf(term).IsZero() {
			c.JSON(400, gin.H{"msg": "Not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"word": term.Words,
			"reading": term.Readings,
			"definition": term.Definitions,
		})
	})

	router.Run()
}
