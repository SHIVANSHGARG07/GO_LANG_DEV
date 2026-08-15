package internal

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Commit struct {
	Hash    string `json: "hash"`
	Message string `json: "message"`
	Author  string `json: "author"`
	Date    string `json: "date"`
}

// store and read
type CommitStore struct {
	Commits  []Commit `json: "commits"`
	filepath string
}

// constructor
func NewCommitStore() (*CommitStore, error) {

	// check if we are in initialized repo
	repoPath, err := GetRepoPath()
	if err != nil {
		return nil, err
	}

	// repoPath will be .gitcli

	// build file path
	commitFile := filepath.Join(repoPath, "commits.json")

	return &CommitStore{
		Commits:  []Commit{},
		filepath: commitFile,
	}, nil

}

// all these are methods to commitStore

// read existing commit.json file into memory
func (cs *CommitStore) Load() error {
	data, err := os.ReadFile(cs.filepath)

	if err != nil {
		return err
	}

	// handle empty file
	if len(data) == 0 || string(data) == "[]" {
		cs.Commits = []Commit{}
		return nil
	}

	// convert json to go struct

	// data json
	// cs go struct
	return json.Unmarshal(data, cs)
}

// save: Write a commit to file
// convert commit store to json and write to file
func (cs *CommitStore) Save() error {
	data, err := json.MarshalIndent(cs, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(cs.filepath, data, 0644)
}

// add commit
// adds it into memory not file yet
func (cs *CommitStore) AddCommit(message, author string) string {

	// generate hash
	timeStamp := time.Now()
	hashInput := fmt.Sprintf("%s%s%v", message, author, timeStamp)

	// 4 things
	// sha1.sum expects []byte so convert string to bytes
	// %x will convert to hexadecimal (human readable)
	// :7 will return first 7 characters of hash
	hash := fmt.Sprintf("%x", sha1.Sum([]byte(hashInput)))[:7]

	// create commit
	commit := Commit{
		Hash:    hash,
		Message: message,
		Author:  author,
		Date:    timeStamp.Format("2006-01-02 15:04:05"),
	}

	// add to memory
	cs.Commits = append(cs.Commits, commit)

	return hash
}
