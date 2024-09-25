package lib

// Repo represents a repository with a name and URL.
type Repo struct {
	Worktree string
	GitDir   string
	Config  string
}

func NewRepo(worktree string, gitDir string) *Repo {
	return &Repo{
		Worktree: worktree,
		GitDir:   gitDir,
	}
}

func (r *Repo) String() string {
	return r.Worktree + " " + r.GitDir
}