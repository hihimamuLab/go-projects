package main

import "fmt"

type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}

type IntSlice []int

func (i IntSlice) Len() int {
	return len(i)
}

type Slice interface {
	Len() int
}

type FloatSlice []float64

func (f FloatSlice) Len() int {
	return len(f)
}

func print(s Slice) {
	for i := 0; i < s.Len(); i++ {
		fmt.Println(s)
	}
}

func main() {
	s := StringSlice{"foo", "bar", "fuga", "hoge", "piyo"}
	print(s)
	i := IntSlice{1, 2, 3, 4, 5}
	print(i)
	f := FloatSlice{1, 2, 3, 4, 5}
	print(f)
}
