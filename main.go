package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const envFile = ".env"

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	baseCmd := os.Args[1:]
	if len(baseCmd) < 1 {
		panic("cmd is required")
	}
	cmd := exec.Command(baseCmd[0], baseCmd[1:]...)
	if _, err := os.Stat(envFile); err == nil {
		envConfig, err := ioutil.ReadFile(envFile)
		panicErr(err)
		envs := strings.Split(string(envConfig), "\n")
		env := os.Environ()
		env = append(env, envs...)
		cmd.Env = env
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
