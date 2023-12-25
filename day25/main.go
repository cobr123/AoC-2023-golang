package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"
)

func main() {
	part1()
}

func part1() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := Part1GetMult(scanner)
	fmt.Println(sum)
}

type Link struct {
	from string
	to   string
}

func (l *Link) Same(other *Link) bool {
	return (l.from == other.from && l.to == other.to) || (l.from == other.to && l.to == other.from)
}

func Part1GetLinks(scanner *bufio.Scanner) []Link {
	links := []Link{}
	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, ": ")
		from := parts[0]
		otherParts := strings.Split(parts[1], " ")
		for _, item := range otherParts {
			links = append(links, Link{from, item})
		}
	}
	return links
}

func Part1GetMult(scanner *bufio.Scanner) int {
	links := Part1GetLinks(scanner)
	initLinks := slices.Clone(links)
	slices.SortFunc(links, func(a, b Link) int {
		if a.from != b.from {
			return strings.Compare(a.from, b.from)
		}
		return strings.Compare(a.to, b.to)
	})
	for i, _ := range links {
		if strings.Compare(links[i].from, links[i].to) >= 0 {
			links[i] = Link{links[i].from, links[i].to}
		} else {
			links[i] = Link{links[i].to, links[i].from}
		}
	}
	ll := len(links)
	for i := 0; i < ll-2; i++ {
		link1 := links[i]
		fmt.Println("checking", link1, i, ll-2)
		for j := i + 1; j < ll-1; j++ {
			link2 := links[j]
			for k := j + 1; k < ll; k++ {
				link3 := links[k]
				n := 3
				res := make([]bool, n, n)
				wg := new(sync.WaitGroup)
				wg.Add(n)

				tmp := []Link{}
				tmp = append(tmp, links[0:i]...)
				tmp = append(tmp, links[i+1:j]...)
				tmp = append(tmp, links[j+1:k]...)
				tmp = append(tmp, links[k+1:]...)
				ctx := context.Background()

				go func(ctx context.Context, links []Link, link1 Link, link2 Link, link3 Link, wg *sync.WaitGroup, idx int, arr []bool) {
					defer wg.Done()
					arr[idx] = Part1Check3Links(ctx, links, link1, link2, link3)
				}(ctx, tmp, link1, link2, link3, wg, 0, res)

				go func(ctx context.Context, links []Link, link1 Link, link2 Link, link3 Link, wg *sync.WaitGroup, idx int, arr []bool) {
					defer wg.Done()
					arr[idx] = Part1Check3Links(ctx, links, link1, link2, link3)
				}(ctx, tmp, link2, link3, link1, wg, 1, res)

				go func(ctx context.Context, links []Link, link1 Link, link2 Link, link3 Link, wg *sync.WaitGroup, idx int, arr []bool) {
					defer wg.Done()
					arr[idx] = Part1Check3Links(ctx, links, link1, link2, link3)
				}(ctx, tmp, link3, link1, link2, wg, 2, res)

				wg.Wait()
				found := true
				for _, item := range res {
					if !item {
						found = false
						break
					}
				}
				if found {
					fmt.Println("found", link1, link2, link3)
					return Part1CountPaths(initLinks, link1, link2, link3)
				}
			}
		}
	}
	return 0
}

func Part1CountPaths(links []Link, link1 Link, link2 Link, link3 Link) int {
	idxs := make([]int, 3, 3)
	idxs[0] = slices.IndexFunc(links, func(link Link) bool {
		return link.Same(&link1)
	})
	idxs[1] = slices.IndexFunc(links, func(link Link) bool {
		return link.Same(&link2)
	})
	idxs[2] = slices.IndexFunc(links, func(link Link) bool {
		return link.Same(&link3)
	})
	slices.Sort(idxs)
	newLinks := append(links[0:idxs[0]], links[idxs[0]+1:idxs[1]]...)
	newLinks = append(newLinks, links[idxs[1]+1:idxs[2]]...)
	newLinks = append(newLinks, links[idxs[2]+1:]...)

	var minSize int = 1e6
	var maxSize int = -1e6
	for _, link := range newLinks {
		res := Part1CountLink(links, link)
		if res > maxSize {
			maxSize = res
		}
		if res < minSize {
			minSize = res
		}
	}
	return minSize * maxSize
}

func Part1CountLink(links []Link, link Link) int {
	cnt := 0
	found := true
	for found {
		found = false
		for i, item := range links {
			if item.from == link.from || item.from == link.to {
				link = item
				links = append(links[0:i], links[i+1:]...)
				found = true
				cnt++
				break
			}
		}
	}
	return cnt
}

func Part1Check3Links(ctx context.Context, links []Link, link1 Link, link2 Link, link3 Link) bool {
	n := 3
	res := make([]bool, n, n)
	wg := new(sync.WaitGroup)
	wg.Add(n)

	localCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func(ctx context.Context, idx int, arr []bool) {
		defer wg.Done()
		arr[idx] = !Part1Find(ctx, links, link1.from, link1.to, []int{})
		if !arr[idx] {
			ctx.Done()
		}
	}(localCtx, 0, res)

	go func(ctx context.Context, idx int, arr []bool) {
		defer wg.Done()
		arr[idx] = Part1CheckOrLinks(ctx, links, link1, link2)
		if !arr[idx] {
			ctx.Done()
		}
	}(localCtx, 1, res)

	go func(ctx context.Context, idx int, arr []bool) {
		defer wg.Done()
		arr[idx] = Part1CheckOrLinks(ctx, links, link1, link3)
		if !arr[idx] {
			ctx.Done()
		}
	}(localCtx, 2, res)

	wg.Wait()
	found := true
	for _, item := range res {
		if !item {
			found = false
			break
		}
	}
	return found
}

func Part1CheckOrLinks(ctx context.Context, links []Link, link1 Link, link2 Link) bool {
	n := 2
	res := make([]bool, n, n)
	wg := new(sync.WaitGroup)
	wg.Add(n)

	localCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func(ctx context.Context, idx int, arr []bool) {
		defer wg.Done()
		arr[idx] = Part1Find(ctx, links, link1.from, link2.from, []int{}) && !Part1Find(ctx, links, link1.from, link2.to, []int{}) && !Part1Find(ctx, links, link1.to, link2.from, []int{}) && Part1Find(ctx, links, link1.to, link2.to, []int{})
		if arr[idx] {
			ctx.Done()
		}
	}(localCtx, 0, res)

	go func(ctx context.Context, idx int, arr []bool) {
		defer wg.Done()
		arr[idx] = Part1Find(ctx, links, link1.from, link2.to, []int{}) && !Part1Find(ctx, links, link1.from, link2.from, []int{}) && !Part1Find(ctx, links, link1.to, link2.to, []int{}) && Part1Find(ctx, links, link1.to, link2.from, []int{})
		if arr[idx] {
			ctx.Done()
		}
	}(localCtx, 1, res)

	wg.Wait()

	for _, item := range res {
		if item {
			return true
		}
	}
	return false
}

func Part1Find(ctx context.Context, links []Link, from string, to string, visited []int) bool {
	if ctx.Err() != nil {
		return false
	}
	for i, item := range links {
		if (item.from == from || item.to == from) && !slices.Contains(visited, i) {
			newFrom := item.from
			if item.from == from {
				newFrom = item.to
			}
			if newFrom == to {
				return true
			}
			newVisited := append(visited, i)
			if Part1Find(ctx, links, newFrom, to, newVisited) {
				return true
			}
		}
	}
	return false
}
