package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {

	dockerfileContent := `
		FROM base_image1
		
		COPY . /app
		WORKDIR /app

		CMD ["python", "train1.py"]
	`

	err := ioutil.WriteFile("Dockerfile", []byte(dockerfileContent), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("**************Dockerfile created****************")

	runCommand("docker", "build", "-t", "t1", ".")
	fmt.Println("**************Image created****************")
	runCommandWithStdIO("docker", "run", "-it", "--name", "t1", "t1")
	fmt.Println("**************Container created and Running****************")
	runCommand("docker", "rm", "t1")
	runCommand("docker", "rmi", "t1")
	fmt.Println("**************Container and Image Deleted****************")
}
func runCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}
	fmt.Printf("Command output:\n%s\n", output)
}
func runCommandWithStdIO(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}
}
