package main

import "fmt"
import "github.com/lapingvino/etime"

func main() {
	now := etime.Now()
	fmt.Println(now.Element(), now)
}
