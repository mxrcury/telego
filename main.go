package main

import (
	"authtg/pkg/cmd"
	"authtg/pkg/telegram"
	"fmt"
	"reflect"
	"strings"
	"time"

	"bufio"
	"context"
	"os"
)
const (
  NAMES_FILE = "names.txt"
)

// TODO(proxy): add proxy support
const (
  ADDRESS = "127.0.0.1"
  PORT = 4332
)


func main() {
  fmt.Println("WELCOME TO _TELEGAUTH_") // TODO(text): change text and make it ASCII ART LOOKS LIKE

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

      _, err = tgClient.Auth().SignIn(clientContext, phoneNumber, authCode, resp.PhoneCodeHash)
      if telegram.Is2FAError(err) {
        fmt.Printf("Please enter a password for 2FA:")
        pass, err := reader.ReadString('\n')
        pass = strings.ReplaceAll(pass, "\n", "")
        if err != nil {
          return err
        }
        if _, err = tgClient.Auth().Password(clientContext, pass); err != nil {
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
      fmt.Println("[SUCCESS!]")
      time.Sleep(time.Second * 1)

      cmd.Init(tgClient, clientContext)
      
      return nil
    }); err != nil {
      fmt.Printf("error: %s\n", err)
      continue
    }
  }
}
