package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	".."
	"../../lsp"
)

func main() {
	const numArgs = 4
	if len(os.Args) != numArgs {
		fmt.Printf("Usage: ./%s <hostport> <message> <maxNonce>", os.Args[0])
		return
	}
	hostport := os.Args[1]
	message := os.Args[2]
	maxNonce, err := strconv.ParseUint(os.Args[3], 10, 64)
	if err != nil {
		fmt.Printf("%s is not a number.\n", os.Args[3])
		return
	}

	client, err := lsp.NewClient(hostport, lsp.NewParams())
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}

	defer client.Close()

	r := bitcoin.NewRequest(message, 0, maxNonce)
	b, _ := json.Marshal(r)
	err = client.Write(b)
	if err != nil {
		printDisconnected()
	}

	b, err = client.Read()
	if err != nil {
		printDisconnected()
	}
	var m bitcoin.Message
	json.Unmarshal(b, &m)

	printResult(m.Hash, m.Nonce)
}

// printResult prints the final result to stdout.
func printResult(hash, nonce uint64) {
	fmt.Println("Result", hash, nonce)
}

// printDisconnected prints a disconnected message to stdout.
func printDisconnected() {
	fmt.Println("Disconnected")
}
