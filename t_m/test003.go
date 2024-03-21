package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func getMachineUniqueInfo() (map[string]string, error) {
	info := make(map[string]string)

	// Get MAC address
	macAddress, err := getMACAddress()
	if err == nil {
		info["MACAddress"] = macAddress
	}

	// Get Hard Drive Serial Number
	hardDriveSerialNumber, err := getHardDriveSerialNumber()
	if err == nil {
		info["HardDriveSerialNumber"] = hardDriveSerialNumber
	}



	return info, nil
}

func getMACAddress() (string, error) {
	cmd := exec.Command("ifconfig")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(output), "\n")
	var macAddress string
	for _, line := range lines {
		if strings.Contains(line, "ether ") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				macAddress = parts[1]
				break
			}
		}
	}
	return macAddress, nil
}

func getHardDriveSerialNumber() (string, error) {
	cmd := exec.Command("lsblk", "--output", "SERIAL", "--noheadings", "--list")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	serialNumber := strings.TrimSpace(string(output))
	return serialNumber, nil
}

func main() {
	machineInfo, err := getMachineUniqueInfo()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for key, value := range machineInfo {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}
