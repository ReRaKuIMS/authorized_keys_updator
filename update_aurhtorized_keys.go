package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type AuthorizedKeysFile struct {
	BeforeSeparator string
	AfterSeparator  string
}

const SeparatorOfBefore = "\n#### PUBKEY_MASTER EDIT START ####\n"
const SeparatorOfAfter = "\n#### PUBKEY_MASTER EDIT END ####\n"

func main() {
	var authorizedKeysPath string
	var pubkeysFileUrl string
	flag.StringVar(&authorizedKeysPath, "f", "", "path to authorized_keys")
	flag.StringVar(&pubkeysFileUrl, "s", "", "url of pubkeys file")

	flag.Parse()

	if authorizedKeysPath == "" || pubkeysFileUrl == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	authorizedKeys := readAuthorizedKeys(authorizedKeysPath)

	pubKeys := getPubkeys(pubkeysFileUrl)

	rewritted := authorizedKeys.inject(pubKeys)
	ioutil.WriteFile(authorizedKeysPath, []byte(rewritted), os.ModePerm)
}

func (f *AuthorizedKeysFile) inject(authorizedKeys string) string {
	return f.BeforeSeparator +
		SeparatorOfBefore +
		authorizedKeys +
		SeparatorOfAfter +
		f.AfterSeparator
}

func readAuthorizedKeys(path string) *AuthorizedKeysFile {
	body, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't read authorized_keys file: %s\n", err)
		os.Exit(2)
	}

	befores := strings.Split(string(body), SeparatorOfBefore)
	before := befores[0]

	afters := strings.Split(string(body), SeparatorOfAfter)
	var after string
	if len(afters) < 2 {
		after = ""
	} else {
		after = afters[1]
	}

	return &AuthorizedKeysFile{BeforeSeparator: before, AfterSeparator: after}
}

func getPubkeys(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "request failed: %s\n", err)
		os.Exit(3)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "request failed")
		os.Exit(4)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't read response body")
		os.Exit(5)
	}

	return string(body)
}
