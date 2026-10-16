package main

import (
	"fmt"
	"os"
	"strconv"
)

func digitValue(c rune) int {
	if c >= '0' && c <= '9' {
		return int(c - '0')
	}
	if c >= 'A' && c <= 'F' {
		return int(c-'A') + 10
	}
	if c >= 'a' && c <= 'f' {
		return int(c-'a') + 10
	}
	return -1
}

func FromBase(s string, base int) (int, bool) {
	if base < 2 || base > 16 || s == "" {
		return 0, false
	}

	n := 0
	for _, c := range s {
		d := digitValue(c)
		if d < 0 || d >= base {
			return 0, false
		}
		n = n*base + d
	}
	return n, true
}

func ToBase(n, base int) (string, bool) {
	if base < 2 || base > 16 || n < 0 {
		return "", false
	}
	if n == 0 {
		return "0", true
	}

	const digits = "0123456789ABCDEF"
	result := ""

	for n > 0 {
		remainder := n % base
		result = string(digits[remainder]) + result
		n /= base
	}
	return result, true
}

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr,
			"Usage: baseconv <number> <from-base> <to-base>")
		os.Exit(1)
	}

	number := os.Args[1]

	fromBase, err := strconv.Atoi(os.Args[2])
	if err != nil || fromBase < 2 || fromBase > 16 {
		fmt.Fprintln(os.Stderr,
			"Error: source base must be between 2 and 16")
		os.Exit(1)
	}

	toBase, err := strconv.Atoi(os.Args[3])
	if err != nil || toBase < 2 || toBase > 16 {
		fmt.Fprintln(os.Stderr,
			"Error: target base must be between 2 and 16")
		os.Exit(1)
	}

	decimal, ok := FromBase(number, fromBase)
	if !ok {
		fmt.Fprintln(os.Stderr,
			"Error: invalid number for the source base")
		os.Exit(1)
	}

	result, ok := ToBase(decimal, toBase)
	if !ok {
		fmt.Fprintln(os.Stderr, "Error: conversion failed")
		os.Exit(1)
	}

	fmt.Println(result)
}
