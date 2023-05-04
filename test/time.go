package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now().Add(time.Hour)
	// dayBeneran := day.Unix() % 60
	fmt.Println(day.Unix())
}
