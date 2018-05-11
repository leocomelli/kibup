package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"

	es "github.com/leocomelli/kibup/elasticsearch"
	"github.com/leocomelli/kibup/github"
	"github.com/spf13/cobra"
)

var (
	esHost  string
	esTypes []string
	esIndex string
	esSort  string
	esSize  int
)

var (
	ghToken       string
	ghHost        string
	ghRepo        string
	ghPath        string
	ghFilename    string
	ghAuthorName  string
	ghAuthorEmail string
)

var local bool

var rootCmd = &cobra.Command{
	Use:   "kibup",
	Short: "a simple and smart way to back up kibana objects",
	Run: func(cmd *cobra.Command, args []string) {
		esOpts := es.ESQueryOptions{
			Host:  esHost,
			Index: esIndex,
			Types: esTypes,
			Sort:  esSort,
			Size:  esSize,
		}

		b, err := es.Query(&esOpts)
		if err != nil {
			logrus.Error(err)
			return
		}

		if local {
			err := writeFile(ghFilename, b)
			if err != nil {
				logrus.Error(err)
			}
		}

		if ghRepo != "" {
			ghOpts := github.GithubOptions{
				APIHost:             ghHost,
				PersonalAccessToken: ghToken,
				RepositoryName:      ghRepo,
				Path:                ghPath,
				Filename:            ghFilename,
				AuthorName:          ghAuthorName,
				AuthorEmail:         ghAuthorEmail,
			}
			err = github.UpdateFile(b, &ghOpts)
			if err != nil {
				logrus.Error(err)
				return
			}
		}
	},
}

func Execute() {
	rootCmd.Flags().StringVar(&ghHost, "github", "https://api.github.com/", "github api url")
	rootCmd.Flags().StringVar(&ghRepo, "repo", "", "repository name (owner/repo)")
	rootCmd.Flags().StringVar(&ghFilename, "file", "kibana.json", "backup filename")
	rootCmd.Flags().StringVar(&ghPath, "path", ".", "backup file location")
	rootCmd.Flags().StringVar(&ghAuthorName, "author-name", "kibup@kibup.com", "github author name")
	rootCmd.Flags().StringVar(&ghAuthorEmail, "author-email", "kibup", "github author email")
	rootCmd.Flags().StringVar(&ghToken, "token", "", "github personal access token")

	rootCmd.Flags().StringVar(&esHost, "host", "http://127.0.0.1:9200", "elasticsearch host:port")
	rootCmd.Flags().StringVar(&esIndex, "index", ".kibana", "kibana index name")
	rootCmd.Flags().StringArrayVar(&esTypes, "types", []string{"dashboard", "visualization", "search", "config", "index-pattern"}, "kibana object types")
	rootCmd.Flags().IntVar(&esSize, "size", 10000, "elastisearch query result size")
	rootCmd.Flags().StringVar(&esSort, "sort", "_type", "field to sort the elasticsearch query")

	rootCmd.Flags().BoolVar(&local, "local", false, "create file locally")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Execute()
}

func writeFile(filename string, b []byte) error {
	logrus.WithField("filename", filename).Info("writting file locally")
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(b)

	return nil
}
