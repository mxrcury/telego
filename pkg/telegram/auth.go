package telegram

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"
	"go.uber.org/zap"
)

// FIX: needs to be changed due to client or ask in manager what about it
const (
  API_ID = 29763419
  API_HASH = "b8b953e3ffde464a8c3b46712db4db0e"
)

type SignUpRequest struct{
  PhoneNumber string

  Code string 
  CodeHash  string

  FirstName string
  LastName string
}

type TelegramAPI struct {
  Client *tg.Client
  Ctx context.Context
}

type TelegramClientOptions struct {
  Address string
  Port  int
  TdFilesDir  string
  TdDbDir string
  PhoneNumber string
}

type TelegramStorage struct{
  DBDir string
  FileDir string
}

func sessionFolder(phone string) string {
	var out []rune
	for _, r := range phone {
		if r >= '0' && r <= '9' {
			out = append(out, r)
		}
	}
	return "phone-" + string(out)
}

func NewTelegramClient(options *TelegramClientOptions) *telegram.Client {
 sessionPath := filepath.Join("sessions", sessionFolder(options.PhoneNumber))

 return telegram.NewClient(
    API_ID,
    API_HASH,
    telegram.Options{
      /*
      Device: telegram.DeviceConfig{
        Proxy: tg.InputClientProxy{
          Address: options.Address,
          Port: options.Port,
        },
      },
      */
      SessionStorage: &session.FileStorage{Path: filepath.Join(sessionPath, "session.json")},
      Logger: zap.L(),
    },
  )
}

func NewSession( client *telegram.Client, f func(clientContext context.Context) error) error {
  if err := client.Run(context.Background(), f); err != nil {
    return err
	}
  return nil
}

func GetAuthCode(api *TelegramAPI, phoneNumber string) (*tg.AuthSentCode, error) {
    requestData := &tg.AuthSendCodeRequest{PhoneNumber: phoneNumber, APIID: API_ID, APIHash: API_HASH}
    resp, err := api.Client.AuthSendCode(api.Ctx, requestData)

    if resp, ok := resp.(*tg.AuthSentCode); ok {
      fmt.Println("code was sent", ok)
      return resp, nil
    }
    if err != nil {
      return nil, err
    }

		return nil, err
}

func SignUp(api *TelegramAPI, body *SignUpRequest) error {

  resp, err := api.Client.AuthSignIn(api.Ctx, &tg.AuthSignInRequest{PhoneNumber: body.PhoneNumber, PhoneCodeHash: body.CodeHash })
  if err != nil {
    return err
  }
  fmt.Println("Response: ", resp)
  return nil

  /*
  resp, err := api.Client.AuthSignUp(api.Ctx, &tg.AuthSignUpRequest{PhoneNumber: body.PhoneNumber, PhoneCodeHash: body.CodeHash, FirstName: body.FirstName, LastName: body.LastName})
  if err != nil {
    return err
  }
  fmt.Println("Response: ", resp)
  return nil
  */
}

