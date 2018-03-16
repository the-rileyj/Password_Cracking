package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func getMD5Hash(text string) (string, error) {
	hasher := md5.New()
	_, err := hasher.Write([]byte(text))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func checkHash(hash, text string, found chan string) {
	if thash, _ := getMD5Hash(text); hash == thash {
		found <- text
	} else {
		found <- ""
	}
}

func main() {
	//file := "lists/rockyou.txt"
	hash := "7c6483ddcd99eb112c060ecbe0543e86"
	//hash := "5f4dcc3b5aa765d61d8327deb882cf99"
	file := "test.txt"
	fi, _ := os.Open(file)
	defer fi.Close()
	fiScan := bufio.NewScanner(fi)
	foundChannel := make(chan string)
	found := ""
	var sent uint
	for fiScan.Scan() {
		select {
		case found = <-foundChannel:
			if found != "" {
				fmt.Printf("Found Password for the hash: %s\n", found)
				sent = 0
				break
			} else {
				sent--
			}
		default:
			sent++
			go checkHash(hash, fiScan.Text(), foundChannel)
		}
	}
	for found == "" {
		select {
		case found = <-foundChannel:
			if found != "" {
				fmt.Printf("Found Password for the hash: %s\n", found)
				break
			} else {
				sent--
				if sent == 0 {
					fmt.Printf("Could not find password for hash %s\n", hash)
					break
				}
			}
		default:
			if sent == 0 {
				fmt.Printf("Could not find password for hash %s\n", hash)
				break
			}
		}
	}
}
