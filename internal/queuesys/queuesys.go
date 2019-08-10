package queuesys // import "git.shymega.org.uk/shymega/git-fwdd/internal/queuesys"

import (
	"fmt"
)

type QueuePayload struct {
	RepoName string
	BranchName string
	CommitCount int
}

func Add2Queue(q QueuePayload) {
	fmt.Printf("Repo: %s\nCommit count: %d\nBranch: %s\n",
		q.RepoName,
		q.CommitCount,
		q.BranchName)
}
