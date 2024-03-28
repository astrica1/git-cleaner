package utils

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

func Execute(command string, flags ...string) (string, error) {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/c", command)
		log.Println("It is better to run this application on Linux based systems or containers")
	} else {
		cmd = exec.Command("/bin/bash", "-c", command)
	}

	cmd.Args = append(cmd.Args, flags...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		log.Printf("Couldn't Run Command << %s >>\n", command)
	}

	return string(output), err
}

func Clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal("Couldn't Clear Terminal: ", err)
	}
}
