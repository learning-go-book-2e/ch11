# Exercise 3

# Question
Cross-compile your program for ARM64 on Windows. If you are using an ARM64 Windows computer, cross-compile for AMD64 on Linux.

# Solution

Run the following commands:

```shell
$ cd ch11/exercise_solutions/ex1
$ GOOS=windows GOARCH=arm64 go build
```

This produces a binary called `ex1.exe`. On Mac OS and various other Unix work-alikes, use the `file` command to validate the right binary was created:

```shell
$ file ex1.exe
ex1.exe: PE32+ executable (console) Aarch64 (stripped to external PDB), for MS Windows
```

If you are running on ARM64 Windows, use the following commands to build for Linux on AMD64:

```shell
$ GOOS=linux GOARCH=amd64 go build
```

This produces a binary called `ex1`. 

Windows doesn't have a command like `file`, but you can install the `fil` command from https://github.com/joeky888/fil using:

```shell
$ go install github.com/joeky888/fil@latest
```

running it will give you:

```shell
$ fil ex1
ex1: Elf file executable, 64bit LSB x86-64, statically linked
```
