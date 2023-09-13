package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

    flag "github.com/spf13/pflag"
)

func main() {
    configFile := flag.StringP("config", "c", "", "Path to the config file")
	key := flag.StringP("key", "k", "", "Key to update or add")
	value := flag.StringP("value", "v", "", "Value to set for the key")
	flag.Parse()

	if *configFile == "" || *key == "" || *value == "" {
		fmt.Println("Usage: config-updater --config/-c <config file> --key/-k <key> --value/-v <value>")
		return
	}

	// Read the config file
	file, err := os.OpenFile(*configFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening config file: %v\n", err)
		return
	}
	defer file.Close()

    scanner := bufio.NewScanner(file)
    configLines := []string{}

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.SplitN(line, "=", 2)
        if len(parts) == 2 {
            configKey := strings.TrimSpace(parts[0])
            configValue := strings.TrimSpace(parts[1])
            if configKey == *key {
                // Check if the value has quotes, and if so, preserve them
                if strings.HasPrefix(*value, `"`) && strings.HasSuffix(*value, `"`) {
                    configValue = fmt.Sprintf(`"%s"`, *value)
                } else {
                    configValue = *value
                }
            }
            configLines = append(configLines, fmt.Sprintf("%s=%s", configKey, configValue))
        } else {
            // Preserve lines that don't have the key-value format
            configLines = append(configLines, line)
        }
    }

    // Write the updated config back to the file
    file.Truncate(0)
    file.Seek(0, 0)

    for _, line := range configLines {
        fmt.Fprintln(file, line)
    }

    fmt.Printf("Config file updated: %s=%s\n", *key, *value)
}
