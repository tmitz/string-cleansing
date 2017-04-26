package stringcleansing

import (
	"regexp"
)

func ReplaceHalfHyphen(s string) string {
	re := regexp.MustCompile("[\u002D\u30FC\u2010\u2011\u2012\u2013\u2014\u2015\u2212\uFF0D\u2500\u2501]")
	res := re.ReplaceAllString(s, "-")

	return res
}
