# Go CLI Password Generator

## Overview
A minimal command-line interface (CLI) tool for generating secure random passwords, written in Go.

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/BryceWayne/password-generator.git
   ```
2. Navigate into the project directory:
   ```
   cd path/to/directory
   ```
3. Build the project:
   ```
   go build
   ```

## Usage

Run the program using:

```bash
./password -length=16
```

### Options

- `-length`: The length of the generated password (default is 12).

## Example

```bash
$ ./password -length=16
Generated password: U3VwZXJTZWN1cmU=
```

## Dependencies

- Go v1.x+
- Crypto library (built-in)

## License

MIT
