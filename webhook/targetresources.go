package webhook

import (
	"fmt"
	"strings"
)

type ArrayFlags []string

func (i *ArrayFlags) String() string {
	return strings.Trim(fmt.Sprint(i), "[]")
}

func (i *ArrayFlags) Set(value string) error {
	*i = append(*i, strings.TrimSpace(value))
	return nil
}

type set struct {
	data map[string]struct{}
}

func NewSet() *set {
	s := &set{}
	s.data = make(map[string]struct{})
	return s
}

func (s *set) Add(value string) {
	s.data[value] = struct{}{}
}

func (s *set) Remove(value string) {
	delete(s.data, value)
}

func (s *set) Contains(value string) bool {
	_, c := s.data[value]
	return c
}

func (s *set) GetStringArr() []string {
	var arr []string

	for key := range s.data {
		arr = append(arr, key)
	}

	return arr
}

var targetResourcesSet *set

func SetTargetResourcesSet(targetResources ArrayFlags) {
	targetResourcesSet := NewSet()

	for _, resource := range targetResources {
		targetResourcesSet.Add(resource)
	}
}
