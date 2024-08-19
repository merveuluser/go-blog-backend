package helpers

import (
	"github.com/JesusIslam/tldr"
	"strings"
)

func PostContentSummarizer(postContent string) (string, error) {
	intoSentences := 3
	bag := tldr.New()
	summary, err := bag.Summarize(postContent, intoSentences)
	if err != nil {
		return " ", err
	}

	return strings.Join(summary, " "), nil
}
