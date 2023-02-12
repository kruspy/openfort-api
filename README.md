# Openfort Cypher REST API Example
A REST API example application to encrypt and decrypt messages using asymmetric cryptography built in Go.

The API allows users to encrypt and decrypt messages. Each user has a unique API key that identifies them. A RSA
key pair is generated for each of the users and the private key is encrypted using AES along with a provided password.

Messages are encrypted using the public key of the user and are decrypted using the private key. Before decrypting a message the private key also needs to be decrypted with the password. 

## Installation & Run
The API can be run using either Docker or via building a native Go binary.

### Docker
Docker is the **recommended** way of running the API as it will automatically spin up a Postgres database and all its dependencies.
```bash
# Start the docker container
make start
```

### Running the API locally
In order to run the application locally it has to be built first.
```bash
# Build the application binary
make build
```
This will create a binary file named **openfort-api** in the out/bin directory. Which will be used to run the application.
```bash
# Run the application locally
make run
```
The application can be configured via the [config file](./conf/example.config.json).

## Tests

The project tests can be run using the Makefile.
```bash
# Run the application tests
make test
```

An [example application](./api-example.go) is provided, which interacts with the api to encrypt and decrypt a message.
```bash
# Run the example application
make run-example
```


## Structure
```
├── api                    main application logic
│   ├── handlers           call handlers
│   ├── models             db models
│   └── router             entrypoints
│       └── middleware     auth middleware
├── cmd
│   └── openfort-api       binary
│       ├── config         application config
│       └── logger
├── conf                   configuration files
├── pkg
    ├── aes-util           library for AES encryption
    └── rsa-util           library for RSA encryption
```

## API Endpoints

### Encrypt

#### Request

`POST /encrypt`

Encrypts a message using RSA asymmetric encryption and returns the result encoded in base64. 

The API key passed in the request is used to retrieve the public key used to encrypt the message.
    
    curl -i -H 'Accept: application/json' -H 'X-Api-Key: 156ed24e-f594-4f28-9b2a-b378802a37eb' -d '{"message":"hello"}' http://localhost:8080/encrypt

### Response

    HTTP/1.1 200 OK
    Date: Sun,12 Feb 2023 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 698

    {"message": "We6AeJy9qC7wV1UiTKCxHTaX5x1YrwwSfpaoQ+4MkSWzgxo5XtyzqV8IhoO5NobYbuHhcN/shLVAa32nNqmnJ5K+YtO7MgUQRnqYHUB6tbUatmXWvgssS0kOK9tS2lUEg9IZ5JZxx624S4IobrFIetXJNlzaA+bLfuV6nN3sBGotiS3U9Rt15q/KsJaLZ1KBLdjJaVpxQXuDJc5K9GOvGCtUt3oJdjIGXbHWbUWX6+gGv4cXy8JwkTl2KQLeb2FRMnmFrYXJzP+RvSaPiyG7yw8tg34C5Ajsqk25SgkQ0V8dqQIuHq2ss+lLV0VeTlZfnfKygCzUX1VM54fMaGnh+r5SrnJI6HS8oRv6GZd52MmuChtIZcGmT1kEu0+agvqbsRlCRy8D29aOVzmfYxVKotxBpiAtDsGItl3XUOSrnbDOQY9J4b8I4tYIcLO+so6OFIPLgUPcY9DLqqD9AQAEJdpRiK9ZsdPELIG0wTX+xspvQvrdTOyDHz6ksp/dXe+IoWewc5Ml3dAQZjss4oRdAbrEynd4aaj2elR+8k+Mczs8w9d40TSvv68azVAug6wmNZoRYbDI24otEPIpOCD1xm6iuLsyWvhaG0+WpJyOinKM29TgD3iYp2rUFVJvr9Mgi24GAPqa3wtl6xpHcWe4KMe5oRlkIxTHZOQ0RvhmM4A="}


### Decrypt

#### Request

`POST /decrypt`

Decrypts a base64 encoded message using RSA asymmetric encryption and returns the original text.

The API key passed in the request is used to retrieve the private key used to decrypt the message.

    curl -i -H 'Accept: application/json' -H 'X-Api-Key: 156ed24e-f594-4f28-9b2a-b378802a37eb' -d '{"password":"1234", "message":"We6AeJy9qC7wV1UiTKCxHTaX5x1YrwwSfpaoQ+4MkSWzgxo5XtyzqV8IhoO5NobYbuHhcN/shLVAa32nNqmnJ5K+YtO7MgUQRnqYHUB6tbUatmXWvgssS0kOK9tS2lUEg9IZ5JZxx624S4IobrFIetXJNlzaA+bLfuV6nN3sBGotiS3U9Rt15q/KsJaLZ1KBLdjJaVpxQXuDJc5K9GOvGCtUt3oJdjIGXbHWbUWX6+gGv4cXy8JwkTl2KQLeb2FRMnmFrYXJzP+RvSaPiyG7yw8tg34C5Ajsqk25SgkQ0V8dqQIuHq2ss+lLV0VeTlZfnfKygCzUX1VM54fMaGnh+r5SrnJI6HS8oRv6GZd52MmuChtIZcGmT1kEu0+agvqbsRlCRy8D29aOVzmfYxVKotxBpiAtDsGItl3XUOSrnbDOQY9J4b8I4tYIcLO+so6OFIPLgUPcY9DLqqD9AQAEJdpRiK9ZsdPELIG0wTX+xspvQvrdTOyDHz6ksp/dXe+IoWewc5Ml3dAQZjss4oRdAbrEynd4aaj2elR+8k+Mczs8w9d40TSvv68azVAug6wmNZoRYbDI24otEPIpOCD1xm6iuLsyWvhaG0+WpJyOinKM29TgD3iYp2rUFVJvr9Mgi24GAPqa3wtl6xpHcWe4KMe5oRlkIxTHZOQ0RvhmM4A="}' http://localhost:8080/decrypt

### Response

    HTTP/1.1 200 OK
    Date: Sun,12 Feb 2023 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 142

    {"message": "hello"}


NOTE: The example request won't work as the generated keys will be different.

## Mocking

* A set of mock keys is created and is completely customisable using the mocking [config file](./conf/mock.json).
* For each mocked set of keys, the API key and the password to encrypt the private key need to be specified.
* A user is also created as part of the mocks to keep consistency in the DB model.

### ToDo

- [ ] Improve testing suite and create tests for code outside business logic.
- [ ] Create tests for the encryption libraries.
- [ ] Add API calls to create users which will generate api keys and key pairs.
