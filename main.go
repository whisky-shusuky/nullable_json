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

		type NullableJSON struct {
			NullableIntInit    *int `json:nullable_id_init`
			NullableIntNotInit *int `json:nullable_id_not_init`
			NotNullIntInit     int  `json:not_null_id_init`
			NotNullIntNotInit  int  `json:not_null_id_not_init`
		}

		nullableJSON := new(NullableJSON)

		nullableJSON.NullableIntInit = EvenMinute()
		nullableJSON.NotNullIntInit = 1

		c.JSON(200, nullableJSON)
	})
	r.Run()
}
