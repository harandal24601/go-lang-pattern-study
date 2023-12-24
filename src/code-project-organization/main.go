package main

import (
	"log"
	"net/http"
)

func main() {
	variableShadowing1()
	variableShadowing2()
	variableShadowing3()
	variableShadowing4()
}

/*
* variableShadowing1
(1) Declare a 'client' variable.
(2) When tracing is true (if tracing is turned on), it creates an HTTP client.
(This is where the client variable is obscured.)
(3) Create a default HTTP client.
(The client variable is also obscured in this block.)

In (2) and (3), we use the variable declarator(:=) to assign the result of the function call to the client variable.
Instead of assigning it to the client of the outer block, it is assigned to the client variable in the lower block.
Thus, the client variable in the outer block will remain in the nil state.
*/
func variableShadowing1() error {
	var client *http.Client // (1)
	if tracing {
		client, err := createClientWithTracing() // (2)
		if err != nil {
			return err
		}
		log.Println(client)
	} else {
		client, err := createDefaultClient() // (3)
		if err != nil {
			return err
		}
		log.Println(client)
	}

	_ = client
	return nil
}

/*
* variableShadowing2
(1) Create a temporary variable 'c'.
(2) Assign 'c' to the original variable 'client'.

This saves the result of execution in a temporary variable, c.
The scope of c is limited to the if block.
Then, assign the saved execution results back to the client.
*/
func variableShadowing2() error {
	var client *http.Client
	if tracing {
		c, err := createClientWithTracing() // (1)
		if err != nil {
			return err
		}
		client = c // (2)
	} else {
		c, err := createDefaultClient()
		if err != nil {
			return err
		}
		client = c
	}

	_ = client
	return nil
}

/*
* variableShadowing3
(1) Declare the 'err' variable.
(2) Use the assignment operator to create '*http.Client' directly into the 'client' variable.

This method is used to use the assignment operator(=) in the subblock to assign the result of the function
directly to the client variable.
However, to do this, we need to create an error variable.
This is because assignment operators only work when the variable name is already declared.
*/
func variableShadowing3() error {
	var client *http.Client
	var err error // (1)
	if tracing {
		client, err = createClientWithTracing() // (2)
		if err != nil {
			return err
		}
	} else {
		client, err = createDefaultClient()
		if err != nil {
			return err
		}
	}

	_ = client
	return nil
}

/*
* variableShadowing4
variableShadowing3 allows us to leave the error handling code out of the if/else statement,
so we don't have to write it separately for both cases.
*/
func variableShadowing4() error {
	var client *http.Client
	var err error
	if tracing {
		client, err = createClientWithTracing()
	} else {
		client, err = createDefaultClient()
	}
	if err != nil {
		return err
	}

	_ = client
	return nil
}

var tracing bool

func createClientWithTracing() (*http.Client, error) {
	return nil, nil
}

func createDefaultClient() (*http.Client, error) {
	return nil, nil
}
