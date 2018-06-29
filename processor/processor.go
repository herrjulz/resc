package processor

import (
	"regexp"
	"strings"

	"github.com/fatih/color"
)

const (
	header    = `#`
	subheader = `##`
	bold      = `(\*\*|\+\+)([-\_\.\/\w\W\s\p{L}\/]+)(\*\*|\+\+)`
	italic    = `(\_|\+\+)([-\_\.\/\w\W\s\p{L}\/]+)(\_|\+\+)`
)

type Processor struct {
	header    *color.Color
	subheader *color.Color
	bold      *color.Color
	italic    *color.Color

	headerRegexp    *regexp.Regexp
	subheaderRegexp *regexp.Regexp
	boldRegexp      *regexp.Regexp
	italicRegexp    *regexp.Regexp
}

func New() *Processor {
	return &Processor{
		header:    color.New(color.FgCyan, color.Bold, color.Underline),
		subheader: color.New(color.FgWhite, color.Bold),
		bold:      color.New(color.Bold),
		italic:    color.New(color.Italic),

		headerRegexp:    regexp.MustCompile(header),
		subheaderRegexp: regexp.MustCompile(subheader),
		boldRegexp:      regexp.MustCompile(bold),
		italicRegexp:    regexp.MustCompile(italic),
	}
}

func (p *Processor) Process(text string) string {
	text = p.processSubHeader(text)
	text = p.processHeader(text)
	text = p.processBold(text)
	text = p.processItalic(text)
	return text
}

func (p *Processor) toTitle(text string) string {
	text = strings.ToTitle(text)
	return p.header.Sprintf("%s", text)
}

func (p *Processor) processHeader(text string) string {
	if p.headerRegexp.MatchString(text) {
		text = p.headerRegexp.ReplaceAllString(text, "")
		text = strings.TrimSpace(text)
		text = p.toTitle(text)
	}
	return text
}

func (p *Processor) processSubHeader(text string) string {
	if p.subheaderRegexp.MatchString(text) {
		text = p.subheaderRegexp.ReplaceAllString(text, "")
		text = strings.TrimSpace(text)
		text = p.subheader.Sprintf("%s", text)
	}
	return text
}

func (p *Processor) processBold(text string) string {
	if p.boldRegexp.MatchString(text) {
		text = p.boldRegexp.ReplaceAllString(text, p.bold.Sprintf("%s", "$2"))
	}
	return text
}

func (p *Processor) processItalic(text string) string {
	if p.italicRegexp.MatchString(text) {
		text = p.italicRegexp.ReplaceAllString(text, p.italic.Sprintf("%s", "$2"))
	}
	return text
}
