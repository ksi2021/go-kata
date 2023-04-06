package main

import (
	"regexp"
	"testing"
)

var data = []string{
	`üÜöÖäÄèàùÈÀÙêâôîûÊÂÔÎÛëïÿËÏŸçÇéÉ Lorem ipsum ipod sit amet, consectetur adipiscing elit. Aliquam lobortis,!  (ante) vitae rutrum mollis, augue ligula luctus nunc, vel blandit massa lacus eu magna. Nulla sit amet molestie velit, ac auctor velit. Morbi erat quam, varius id augue sed, molestie vestibulum urna. Pellentesque sit amet ---- ligula libero. Nam scelerisque urna sit amet laoreet scelerisque. Duis id sapien enim. Praesent fermentum pharetra mauris id tincidunt. Fusce porta augue ut sodales euismod. Praesent eget ligula imac, hendrerit nulla et, interdum turpis. Etiam rhoncus efficitur libero at tempus. Integer blandit mi et ornare rhoncus. Nulla in odio et dui sagittis hendrerit eu iphone lorem. Suspendisse consectetur pharetra tellus, fringilla pulvinar ex egestas vel.`,

	`üÜöÖäÄèàùÈÀÙêâôîûÊÂÔÎÛëïÿËÏŸçÇéÉ Morbi volutpat risus porta orci finibus porttitor. Integer ipsum nisl, feugiat eu velit id, venenatis fermentum ex. Suspendisse potenti. Maecenas viverra viverra posuere. Nulla vitae pulvinar est. In magna eros, sagittis in ex cursus, consequat tempus quam. Mauris molestie, ex vitae aliquam facilisis,----- purus diam pharetra felis, ut tristique augue velit in purus. Fusce laoreet, sapien eget imperdiet sollicitudin, dui magna sagittis lorem, sed vulputate lectus elit nec justo. Vestibulum elit risus, fringilla non nunc id, ornare blandit mi. Pellentesque ac tincidunt felis. Mauris aliquet scelerisque orci, ac molestie tortor imperdiet eget. Integer rhoncus lectus in augue feugiat sollicitudin. Nunc porta ut purus et molestie.`,

	`üÜöÖäÄèàùÈÀÙêâôîûÊÂÔÎÛëïÿËÏŸçÇéÉ Lorem ipsum ipod sit amet, consectetur adipiscing elit. Pellentesque eu libero vel tellus vestibulum ipod iphone et lectus. Praesent sollicitudin placerat risus iphone dapibus. Proin non ipod eget enim tempus malesuada. iphoneque id diam eu eros viverra pellentesque. Donec iphone ante dapibus (lectus) faucibus blandit. Suspendisse euismod blandit magna, non facilisis arcu tincidunt eu. Donec pharetra in nisl non efficitur.`,

	`üÜöÖäÄèàùÈÀÙêâôîûÊÂÔÎÛëïÿËÏŸçÇéÉ Duis massa orci, pellentesque at ----- nulla at, venenatis - consequat eros. Donec sit amet vehicula urna. Vivamus imac leo non aliquam faucibus. Vestibulum vel varius erat. Fusce dapibus posuere orci, vitae vulputate mauris tristique aliquam. In vel sapien non felis rutrum tincidunt. Nullam auctor purus sed nulla sollicitudin euismod.`,

	`üÜöÖäÄèàùÈÀÙêâôîûÊÂÔÎÛëïÿËÏŸçÇéÉ Donec semper dui iphone ex sollicitudin, non imac felis {dignissim}. Integer macbook aliquam magna, ipod commodo orci posuere a. iphoneque accumsan felis sit amet lacinia vehicula. Morbi velit libero, bibendum non sapien id, molestie maximus arcu. Maecenas aliquam, justo eget mollis porta, nisi ipod sagittis justo, in vehicula augue elit iphone purus. Maecenas magna elit, pulvinar vel luctus fermentum, consequat suscipit est. Morbi laoreet sem eu massa convallis pharetra. Ut iaculis dui eget elit hendrerit venenatis. Sed sed lacus ut turpis dignissim ipod iphone sed tortor. Nulla rhoncus turpis a nulla congue, eget posuere metus porta. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Donec vel est sem. Nulla vitae tortor sit amet ipsum sodales ornare eu a nunc.`,
}

var umlauts = "üÜöÖäÄèàùÈÀÙêâôîûÊÂÔÎÛëïÿËÏŸçÇéÉ"

var replaceWords = map[string]string{
	"(?i)\\bimac\\b":    "iMac",
	"(?i)\\biphone\\b":  "iPhone",
	"(?i)\\bipad\\b":    "iPad",
	"(?i)\\bipod\\b":    "iPod",
	"(?i)\\bmacbook\\b": "MacBook",
}

var replacements = []string{
	`(?i)(\\?\\<\\=\\^|\\s)[.*#=!]+([0-9A-ZА-ЯЁҐЄIЇ\x{0456}\x{0457}` + umlauts + `]+)[.*#=!]+(\\?\\=\\s|$)`,
	`(?i)[^()+\/\\\!;:, \."«»*0-9A-ZА-Я\\–\\—#№ЁҐЄIЇ` + umlauts + `\x{0456}\x{0457}²’‘“”\\'&-]`,
	`^(?i)[^«»'0-9A-ZА-Я\\–\\—ЁҐЄIЇ` + umlauts + "]",
	"^(?i)[^0-9A-ZА-ЯҐЄIЇ" + umlauts + "]$",
	"[<\\[{]",
	"[>\\]}]",
	"[_]",
	"^[() +\\/,.-]+",
	"[(\\/ ,.-]+$",
	"!+",
	umlauts,
}

var replacementsMap = map[string]string{
	`(?i)(\\?\\<\\=\\^|\\s)[.*#=!]+([0-9A-ZА-ЯЁҐЄIЇ\x{0456}\x{0457}` + umlauts + `]+)[.*#=!]+(\\?\\=\\s|$)`: "${1}",
	`(?i)[^()+\/\\\!;:, \."«»*0-9A-ZА-Я\\–\\—#№ЁҐЄIЇ` + umlauts + `\x{0456}\x{0457}²’‘“”\\'&-]`:             "",
	`^(?i)[^«»'0-9A-ZА-Я\\–\\—ЁҐЄIЇ` + umlauts + "]":                                                        "",
	"^(?i)[^0-9A-ZА-ЯҐЄIЇ" + umlauts + "]$":                                                                 "",
	"[<\\[{]":        "(",
	"[>\\]}]":        ")",
	"[_]":            " ",
	"^[() +\\/,.-]+": "",
	"[(\\/ ,.-]+$":   "",
	"!+":             "",
	umlauts:          "",
}

var (
	replacementMapRe = make(map[string]*regexp.Regexp, len(replacementsMap))
)

func init() {
	for _, pattern := range replacements {
		replacementMapRe[pattern] = regexp.MustCompile(pattern)
	}
}

type FilterText struct {
	replacements         []string
	replacementsMap      map[string]string
	replacementMapRe     map[string]*regexp.Regexp
	replacementWordMapRe map[string]*regexp.Regexp
}

func NewFilterText() *FilterText {
	repMap := map[string]string{
		`(?i)(\\?\\<\\=\\^|\\s)[.*#=!]+([0-9A-ZА-ЯЁҐЄIЇ\x{0456}\x{0457}` + umlauts + `]+)[.*#=!]+(\\?\\=\\s|$)`: "${1}",
		`(?i)[^()+\/\\\!;:, \."«»*0-9A-ZА-Я\\–\\—#№ЁҐЄIЇ` + umlauts + `\x{0456}\x{0457}²’‘“”\\'&-]`:             "",
		`^(?i)[^«»'0-9A-ZА-Я\\–\\—ЁҐЄIЇ` + umlauts + "]":                                                        "",
		"^(?i)[^0-9A-ZА-ЯҐЄIЇ" + umlauts + "]$":                                                                 "",
		"[<\\[{]":        "(",
		"[>\\]}]":        ")",
		"[_]":            " ",
		"^[() +\\/,.-]+": "",
		"[(\\/ ,.-]+$":   "",
		"!+":             "",
		umlauts:          "",
	}
	filterText := &FilterText{
		replacements: []string{
			`(?i)(\\?\\<\\=\\^|\\s)[.*#=!]+([0-9A-ZА-ЯЁҐЄIЇ\x{0456}\x{0457}` + umlauts + `]+)[.*#=!]+(\\?\\=\\s|$)`,
			`(?i)[^()+\/\\\!;:, \."«»*0-9A-ZА-Я\\–\\—#№ЁҐЄIЇ` + umlauts + `\x{0456}\x{0457}²’‘“”\\'&-]`,
			`^(?i)[^«»'0-9A-ZА-Я\\–\\—ЁҐЄIЇ` + umlauts + "]",
			"^(?i)[^0-9A-ZА-ЯҐЄIЇ" + umlauts + "]$",
			"[<\\[{]",
			"[>\\]}]",
			"[_]",
			"^[() +\\/,.-]+",
			"[(\\/ ,.-]+$",
			"!+",
			umlauts,
		},
		replacementsMap:      repMap,
		replacementMapRe:     make(map[string]*regexp.Regexp, len(repMap)),
		replacementWordMapRe: make(map[string]*regexp.Regexp, len(repMap)),
	}

	for _, pattern := range filterText.replacements {
		filterText.replacementMapRe[pattern] = regexp.MustCompile(pattern)
	}

	return filterText
}

func (ft *FilterText) SanitizeText(text string) string {
	for _, pattern := range replacements {
		text = ft.replacementMapRe[pattern].ReplaceAllString(text, replacementsMap[pattern])
	}
	for pattern, replacement := range replaceWords {
		re := regexp.MustCompile(pattern)
		text = re.ReplaceAllString(text, replacement)
	}
	return text
}

func SanitizeText(text string) string {
	for _, pattern := range replacements {
		re := regexp.MustCompile(pattern)
		text = re.ReplaceAllString(text, replacementsMap[pattern])
	}

	for pattern, replacement := range replaceWords {
		re := regexp.MustCompile(pattern)
		text = re.ReplaceAllString(text, replacement)
	}
	return text
}

func BenchmarkSanitizeText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := range data {
			SanitizeText(data[i])
		}
	}
}

func BenchmarkSanitizeText2(b *testing.B) {
	ft := NewFilterText()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := range data {
			ft.SanitizeText(data[i])
		}
	}
}
