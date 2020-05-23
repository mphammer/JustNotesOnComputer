package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/shurcooL/github_flavored_markdown"
	"github.com/shurcooL/github_flavored_markdown/gfmstyle"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(viewCmd)
}

// https://github.com/jmcfarlane/markdown/blob/master/main.go

var viewCmd = &cobra.Command{
	Use:     "view [NOTE]",
	Aliases: []string{"v"},
	Short:   "View a file in your browser",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("this command takes up to 1 argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		notePath := args[0]
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			dir, file := filepath.Split(notePath)
			if r.URL.Path[1:] != "" {
				nDir, nFile := filepath.Split(r.URL.Path[1:])
				if nDir != "" {
					dir = fmt.Sprintf("%s%s", dir, nDir)
				}
				file = nFile
			}
			newLink := fmt.Sprintf("%s%s", dir, file)
			b, err := ioutil.ReadFile(newLink)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			io.WriteString(w, `<html><head><meta charset="utf-8">
				<link href="/assets/gfm.css" media="all" rel="stylesheet" type="text/css" />
				<link href="https://cdnjs.cloudflare.com/ajax/libs/octicons/2.1.2/octicons.css" media="all" rel="stylesheet" type="text/css" />
			</head><body><article class="markdown-body entry-content" style="padding: 30px;">`)
			w.Write(github_flavored_markdown.Markdown(b))
			io.WriteString(w, `</article></body></html>`)
		})
		http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(gfmstyle.Assets)))
		fmt.Printf("Go to: http://localhost:8080\n")
		http.ListenAndServe(":8080", nil)
		return nil
	},
}
