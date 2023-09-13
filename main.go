package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) != 4 {
        fmt.Println("Usage: go run main.go <config file> <key> <value>")
        return
    }

    configFile := os.Args[1]
    key := os.Args[2]
    value := os.Args[3]

    // Read the config file
    file, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0644)
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
            if configKey == key {
                configValue = value
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

    fmt.Printf("Config file updated: %s=%s\n", key, value)
}
