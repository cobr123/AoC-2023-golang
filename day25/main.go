package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"slices"
	"strings"
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
	//slices.SortFunc(links, func(a, b Link) int {
	//	if a.from != b.from {
	//		return strings.Compare(a.from, b.from)
	//	}
	//	return strings.Compare(a.to, b.to)
	//})
	for i, _ := range links {
		if strings.Compare(links[i].from, links[i].to) >= 0 {
			links[i] = Link{links[i].from, links[i].to}
		} else {
			links[i] = Link{links[i].to, links[i].from}
		}
	}
	queue := make(chan Job)
	pathCounts := make(chan int)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(links []Link, initLinks []Link, in <-chan Job, out chan<- int) {
			for job := range in {
				out <- Part1CheckLinks(links, initLinks, job)
			}
		}(links, initLinks, queue, pathCounts)
	}

	go func(links []Link, out chan<- Job) {
		ll := len(links)
		for i := 0; i < ll-2; i++ {
			for j := i + 1; j < ll-1; j++ {
				fmt.Println("checking", links[i], i, j, ll)
				for k := j + 1; k < ll; k++ {
					out <- Job{i, j, k}
				}
			}
		}
		close(out)
		pathCounts <- -1
	}(links, queue)

	for cnt := range pathCounts {
		if cnt > 0 {
			return cnt
		}
		if cnt == -1 {
			break
		}
	}
	return 0
}

type Job struct {
	i int
	j int
	k int
}

func Part1CheckLinks(links []Link, initLinks []Link, job Job) int {
	link1 := links[job.i]
	link2 := links[job.j]
	link3 := links[job.k]

	tmp := []Link{}
	tmp = append(tmp, links[0:job.i]...)
	tmp = append(tmp, links[job.i+1:job.j]...)
	tmp = append(tmp, links[job.j+1:job.k]...)
	tmp = append(tmp, links[job.k+1:]...)

	found := Part1Check3Links(tmp, link1, link2, link3) &&
		Part1Check3Links(tmp, link2, link3, link1) &&
		Part1Check3Links(tmp, link3, link1, link2)

	if found {
		fmt.Println("found", link1, link2, link3)
		return Part1CountPaths(initLinks, link1, link2, link3)
	} else {
		return 0
	}
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

func Part1Check3Links(links []Link, link1 Link, link2 Link, link3 Link) bool {
	return Part1OneOfPathsExists(links, link1, link2) && !Part1PathExists(links, link1.from, link1.to, []int{}) && Part1OneOfPathsExists(links, link1, link3)

}

func Part1OneOfPathsExists(links []Link, link1 Link, link2 Link) bool {
	return (Part1PathExists(links, link1.from, link2.from, []int{}) && !Part1PathExists(links, link1.from, link2.to, []int{}) && !Part1PathExists(links, link1.to, link2.from, []int{}) && Part1PathExists(links, link1.to, link2.to, []int{})) ||
		(Part1PathExists(links, link1.from, link2.to, []int{}) && !Part1PathExists(links, link1.from, link2.from, []int{}) && !Part1PathExists(links, link1.to, link2.to, []int{}) && Part1PathExists(links, link1.to, link2.from, []int{}))
}

func Part1PathExists(links []Link, from string, to string, visited []int) bool {
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
			if Part1PathExists(links, newFrom, to, newVisited) {
				return true
			}
		}
	}
	return false
}
