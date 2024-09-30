package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

const FILENAME = "index.txt"
const TIMEFORMAT_SHORT = "060102150405Z"
const TIMEFORMAT_LONG = "20060102150405Z"

const (
	CERT_STATUS = iota
	CERT_EXPIRED
	CERT_REVOKED
	CERT_SERIAL
	CERT_FILENAME
	CERT_CN
)

var TIME_90D = time.Now().Add(90 * 24 * time.Hour)
var TIME_30D = time.Now().Add(30 * 24 * time.Hour)
var TIME_NOW = time.Now()

type Certificate struct {
	Status   string
	Expired  *time.Time
	Revoked  *time.Time
	Serial   *string
	Filename *string
	CN       *string
}

type Certificates struct {
	Certificates []Certificate
}

func (c Certificates) parse_time(timestring string) (*time.Time, error) {
	date, err := time.Parse(TIMEFORMAT_SHORT, timestring)
	if err == nil {
		return &date, nil
	}
	date, err = time.Parse(TIMEFORMAT_LONG, timestring)
	if err != nil {
		return nil, err
	}
	return &date, nil
}

func (c *Certificates) process_line(line []string) {
	cert := Certificate{}
	expired, err := c.parse_time(line[CERT_EXPIRED])
	if err != nil {
		panic(err)
	}
	cert.Expired = expired
	revoked_string := line[CERT_REVOKED]
	if revoked_string != "" {
		revoked, err := c.parse_time(line[CERT_REVOKED])
		if err != nil {
			panic(err)
		}
		cert.Revoked = revoked
	}
	cert.Status = line[CERT_STATUS]
	filename := &line[CERT_FILENAME]
	if *filename != "unknown" {
		cert.Filename = &line[CERT_FILENAME]
	}
	cert.Serial = &line[CERT_SERIAL]
	cert.CN = &line[CERT_CN]
	switch cert.Status {
	case "V":
		c.Certificates = append(c.Certificates, cert)
	}
}

func (c *Certificates) print_expired() {
	for _, cert := range c.Certificates {
		if cert.Expired.Before(TIME_NOW) {
		} else if cert.Expired.Before(TIME_30D) {
			fmt.Println("Oh Shit")
		} else if cert.Expired.Before(TIME_90D) {
			fmt.Println("Shitty Shit!")
			fmt.Println(cert)
			fmt.Println(*cert.CN)
		}
	}
}

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
	cert_store.print_expired()
}
