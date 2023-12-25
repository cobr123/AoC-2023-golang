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

				go func(ctx context.Context, links []string, from1 string, to1 string, from2 string, to2 string, from3 string, to3 string, wg *sync.WaitGroup, idx int, arr []bool) {
					defer wg.Done()
					arr[idx] = Part1CheckLinks(ctx, links, from1, to1, from2, to2, from3, to3)
				}(ctx, tmp, fromTo1[0], fromTo1[1], fromTo2[0], fromTo2[1], fromTo3[0], fromTo3[1], wg, 0, res)

				go func(ctx context.Context, links []string, from1 string, to1 string, from2 string, to2 string, from3 string, to3 string, wg *sync.WaitGroup, idx int, arr []bool) {
					defer wg.Done()
					arr[idx] = Part1CheckLinks(ctx, links, from1, to1, from2, to2, from3, to3)
				}(ctx, tmp, fromTo2[0], fromTo2[1], fromTo3[0], fromTo3[1], fromTo1[0], fromTo1[1], wg, 1, res)

				go func(ctx context.Context, links []string, from1 string, to1 string, from2 string, to2 string, from3 string, to3 string, wg *sync.WaitGroup, idx int, arr []bool) {
					defer wg.Done()
					arr[idx] = Part1CheckLinks(ctx, links, from1, to1, from2, to2, from3, to3)
				}(ctx, tmp, fromTo3[0], fromTo3[1], fromTo1[0], fromTo1[1], fromTo2[0], fromTo2[1], wg, 2, res)

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
				}
			}
		}
	}
	return 0
}

func Part1CheckLinks(ctx context.Context, links []string, from1 string, to1 string, from2 string, to2 string, from3 string, to3 string) bool {
	n := 3
	res := make([]bool, n, n)
	wg := new(sync.WaitGroup)
	wg.Add(n)

	ll := len(links)

	go func(ctx context.Context, idx int, arr []bool) {
		defer wg.Done()
		arr[idx] = !Part1Find(ctx, links, from1, to1, ll)
		if !arr[idx] {
			ctx.Done()
		}
	}(ctx, 0, res)

	go func(ctx context.Context, idx int, arr []bool) {
		defer wg.Done()
		arr[idx] = (Part1Find(ctx, links, from1, from2, ll) && !Part1Find(ctx, links, from1, to2, ll) && !Part1Find(ctx, links, to1, from2, ll) && Part1Find(ctx, links, to1, to2, ll)) ||
			(Part1Find(ctx, links, from1, to2, ll) && !Part1Find(ctx, links, from1, from2, ll) && !Part1Find(ctx, links, to1, to2, ll) && Part1Find(ctx, links, to1, from2, ll))
		if !arr[idx] {
			ctx.Done()
		}
	}(ctx, 1, res)

	go func(ctx context.Context, idx int, arr []bool) {
		defer wg.Done()
		arr[idx] = (Part1Find(ctx, links, from1, from3, ll) && !Part1Find(ctx, links, from1, to3, ll) && !Part1Find(ctx, links, to1, from3, ll) && Part1Find(ctx, links, to1, to3, ll)) ||
			(Part1Find(ctx, links, from1, to3, ll) && !Part1Find(ctx, links, from1, from3, ll) && !Part1Find(ctx, links, to1, to3, ll) && Part1Find(ctx, links, to1, from3, ll))
		if !arr[idx] {
			ctx.Done()
		}
	}(ctx, 2, res)

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
