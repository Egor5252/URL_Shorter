package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func RandomWord() (string, error) {
	file, err := os.Open("pkg\\utils\\words.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	num := rand.Intn(318946) + 1

	scanner := bufio.NewScanner(file)
	current := 1
	for scanner.Scan() {
		if current == num {
			return scanner.Text(), nil
		}
		current++
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("строка номер %d не найдена", num)
}
