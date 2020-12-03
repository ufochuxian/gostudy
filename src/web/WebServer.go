package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/gomodule/redigo/redis"
	"log"
	"myutil"
	"net/http"
	"time"
)

var callPackSubLesson = func(c *gin.Context) {
	sublids := c.QueryArray("sublid")
	fmt.Println(sublids)
	prepareSend()
	myutil.Receive()
}

func prepareSend() {
	fmt.Println("begin prepareSend")
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("begin send")
	myutil.Send()
}

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		println("welcome to use gin")
	})
	router.GET("/callPackSubLesson", callPackSubLesson)
	log.Fatal(http.ListenAndServe(":9090", router))
}
