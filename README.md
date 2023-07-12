# IPster

IPster is a CLI application that stores IP addresses and their associated details in the terminal for quick access. It allows you to easily add, list, remove, and generate SSH commands for saved IP addresses.

## Installation

To install IPster, you can use `go install`. Make sure you have Go installed on your system before proceeding with the installation.

```shell
go install github.com/charlesdevelops/IPster@latest
```

## Usage

The basic usage of IPster is as follows:

```shell
ipster [command]
```

### Available Commands

- `add`: Add an IP address
- `completion`: Generate the autocompletion script for the specified shell
- `help`: Help about any command
- `ls`: List all saved addresses
- `rm`: Remove entries from the application
- `ssh`: Generates an SSH command to a saved IP address

You can use the `--help` flag with any command to get more information about it. For example:

```shell
ipster add --help
```

### Examples

- Add an IP address:

  ```shell
  ipster add <IP address> -d <description> -k <key location>
  ```

- List all saved addresses:

  ```shell
  ipster ls
  ```

- Remove entries from the application:

  ```shell
  ipster rm <IP address>
  ```

- Generate an SSH command for a saved IP address:

  ```shell
  ipster ssh <IP address>
  ```

### Dependencies

IPster depends on the following libraries:

- `sqlite3`
- `cobra` by spf13 (the framework used to build the CLI app)
- `github.com/atotto/clipboard` (for copying strings to the clipboard from the SSH command)
- `modernc.org/sqlite` (for the non-cgo SQLite3 driver)

## Acknowledgements

IPster was written in Go. Many thanks to the following projects and their contributors for their work:

- [sqlite3](https://www.sqlite.org/)
- [cobra by spf13](https://github.com/spf13/cobra)
- [github.com/atotto/clipboard](https://github.com/atotto/clipboard)
- [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)
- [Go](https://go.dev/)

## License

This project is licensed under the [MIT License](LICENSE).

---

Thank you for using IPster! If you encounter any issues or have suggestions for improvements, please open an issue on the [GitHub repository](https://github.com/charlesdevelops/IPster).
