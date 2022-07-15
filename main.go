package main

import (
	"log"
	"os"
	"strings"

	"github.com/leodido/go-conventionalcommits"
	"github.com/leodido/go-conventionalcommits/parser"
)

func getLastCommitMsg() (string, error) {
	ret, err := os.ReadFile(".git/COMMIT_EDITMSG")
	if err != nil {
		return "", err
	}

	// remove the last new line that is added automatically in the file
	str := strings.TrimRight(string(ret), "\n")
	return str, nil
}

func main() {
	commitMsg, err := getLastCommitMsg()
	if err != nil {
		log.Fatal(err)
	}

	res, err := parser.NewMachine(conventionalcommits.WithTypes(conventionalcommits.TypesConventional)).Parse([]byte(commitMsg))
	if err != nil {
		log.Fatal(err)
	}
	if !res.Ok() {
		log.Fatal("Commit was not correctly formatted.")
	}

}
