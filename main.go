package main

import (
	"authtg/pkg/telegram"
	"strings"

	"bufio"
	"context"
	"fmt"
	"log"
	"os"
)

func main() {
  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("[PHONE NUMBER] Please enter phone number:")
    phoneNumber, err := reader.ReadString('\n')
    if err != nil {
      log.Printf("error:%v\n", err)
      continue
    }

    validPhoneNumberRegex := "+7"
    if isValidPhoneNumber := strings.HasPrefix(phoneNumber, validPhoneNumberRegex); !isValidPhoneNumber {
      log.Println("phone number format should starts with +7")
      continue
    }
    phoneNumber = strings.ReplaceAll(phoneNumber, "\n", "")
    tgClient := telegram.NewTelegramClient()
    if err := telegram.NewSession(tgClient, func(clientContext context.Context) error {
      api := tgClient.API()

      if err := telegram.GetAuthCode(api, clientContext, phoneNumber); err != nil {
        return err
      }
    
      fmt.Printf("[CODE] Please enter received code:")

      authCode, err := reader.ReadString('\n')
      if err != nil {
        log.Printf("error:%s\n", err)
        return err
      }
      authCode = strings.ReplaceAll(authCode, "\n", "")

      requestData := &telegram.SignUpRequest{PhoneNumber: phoneNumber, Code: authCode}
      if err := telegram.SignUp(api, clientContext, requestData); err != nil {
        return err
      }

      return nil
    }); err != nil {
      log.Printf("error: %s [%s]", err, phoneNumber)
    }
 }
}
