package utils

import (
	"strings"
)

func ValidatePassword(str string) bool {
	//正序倒序不能一样
	if str == Reverse(str) {
		return false
	}

	//至少包含一个大写字符
	if !HasUpper(str) {
		return false
	}

	//至少包含一个数字
	if !HasNumber(str) {
		return false
	}

	//至少包含两个特殊字符 ~!@#$%^&*()_+
	if !HasSpecial(str, 2) {
		return false
	}

	//不能包含任何字符两次以上
	if MoreThan(str, 2) {
		return false
	}

	return true
}

func HasUpper(s string) (b bool) {
	for _, v := range s {
		if strings.ContainsRune("ABCDEFGHIJKLMNOPQRSTUVWXYZ", v) {
			b = true
		}
	}
	return b
}

func HasNumber(s string) (b bool) {
	for _, v := range s {
		if strings.ContainsRune("0123456789", v) {
			b = true
		}
	}
	return b
}

// TODO
func MoreThan(s string, n int) (b bool) {
	var m = make(map[string]int)

	for _, v := range s {
		if m[string(v)] > 0 {
			m[string(v)] += 1
		} else {
			m[string(v)] = 1
		}
	}

	for _, mv := range m {
		if mv > n {
			b = true
		}
	}
	return b
}

// ~!@#$%^&*()_+
func HasSpecial(s string, n int) (b bool) {
	count := 0

	for _, v := range s {
		if strings.ContainsRune("~!@#$%^&*()_+", v) {
			count++
		}
	}

	if count >= n {
		b = true
	}
	return b
}
