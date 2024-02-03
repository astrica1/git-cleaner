package utils

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
