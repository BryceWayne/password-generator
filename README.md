# Go CLI Password Generator with Restrictions

## Overview
A command-line interface (CLI) tool for generating secure random passwords with specific restrictions, written in Go. The generated passwords will not contain three consecutive characters of the same type (uppercase, lowercase, digit, or symbol).

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/BryceWayne/password-generator.git>
   ```
2. Navigate into the project directory:
   ```
   cd ./password-generator
   ```
3. Build the project:
   ```
   go build
   ```

## Usage

Run the program using:

```bash
./password-generator -length=16
```

### Options

- `-length`: The length of the generated password (default is 12).

## Example

```bash
$ ./password-generator -length=16
Generated password: Aa1!Bb2@Cc3#
```

## Restrictions

The generated passwords will not contain three consecutive characters of the same type:

- Lowercase: `a-z`
- Uppercase: `A-Z`
- Digit: `0-9`
- Symbol: `!@#$%^&*()-_=+[]{}|;:,.<>?`

## Dependencies

- Go v1.x+
- Crypto library (built-in)
- Math/Big library (built-in)

## License

MIT