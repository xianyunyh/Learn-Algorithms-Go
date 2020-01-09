package hash

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {

	table := make(map[string]int)

	for _, v := range cpdomains {
		items := strings.Split(v, " ")
		if len(items) <= 1 {
			continue
		}
		n, _ := strconv.Atoi(items[0])
		t := strings.Split(items[1], ".")
		if len(t) > 2 {
			k := strings.Join(t[1:], ".")
			if _, ok := table[k]; ok {
				table[k] += n
			} else {
				table[k] = n
			}
		}

		//后缀域名
		suffix := t[len(t)-1]
		if _, ok := table[suffix]; ok {
			table[suffix] += n
		} else {
			table[suffix] = n
		}

		//全域名
		if _, ok := table[items[1]]; ok {
			table[items[1]] += n
		} else {
			table[items[1]] = n
		}

	}

	result := []string{}
	for k, v := range table {
		result = append(result, strconv.Itoa(v)+" "+k)
	}
	return result
}
