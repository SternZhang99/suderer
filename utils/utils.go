package utils

import "strings"

func UrlMaker(base, toJoin string) string {
	ret := []string{base, toJoin}
	return strings.Join(ret, "")
}
