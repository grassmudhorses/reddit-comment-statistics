package main

import (
	"fmt"
	"os"
	"time"

	"github.com/grassmudhorses/reddit-comment-statistics/commentparser"
	commentanalyzer "github.com/grassmudhorses/reddit-comment-statistics/sentenceanalyzer"
	"github.com/turnage/graw/reddit"
)

func main() {
	args := os.Args[1:]
	rate := 5 * time.Second
	script, err := reddit.NewScript("graw:doc_script:0.3.1", rate)
	if err != nil {
		fmt.Println(err)
		return
	}

	comments, err := commentparser.GetAnonymizedCommentsForURL("/"+args[0], script)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Time, Flair, Upvotes, Downvotes, FK Grade Level, SMOG, Positive Sentiment, Neutral Sentiment, Negative Sentiment\n")
	for _, comment := range comments {
		analysis := commentanalyzer.Analyze(*comment)
		fmt.Printf("%v, %v, %v, %v, %v, %v, %v, %v, %v \n", comment.Time, comment.Flair, comment.Upvotes, comment.Downvotes, analysis.FKGradeLevel, analysis.SMOG, analysis.Positive, analysis.Neutral, analysis.Negative)
	}
}
