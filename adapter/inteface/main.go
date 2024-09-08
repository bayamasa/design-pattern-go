package main

func main() {
    client := &Client{}
    wifi := &WifiAdapter{}
    wired := &WiredAdapter{}

    client.connect(wifi)
    client.connect(wired)
}