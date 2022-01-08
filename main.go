package main

import (
	"bufio"
	"fmt"
	"mangtas-test/src/service"
	"os"
	"strings"
)

func main() {
	svc := service.NewService()
	fmt.Println("Top Ten Words")
	fmt.Println("Enter [exit] to stop the service-")
	fmt.Println("---------------------")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input-> ")
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println()
		if strings.EqualFold(text, "exit") {
			break
		} else {
			res := svc.GetTopTenWords(text)
			svc.PrintWords(res)
		}
		fmt.Print("\nInput-> ")
	}

}
