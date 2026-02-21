package clients

type Clients struct {
	Git *GitClient
}

func NewClients() *Clients {
	return &Clients{
		Git: NewGitClient(),
	}
}
