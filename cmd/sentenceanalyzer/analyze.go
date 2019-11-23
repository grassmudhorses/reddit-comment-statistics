package commentanalyzer

import (
	read "github.com/BluntSporks/readability"
	"github.com/grassmudhorses/reddit-comment-statistics/pkg/commentparser"
)

type CommentMeta struct {
	Flair        string
	FKGradeLevel float64
	SMOG         float64
}

func Analyze(comment commentparser.AnonymizedComment) *CommentMeta {
	fk := read.Fk(comment.Body)
	smog := read.Smog(comment.Body)
	return nil
}
