package main

import (
	"fmt"
	"log"
	"os"

	gogithub "github.com/sony-nurdianto/GoRust/github/go_github"
)

func main() {
	result, err := gogithub.SearchIssues(os.Args[:1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
