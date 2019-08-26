git-object-trainer
======

## Description

You can learn git object and git internal by `git-object-trainer`.

`git-object-trainer` dosen't allow you to **Porcelain** git commands(add, commit, show, diff, checkout, tag, log, diff). You have to use **Plumbing** git commands(hash-object, cat-file, update-index, write-tree) instead of **Porcelain** git commands.

If you exec `git add`, `git-object-trainer` teach you how to construct `git add` with Plumbing git command.

- [10. Git Internals](https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain)

## Getting Started

### Install
```shell
go get github.com/sonatard/git-object-trainer/
```

### Settings

Add alias to your `.bash_profile`

- .bash_profile

```
alias git=git-object-trainer
```


## Example

```shell
$ git add .
Error: Exec following commands instead of git add

Basic way
# 1. Create . objects
$ git hash-object -w .
[FILE-HASH-STRING]
# 3. Check created object
$ git cat-file -p [FILE-HASH-STRING]
[FILE STRINGS]
# 4. Add or update file to staging area
## 4.1 add
$ git update-index --add --caechinfo 10644 [FILE-HASH-STRING] .
## 4.2 update
$ git update-index --caechinfo 10644 [FILE-HASH-STRING] .
# 5. Create tree
$ git write-tree
[TREE-HASH-STRING]
Remember [TREE-HASH-STRING] for commit
# 6. Check tree object
$ git cat-file -p [TREE-HASH-STRING]
[PERMISSION] blob [HASH] .

Easy way. You need not to remember [FILE-HASH-STRING]
# 1. Create tree
$ git write-tree
[TREE-HASH-STRING]
Remember [TREE-HASH-STRING] for commit
# 2. Add or update file to staging area
## 4.1 add
$ git update-index --add .
## 4.2 update
$ git update-index  .
# 6. Check tree object
$ git cat-file -p [TREE-HASH-STRING]
[PERMISSION] blob [HASH] .
```
