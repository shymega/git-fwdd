package queuesys // import "github.com/shymega/git-fwdd/internal/queuesys"

type QueuePayload struct {
	RepoName    string
	BranchName  string
	CommitCount int
}
