package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/gomodule/redigo/redis"
	"log"
	"myutil"
	"net/http"
)

var callPackSubLesson = func(c *gin.Context) {
	sublids := c.QueryArray("sublid")
	fmt.Println(sublids)
	prepareSend(sublids)
}

func prepareSend(sublids []string) {
	//fmt.Println("begin prepareSend")
	//timer1 := time.NewTimer(1 * time.Second)
	//<-timer1.C
	//fmt.Println("begin send")
	myutil.Send(sublids)
}

func main() {
	myutil.Receive()
	router := gin.Default()
	router.GET("/callPackSubLesson", callPackSubLesson)
	log.Fatal(http.ListenAndServe(":9090", router))
}
