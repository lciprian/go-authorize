package authorize

import (
	"encoding/xml"
	"fmt"

	"github.com/lciprian/go-authorize/client"
	"github.com/lciprian/go-authorize/models"
)

type (
	Authorize struct {
		cli *client.Client

		apiLoginId     string
		transactionKey string
		key            string
		sandbox        bool
	}
)

func New(apiLoginId, transactionKey, key string, sandbox bool) (*Authorize, error) {
	if apiLoginId == "" {
		return nil, fmt.Errorf("Empty field: apiLoginId")
	}

	if transactionKey == "" {
		return nil, fmt.Errorf("Empty field: TransactionKey")
	}

	if key == "" {
		return nil, fmt.Errorf("Empty field: key")
	}

	return &Authorize{
		apiLoginId:     apiLoginId,
		transactionKey: transactionKey,
		key:            key,
		sandbox:        sandbox,
		cli:            client.NewClient(sandbox),
	}, nil
}

func (a *Authorize) getMerchantAuth() *models.MerchantAuth {
	return &models.MerchantAuth{
		ApiLoginId:     a.apiLoginId,
		TransactionKey: a.transactionKey,
	}
}

func (a *Authorize) getErrorMessage(res []byte) (string, error) {
	errResponse := &models.ErrorResponse{}

	if err := xml.Unmarshal(res, errResponse); err != nil {
		return "", err
	}

	if len(errResponse.Messages) == 0 {
		return "", fmt.Errorf("Empty error message")
	}

	return errResponse.Messages[0].Text, nil
}

// Test Authentication Credentials
func (a *Authorize) AuthenticationTest() error {
	req := &models.AuthenticateTestRequest{
		MerchantAuth: a.getMerchantAuth(),
	}

	res, err := a.cli.PostRequest(req)
	if err != nil {
		return err
	}

	response := &models.AuthenticateTestResponse{}
	if err = xml.Unmarshal(res, response); err != nil {
		if message, err := a.getErrorMessage(res); err == nil {
			return fmt.Errorf(message)
		}

		return err
	}

	if len(response.Messages) > 0 && response.Messages[0].ResultCode == "Error" {
		return fmt.Errorf(response.Messages[0].Text)
	}

	return nil
}
