



// branch store
type BranchStore struct {
	Branches      []string `json:"branches"`
	filepath      string
	currentBranch string
}

// constructor
func NewBranchStore() (*BranchStore, error) {

	repoPath, err := GetRepoPath()
	if err != nil {
		return nil, err
	}

	branchFile := filepath.Join(repoPath, "branches.json")
	currentBranchFile := filepath.Join(repoPath, "current-branch.txt")

	return &BranchStore{
		Branches:      []string{},
		filepath:      branchFile,
		currentBranch: currentBranchFile,
	}, nil

}

// load branches from file using unmarshal
func (bs *BranchStore) Load() error {

	// read file
	data, err := os.ReadFile(bs.filepath)

	if err != nil {
		if os.IsNotExist(err) {
			bs.Branches = []string{}
			return nil
		}
		return err
	}

	// check for empty or null
	if len(data) == 0 || string(data) == "[]" {
		bs.Branches = []string{}
		return nil
	}

	// return data from json to struct (go)
	// if data found
	return json.Unmarshal(data, bs)

}

// save to file using marshal indent
func (bs *BranchStore) Save() error {

	data, err := json.MarshalIndent(bs, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(bs.filepath, data, 0644)

}

// get current branch
// read current-branch.txt
func (bs *BranchStore) GetCurrentBranch() (string, error) {

	data, err := os.ReadFile(bs.currentBranch)
	if err != nil {
		return "", err
	}

	return string(data), nil

}

// set current branch
// write to current-branch.txt
func (bs *BranchStore) SetCurrentBranch(branch string) error {

	// check if branch exists or not
	if !bs.Exists(branch) {
		return fmt.Errorf("branch '%s' does not exist", branch)
	}

	// write to file
	return os.WriteFile(bs.currentBranch, []byte(branch), 0644)
}

// check if branch searching is there or not in memory
func (bs *BranchStore) Exists(branch string) bool {
	for _, b := range bs.Branches {
		if b == branch {
			return true
		}
	}
	return false
}

// create new branch in memory
func (bs *BranchStore) Create(name string) error {
	// Validate name
	if name == "" {
		return fmt.Errorf("branch name cannot be empty")
	}

	// Check if already exists
	if bs.Exists(name) {
		return fmt.Errorf("branch '%s' already exists", name)
	}

	// Add to list
	bs.Branches = append(bs.Branches, name)
	return nil
}

// delete new bracnh in memory
func (bs *BranchStore) Delete(name string) error {
	// Check if exists
	if !bs.Exists(name) {
		return fmt.Errorf("branch '%s' not found", name)
	}

	// Find and remove
	for i, branch := range bs.Branches {
		if branch == name {
			bs.Branches = append(bs.Branches[:i], bs.Branches[i+1:]...)
			return nil
		}
	}

	return nil
}