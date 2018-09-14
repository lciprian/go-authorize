package models

import (
	"encoding/xml"
)

type (
	AuthenticateTestRequest struct {
		XMLName      xml.Name      `xml:"AnetApi/xml/v1/schema/AnetApiSchema.xsd authenticateTestRequest"`
		MerchantAuth *MerchantAuth `xml:"merchantAuthentication"`
	}

	AuthenticateTestResponse struct {
		XMLName  xml.Name            `xml:"authenticateTestResponse"`
		Messages []*ResponseMessages `xml:"messages"`
	}

	MerchantAuth struct {
		XMLName        xml.Name `xml:"merchantAuthentication"`
		ApiLoginId     string   `xml:"name"`
		TransactionKey string   `xml:"transactionKey"`
	}

	ErrorResponse struct {
		XMLName  xml.Name            `xml:"ErrorResponse"`
		Messages []*ResponseMessages `xml:"messages"`
	}

	ResponseMessages struct {
		ResultCode  string `xml:"resultCode"`
		Code        string `xml:"message>code"`
		Text        string `xml:"message>text"`
		Description string `xml:"message>description"`
	}

	Errors struct {
		ErrorCode string `xml:"error>errorCode"`
		ErrorText string `xml:"error>errorText"`
	}
)
