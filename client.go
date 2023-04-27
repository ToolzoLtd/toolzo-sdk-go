package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
)

type Payload struct {
	Type       string
	ApiVersion string
	Payload    interface{}
}
type Request struct {
	Type          string
	CorrelationId interface{}
	Payload       Payload
}
type Response struct {
	ApiKey    string `json:"ApiKey"`
	Payload   string `json:"Payload"`
	Signature string `json:"Signature"`
	Encrypted bool   `json:"Encrypted"`
}

func MakeRequest(payload interface{}, method string) []byte {

	u := uuid.New()
	mainMessage := Request{
		Type:          RequestTypeRequest,
		CorrelationId: u,
		Payload: Payload{
			Type:       method,
			ApiVersion: ApiVersion,
			Payload:    payload,
		},
	}

	// Form mainMessage in JSON format as string;
	mainMessageByte, _ := json.Marshal(mainMessage)
	// mainMessage in Base64;
	payloadBase64 := base64.StdEncoding.EncodeToString(mainMessageByte)
	// Convert hmacKey from Base64 string to byte array;
	secret, _ := base64.StdEncoding.DecodeString(Secret)
	// Generate a signature using the HMAC algorithm based
	// on the mainMessage and the secret;
	h := hmac.New(sha1.New, secret)
	h.Write(mainMessageByte)
	// Signature into a Base64 string;
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	type Message struct {
		ApiKey    string
		Payload   string
		Signature string
	}

	httpMessage := Message{
		ApiKey:    ApiKey,
		Payload:   payloadBase64,
		Signature: signature,
	}
	httpMessageFinal, _ := json.Marshal(httpMessage)
	return httpMessageFinal
}

func SendRequest(httpMessageFinal []byte) Response {
	req, err := http.NewRequest("POST", ApiUrl, bytes.NewBuffer(httpMessageFinal))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println(res.Status)
	body, _ := io.ReadAll(res.Body)

	result := Response{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	payloadRes, _ := base64.StdEncoding.DecodeString(result.Payload)
	result.Payload = string(payloadRes)
	return result
}
