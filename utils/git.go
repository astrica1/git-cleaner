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

func (r Repository) RemoveTag(tagName string) error {
	cmd := "git push --delete origin " + tagName
	_, err := Execute(cmd)
	return err
}

func (r Repository) GetTagBranch(tagID string) []string {
	cmd := "git branch -a --contains " + tagID
	branches, err := Execute(cmd)
	if err != nil || branches == "" {
		return nil
	}

	branchesList := strings.Split(branches, "\n")
	branchesList = branchesList[:len(branchesList)-1]

	for i, b := range branchesList {
		println(b)
		if strings.HasSuffix(b, "main") && len(b) < 8 {
			branchesList[i] = "main"
			continue
		}

		if strings.Contains(b, "HEAD ->") {
			branchesList = append(branchesList[:i], branchesList[i+1:]...)
		}

		tmp := strings.Split(b, "origin/")
		if len(tmp) > 1 {
			branchesList[i] = strings.Join(tmp[1:], "origin/")
		}
	}

	return branchesList
}
