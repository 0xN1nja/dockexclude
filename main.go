package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"

	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v3"
)

type Service struct {
	ContainerName string `yaml:"container_name,omitempty"`
}

type DockerComposeConfig struct {
	Services map[string]Service `yaml:"services"`
}

func runDockerCompose(dockerCommand string, excludedContainers []string) {
	composeFilePath := "docker-compose.yaml"

	fileContent, err := os.ReadFile(composeFilePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var composeConfig DockerComposeConfig
	err = yaml.Unmarshal(fileContent, &composeConfig)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	commandArgs := []string{"compose"}
	if dockerCommand == "up" {
		commandArgs = append(commandArgs, "up", "-d")
	} else {
		commandArgs = append(commandArgs, dockerCommand)
	}

	for serviceName := range composeConfig.Services {
		if slices.Contains(excludedContainers, serviceName) {
			commandArgs = append(commandArgs, serviceName)
		}
	}

	dockerCmd := exec.Command("docker", commandArgs...)
	commandOutput, err := dockerCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing docker command: %v\nOutput: %s", err, string(commandOutput))
	}

	fmt.Println(string(commandOutput))

	_, err = yaml.Marshal(&composeConfig)
	if err != nil {
		log.Fatalf("Error marshalling back to YAML: %v", err)
	}
}

func main() {
	excludeFlag := &cli.StringSliceFlag{
		Name:    "exclude",
		Aliases: []string{"e"},
		Usage:   "services to exclude",
	}

	cmd := &cli.Command{
		Name:        "Dockexclude",
		Usage:       "Manage your Docker Compose stack with the ability to exclude services.",
		UsageText:   "dockexclude <command> [--exclude <service>...]",
		Description: "Manage your Docker Compose stack with the ability to exclude services.",
		Commands: []*cli.Command{
			{
				Name:      "up",
				Usage:     "Start and create containers",
				UsageText: "dockexclude up [--exclude <service>...]",
				Flags:     []cli.Flag{excludeFlag},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					excludedServices := cmd.StringSlice("exclude")
					runDockerCompose("up", excludedServices)
					return nil
				},
			},
			{
				Name:      "start",
				Usage:     "Start existing containers",
				UsageText: "dockexclude start [--exclude <service>...]",
				Flags:     []cli.Flag{excludeFlag},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					excludedServices := cmd.StringSlice("exclude")
					runDockerCompose("start", excludedServices)
					return nil
				},
			},
			{
				Name:      "down",
				Usage:     "Stop and remove containers",
				UsageText: "dockexclude down [--exclude <service>...]",
				Flags:     []cli.Flag{excludeFlag},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					excludedServices := cmd.StringSlice("exclude")
					runDockerCompose("down", excludedServices)
					return nil
				},
			},
			{
				Name:      "stop",
				Usage:     "Stop running containers",
				UsageText: "dockexclude stop [--exclude <service>...]",
				Flags:     []cli.Flag{excludeFlag},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					excludedServices := cmd.StringSlice("exclude")
					runDockerCompose("stop", excludedServices)
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
