package main

import "fmt"

type Response struct {
	Code   int
	Status string
}

func f1(a ...int) (s int) {
	fmt.Println(a[:2])
	return 0

}
