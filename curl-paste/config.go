package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Options struct {
	confFile string
}

type Config struct {
	server    string
	address   string
	port      string
	directory string
	size      uint32
	logfile   string
}

func (c Config) String() string {

	var s string

	s += "Server name: " + c.server + "\n"
	s += "Listening on: " + c.address + ":" + c.port + "\n"
	s += "Paste directory: " + c.directory + "\n"
	s += "Max size: " + string(c.size) + "\n"
	s += "Log file: " + c.logfile + "\n"

	return s

}

func parseConfig(fName string, c *Config) error {

	f, err := os.Open(fName)
	if err != nil {
		return err
	}

	r := bufio.NewScanner(f)

	line := 0
	for r.Scan() {
		s := r.Text()
		line += 1
		if matched, _ := regexp.MatchString("^([ \t]*)$", s); matched != true {
			if matched, _ := regexp.MatchString("^#", s); matched != true {
				if matched, _ := regexp.MatchString("^([a-zA-Z]+)=.*", s); matched == true {
					fields := strings.Split(s, "=")
					switch strings.Trim(fields[0], " \t\"") {
					case "server":
						c.server = strings.Trim(fields[1], " \t\"")
					case "address":
						c.address = strings.Trim(fields[1], " \t\"")
					case "port":
						c.port = strings.Trim(fields[1], " \t\"")
					case "directory":
						c.directory = strings.Trim(fields[1], " \t\"")
					case "logfile":
						c.logfile = strings.Trim(fields[1], " \t\"")
					default:
						fmt.Fprintf(os.Stderr, "Error reading config file %s at line %d: unknown variable '%s'\n",
							fName, line, fields[0])
					}
				} else {
					fmt.Fprintf(os.Stderr, "Error reading config file %s at line %d: unknown statement '%s'\n",
						fName, line, s)
				}
			}
		}
	}
	return nil
}
