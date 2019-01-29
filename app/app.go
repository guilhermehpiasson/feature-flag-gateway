package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func GetFFState(c *gin.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	id := c.Params.ByName("id")

	val, err := client.Get("ff:" + id).Result()
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(500, gin.H{"nok - Error": "GET api/v1/ff: - id: " + id})
	} else {
		fmt.Println("ff:ff-test", val)
		c.JSON(200, gin.H{"ok": "GET api/v1/ff: - id: " + id + " - value: " + val})
	}
}
