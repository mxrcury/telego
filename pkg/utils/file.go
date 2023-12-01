package utils

import (
	"bufio"
	"os"
)

func ReadDocumentByLine(name string) ([]string, error) {
  file, err := os.OpenFile(name, os.O_RDONLY, os.ModeDevice)
  if err != nil {
    return nil, err
  }
  var lines []string
  scanner := bufio.NewScanner(file)
  var line string
  for scanner.Scan() {
    line = scanner.Text()
    lines = append(lines, line)
  }
  return lines, nil
}
