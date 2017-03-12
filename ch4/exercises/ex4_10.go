// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.10, p. 112: Modify issues to report the results in age categories,
// say less than a month old, less than a year old, and more than a year old.
// Based on issues.

// $ go run ex4_10.go -- repo:golang/go is:open json decoder

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d total issues\n", result.TotalCount)
	// Group issues into three age categories.
	thisMonth, thisYear, older := []*github.Issue{}, []*github.Issue{}, []*github.Issue{}
	now := time.Now()
	// One month defined as 30 days.
	oneMonth, _ := time.ParseDuration("-720h")
	oneMonthAgo := now.Add(oneMonth)
	// One year defined as 365 days.
	oneYear, _ := time.ParseDuration("-8760h")
	oneYearAgo := now.Add(oneYear)
	// Group issues by age.
	for _, item := range result.Items {
		if item.CreatedAt.Before(oneYearAgo) {
			older = append(older, item)
		} else if item.CreatedAt.Before(oneMonthAgo) {
			thisYear = append(thisYear, item)
		} else {
			thisMonth = append(thisMonth, item)
		}
	}
	fmt.Println("Last month issues:")
	for _, item := range thisMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("Last year issues:")
	for _, item := range thisYear {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("Older issues:")
	for _, item := range older {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
