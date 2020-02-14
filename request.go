package goreq

import (
	"net/url"
	"strconv"
	"strings"
)

func queryValue(k string, v string, params url.Values) string {
	if _v := params.Get(k); _v != "" {
		return _v
	}
	return v
}

func queryValueInt(k string, v int, params url.Values) int {
	if _v := queryValue(k, "", params); _v != "" {
		if i, err := strconv.Atoi(_v); err == nil {
			return i
		}
	}
	return v
}
func queryValueStringSlice(k string, del string, params url.Values) []string {
	var vals []string
	if _v := queryValue(k, "", params); _v != "" {
		for _, s := range strings.Split(_v, del) {
			vals = append(vals, strings.TrimSpace(s))
		}
	}
	return vals
}

func Query(k string, params url.Values) string {
	return queryValue(k, "", params)
}

func QueryInt(k string, params url.Values) int {
	return queryValueInt(k, 0, params)
}

func QueryStringSlice(k string, del string, params url.Values) []string {
	return queryValueStringSlice(k, del, params)
}
