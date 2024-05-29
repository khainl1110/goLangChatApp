package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type TextMessageBody struct {
	Message string `json:"message"`
}

func dump() {
	// member := redis.Z{
	// 	Score:  34,
	// 	Member: "Test member",
	// }
	//redisClient.ZAdd("hello", member)

	// redisClient.Do("ZADD", "hello", 23, "wassup")
	// redisClient.Do("ZADD", "hello", 45454, "wtf")
}

func getMsg(c *gin.Context) {
	result := pullMsg("api-project-70002766628", "test-subscription")
	c.JSON(http.StatusOK, gin.H{"message": result})
}

func postMsgReq(c *gin.Context) {
	var requestBody TextMessageBody

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Print("haha")
		fmt.Println(err)
	}

	publishMsg("api-project-70002766628", "test-topic", requestBody.Message)

	c.JSON(http.StatusOK, gin.H{"message": "Message sent"})
}

func main() {
	fmt.Println("Hello, World!")
	// publishMsg("api-project-70002766628", "test-topic", "Hello Whuniohorld:3453fd")
	// print("____________")
	// pullMsg("api-project-70002766628", "test-subscription")
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/sendMsg", postMsgReq)
	router.GET("/getMsg", getMsg)

	router.Run()
}
