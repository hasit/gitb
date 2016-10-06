package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/gizak/termui"
)

func main() {
	out, err := exec.Command("git", "branch").Output()
	if err != nil {
		log.Fatal(err)
	}
	branches := strings.Split(string(out), "\n")
	branches = branches[:len(branches)-1]

	branchList := []string{}
	width := 0
	currentSel := 0
	currentBranch := 0

	for index, branch := range branches {
		if branch[0] == 42 {
			branchList = append(branchList, "["+branch+"]"+"(fg-red)")
			currentBranch = index
			currentSel = index
		} else {
			branchList = append(branchList, "["+branch+"]"+"(fg-white)")
		}
		if len(branch) > width {
			width = len(branch)
		}
	}

	err = termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	ls := termui.NewList()
	ls.ItemFgColor = termui.ColorYellow
	ls.BorderLabel = "Branches"
	ls.Height = len(branches) + 2
	ls.Width = width + 3
	ls.Y = 0

	ls.Items = branchList
	termui.Render(ls)

	handleUpDown := func(p, n int) {
		if p != n {
			branchList[n] = strings.Replace(branchList[n], "white", "red", -1)
			branchList[p] = strings.Replace(branchList[p], "red", "white", -1)
		}
		ls.Items = branchList
		termui.Render(ls)
	}
	handleEnter := func(currentSel int) {
		if currentBranch != currentSel {
			branchList[currentBranch] = strings.Replace(branchList[currentBranch], "*", " ", 1)
			branchList[currentSel] = strings.Replace(branchList[currentSel], " ", "*", 1)
		}
		currentBranch = currentSel
		ls.Items = branchList
		termui.Render(ls)
	}
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})
	termui.Handle("/sys/kbd/<down>", func(e termui.Event) {
		last := currentSel
		currentSel = (currentSel + 1) % len(branchList)
		handleUpDown(last, currentSel)
	})
	termui.Handle("/sys/kbd/<up>", func(termui.Event) {
		last := currentSel
		currentSel = (((currentSel - 1) % len(branchList)) + len(branchList)) % len(branchList)
		handleUpDown(last, currentSel)
	})
	termui.Handle("/sys/kbd/<enter>", func(termui.Event) {
		branchListItem := branchList[currentSel]
		branchListItem = strings.TrimPrefix(branchListItem, "[  ")
		branchListItem = strings.TrimSuffix(branchListItem, "](fg-red)")
		exec.Command("git", "checkout", branchListItem).Output()
		if err != nil {
			panic(err)
		}
		handleEnter(currentSel)
	})
	termui.Loop()
}
