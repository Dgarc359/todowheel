package main;

import(
    "os"
    "todowheel-backend/client"
)

func main() {
    args := os.Args[1:]
    if args[0] == "--server" || args[0] == "-s" {
        startServer()
    } else if args[0] == "--client" || args[0] == "-c" {
        client.CreateClient()
    }
}
