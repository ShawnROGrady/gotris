package main

import "strings"

type stringArrayFlag []string

func (s *stringArrayFlag) String() string {
	return strings.Join(*s, ", ")
}

func (s *stringArrayFlag) Set(val string) error {
	*s = append(*s, val)
	return nil
}
