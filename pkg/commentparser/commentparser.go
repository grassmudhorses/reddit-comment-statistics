package commentparser

import (
	"time"

	"github.com/turnage/graw/reddit"
)

type AnonymizedComment struct {
	Time  time.Time
	Body  string
	Flair string
}

func GetAnonymizedCommentsForURL(url string) (comments []*AnonymizedComment, err error) {
	rate := 5 * time.Second
	script, err := reddit.NewScript("graw:doc_script:0.3.1 by /u/Academic_Tune", rate)
	if err != nil {
		return
	}
	post, err := script.Thread(url)
	if err != nil {
		return
	}
	comments = make([]*AnonymizedComment, len(post.Replies))
	for i, v := range post.Replies {
		comments[i].Time = time.Unix(int64(v.CreatedUTC), 0)
		comments[i].Body = v.Body
		comments[i].Flair = v.AuthorFlairText
	}
	return comments, nil
}
