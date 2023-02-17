package main

import (
	api "Hello/api"
	es "Hello/error"
	chat "Hello/room"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func main() {
	//poe.ParseXML()
	api.TestSwitch()
	api.Test_printChanNums()
	api.Test_threeNumberSum()

	//api.TestMySQLAPI()
}

func startChatRoom() {
	chat.Run()

	router := gin.New()

	indexAbs, err := filepath.Abs("resource/index.html")
	if err != nil {
		es.ErrorToString(err)
		return
	}

	router.LoadHTMLFiles(indexAbs)

	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		chat.ServeWs(c.Writer, c.Request, roomId)
	})

	err = router.Run("0.0.0.0:9898")
	if err != nil {
		es.ErrorToString(err)
		return
	}

}
