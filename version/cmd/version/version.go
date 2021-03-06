package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/sensu/sensu-go/version"
)

var (
	fFullVersion = flag.Bool("f", false, "output the version of the build with iteration")
	fHighest     = flag.Bool("h", false, "output the highest tagged version")
	fIteration   = flag.Bool("i", false, "output the iteration of the build")
	fPrerelease  = flag.Bool("p", false, "output the prerelease version of the build")
	fBuildType   = flag.Bool("t", false, "output the type of build this is")
	fVersion     = flag.Bool("v", false, "output the version of the build without iteration")
)

func buildType(tag string) (string, error) {
	return string(version.BuildTypeFromTag(tag)), nil
}

func main() {
	flag.Parse()
	tag, err := getTag()
	if err != nil {
		log.Fatal(err)
	}
	var fn func(string) (string, error)
	if *fFullVersion {
		fn = version.FullVersion
	} else if *fHighest {
		tags, err := getTags()
		if err != nil {
			log.Fatal(err)
		}
		result, err := version.HighestVersion(tags)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
		return
	} else if *fIteration {
		fn = version.Iteration
	} else if *fPrerelease {
		fn = version.GetPrereleaseVersion
	} else if *fBuildType {
		fn = buildType
	} else if *fVersion {
		fn = version.GetVersion
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
	result, err := fn(tag)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func getTag() (string, error) {
	cmd := exec.Command("git", "describe", "--exact-match", "--tags", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			return "", err
		}
	}
	return string(out), nil
}

func getTags() ([]string, error) {
	cmd := exec.Command("git", "tag", "-l")
	out, err := cmd.Output()
	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			return nil, err
		}
	}
	output := strings.Trim(string(out), "\n")
	tags := strings.Split(output, "\n")
	return tags, nil
}
