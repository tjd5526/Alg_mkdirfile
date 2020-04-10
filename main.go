package main

import (
	"os"

	"github.com/alexflint/go-arg"
)

var args struct {
	SITE string `arg:"-s, required" help:"Site name"`
	DIFF string `arg:"-d, required" help:"Difficulty"`
	PROB string `arg:"-p, required" help:"Problem Name"`
	LANG string `arg:"-l, required" help:"Language Name"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	arg.MustParse(&args)
	path := "./" + args.SITE + "/" + args.DIFF + "/" + args.PROB
	os.MkdirAll(path, os.ModeDir)
	if args.LANG == "go" {
		f1, err := os.Create(path + "/main.go")
		check(err)
		f2, err := os.Create(path + "/README.md")
		check(err)
		defer f1.Close()
		defer f2.Close()

		f1.WriteString("package main\n\nfunc main() {\n}")
		f2.WriteString("[]()\n```go\n```")
	}
	if args.LANG == "python" {
		f1, err := os.Create(path + "/main.py")
		check(err)
		f2, err := os.Create(path + "/README.md")
		check(err)
		defer f1.Close()
		defer f2.Close()

		f2.WriteString("[]()\n```python\n```")
	}

}
