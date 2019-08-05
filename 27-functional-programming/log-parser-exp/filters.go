package main

import "strings"

func noopFilter(r result) bool {
	return true
}

func not(filter filterFunc) filterFunc {
	return func(r result) bool {
		return !filter(r)
	}
}

func domainExtFilter(domains ...string) filterFunc {
	return func(r result) bool {
		for _, domain := range domains {
			if strings.HasSuffix(r.domain, "."+domain) {
				return true
			}
		}
		return false
	}
}

func domainFilter(domain string) filterFunc {
	return func(r result) bool {
		return strings.Contains(r.domain, domain)
	}
}

func orgDomainsFilter(r result) bool {
	return strings.HasSuffix(r.domain, ".org")
}
