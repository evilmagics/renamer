package main

import "fmt"

var version = "Dev"

func main() {
	args := parseArgs()

	res := Rename(args.Dir, WithPrefix(args.Prefix), WithSuffix(args.Suffix))

	fmt.Println()
	fmt.Println("Dir     :", args.Dir)
	fmt.Println("Count   :", res.Count)
	fmt.Println("Success :", res.Success)
	fmt.Println("Failed  :", res.Failed)
	fmt.Println("Elapsed :", res.Elapsed)
	if len(res.Errors) > 0 {
		fmt.Println("Errors  :")
		for i, e := range res.Errors {
			fmt.Printf("%3d. %-25s : %s\n", i+1, e.Name, e.Err)
		}
	}

}
