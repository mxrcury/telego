package utils

import (
	"fmt"
	"strconv"
	"strings"
)

const (
  HELP_CMD = ":help"
  ME_CMD = ":me"

  CHAT_INFO_CMD = ":chat"
  CHATS_CMD = ":chats"

  EXIT_CMD = ":exit"
  LOGOUT_CMD = ":logout"
)

func IsHelpCommand(command string) bool {
  return strings.HasPrefix(command, HELP_CMD)
}

func IsChatsCommand(command string) bool {
  return strings.HasPrefix(command, CHATS_CMD)
}

func IsExitCommand(command string) bool {
  return strings.HasPrefix(command, EXIT_CMD)
}

func IsLogOutCommand(command string) bool {
  return strings.HasPrefix(command, LOGOUT_CMD)
}

func IsGetMeCommand(command string) bool {
  return strings.HasPrefix(command, ME_CMD)
}

func IsChatInfoCommand(command string) (int, bool) {
  command = strings.ReplaceAll(command, "\n", "")
  hasPrefix := strings.HasPrefix(command, CHAT_INFO_CMD)
  if !hasPrefix {
    return 0, false
  }
  fmt.Println("HAST PREFIX", hasPrefix)
  splitedCommand := strings.Split(command, " ")
  fmt.Println("SPLITTED:", splitedCommand)
  if len(splitedCommand) < 2 {
    return 0, false
  }
  parsedNumber, err := strconv.ParseInt(splitedCommand[1], 10, 64)
  fmt.Printf("PARSED INT:%d, ERROR:%s\n", parsedNumber, err)
  if err != nil {
    return 0, false
  }
  return int(parsedNumber), true
}
