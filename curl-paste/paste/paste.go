package paste

import (
	"bufio"
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Store(title, date, content, destDir string) (string, error) {

	h := sha256.New()

	h.Write([]byte(title))
	h.Write([]byte(date))
	h.Write([]byte(content))

	paste := fmt.Sprintf("# Title: %s\n# Date: %s\n%s", title, date, content)

	pasteHash := fmt.Sprintf("%x", h.Sum(nil))
	log.Printf("  -- hash: %s\n", pasteHash)
	directory := destDir

	for i := 0; i < len(pasteHash)-16; i++ {
		pasteName := pasteHash[i : i+16]
		if _, err := os.Stat(directory + pasteName); os.IsNotExist(err) {
			if err := ioutil.WriteFile(directory+pasteName, []byte(paste), 0644); err == nil {
				log.Printf("  -- saving new paste to : %s", directory+pasteName)
				return pasteName, nil
			} else {
				log.Printf("Cannot create the paste: %s!\n", directory+pasteName)
			}
		}
	}
	return "", errors.New("Could not store the paste!")
}

func Retrieve(URI string) (title, date, content string, err error) {

	fCont, err := os.Open(URI)
	defer fCont.Close()

	if err == nil {
		stuff := bufio.NewScanner(fCont)
		stuff.Scan()
		title = strings.Trim(strings.Split(stuff.Text(), ":")[1], " ")
		stuff.Scan()
		date = strings.Trim(strings.Join(strings.Split(stuff.Text(), ":")[1:], ":"), " ")
		for stuff.Scan() {
			content += stuff.Text() + "\n"
		}
	} else {

		return "", "", "", errors.New("No data to retrieve with that ID!")
	}

	return title, date, content, nil
}
