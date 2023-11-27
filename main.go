package main

import (
	"authtg/pkg/telegram"
	"authtg/pkg/utils"
	"fmt"
	"log"
	"strings"

	"bufio"
	"context"
	"os"
)

const (
  NAMES_FILE = "names.txt"
)

func main() {
  names, err := utils.ReadDocumentByLine(NAMES_FILE)
  if err != nil {
    log.Fatalln(err)
  }

   for {
    reader := bufio.NewReader(os.Stdin)

    fmt.Printf("[PHONE NUMBER]:")
    phoneNumber, err := reader.ReadString('\n')
    if err != nil {
      log.Printf("error:%v\n", err)
      continue
    }

    /*
    validPhoneNumberRegex := "+7"
    if isValidPhoneNumber := strings.HasPrefix(phoneNumber, validPhoneNumberRegex); !isValidPhoneNumber {
      log.Println("phone number format should starts with +7")
      continue
    }*/
    phoneNumber = strings.ReplaceAll(phoneNumber, "\n", "")

    options := telegram.TelegramClientOptions{Address: "127.0.0.1", Port: 3000, PhoneNumber: phoneNumber}
    tgClient := telegram.NewTelegramClient(&options)
    if err := telegram.NewSession(tgClient, func(clientContext context.Context) error {
      firstName, lastName := utils.GetRandomName(names)
      api := tgClient.API()

      resp, err := telegram.GetAuthCode(&telegram.TelegramAPI{api, clientContext}, phoneNumber); 
      if err != nil {
        return err
      }
    
      fmt.Printf("[CODE]:")

      authCode, err := reader.ReadString('\n')
      if err != nil {
        log.Printf("error:%s\n", err)
        return err
      }
      authCode = strings.ReplaceAll(authCode, "\n", "")
      codeHash := resp.PhoneCodeHash
      fmt.Println(codeHash, authCode, len(authCode))
      
      requestData := &telegram.SignUpRequest{PhoneNumber: phoneNumber, Code: authCode, FirstName: firstName, LastName: lastName, CodeHash: codeHash}
      if err := telegram.SignUp(&telegram.TelegramAPI{api, clientContext}, requestData); err != nil {
        return err
      }
      fmt.Println("[SUCCESS!] account created.")
      return nil
    }); err != nil {
      log.Printf("error: %s [%s]", err, phoneNumber)
    }
  }
}
