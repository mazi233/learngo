package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(
		`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	genderRe = regexp.MustCompile(
		`<td[^>]*><span class="grayL">性别：</span>([^<]+)</td>`)
	cityUrlRe = regexp.MustCompile(
		`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	genderMatch := genderRe.FindSubmatch(contents)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		gender := string(genderMatch[1])
		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name, gender)
			},
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
