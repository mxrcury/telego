package utils

import (
	"math/rand"
	"strings"
)

func GetRandomName(names []string ) (string, string) {
  randomIndex := rand.Intn(len(names))
  randomName := strings.Split(names[randomIndex], " ")
  return strings.Trim(randomName[0], " "), strings.Trim(randomName[1], " ")
}
