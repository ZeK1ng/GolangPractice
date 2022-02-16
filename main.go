package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func fibbonaci(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return fibbonaci(x-1) + fibbonaci(x-2)
}
func factorial(x int) int {
	if x <= 0 {
		return 1
	}
	return x * factorial(x-1)
}
func changeResponse(x *Response) {
	x.Status = "111"
	x.Code = 20
}
func main() {
	x := Response{10, "22"}
	changeResponse(&x)
	fmt.Println(x)

	//setupServer()

}

func setupServer() {
	router := gin.Default()
	router.GET("/test", handle)

	router.Run("localhost:8080")
}

func handle(context *gin.Context) {
	fmt.Println("get")
	resp := Response{200, "abc"}
	params := context.QueryArray("id")
	fmt.Println(params)
	context.IndentedJSON(http.StatusCreated, resp)
}
