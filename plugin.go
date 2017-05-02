package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type (
	Config struct {
		ApiKey      string
		Gemspec     string
		Host        string
		Gemname     string
		Password    string
		Username    string
		SkipCleanup bool
	}

	Repo struct {
		Name    string
	}

	Plugin struct {
		Repo   Repo
		Config Config
	}
)

func (p Plugin) Exec() error {
	if (p.Config.Username == "" || p.Config.Password == "") && p.Config.ApiKey == "" {
		return errors.New("Please provide an API key or username/password credentials")
	}

	dpl := buildDpl(p)
	var err error

	dpl.Dir, err = os.Getwd()
	if err != nil {
		return err
	}
	dpl.Stderr = os.Stderr
	dpl.Stdout = os.Stdout

	return dpl.Run()
}

func buildDpl(p Plugin) *exec.Cmd {
	args := []string{
		"--provider=rubygems",
	}

	if p.Config.SkipCleanup {
		args = append(args, "--skip-cleanup")
	}

	if p.Config.ApiKey != "" {
		args = append(args, fmt.Sprintf(
			"--api-key=%s",
			p.Config.ApiKey))
	}

	if p.Config.Username != "" {
		args = append(args, fmt.Sprintf(
			"--user=%s",
			p.Config.Username))
	}

	if p.Config.Password != "" {
		args = append(args, fmt.Sprintf(
			"--password=%s",
			p.Config.Password))
	}

	if p.Config.Host != "" {
		args = append(args, fmt.Sprintf(
			"--host=%s",
			p.Config.Host))
	}

	if p.Config.Gemname != "" {
		args = append(args, fmt.Sprintf(
			"--app=%s",
			p.Config.Gemname))
	} else {
		args = append(args, fmt.Sprintf(
			"--app=%s",
			p.Repo.Name))
	}

	if p.Config.Gemspec != "" {
		args = append(args, fmt.Sprintf(
			"--gemspec=%s",
			p.Config.Gemspec))
	} else {
		args = append(args, fmt.Sprintf(
			"--gemspec=%s",
			p.Repo.Name))
	}

	return exec.Command("dpl", args...)
}
