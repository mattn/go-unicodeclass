package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mattn/go-unicodeclass"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(unicodeclass.SplitClass)
	for scan.Scan() {
		fmt.Println(scan.Text())
	}
	if scan.Err() != nil {
		log.Fatal(scan.Err())
	}
}
