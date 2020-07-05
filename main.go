package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	var available []string
	// Check per 50 time
	for i := 0; i < 50; i++ {
		//Username length is 5
		randomdata, isOK := random(5)
		if !isOK {
			fmt.Println("🛑 Username length must be between 1 and 39")
			break
		}
		status, name := findPositive(randomdata)
		// Status 1 : available, Status: 0 Taken
		if status == 1 {
			fmt.Println(i, "-", "github.com/"+name, "is available ✔")
			available = append(available, name)
		} else {
			fmt.Println(i, "-", "github.com/"+name, "was taken ❌")
		}
		time.Sleep(1500 * time.Millisecond)
	}
	fmt.Println("Available Username List: ", available)
}

func findPositive(userName string) (int, string) {
	counter := 0
	res, err := http.Get("https://github.com/" + userName)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
	} else {
		counter++
	}
	return counter, userName
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func random(length int) (string, bool) {
	var isOk bool
	if length >= 1 && length <= 39 {
		const charset = "abcdefghijklmnopqrstuvwxyz"
		b := make([]byte, length)
		for i := range b {
			b[i] = charset[seededRand.Intn(len(charset))]
		}
		isOk = true
		return string(b), isOk
	} else {
		isOk = false
	}
	return "", isOk
}
