package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Remote struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type RemoteStore struct {
	Remotes  []Remote `json:"remotes"`
	filepath string
}

// constructor
func NewRemoteStore() (*RemoteStore, error) {
	// check if inside repo or not

	repoPath, err := GetRepoPath()

	if err != nil {
		return nil, err
	}

	remoteFile := filepath.Join(repoPath, "remotes.json")
	return &RemoteStore{
		Remotes:  []Remote{},
		filepath: remoteFile,
	}, nil
}

func (rs *RemoteStore) Load() error {

	// read file
	data, err := os.ReadFile(rs.filepath)

	if err != nil {
		return err
	}

	// check for empty or null
	if len(data) == 0 || string(data) == "[]" {
		rs.Remotes = []Remote{}
		return nil
	}

	// return data from json to struct (go)
	// if data found
	return json.Unmarshal(data, rs)

}

func (rs *RemoteStore) Save() error {
	data, err := json.MarshalIndent(rs, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(rs.filepath, data, 0644)
}

// add the remote in list in memory
func (rs *RemoteStore) Add(name, url string) error {

	// check for duplicates remote
	for _, remote := range rs.Remotes {
		if remote.Name == name {
			return fmt.Errorf("remote '%s' already exists", name)
		}
	}

	// add in memory
	rs.Remotes = append(rs.Remotes, Remote{Name: name, URL: url})
	return nil

}

// remove in memory
func (rs *RemoteStore) Remove(name string) error {

	// find and match which to remove
	// when found using slices remove the required one
	for i, remote := range rs.Remotes {
		if remote.Name == name {
			rs.Remotes = append(rs.Remotes[:i], rs.Remotes[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("remote '%s' not found", name)
}
