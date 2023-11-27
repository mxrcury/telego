package utils

import (
	"io"
	"os"
	"strings"
)

func ReadDocumentByLine(name string) ([]string, error) {
        file, err := os.OpenFile(name, os.O_RDONLY, os.ModeDevice)
        if err != nil {
          return nil, err
        }
        var lines []string
        for {
          b := make([]byte, 80)
          n, err := file.Read(b)
          if err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
          }
          str := string(b[:n])
          splitted := strings.Split(str, "\n")
          for _, line := range splitted {
            if strings.Trim(line, " ") != "" {
              lines = append(lines, line)
            }
          }
          
 }
  return lines, nil
}
