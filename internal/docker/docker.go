package docker

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func RunContainer(containerName, labelValue string) error {
	cmd := exec.Command("docker", "run", "-d", "--name", containerName, "--label", "value="+labelValue, "alpine", "sh", "-c", "sleep 9999999")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could not run container: %v: %s", err, stderr.String())
	}
	return nil
}

func GetLabelValue(containerName string) (string, error) {
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

func DeleteContainer(containerName string) error {
	cmdStop := exec.Command("docker", "stop", containerName)
	var stderrStop bytes.Buffer
	cmdStop.Stderr = &stderrStop
	if err := cmdStop.Run(); err != nil {
		fmt.Printf("Could not stop container (it might be already stopped): %v: %s\n", err, stderrStop.String())
	}

	cmdRm := exec.Command("docker", "rm", containerName)
	var stderrRm bytes.Buffer
	cmdRm.Stderr = &stderrRm
	if err := cmdRm.Run(); err != nil {
		return fmt.Errorf("could not remove container: %v: %s", err, stderrRm.String())
	}

	return nil
}
