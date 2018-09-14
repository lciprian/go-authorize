go-authorize
==========

go-authorize is an implementation of the Authorize.Net API in Golang.


## Import
    import "github.com/lciprian/go-authorize"

## Usage
~~~ go
package main

import (
    "fmt"
    "github.com/lciprian/go-authorize"
)

const (
    API_LOGIN_ID    = "YOUR_API_LOGIN_ID"
    TRANSACTION_KEY = "YOUR_TRANSACTION_KEY"
    KEY             = "YOUR_KEY"
)

func main() {
    // Authorize client
    authorize := authorize.New(API_LOGIN_ID, TRANSACTION_KEY, KEY)

    // Test Authentication Credentials
    err := authorize.AuthenticationTest()
    fmt.Println(err)
}
~~~
