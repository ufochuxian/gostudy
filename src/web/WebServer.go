package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/gomodule/redigo/redis"
	"log"
	"net/http"
)

var callPackSubLesson = func(c *gin.Context) {
	sublids := c.QueryArray("sublid")
	fmt.Println(sublids)
}

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		println("welcome to use gin")
	})
	router.GET("/callPackSubLesson", callPackSubLesson)
	log.Fatal(http.ListenAndServe(":9090", router))
}
