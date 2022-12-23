package main

import (
	"fmt"
	"log"
	"net/http"
)

type calc struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

func main() {
	http.HandleFunc("/", homePage)//router
	log.Fatal(http.ListenAndServe(":10000", nil))

}
func addition(num1 int, num2 int) int { //service
	return num1 + num2
}
func subtraction(num1 int, num2 int)int{
    return num2 - num1
}
func homePage(w http.ResponseWriter, r *http.Request) {//controller
	calculation := calc{Num1: 1, Num2: 6} //request

	sum := addition(calculation.Num1, calculation.Num2)
    sub := subtraction(calculation.Num1, calculation.Num2)
	fmt.Fprintf(w, "addtion: %d\n", sum)
    fmt.Fprintf(w, "subtraction: %d",sub)
}
