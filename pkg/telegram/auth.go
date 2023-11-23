package telegram

import (
	"context"
	"fmt"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"
)


// FIX: needs to be changed due to client or ask in manager what about it
const (
  API_ID = 29763419
  API_HASH = "b8b953e3ffde464a8c3b46712db4db0e"
)

type SignUpRequest struct{
  PhoneNumber string
  Code string 
}


func NewTelegramClient() *telegram.Client {
 return telegram.NewClient(API_ID, API_HASH, telegram.Options{})
}

func NewSession( client *telegram.Client, f func(clientContext context.Context) error) error {
  if err := client.Run(context.Background(), f); err != nil {
    return err
	}
  return nil
}

// TODO: mb refactor to one structure with client and context
func GetAuthCode(client *tg.Client, ctx context.Context, phoneNumber string) error {
    requestData := &tg.AuthSendCodeRequest{PhoneNumber: phoneNumber, APIID: API_ID, APIHash: API_HASH}
    if _, err := client.AuthSendCode(ctx, requestData); err != nil {
      return err
    }
		return nil
}

func SignUp(client *tg.Client, ctx context.Context, body *SignUpRequest) error {
  resp, err := client.AuthSignUp(ctx, &tg.AuthSignUpRequest{PhoneNumber: body.PhoneNumber, PhoneCodeHash: body.Code })
  if err != nil {
    return err
  }
  fmt.Println("Response: ", resp)
  return nil
}
