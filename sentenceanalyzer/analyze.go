package commentanalyzer

import (
	"regexp"

	read "github.com/BluntSporks/readability"
	"github.com/grassmudhorses/reddit-comment-statistics/commentparser"
	"github.com/grassmudhorses/vader-go/sentitext"
)

// Emoji simple regex to match only emoji
var SenteceSeparator *regexp.Regexp

func init() {
	SenteceSeparator = regexp.MustCompile(`[^\p{Zl}\p{Zp}.··۔。︒﹒．｡]+`)
}

//CommentMeta anonymized metadata for a comment
type CommentMeta struct {
	FKGradeLevel float64 `csv:"fk"`
	SMOG         float64 `csv:"smog"`
	Positive     float64 `csv:"pos"`
	Neutral      float64 `csv:"neu"`
	Negative     float64 `csv:"neg"`
}

//Analyze runs all available analysis on the AnonymizedComment
func Analyze(comment commentparser.AnonymizedComment) CommentMeta {
	fk := read.Fk(comment.Body)
	smog := read.Smog(comment.Body)
	//analyze each sentence (separated by periods or newlines)
	sentences := []string{}
	for _, sentence := range SenteceSeparator.FindAllString(comment.Body, -1) {
		if len(sentence) > 0 {
			sentences = append(sentences, sentence)
		}
	}
	allPositive := 0.0
	allNeutral := 0.0
	allNegative := 0.0
	allWeight := 0.0

	for _, sentence := range sentences {
		senLen := float64(len(sentence))
		allWeight += senLen
		vader := sentitext.PolarityScore(sentence)
		allPositive += vader.Positive * senLen
		allNeutral += vader.Neutral * senLen
		allNegative += vader.Negative * senLen
	}
	if allWeight < 1 {
		allWeight = 1
	}
	return CommentMeta{
		FKGradeLevel: fk,
		SMOG:         smog,
		Positive:     allPositive / allWeight,
		Neutral:      allNeutral / allWeight,
		Negative:     allNegative / allWeight,
	}

}
