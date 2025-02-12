package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/gizak/termui"
)

// main initializes a CLI for managing Git branches using termui.
// It executes the "git branch" command to obtain the list of branches, identifies the current branch (marked with an asterisk), and formats branch names with color coding based on their selection status.
// The function sets up a terminal UI that displays the branch list, a status message, and a legend of keyboard shortcuts (q for quit, <up>/<down> for navigation, and <enter> for switching branches).
// It registers event handlers to update the highlighted branch on navigation and to execute a "git checkout" for branch switching, updating the UI accordingly with confirmation messages.
// Errors during command execution or UI initialization are handled by logging fatal errors or panicking.
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

	message := termui.NewPar("Branches loaded.")
	message.Width = 16
	message.Height = 5
	message.Border = false
	message.Y = ls.Height / 2
	message.X = ls.Width + 2

	legend := termui.NewPar("q       Quit\n<down>  Next branch\n<up>    Previous branch\n<enter> Switch branch")
	legend.Height = 6
	legend.Width = 25
	legend.Y = ls.Height + 2
	legend.BorderLabel = "Shortcuts"

	ls.Items = branchList
	termui.Render(message, legend, ls)

	handleUpDown := func(p, n int) {
		if p != n {
			branchList[n] = strings.Replace(branchList[n], "white", "red", -1)
			branchList[p] = strings.Replace(branchList[p], "red", "white", -1)
		}
		ls.Items = branchList
		termui.Render(message, legend, ls)
	}

	handleEnter := func(out string, currentSel int) {
		if currentBranch != currentSel {
			branchList[currentBranch] = strings.Replace(branchList[currentBranch], "*", " ", 1)
			branchList[currentSel] = strings.Replace(branchList[currentSel], " ", "*", 1)
		}
		currentBranch = currentSel
		ls.Items = branchList
		m := "Switched branch to '" + strings.Replace(branchList[currentSel], "* ", "", -1) + "'\n" + out
		message.Text = m
		message.Width = len(m) + 3
		termui.Clear()
		termui.Render(message, legend, ls)
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
		branchListItem = branchListItem[3:]
		branchListItem = strings.TrimSuffix(branchListItem, "](fg-red)")
		out, err := exec.Command("git", "checkout", branchListItem).Output()
		if err != nil {
			panic(err)
		}
		handleEnter(string(out), currentSel)
	})
	termui.Loop()
}

// HandleResize processes window resize events for the terminal user interface.
// This placeholder implementation does not adjust the layout and always returns nil.
// In a future update, it should recalculate UI dimensions and return an error if the operation fails.
func handleResize() error {
	x := 10
	return nil
}
