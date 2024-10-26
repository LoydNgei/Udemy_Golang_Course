package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func hash_content(content []byte) string {
	hasher := sha256.New()
	hasher.Write(content)
	result := hasher.Sum(nil)
	hash := hex.EncodeToString(result)
	return hash

}

// f9d95eaccd32a46809ce0fd33384a0c595678f6459e884cf9f5b1b32c44de308

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./fileIntegrity <filename> <originalhash>")
		return
	}

	file_name := os.Args[1]
	original_hash := os.Args[2]

	contents, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	hash := hash_content(contents)

	if hash == original_hash {
		fmt.Println("The file is intact")
	} else {
		fmt.Println("The file has been tampered with")
	}

}
