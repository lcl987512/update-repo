package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func IsGitDir(path string) int {
	if strings.Compare(path, ".git") == 0 {
		return 1
	}
	if strings.Contains(path, ".git") {
		return 2
	}
	return 0
}

func IsSvnDir(path string) int {
	if strings.Compare(path, ".git") == 0 {
		return 1
	}
	if strings.Contains(path, ".git") {
		return 2
	}
	return 0
}

func RunGitCommand(path string, gitCmd string) {
	cmd := exec.Command("CMD", "/C", "git -C", path, gitCmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()

	fmt.Println(out.String())
}

func UpdateRepos(path string) {
	dir, _ := ioutil.ReadDir(path)
	for _, fi := range dir {
		if fi.IsDir() {
			if IsGitDir(fi.Name()) == 1 {

				fmt.Printf("call git path=%v is a git dir=%v EXEC pull\n", path, fi.Name())
				RunGitCommand(path, "pull")
			} else if IsGitDir(fi.Name()) == 2 {

				fmt.Printf("call git path=%v is a bare git dir=%v EXEC fetch\n", path, fi.Name())
				RunGitCommand(path, "fetch")
			} else {
				//fmt.Printf("call next dir %v\n", fi.Name())
				UpdateRepos(path + "/" + fi.Name())
			}
		}
	}
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	UpdateRepos(root)
}
