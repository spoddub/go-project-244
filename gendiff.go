package code

import (
	"code/formatters"
	"code/parsers"
)

func GenDiff(path1, path2, format string) (string, error) {
	data1, err := parsers.Parse(path1)
	if err != nil {
		return "", err
	}

	data2, err := parsers.Parse(path2)
	if err != nil {
		return "", err
	}

	if format == "" {
		format = "stylish"
	}

	return formatters.Format(data1, data2, format)
}
