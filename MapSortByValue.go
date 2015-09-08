package main

import (
    "fmt"
    "sort"
)

func rankByMapCount(mapFrequencies map[int]int) PairList{
    pl := make(PairList, len(mapFrequencies))
    i := 0
    for k, v := range mapFrequencies {
      pl[i] = Pair{k, v}
      i++
    }
    sort.Sort(sort.Reverse(pl))
    return pl
}

type Pair struct {
    Key   int
    Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }


func main() {
    a := make(map[int]int)
    a[5] = 1
    a[4] = 2 
    a[3] = 3 
    a[2] = 4
    a[1] = 5
    fmt.Println(rankByMapCount(a))
