package main

import (
	"fmt"
	"log"
	"os"

	"github.com/leodido/go-conventionalcommits"
	"github.com/leodido/go-conventionalcommits/parser"
)

func getLastCommitMsg() (string, error) {
	ret, err := os.ReadFile(".git/COMMIT_EDITMSG")
	return string(ret), err
}

func main() {
	commitMsg, err := getLastCommitMsg()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", commitMsg)
	res, err := parser.NewMachine(conventionalcommits.WithTypes(conventionalcommits.TypesConventional)).Parse([]byte(commitMsg))
	if err != nil {
		log.Fatal(err)
	}
	if !res.Ok() {
		log.Fatal("Commit was not correctly formatted.")
	}
}