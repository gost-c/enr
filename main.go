package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var configFile string
	var help bool

	flag.StringVar(&configFile, "config", ".env", "custom config file")
	flag.StringVar(&configFile, "c", ".env", "custom config file (short)")

	flag.BoolVar(&help, "help", false, "show help")
	flag.BoolVar(&help, "h", false, "show help(short)")

	if help {
		showHelp()
	}

	flag.Parse()
	baseCmd := flag.Args()
	if len(baseCmd) < 1 {
		showHelp()
	}
	cmd := exec.Command(baseCmd[0], baseCmd[1:]...)
	if _, err := os.Stat(configFile); err == nil {
		envConfig, err := ioutil.ReadFile(configFile)
		if err != nil {
			log.Fatal(err)
		}
		envs := strings.Split(string(envConfig), "\n")
		env := os.Environ()
		env = append(env, envs...)
		cmd.Env = env
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		os.Exit(0)
	}
	log.Fatal("enr config file not exists. create .env file or use -c=configfile option")
}

func showHelp() {
	fmt.Println(`

Enr is a command line env wrapper. Can run behind cmds with envs of config file.

Usage:

	enr [options] <your cmds>

Options:

	-config="yourConfigFile", -c="yourConfigFile" 	Your custom config file, default is '.env'
	-help, -h                                       Show help

Example:

	enr -c=".env.example" bash test.sh
`)
	os.Exit(0)
}
