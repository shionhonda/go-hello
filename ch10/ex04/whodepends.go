package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type packageDep struct {
	Name string
	Deps []string
}

func listPackages() ([]string, error) {
	out, err := exec.Command("go", "list", "...").Output()
	if err != nil {
		return nil, err
	}
	var pkgs []string
	for _, pkg := range strings.Split(string(out), "\n") {
		if pkg != "" {
			pkgs = append(pkgs, pkg)
		}
	}
	return pkgs, nil
}

func main() {
	var err error
	query := os.Args[1]
	pkgs, err := listPackages()
	if err != nil {
		log.Fatal(err)
	}
	for _, pkg := range pkgs {
		out, err := exec.Command("go", "list", "-json", pkg).Output()
		if err != nil {
			log.Fatal(err)
		}
		var pkgDeps packageDep
		err = json.Unmarshal(out, &pkgDeps)
		if err != nil {
			log.Fatalf("Cannot unmarshal: %v", err)
		}
		for _, dep := range pkgDeps.Deps {
			if dep == query {
				fmt.Println(pkg)
				break
			}
		}
	}

}
