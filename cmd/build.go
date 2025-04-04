/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var agentImage string
var workingDirectory string
var imageTag string

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build your agent",
	Long: `Build your agent

Example:
    hive-cli build -d "/home/user/my-agent/src" -t "my-agent:latest"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return buildAgent(workingDirectory, imageTag)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	buildCmd.PersistentFlags().StringVarP(&workingDirectory, "dir", "d", "", "directory with agent code")
	buildCmd.PersistentFlags().StringVarP(&imageTag, "tag", "t", "", "image tag")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//go:embed Dockerfile
var d []byte

func buildAgent(dir, tag string) error {

	// 1. Validate that docker is on system
	if err := checkDocker(); err != nil {
		return fmt.Errorf("%w", err)
	}

	// 2. Validate directory exists (or use PWD if no value)
	if dir == "" {
		dir, _ = os.Getwd()
	}
	if _, err := os.Stat(dir); err != nil {
		return fmt.Errorf("%w", err)
	}

	// 3. Validate requirements.txt and run_agent.py files exist in dir
	if _, err := os.Stat(dir + "/requirements.txt"); err != nil {
		return fmt.Errorf("Missing file: requirements.txt required")
	}
	if _, err := os.Stat(dir + "/run_agent.py"); err != nil {
		return fmt.Errorf("Missing file: run_agent.py required")
	}

	// 4. Exec docker build command
	// create tempdir for Dockerfile
	dockerFile, _ := os.CreateTemp("", "Dockerfile")
	defer dockerFile.Close()
	if _, err := dockerFile.Write(d); err != nil {
		return fmt.Errorf("%w", err)
	}

	buildContainer := exec.Command("docker", "build", "-t", tag, "-f", dockerFile.Name(), dir)
	if err := runCommandWithOutput(buildContainer); err != nil {
		return fmt.Errorf("Error building agent: %w", err)
	}

	return nil
}

// checkDocker checks that Docker is running on the users local system.
func checkDocker() error {
	dockerCheck := exec.Command("docker", "stats", "--no-stream")
	if err := dockerCheck.Run(); err != nil {
		return fmt.Errorf("docker not running")
	}
	return nil
}

func runCommandWithOutput(c *exec.Cmd) error {
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return fmt.Errorf("piping output: %w", err)
	}
	fmt.Print("\n")
	return nil
}
