package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	url         = "https://www.dictionary.com/e/crb-ajax/cached.php?page=%d&wordLength=5&letter=%s&action=get_wf_widget_page&pageType=4&nonce=83fdfeeef4"
	aRune       = 97
	zRune       = 122
	toClipboard bool
)

func init() {
	crawlCmd.Flags().BoolVarP(&toClipboard, "clipboard", "", false, "output words to clipboard")
}

var crawlCmd = &cobra.Command{
	Use:     "crawl",
	Short:   "crawls dictionary.com for 5 letter words.",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		words := []string{}
		for i := aRune; i < zRune+1; i++ {
			letterWords, err := fetchLetterWords(strings.ToUpper(string(rune(i))))
			if err != nil {
				return errors.Wrap(err, "fetchLetterWords")
			}
			words = append(words, letterWords...)
		}
		if toClipboard {
			return clipboard.WriteAll(strings.Join(words, "\n"))
		}
		for _, word := range words {
			fmt.Println(word)
		}
		return nil
	},
}

type dictionaryData struct {
	Success bool `json:"success"`
	Data    struct {
		Words []string `json:"words"`
	} `json:"data"`
}

func fetchLetterWords(letter string) ([]string, error) {
	var page = 1
	out := []string{}
	for {

		uri := fmt.Sprintf(url, page, letter)
		req, err := http.NewRequest(http.MethodGet, uri, nil)
		if err != nil {
			return nil, errors.Wrap(err, "http.NewRequest")
		}
		client := http.Client{}
		fmt.Println(letter, page)
		resp, err := client.Do(req)
		if err != nil {
			return nil, errors.Wrap(err, "client.Do")
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "ioutil.ReadAll")
		}
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("bad status code: %d %s", resp.StatusCode, string(data))
		}
		dict := &dictionaryData{}
		if err := json.Unmarshal(data, dict); err != nil {
			fmt.Println(string(data))
			break
		}
		out = append(out, dict.Data.Words...)
		page++
	}
	return out, nil
}
