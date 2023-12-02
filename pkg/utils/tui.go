package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)



func ClearTerminal(){
  var cmd *exec.Cmd

  if runtime.GOOS == "windows" {
    cmd = exec.Command("cmd", "/c", "cls")
  } else {
    cmd = exec.Command("clear")
  }
  cmd.Stdout = os.Stdout
  cmd.Run()
}

func PrintInput(){
  fmt.Printf("──────────────────\n")
  fmt.Printf("$ ")
}
