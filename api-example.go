package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	EncryptURL string = "http://localhost:8080/encrypt"
	DecryptURL string = "http://localhost:8080/decrypt"

	Message  string = "secret message" // THIS IS THE ENCRYPTED MESSAGE.
	Password string = "1234"
)

type EncryptRequest struct {
	Message string `json:"message"`
}

type EncryptResponse struct {
	Message string `json:"message"`
}

type DecryptRequest struct {
	Message  string `json:"message"`
	Password string `json:"password"`
}

type DecryptResponse struct {
	Message string `json:"message"`
}

func main() {
	client := &http.Client{}

	// Send a request to the /encrypt endpoint with a message to encrypt.
	log.Println(fmt.Sprintf("Sending request to encrypt message: %s", Message))
	res, err := client.Do(encryptRequest())
	if err != nil {
		log.Fatal(err)
	}

	// Parse the response to get the encrypted message encoded to base64.
	encryptResponse := &EncryptResponse{}
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(encryptResponse)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("Message encrypted"))

	// Use the encrypted message along with the password to send a request to the /decrypt endpoint.
	log.Println("Sending request to decrypt the message")
	res, err = client.Do(decryptRequest(encryptResponse.Message, Password))
	if err != nil {
		log.Fatal(err)
	}

	// Parse the response and get the original unencrypted message.
	decryptResponse := &DecryptResponse{}
	dec = json.NewDecoder(res.Body)
	err = dec.Decode(decryptResponse)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("Message decrypted. Original message is: %s", decryptResponse.Message))
}

func encryptRequest() *http.Request {
	reqBody, err := json.Marshal(&EncryptRequest{Message: Message})
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", EncryptURL, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", "156ed24e-f594-4f28-9b2a-b378802a37eb")

	return req
}

func decryptRequest(msg, pwd string) *http.Request {
	reqBody, err := json.Marshal(&DecryptRequest{Message: msg, Password: pwd})
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", DecryptURL, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", "156ed24e-f594-4f28-9b2a-b378802a37eb")

	return req
}
