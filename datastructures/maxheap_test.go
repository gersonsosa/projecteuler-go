package datastructures

import (
    "testing"
    "sort"
)

func TestMaxHeap_Peek(t *testing.T) {
    cases := []struct {
        in []int
        want int
    }{
        {[]int{1, 3, 40, 199, 200}, 200},
        {[]int{1}, 1},
        {[]int{0, 100, 90, 150}, 150},
    }
    for _, c := range cases {
        m := NewMaxHeap()
        for _, i := range c.in {
            m.Add(i)
        }
        got,err := m.Peek()
        if err != nil {
            t.Errorf("Peek error(%v): error was: %v", c.in, err)
        }
        if got != c.want {
            t.Errorf("Peek(%v) == %v, want %v", c.in, got, c.want)
        }
    }
}

func TestMaxHeap_Poll(t *testing.T) {
    cases := []struct {
        in []int
    }{
        {[]int{1, 3, 40, 199, 200}},
        {[]int{1}},
        {[]int{0, 100, 90, 150}},
        {[]int{1000, 1100, 90, 1500}},
    }
    for _, c := range cases {
        m := NewMaxHeap()

        for _, i := range c.in {
            m.Add(i)
        }

        sort.Sort(sort.Reverse(sort.IntSlice(c.in)))

        for _, want := range c.in {
            got, err := m.Poll()
            if err != nil {
                t.Errorf("Peek error(%v): error was: %v", c.in, err)
            }
            if got != want {
                t.Errorf("Peek(%v) == %v, want %v", c.in, got, want)
            }
        }
    }
}
