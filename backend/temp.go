package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"

// )

// type Author struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// var redisClient *redis.Client

// func init() {
// 	redisClient = redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "",
// 		DB:       0,
// 	})
// }

// // func main() {

// // 	router := gin.Default()

// // 	router.GET("/getResource", getResource)
// // 	router.GET("/createMember", createMember)
// // 	router.GET("/getLeaderboard", getLeaderboard)
// // 	router.GET("/testGet/:name", testGet)

// // 	router.Run()

// // 	// redisClient.Do("ZADD", "leaderboard", 33, "user1")
// // 	// redisClient.Do("ZADD", "leaderboard", 454, "user2")

// // 	// total, err := redisClient.Do("ZREVRANK", "leaderboard", "user1").Result()

// // 	// if err == nil {
// // 	// 	fmt.Println("haha")
// // 	// 	// total1 := *total
// // 	// 	fmt.Print(total)
// // 	// }

// // }

// func getResource(c *gin.Context) {

// 	ip := c.ClientIP()

// 	redisClient.SetNX(ip, 0, 5*1000000000)
// 	val, err := redisClient.Get(ip).Result()

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	temp, _ := strconv.Atoi(val)
// 	temp += 1
// 	redisClient.Set(ip, temp, 5*1000000000)

// 	if temp >= 6 {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "You reached the threshold"})
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"message": temp})
// 	}

// }

// func createMember(c *gin.Context) {

// 	// member := redis.Z{
// 	// 	Score:  34,
// 	// 	Member: "Test member",
// 	// }
// 	//redisClient.ZAdd("hello", member)

// 	redisClient.Do("ZADD", "hello", 23, "wassup")
// 	redisClient.Do("ZADD", "hello", 45454, "wtf")

// 	c.JSON(http.StatusOK, gin.H{"message": "Data added"})
// }

// func getLeaderboard(c *gin.Context) {
// 	// scores := redisClient.ZRange("hello", 0, 9)

// 	//score := redisClient.ZScore("hello", "wassup")

// 	score, _ := redisClient.Do("ZREVRANK", "hello", "wassup").Result()

// 	c.JSON(http.StatusOK, gin.H{"message": score})

// 	//c.JSON(http.StatusOK, scores)
// }

// func testGet(c *gin.Context) {
// 	name := c.Param("name")
// 	c.JSON(http.StatusOK, gin.H{"test": name})
// }
