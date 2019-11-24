package commentparser

import (
	"time"

	"github.com/turnage/graw/reddit"
)

type AnonymizedComment struct {
	Time      time.Time `csv:"time"`
	Body      string    `csv:"-"`
	Flair     string    `csv:"flair"`
	Upvotes   int       `csv:"ups"`
	Downvotes int       `csv:"downs"`
}

func GetAnonymizedCommentsForURL(url string, script reddit.Script) (comments []*AnonymizedComment, err error) {
	post, err := script.Thread(url)
	if err != nil {
		return
	}
	comments = []*AnonymizedComment{}
	for _, v := range post.Replies {
		comments = append(comments, anonymize(v)...)
	}
	return comments, nil
}

func anonymize(v *reddit.Comment) []*AnonymizedComment {

	comment := AnonymizedComment{}
	comment.Time = time.Unix(int64(v.CreatedUTC), 0)
	comment.Body = v.Body
	comment.Flair = v.AuthorFlairText
	comment.Downvotes = int(v.Downs)
	comment.Upvotes = int(v.Ups)

	comments := []*AnonymizedComment{&comment}
	for _, u := range v.Replies {
		comments = append(comments, anonymize(u)...)
	}
	return comments
}
