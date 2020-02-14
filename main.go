package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

// EvenMinute returns minute if minute is Even number.
// 現在分が偶数であれば分を返す。そうでなければ奇数を返す。
func EvenMinute() *int {
	t := time.Now()
	minute := t.Minute()
	if minute%2 == 0 {
		return &minute
	}
	return nil
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		type NullableJson struct {
			NullableIDInit    *int `json:nullable_id_init`
			NullableIDNotInit *int `json:nullable_id_not_init`
			NotNullIDInit     int  `json:not_null_id_init`
			NotNullIDNotInit  int  `json:not_null_id_not_init`
		}

		nullableJson := new(NullableJson)

		nullableJson.NullableIDInit = EvenMinute()
		nullableJson.NotNullIDInit = 1

		c.JSON(200, nullableJson)
	})
	r.Run()
}
