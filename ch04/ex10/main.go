// Reports GitHub issues in the last month, year, and all-time.

package main

import (
	"fmt"
	"go-hello/ch04/ex10/github"
	"log"
	"time"
)

func main() {
	res, err := github.SearchIssues([]string{"repo:golang/go is:open json decoder"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %v issues.\n", res.TotalCount)
	now := time.Now()

	var iMonth, iYear, iAll []*github.Issue
	for _, item := range res.Items {
		days := now.Sub(item.CreatedAt).Hours() / 24
		if days <= 30. {
			iMonth = append(iMonth, item)
		} else if days <= 365. {
			iYear = append(iYear, item)
		} else {
			iAll = append(iAll, item)
		}
	}

	fmt.Println("\nIssues created in the last month:")
	for _, item := range iMonth {
		fmt.Printf("#%-5d %9.9s %55s\n", item.Number, item.User.Login, item.Title)
	}

	fmt.Println("\nIssues created in the last year:")
	for _, item := range iYear {
		fmt.Printf("#%-5d %9.9s %55s\n", item.Number, item.User.Login, item.Title)
	}

	fmt.Println("\nIssues Older:")
	for _, item := range iAll {
		fmt.Printf("#%-5d %9.9s %55s\n", item.Number, item.User.Login, item.Title)
	}
}
