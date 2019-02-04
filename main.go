package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // initialize global pseudo random generator
	words := loadWords()
	num := numWordsToGet()
	for i := 0; i < num; i++ {
		fmt.Println(randomChoice(words))
	}
}

func randomChoice(l []string) string {
	return l[rand.Int()%len(l)]
}

func numWordsToGet() int {
	if len(os.Args) == 1 {
		return 1
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func loadWords() []string {
	var w []string

	file, err := os.Open("./common_words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		w = append(w, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return w
}
