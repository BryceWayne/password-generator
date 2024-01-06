# Password Generator

## Overview
This command-line tool, implemented in Go and contained within `passgen.go`, generates secure, random passwords. It ensures that no more than two consecutive characters are the same type (uppercase, lowercase, digit, or symbol).

## Installation
Requires Go v1.x+. Clone the repository and build the project.

## Usage
`go run passgen.go -length [desired length]`

Specify the password length using the `-length` flag. The default length is 16 characters.

## Example
`go run passgen.go -length 20`

Generates a 20-character password.

## How it Works
`passgen.go` uses cryptographic randomness and character type checking to generate passwords according to the specified length and restrictions.

## Dependencies
- Go standard library packages: `crypto/rand`, `flag`, `fmt`, `math/big`, `unicode`.

## License
MIT License
