package main

import "fmt"

func main() {

	list := []any{AccountID, ActiveConnections}

	if contains(list, AccountID) {
		fmt.Println("AccountID is in the list")
	} else {
		fmt.Println("AccountID is not in the list")
	}
}

type MeasuredType string

const (
	AccountID         MeasuredType = "ACCOUNT_ID"
	ActiveConnections MeasuredType = "ACTIVE_CONNECTIONS"
)

func contains(list []any, item any) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}
