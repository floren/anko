// Package sort implements sort interface for anko script.
package sort

import (
	"github.com/mattn/anko/vm"
	"reflect"
	s "sort"
)

type is []interface{}

func (p is) Len() int           { return len(p) }
func (p is) Less(i, j int) bool { return p[i].(int64) < p[j].(int64) }
func (p is) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type fs []interface{}

func (p fs) Len() int           { return len(p) }
func (p fs) Less(i, j int) bool { return p[i].(float64) < p[j].(float64) }
func (p fs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type ss []interface{}

func (p ss) Len() int           { return len(p) }
func (p ss) Less(i, j int) bool { return p[i].(string) < p[j].(string) }
func (p ss) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Import(env *vm.Env) {
	m := env.NewModule("sort")
	m.Define("Ints", reflect.ValueOf(func(ints []interface{}) []interface{} {
		s.Sort(is(ints))
		return ints
	}))
	m.Define("Float64s", reflect.ValueOf(func(ints []interface{}) []interface{} {
		s.Sort(is(ints))
		return ints
	}))
	m.Define("Strings", reflect.ValueOf(func(ints []interface{}) []interface{} {
		s.Sort(is(ints))
		return ints
	}))

}