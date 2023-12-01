package utils

import "strings"

const (
  HELP_CMD = ":help"
  CHATS_CMD = ":chats"
  ME_CMD = ":me"

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
