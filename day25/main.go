package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
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

func Part1GetMult(scanner *bufio.Scanner) int {
	links := []string{}
	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, ": ")
		from := parts[0]
		otherParts := strings.Split(parts[1], " ")
		for _, item := range otherParts {
			links = append(links, from+"/"+item)
		}
	}
	for i := 0; i < len(links)-2; i++ {
		for j := i + 1; j < len(links)-1; j++ {
			for k := j + 1; k < len(links); k++ {
				fromTo1 := strings.Split(links[i], "/")
				fromTo2 := strings.Split(links[j], "/")
				fromTo3 := strings.Split(links[k], "/")
				link1 := Link{fromTo1[0], fromTo1[1]}
				link2 := Link{fromTo2[0], fromTo2[1]}
				link3 := Link{fromTo3[0], fromTo3[1]}
				fmt.Println("checking", fromTo1, fromTo2, fromTo3)
				n := 3
				res := make([]bool, n, n)
				wg := new(sync.WaitGroup)
				wg.Add(n)

				tmp := []string{}
				tmp = append(tmp, links[i+1:j]...)
				tmp = append(tmp, links[j+1:k]...)
				tmp = append(tmp, links[k+1:]...)
				ctx := context.Background()

				go func(ctx context.Context, links []string, link1 Link, link2 Link, link3 Link, wg *sync.WaitGroup, idx int, arr []bool) {
					defer wg.Done()
					arr[idx] = Part1Check3Links(ctx, links, link1, link2, link3)
				}(ctx, tmp, link1, link2, link3, wg, 0, res)

				go func(ctx context.Context, links []string, link1 Link, link2 Link, link3 Link, wg *sync.WaitGroup, idx int, arr []bool) {
					defer wg.Done()
					arr[idx] = Part1Check3Links(ctx, links, link1, link2, link3)
				}(ctx, tmp, link2, link3, link1, wg, 1, res)

				go func(ctx context.Context, links []string, link1 Link, link2 Link, link3 Link, wg *sync.WaitGroup, idx int, arr []bool) {
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
					fmt.Println("found", fromTo1, fromTo2, fromTo3)
					return 0
				}
			}
		}
	}
	return 0
}

func Part1Check3Links(ctx context.Context, links []string, link1 Link, link2 Link, link3 Link) bool {
	n := 3
	res := make([]bool, n, n)
	wg := new(sync.WaitGroup)
	wg.Add(n)

	ll := len(links)
	localCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func(ctx context.Context, idx int, arr []bool) {
		defer wg.Done()
		arr[idx] = !Part1Find(ctx, links, link1.from, link1.to, ll)
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

func Part1CheckOrLinks(ctx context.Context, links []string, link1 Link, link2 Link) bool {
	n := 2
	res := make([]bool, n, n)
	wg := new(sync.WaitGroup)
	wg.Add(n)

	ll := len(links)
	localCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func(ctx context.Context, idx int, arr []bool) {
		defer wg.Done()
		arr[idx] = Part1Find(ctx, links, link1.from, link2.from, ll) && !Part1Find(ctx, links, link1.from, link2.to, ll) && !Part1Find(ctx, links, link1.to, link2.from, ll) && Part1Find(ctx, links, link1.to, link2.to, ll)
		if arr[idx] {
			ctx.Done()
		}
	}(localCtx, 0, res)

	go func(ctx context.Context, idx int, arr []bool) {
		defer wg.Done()
		arr[idx] = Part1Find(ctx, links, link1.from, link2.to, ll) && !Part1Find(ctx, links, link1.from, link2.from, ll) && !Part1Find(ctx, links, link1.to, link2.to, ll) && Part1Find(ctx, links, link1.to, link2.from, ll)
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

func Part1Find(ctx context.Context, links []string, from string, to string, cnt int) bool {
	if cnt < 0 {
		return false
	}
	if ctx.Err() != nil {
		return false
	}
	for i, item := range links {
		fromTo := strings.Split(item, "/")
		if fromTo[0] == from || fromTo[1] == from {
			newFrom := fromTo[0]
			if fromTo[0] == from {
				newFrom = fromTo[1]
			}
			if newFrom == to {
				return true
			}
			tmp := []string{}
			tmp = append(tmp, links[0:i]...)
			tmp = append(tmp, links[i+1:]...)
			if Part1Find(ctx, tmp, newFrom, to, cnt-1) {
				return true
			}
		}
	}
	return false
}
