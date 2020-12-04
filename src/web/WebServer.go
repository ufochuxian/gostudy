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
	var sublessoninfos []*myutil.Sublessoninfo
	err := c.BindJSON(&sublessoninfos)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(sublessoninfos))
	prepareSend(sublessoninfos)
}

func prepareSend(sublids []*myutil.Sublessoninfo) {
	//fmt.Println("begin prepareSend")
	//timer1 := time.NewTimer(1 * time.Second)
	//<-timer1.C
	fmt.Println("begin send")
	myutil.Send(sublids)
}

func main() {
	myutil.Receive()
	router := gin.Default()
	router.POST("/callPackSubLesson", callPackSubLesson)
	log.Fatal(http.ListenAndServe(":9090", router))
}
