package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	build     string
	buildDate string
)

func main() {
	fmt.Printf("Drone Rubygems Plugin built at %s\n", buildDate)

	workspace := drone.Workspace{}
	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("workspace", &workspace)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if (len(vargs.Username) == 0 || len(vargs.Password) == 0) && len(vargs.APIKey) == 0 {
		fmt.Println("Please provide an API key or username/password credentials")

		os.Exit(1)
		return
	}

	dpl := buildDpl(&workspace, &repo, &build, &vargs)

	dpl.Dir = workspace.Path
	dpl.Stderr = os.Stderr
	dpl.Stdout = os.Stdout

	if err := dpl.Run(); err != nil {
		fmt.Println(err)

		os.Exit(1)
		return
	}
}

func buildDpl(workspace *drone.Workspace, repo *drone.Repo, build *drone.Build, vargs *Params) *exec.Cmd {
	args := []string{
		"--provider=rubygems",
	}

	if vargs.SkipCleanup {
		args = append(args, "--skip-cleanup")
	}

	if len(vargs.APIKey) > 0 {
		args = append(args, fmt.Sprintf(
			"--api-key=%s",
			vargs.APIKey))
	}

	if len(vargs.Username) > 0 {
		args = append(args, fmt.Sprintf(
			"--user=%s",
			vargs.Username))
	}

	if len(vargs.Password) > 0 {
		args = append(args, fmt.Sprintf(
			"--password=%s",
			vargs.Password))
	}

	if len(vargs.Host) > 0 {
		args = append(args, fmt.Sprintf(
			"--host=%s",
			vargs.Host))
	}

	if len(vargs.Name) > 0 {
		args = append(args, fmt.Sprintf(
			"--app=%s",
			vargs.Name))
	} else {
		args = append(args, fmt.Sprintf(
			"--app=%s",
			repo.Name))
	}

	if len(vargs.Gemspec) > 0 {
		args = append(args, fmt.Sprintf(
			"--gemspec=%s",
			vargs.Gemspec))
	} else {
		args = append(args, fmt.Sprintf(
			"--gemspec=%s",
			repo.Name))
	}

	return exec.Command("dpl", args...)
}
