package configs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Loadconfig() ([]string, map[string]string) {

	var profiles []string
	regions := map[string]string{}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home, err")
		return profiles, regions
	}

	configFile := filepath.Join(homeDir, ".aws", "config")

	file, err := os.Open(configFile)
	if err != nil {
		fmt.Println("Error getting config file", err)
		return profiles, regions
	}

	defer file.Close()

	current := ""

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {

			line = strings.Trim(line, "[]")
			line = strings.TrimPrefix(line, "profile ")
			line = strings.TrimSpace(line)

			profiles = append(profiles, line)
			current = line

			continue
		}

		if current == "" {
			continue
		}

		key, value, found := strings.Cut(line, "=")

		if found &&
			strings.TrimSpace(key) == "region" {

			regions[current] = strings.TrimSpace(value)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading config file", err)
	}

	return profiles, regions
}
