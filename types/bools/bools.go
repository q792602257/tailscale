// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Package bools contains the [Int], [Compare], and [IfElse] functions.
package bools

// Int returns 1 for true and 0 for false.
func Int(v bool) int {
	if v {
		return 1
	} else {
		return 0
	}
}

// Compare compares two boolean values as if false is ordered before true.
func Compare[T ~bool](x, y T) int {
	switch {
	case x == false && y == true:
		return -1
	case x == true && y == false:
		return +1
	default:
		return 0
	}
}

// IfElse is a ternary operator that returns trueVal if condExpr is true
// otherwise it returns falseVal.
// IfElse(c, a, b) is roughly equivalent to (c ? a : b) in languages like C.
func IfElse[T any](condExpr bool, trueVal T, falseVal T) T {
	if condExpr {
		return trueVal
	} else {
		return falseVal
	}
}
