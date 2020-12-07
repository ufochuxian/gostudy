package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/gomodule/redigo/redis"
	"log"
	"myutil"
	"net/http"
	"tgmnet"
)

var callPackSubLesson = func(c *gin.Context) {
	var sublessoninfos []*myutil.Sublessoninfo
	err := c.BindJSON(&sublessoninfos)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(sublessoninfos))
	prepareSend(sublessoninfos, func(resp tgmnet.ResponseOld) {
		c.JSON(http.StatusOK, resp)
	})
}

func prepareSend(sublids []*myutil.Sublessoninfo, callback tgmnet.Callback) {
	//fmt.Println("begin prepareSend	")
	//timer1 := time.NewTimer(1 * time.Second)
	//<-timer1.C
	fmt.Println("begin send")
	myutil.Send(sublids, callback)
}

func main() {
	myutil.Receive()
	router := gin.Default()
	router.POST("/call_pack_sub_lesson", callPackSubLesson)
	log.Fatal(http.ListenAndServe(":9090", router))
}
