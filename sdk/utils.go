package sdk

import (
	"fmt"
	"strings"
)

const (
	IsEqualTo            = "isEqualTo"
	IsNotEqualTo         = "isNotEqualTo"
	Includes             = "includes"
	Excludes             = "excludes"
	PropertyExists       = "propertyExists"
	PropertyNotExists    = "propertyNotExists"
	Regex                = "regex"
	GreaterThan          = "greaterThan"
	LessThan             = "LessThan"
	GreaterThanOrEqualTo = "greaterThanOrEqualTo"
	LessThanOrEqualTo    = "LessThanOrEqualTo"
)

type Filter struct {
	Option string
	Field  string
	Value  string
}

type RequestOptions struct {
	Limit     int
	Page      int
	Offset    int
	SortField string
	SortType  string
	Filters   []Filter
}

func MakeRequestURL(url string, reqOpt RequestOptions) string {
	return url + encodeRequestOptions(reqOpt)
}

func encodeRequestOptions(reqOpt RequestOptions) string {
	params := []string{}

	if reqOpt.Limit != 0 {
		params = append(params, fmt.Sprintf("limit=%d", reqOpt.Limit))
	}

	if reqOpt.Page != 0 {
		params = append(params, fmt.Sprintf("page=%d", reqOpt.Page))
	}

	if reqOpt.Offset != 0 {
		params = append(params, fmt.Sprintf("offset=%d", reqOpt.Offset))
	}

	if reqOpt.SortField != "" && reqOpt.SortType != "" {
		params = append(params, fmt.Sprintf("sort=%s:%s", reqOpt.SortField, reqOpt.SortType))
	}

	for _, f := range reqOpt.Filters {
		if f.Option == IsEqualTo || f.Option == Regex {
			params = append(params, fmt.Sprintf("%s=%s", f.Field, f.Value))
		} else if f.Option == IsNotEqualTo {
			params = append(params, fmt.Sprintf("%s!=%s", f.Field, f.Value))
		} else if f.Option == Includes {
			params = append(params, fmt.Sprintf("%s=%s", f.Field, f.Value))
		} else if f.Option == Excludes {
			params = append(params, fmt.Sprintf("%s=%s", f.Field, f.Value))
		} else if f.Option == PropertyExists {
			params = append(params, f.Value)
		} else if f.Option == PropertyNotExists {
			params = append(params, f.Value)
		} else if f.Option == GreaterThan {
			params = append(params, fmt.Sprintf("%s>=%s", f.Field, f.Value))
		} else if f.Option == LessThan {
			params = append(params, fmt.Sprintf("%s<=%s", f.Field, f.Value))
		} else {
			fmt.Printf("ERROR: wrong param name: %s\n", f.Option)
		}
	}

	return "?" + strings.Join(params, "&")
}
