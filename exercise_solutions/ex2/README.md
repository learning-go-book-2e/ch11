# Exercise 2

## Question
Use +go install+ to install +staticcheck+. Run it against your program and fix any problems it finds.

## Solution

Run the following commands:

```shell
go install honnef.co/go/tools/cmd/staticcheck@latest
cd ch11/exercise_solutions/ex1
staticcheck ./...
```
