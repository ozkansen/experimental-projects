package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrEnvironmentNotFound = errors.New("environment not found")

// getDNUrl format: "postgres://postgres:ExamplePassword!@localhost:5432/postgres"
func getDBUrl() (string, error) {
	val, ok := os.LookupEnv("DB_URL")
	if !ok {
		return "", fmt.Errorf("DB_URL %v", ErrEnvironmentNotFound)
	}
	return val, nil
}
