package main

// https://pkg.go.dev/golang.org/x/crypto/bcrypt
import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := []byte("123senha")

	hash, err := bcrypt.GenerateFromPassword(password, 10)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(hash))

	bcrypt.CompareHashAndPassword(hash, []byte("123senha"))

}
