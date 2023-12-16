package main

import (
	"authtg/pkg/cmd"
	"authtg/pkg/telegram"
	"authtg/pkg/utils"
	"fmt"
	"reflect"
	"strings"
	"time"

	"bufio"
	"context"
	"os"
)

const (
  WELCOME_TEXT= `
 _ _ _     _                      _          _____ _____ __    _____ _____ _____ 
| | | |___| |___ ___ _____ ___   | |_ ___   |_   _|   __|  |  |   __|   __|     |
| | | | -_| |  _| . |     | -_|  |  _| . |    | | |   __|  |__|   __|  |  |  |  |
|_____|___|_|___|___|_|_|_|___|  |_| |___|    |_| |_____|_____|_____|_____|_____|
                                                                                 
  `
  GREEN = "\033[32m"
  BLUE = "\033[36m"
  RESET = "\033[0m"
)

func main() {
  SetupConfig()
  utils.ClearTerminal()
  fmt.Println(BLUE + WELCOME_TEXT) // TODO(text): change text and make it ASCII ART LOOKS LIKE
  fmt.Println(RESET)

   for {
    reader := bufio.NewReader(os.Stdin)

    fmt.Printf("Please enter your phone number:")
    phoneNumber, err := reader.ReadString('\n')
    if err != nil {
      fmt.Printf("error:%v\n", err)
      continue
    }

    validPhoneNumberRegex := "+"
    if isValidPhoneNumber := strings.HasPrefix(phoneNumber, validPhoneNumberRegex); !isValidPhoneNumber {
      fmt.Println("phone number format should starts with +")
      continue
    }
    phoneNumber = strings.ReplaceAll(phoneNumber, "\n", "")

    options := telegram.TelegramClientOptions{
        PhoneNumber: phoneNumber,
    }
    tgClient := telegram.NewTelegramClient(&options)
    if err := telegram.NewSession(tgClient, func(clientContext context.Context) error {
      tgAPI := telegram.NewTelegramAPI(tgClient, clientContext)

      resp, err := tgAPI.GetAuthCode(phoneNumber); 
      if err != nil {
        return err
      }

      fmt.Printf("Please enter a received code:")

      authCode, err := reader.ReadString('\n')
      if err != nil {
        return err
      }
      authCode = strings.ReplaceAll(authCode, "\n", "")

      err = tgAPI.SignInWith2FA(phoneNumber, authCode, resp.PhoneCodeHash)
      if telegram.Is2FAError(err) {
        fmt.Printf("Please enter a password for 2FA:")
        pass, err := reader.ReadString('\n')
        if err != nil {
          return err
        }
        pass = strings.ReplaceAll(pass, "\n", "")

        if err = tgAPI.SignInWith2FAPassword(pass); err != nil {
          return err
        }
      } else if err != nil {
        return err
      } else {
        resp, err := tgAPI.SignIn(&telegram.SignInRequest{PhoneNumber: phoneNumber, Code: authCode, CodeHash: resp.PhoneCodeHash})
        if err != nil {
          return err
        }
        fmt.Println("CLEAR AUTH RESP:", resp, reflect.TypeOf(resp)) // TODO: update it
      }
      fmt.Println(GREEN + "[SUCCESS!]")
      fmt.Println(RESET)
      time.Sleep(time.Second * 1)

      cmd.Init(tgAPI)
      
      return nil
    }); err != nil {
      fmt.Printf("error: %s\n", err)
      continue
    }
  }
}
