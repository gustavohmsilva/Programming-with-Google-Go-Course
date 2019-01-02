package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	leMap := map[string]string{"name": "", "address": ""}
	for key := range leMap {
		fmt.Printf("%s: ", strings.Title(key))
		scanBuffer := bufio.NewScanner(os.Stdin)
		if scanBuffer.Scan() {
			leMap[key] = scanBuffer.Text()
		}
	}
	jsonOutput, _ := json.Marshal(leMap)
	fmt.Println(string(jsonOutput))
}
