package main

import "fmt"

type Client struct{}

func (c *Client) connect(connector InternetConnector) {
	fmt.Println("Client connecting to the internet")
	connector.Connect()
}
