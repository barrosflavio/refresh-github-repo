package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execCmd(cmd string) (string, error) {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func main() {

	var commitMessage string

	fmt.Print("Mensagem de commit: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		commitMessage = scanner.Text()
	}
	commitMessage = strings.TrimSpace(commitMessage)

	fmt.Println(execCmd("git add ."))
	fmt.Println(execCmd(fmt.Sprintf("git commit -m \"%s\"", commitMessage)))
	fmt.Println(execCmd("git push -u origin main"))
}
