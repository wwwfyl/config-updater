# config-updater
The "config-updater" program is a command-line utility written in Go that facilitates the manipulation of configuration files. It provides a user-friendly way to modify or add key-value pairs within a configuration file.

# Features:

1. Configuration File Handling: The program takes as input the path to a configuration file, a key, and a value. It reads the contents of the configuration file, allowing you to modify or add key-value pairs.
2. Non-Visible Character Validation: Before making any changes, the program checks whether the provided key and value contain non-visible or control characters. If any are found, it prevents the modification or addition and displays an error message.
3. Key-Value Manipulation: The program updates the key-value pair if the specified key already exists in the configuration file. If the key is absent, it adds the key-value pair to the configuration file.
4. Data Integrity: After updating the configuration, the program ensures that the configuration file maintains its structure and data integrity, preserving the original formatting and layout.
5. User-friendly: With a clear command-line interface and error handling, the "config-updater" program is user-friendly and suitable for both manual and batch configuration updates.
