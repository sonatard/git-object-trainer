package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: git --help")
		os.Exit(1)

	}
	args := os.Args[1:]
	err := porcelainToPlumbing(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Exec following commands instead of git %v\n\n%v\n\n", os.Args[1], err)
		os.Exit(1)
	}

	// git command wrapper
	gitCmd := exec.Command("git", args...)
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	err = gitCmd.Run()

	if err != nil {
		status := 1
		if e2, ok := err.(*exec.ExitError); ok {
			if s, ok := e2.Sys().(syscall.WaitStatus); ok {
				status = s.ExitStatus()
			}
		}
		os.Exit(status)
	}
}

func porcelainToPlumbing(args []string) error {

	switch args[0] {
	case "add":
		if len(args) < 2 {
			return errors.New("Nothing specified, nothing added.\n")
		}
		cmdout := []string{
			"Basic way",
			"# 1. Create " + args[1] + " objects",
			"$ git hash-object -w " + args[1],
			"[FILE-HASH-STRING]",
			"# 3. Check created object",
			"$ git cat-file -p [FILE-HASH-STRING]",
			"[FILE STRINGS]",
			"# 4. Add or update file to staging area",
			"## 4.1 add",
			"$ git update-index --add --caechinfo 10644 [FILE-HASH-STRING] " + args[1],
			"## 4.2 update",
			"$ git update-index --caechinfo 10644 [FILE-HASH-STRING] " + args[1],
			"# 5. Create tree",
			"$ git write-tree",
			"[TREE-HASH-STRING]",
			"Remember [TREE-HASH-STRING] for commit",
			"# 6. Check tree object",
			"$ git cat-file -p [TREE-HASH-STRING]",
			"[PERMISSION] blob [HASH] " + args[1],
			"",
			"Easy way. You need not to remember [FILE-HASH-STRING]",
			"# 1. Create tree",
			"$ git write-tree",
			"[TREE-HASH-STRING]",
			"Remember [TREE-HASH-STRING] for commit",
			"# 2. Add or update file to staging area",
			"## 4.1 add",
			"$ git update-index --add " + args[1],
			"## 4.2 update",
			"$ git update-index  " + args[1],
			"# 6. Check tree object",
			"$ git cat-file -p [TREE-HASH-STRING]",
			"[PERMISSION] blob [HASH] " + args[1],
		}
		return errors.New(strings.Join(cmdout, "\n"))
	case "commit":
		cmdout := []string{
			"# 1. Check current branch",
			"$ cat .git/HEAD",
			"ref: refs/heads/[BRANCH-NAME]",
			"# 2. Check last commit hash",
			"$ cat .git/refs/heads/[BRANCH-NAME]",
			"If first commit, don't exec",
			"[LAST-COMMIT-HASH]",
			"# 3. Commit, connect LATEST-COMMIT",
			"$ echo 'commit message' | git commit-tree [TREE-HASH-STRING] -p [LAST-COMMIT-HASH]",
			"If first commit, remove -p [LAST-COMMIT-HASH]",
			"[COMMIT-HASH-STRING]",
			"# 4. Check commit object",
			"$ git cat-file -p [COMMIT-HASH-STRING]",
			"[COMMIT INFO]",
			"# 5. Add commit to branch",
			"$ git update-ref refs/heads/[BRANCH-NAME] [COMMIT-HASH-STRING]",
		}
		return errors.New(strings.Join(cmdout, "\n"))
	case "checkout":
		if len(args) < 2 {
			return errors.New("Nothing specified, nothing checkout.\n")
		}
		cmdout := []string{
			"# 1. Checkout",
			"$ git symbolic-ref HEAD refs/heads/" + args[1],
		}
		return errors.New(strings.Join(cmdout, "\n"))
	case "tag":
		if len(args) < 2 {
			return errors.New("Nothing specified, nothing tag.\n")
		}
		cmdout := []string{
			"# 1. tag",
			"$ git update-ref refs/tags/" + args[1],
			"# 2. check tag hash",
			"$ cat .git/refs/tags/" + args[1],
			"[TAG-HASH]",
			"# 3. Check tag",
			"$  git cat-file -p [TAG-HASH]",
			"[TAG-INFORMATION]",
		}
		return errors.New(strings.Join(cmdout, "\n"))
	case "log":
		cmdout := []string{
			"# 1. Check current branch",
			"$ cat .git/HEAD",
			"ref: refs/heads/[BRANCH-NAME]",
			"# 2. Check last commit hash",
			"$ cat .git/refs/heads/[BRANCH-NAME]",
			"[LAST-COMMIT-HASH]",
			"# 3. Check last commit log",
			"$ git cat-file -p [LAST-COMMIT-HASH]",
			"tree [TREE-HASH]",
			"parent [BEFORE-COMMIT-HASH]",
			"author [AUTHOR-INFO]",
			"committer [COMMITER-INFO]",
			"[COMMIT-LOG]",
			"# 4. Check before commit log",
			"$ git cat-file -p [BEFORE-COMMIT-HASH]",
			"tree [BEFORE-TREE-HASH]",
			"parent [BEFORE-COMMIT-HASH]",
			"author [BEFORE-AUTHOR-INFO]",
			"committer [BEFORE-COMMITER-INFO]",
			"[BEFORE-COMMIT-LOG]",
			"# 5. Repeate # 4. Check before commit log",
		}
		return errors.New(strings.Join(cmdout, "\n"))
	case "show":
		fallthrough
	case "diff":
		cmdout := []string{
			"# 1. Check current branch",
			"$ cat .git/HEAD",
			"ref: refs/heads/[BRANCH-NAME]",
			"# 2. Check last commit hash",
			"$ cat .git/refs/heads/[BRANCH-NAME]",
			"[LATEST-COMMIT-HASH]",
			"# 3. Check last commit log",
			"$ git cat-file -p [LATEST-COMMIT-HASH]",
			"tree [LATEST-TREE-HASH]",
			"parent [BEFORE-COMMIT-HASH]",
			"author [LATESTAUTHOR-INFO]",
			"committer [LATESTCOMMITER-INFO]",
			"[LATEST-COMMIT-LOG]",
			"# 4. Check before commit log",
			"$ git cat-file -p [BEFORE-COMMIT-HASH]",
			"tree [BEFORE-TREE-HASH]",
			"parent [BEFORE-COMMIT-HASH]",
			"author [BEFORE-AUTHOR-INFO]",
			"committer [BEFORE-COMMITER-INFO]",
			"[BEFORE-COMMIT-LOG]",
			"# 5. Check latest commit files",
			"$ git cat-file -p [LATEST-TREE-HASH]",
			"[PERMISSION1] blob [LATEST-OBJECT-HASH1] FILE-NAME1",
			"[PERMISSION2] blob [LATEST-OBJECT-HASH2] FILE-NAME2",
			"...",
			"# 6. Check before commit files",
			"$ git cat-file -p [BEFORE-TREE-HASH]",
			"[PERMISSION1] blob [BEFORE-OBJECT-HASH1] FILE-NAME1",
			"[PERMISSION2] blob [BEFORE-OBJECT-HASH2] FILE-NAME2",
			"# 7. Show diff",
			"$ git cat-file -p [BEFORE-OBJECT-HASH1] | ( git cat-file -p [LATEST-OBJECT-HASH1] | diff /dev/fd/3 -) 3<&0",
			"[DIFF]",
		}
		return errors.New(strings.Join(cmdout, "\n"))
	}

	return nil
}
