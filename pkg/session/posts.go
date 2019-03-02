package session

import (
	"github.com/jzelinskie/geddit"
)

// GetPosts returns the posts for the current selected subreddit
func (s *Session) GetPosts() ([]*geddit.Submission, error) {
	if s.Subreddit == "" {
		return s.getFrontpage()
	}

	return s.getSubreddit()
}

func (s *Session) getFrontpage() ([]*geddit.Submission, error) {
	if s.LoginSess != nil {
		return s.LoginSess.Frontpage(geddit.DefaultPopularity, geddit.ListingOptions{})
	}

	return s.DefaultSess.DefaultFrontpage(geddit.DefaultPopularity, geddit.ListingOptions{})
}

func (s *Session) getSubreddit() ([]*geddit.Submission, error) {
	if s.LoginSess != nil {
		return s.LoginSess.SubredditSubmissions(s.Subreddit, geddit.DefaultPopularity, geddit.ListingOptions{})
	}

	return s.DefaultSess.SubredditSubmissions(s.Subreddit, geddit.DefaultPopularity, geddit.ListingOptions{})
}
