package main

func main() {
    client := &Client{}
    wired := NewWiredAdapter("localhost")
    wifi := NewWifiAdapter()

    client.connect(wifi)
    client.connect(wired)
}