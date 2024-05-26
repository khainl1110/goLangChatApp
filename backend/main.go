package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TextMessageBody struct {
	Message string `json:"message"`
}

func test(c *gin.Context) {

	// member := redis.Z{
	// 	Score:  34,
	// 	Member: "Test member",
	// }
	//redisClient.ZAdd("hello", member)

	// redisClient.Do("ZADD", "hello", 23, "wassup")
	// redisClient.Do("ZADD", "hello", 45454, "wtf")

	var requestBody TextMessageBody

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Print("haha")
		fmt.Println(err)
	}

	publish("api-project-70002766628", "test-topic", requestBody.Message)
	// pullMsgsCustomAttributes("api-project-70002766628", "test-subscription")
	c.JSON(http.StatusOK, gin.H{"message": "Data addfdfdded"})
}

func main() {
	fmt.Println("Hello, World!")
	publish("api-project-70002766628", "test-topic", "Hello Whuniohorld:3453fd")
	// print(err.Error())
	print("____________")
	pullMsgsCustomAttributes("api-project-70002766628", "test-subscription")
	router := gin.Default()

	router.POST("/test", test)

	router.Run()
}
