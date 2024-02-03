package utils

import "strings"

type Repository string

func (r *Repository) Clone() error {
	cmd := "git clone " + string(*r)
	_, err := Execute(cmd)
	return err
}

func (r *Repository) Fetch() error {
	cmd := "git fetch origin"
	_, err := Execute(cmd)
	return err
}

func (r *Repository) BranchList() []string {
	cmd := "for branch in `git branch -r | grep -v HEAD`;do echo -e `git show --format=\"%ci %cr\" $branch | head -n 1` \t$branch; done | sort -r"
	branches, err := Execute(cmd)
	if err != nil {
		return nil
	}

	if branches == "" {
		return nil
	}

	branchesList := strings.Split(branches, "\n")
	return branchesList
}

func (r *Repository) RemoveBranch(branchName string) error {
	cmd := "git branch --delete " + branchName
	_, err := Execute(cmd)
	if err != nil {
		return err
	}

	cmd = "git push origin --delete " + branchName
	_, err = Execute(cmd)
	return err
}
