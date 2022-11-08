package sdk

import (
	"fmt"
	"strings"
)

const (
	isEqualTo            = "isEqualTo"
	isNotEqualTo         = "isNotEqualTo"
	includes             = "includes"
	excludes             = "excludes"
	propertyExists       = "propertyExists"
	propertyNotExists    = "propertyNotExists"
	regex                = "regex"
	greaterThan          = "greaterThan"
	LessThan             = "LessThan"
	greaterThanOrEqualTo = "greaterThanOrEqualTo"
	LessThanOrEqualTo    = "LessThanOrEqualTo"
)

type Filter struct {
	Option string
	Field  string
	Value  string
}

func MakeFilter(option string, field string, value string) Filter {
	return Filter{
		Option: option,
		Field:  field,
		Value:  value,
	}
}

type RequestOptions struct {
	Limit     int
	Page      int
	Offset    int
	SortField string
	SortType  string
	Filters   []Filter
}

func MakeRequestOptions(limit int, page int, offset int, sortField string, sortType string, filters []Filter) RequestOptions {
	return RequestOptions{
		Limit:     limit,
		Page:      page,
		Offset:    offset,
		SortField: sortField,
		SortType:  sortType,
		Filters:   filters,
	}
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
		if f.Option == isEqualTo || f.Option == regex {
			params = append(params, fmt.Sprintf("%s=%s", f.Field, f.Value))
		} else if f.Option == isNotEqualTo {
			params = append(params, fmt.Sprintf("%s!=%s", f.Field, f.Value))
		} else if f.Option == includes {
			params = append(params, fmt.Sprintf("%s=%s", f.Field, f.Value))
		} else if f.Option == excludes {
			params = append(params, fmt.Sprintf("%s=%s", f.Field, f.Value))
		} else if f.Option == propertyExists {
			params = append(params, f.Value)
		} else if f.Option == propertyNotExists {
			params = append(params, f.Value)
		} else if f.Option == greaterThan {
			params = append(params, fmt.Sprintf("%s>=%s", f.Field, f.Value))
		} else if f.Option == LessThan {
			params = append(params, fmt.Sprintf("%s<=%s", f.Field, f.Value))
		} else {
			fmt.Printf("ERROR: wrong param name: %s\n", f.Option)
		}
	}

	return "?" + strings.Join(params, "&")
}
