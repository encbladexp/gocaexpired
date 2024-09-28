package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Certificate struct {
	Status   string
	Expired  string
	Revoked  string
	Serial   string
	Filename string
	CN       string
}

type Certificates struct {
	Certificates []Certificate
}

func (c *Certificates) process_line(line []string) {
	cert := Certificate{
		line[0],
		line[1],
		line[2],
		line[3],
		line[4],
		line[5],
	}
	c.Certificates = append(c.Certificates, cert)
}

const FILENAME = "index.txt"

func main() {
	file, err := os.Open(FILENAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	csvreader := csv.NewReader(file)
	csvreader.Comma = '\t'
	lines, err := csvreader.ReadAll()
	if err != nil {
		panic(err)
	}
	cert_store := Certificates{}
	for _, line := range lines {
		cert_store.process_line(line)
	}
	fmt.Println(cert_store.Certificates)
}
