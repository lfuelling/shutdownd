package main

import (
	"errors"
	"os/exec"
	"strings"
)

func shutdownSystem(command string, args ...string) (string, error) {
	if err := exec.Command(command, args...).Run(); err != nil {
		return "Failed to trigger shutdown!", err
	} else {
		return "OK", nil
	}
}

func handleShutdown(config Config) (string, error) {
	osType := strings.ToLower(config.osType)
	switch osType {
	case "linux":
	case "bsd":
		if config.useSudo {
			return shutdownSystem("sudo", "shutdown", "-h", "now")
		} else {
			return shutdownSystem("shutdown", "-h", "now")
		}
	case "windows":
		return shutdownSystem("cmd", "/C", "shutdown", "/t", "0", "/s")
	}
	return "Configuration Error!", errors.New("Invalid osType: '" + osType + "'!")
}
