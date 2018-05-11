package github

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"

	"context"

	gh "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	commitMessage = "update %s"
	defaultBranch = "master"
)

type Github struct {
	http   *http.Client
	client *gh.Client
}

type GithubOptions struct {
	APIHost             string
	PersonalAccessToken string
	RepositoryName      string
	Path                string
	Filename            string
	AuthorName          string
	AuthorEmail         string
}

func UpdateFile(b []byte, opts *GithubOptions) error {

	g := new(opts)

	path := fmt.Sprintf("%s/%s", opts.Path, opts.Filename)
	sha, err := g.getSHA(opts.RepositoryName, path)
	if err != nil {
		return err
	}

	message := fmt.Sprintf(commitMessage, opts.Filename)

	repositoryContentsOptions := &gh.RepositoryContentFileOptions{
		Message:   &message,
		Content:   b,
		SHA:       &sha,
		Committer: &gh.CommitAuthor{Name: &opts.AuthorName, Email: &opts.AuthorEmail},
		Branch:    &defaultBranch,
	}

	owner, name := splitRepoName(opts.RepositoryName)
	updateResponse, _, err := g.client.Repositories.UpdateFile(context.Background(), owner, name, path, repositoryContentsOptions)

	logrus.WithField("filename", opts.Filename).WithField("commit", *updateResponse.Commit.SHA).
		Info("file content has been updated successfully")

	return err
}

func new(opts *GithubOptions) *Github {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: opts.PersonalAccessToken},
	)

	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := gh.NewClient(tc)

	url, _ := url.Parse(opts.APIHost)
	client.BaseURL = url

	return &Github{http: tc, client: client}
}

func (g *Github) getSHA(repo, filename string) (string, error) {
	owner, name := splitRepoName(repo)
	fc, _, resp, err := g.client.Repositories.GetContents(context.Background(), owner, name, filename, nil)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			logrus.WithField("filename", filename).Info("file not found (unversioned)")
			return "", nil
		}
		return "", err
	}

	logrus.WithField("filename", filename).WithField("sha", *fc.SHA).Info("getting file metadata")

	return *fc.SHA, nil
}

func splitRepoName(repo string) (string, string) {
	r := strings.Split(repo, "/")
	return r[0], r[1]
}
