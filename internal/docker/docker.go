/*
Copyright © 2025 Elwan Mayencourt <mayencourt@elwan.ch>
*/
package docker

import (
	"bytes"
	"fmt"
	"kvd/internal/utils"
	"log"
	"os/exec"
	"strings"
)

const maxContainerNameLength = 63
const prunePrefix = "prune-"
const pruneSuffixLength = maxContainerNameLength - len(prunePrefix)

func RunContainer(containerName, labelValue string, retries int) error {
	cmd := exec.Command("docker", "run", "-d", "--name", containerName, "--label", "value="+labelValue, "alpine", "sh", "-c", "sleep 9999999")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err == nil {
		return nil
	}

	if retries > 0 && strings.Contains(stderr.String(), "Conflict. The container name") {

		toPruneContainerName := prunePrefix + utils.GenerateRandomString(pruneSuffixLength)

		if renameErr := RenameContainer(containerName, toPruneContainerName); renameErr != nil {
			return renameErr
		}

		return RunContainer(containerName, labelValue, retries-1)
	}

	return fmt.Errorf("could not run container: %v: %s", err, stderr.String())
}

func DeleteContainer(containerName string, retries int) error {
	toPruneContainerName := prunePrefix + utils.GenerateRandomString(pruneSuffixLength)

	if renameErr := RenameContainer(containerName, toPruneContainerName); renameErr != nil {
		if retries > 0 {
			return DeleteContainer(containerName, retries-1)
		}
		return renameErr
	}

	return nil
}

func ForceDeleteContainer(containerName string) error {
	cmdStop := exec.Command("docker", "stop", containerName)
	var stderrStop bytes.Buffer
	cmdStop.Stderr = &stderrStop
	if err := cmdStop.Run(); err != nil {
		log.Printf("Could not stop container (it might be already stopped): %v: %s", err, stderrStop.String())
	}

	cmdRm := exec.Command("docker", "rm", containerName)
	var stderrRm bytes.Buffer
	cmdRm.Stderr = &stderrRm
	if err := cmdRm.Run(); err != nil {
		return fmt.Errorf("could not remove container: %v: %s", err, stderrRm.String())
	}

	return nil
}

func GetContainerLabelValue(containerName string) (string, error) {
	cmd := exec.Command("docker", "inspect", "--format", `'{{index .Config.Labels "value"}}'`, containerName)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("could not inspect container: %v: %s", err, stderr.String())
	}
	return strings.Trim(strings.TrimSpace(out.String()), "'"), nil
}

func RenameContainer(oldName, newName string) error {
	cmd := exec.Command("docker", "rename", oldName, newName)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could not rename container: %v: %s", err, stderr.String())
	}
	return nil
}

func ListPruneContainers() ([]string, error) {
	cmd := exec.Command("docker", "ps", "-a", "--filter", "name=^"+prunePrefix, "--format", "{{.Names}}")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("could not list prune containers: %v: %s", err, stderr.String())
	}

	containerNames := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(containerNames) == 1 && containerNames[0] == "" {
		return []string{}, nil
	}

	return containerNames, nil
}
