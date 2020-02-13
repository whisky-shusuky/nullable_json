package main

import "github.com/gin-gonic/gin"

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

		tmpInt := 1
		nullableJson.NullableIDInit = &tmpInt
		nullableJson.NotNullIDInit = 1

		c.JSON(200, nullableJson)
	})
	r.Run()
}
