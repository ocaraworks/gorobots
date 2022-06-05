package main

import (
	"bufio"
	"io"
	"net/url"
	"regexp"
	"robots/helpers"
	"strings"
)

type Robots struct {
}

//Parse 解析
func (s Robots) Parse(bodyReader io.Reader) (map[string]map[string][]interface{}, []interface{}, error) {
	var err error
	var rules = make(map[string]map[string][]interface{}, 0)
	var sitemaps = make([]interface{}, 0)

	b := bufio.NewReader(bodyReader)
	agentName := ""
	isEof := false

	for {
		content, err := b.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				isEof = true
			} else {
				break
			}
		}

		content = strings.Trim(content, helpers.GetSpace())

		if content != "" {
			rawData := strings.Split(content, "#")
			//注释前面为空
			if len(rawData) <= 0 {
				break
			}

			row := strings.Split(strings.Trim(rawData[0], helpers.GetSpace()), ":")
			length := len(row)

			//没有key时
			if length <= 0 {
				break
			}

			key := strings.ToLower(row[0])
			value := strings.Join(row[1:], ":")

			if length >= 2 {
				value = strings.TrimRight(strings.TrimRight(strings.Trim(value, helpers.GetSpace()), "\n"), "\r")
				if key == "user-agent" {
					agentName = strings.ToLower(value)
					if _, ok := rules[agentName]; ok == false {
						rules[agentName] = make(map[string][]interface{}, 0)
					}
				} else if key == "sitemap" {
					sitemaps = append(sitemaps, value)
				} else {
					if agentName != "" {
						if _, ok := rules[agentName][key]; ok == false {
							rules[agentName][key] = make([]interface{}, 0)
						}
						rules[agentName][key] = append(rules[agentName][key], value)
					}
				}
			}
		}

		if isEof {
			break
		}
	}

	return rules, sitemaps, err
}

//CheckPath 检测路径
func (s Robots) CheckPath(urlPath, spiderName string, authList map[string]map[string][]interface{}) (bool, string, int) {
	var res bool
	var matchRes string
	var delaySeconds int

	res = true
	spiderName = strings.ToLower(spiderName)
	u, _ := url.Parse(urlPath)
	urlPath = u.Path
	urlPath = strings.TrimLeft(strings.Replace(urlPath, "?", "\\?", 0), "/")

	//如果全匹配允许，则检查具体匹配的禁止权限
	if val, ok := authList[spiderName]; ok {
	outerDisallow:
		for k, v := range val {
			if k == "disallow" {
				for _, m := range v {
					if isMatch, matchString := s.MatchAuth(m, urlPath, true); isMatch {
						res = false
						matchRes = matchString
						break outerDisallow
					}
				}
			}
		}
		//如果最终禁止，则检查具体匹配的允许权限
		if res == false {
		outerAllow:
			for k, v := range val {
				if k == "allow" {
					for _, m := range v {
						if isMatch, matchString := s.MatchAuth(m, urlPath, false); isMatch {
							res = true
							matchRes = matchString
							break outerAllow
						}
					}
				} else if k == "crawl-delay" {
					delaySeconds = helpers.ToInt(v)
				}
			}
		}
	} else {
		//先检查全匹配禁止权限
		if val, ok = authList["*"]; ok {
		globalDisallow:
			for k, v := range val {
				if k == "disallow" {
					for _, m := range v {
						if isMatch, matchString := s.MatchAuth(m, urlPath, true); isMatch {
							res = false
							matchRes = matchString
							break globalDisallow
						}
					}
				}
			}
		}
		//如果最终禁止，则检查具体匹配的允许权限
		if res == false {
			if val, ok = authList["*"]; ok {
			globalAllow:
				for k, v := range val {
					if k == "allow" {
						for _, m := range v {
							if isMatch, matchString := s.MatchAuth(m, urlPath, false); isMatch {
								res = true
								matchRes = matchString
								break globalAllow
							}
						}
					} else if k == "crawl-delay" {
						delaySeconds = helpers.ToInt(v)
					}
				}
			}
		}
	}

	return res, matchRes, delaySeconds
}

//MatchAuth 匹配权限
func (s Robots) MatchAuth(matchStr interface{}, urlPath string, allowEmpty bool) (bool, string) {
	if matchStr == "/" {
		return true, ""
	}

	if matchStr == "" {
		return allowEmpty, ""
	}

	ms := helpers.ToString(matchStr)
	ms = strings.TrimLeft(ms, "/")

	//?转/?
	ms = strings.Replace(ms, string(rune(63)), string(rune(92))+string(rune(63)), -1)
	match, _ := regexp.MatchString(ms, urlPath)

	return match, ms
}
