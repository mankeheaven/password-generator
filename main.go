package main

import (
	"fmt"
	"time"

	"password-generator/utils"
)

func main() {

	// 最小长度至少9位
	// TODO 最少包含2位特殊字符
	// 以字母开头
	// 至少包含一个大写字符
	// 至少包含一个数字
	// 顺读倒读不能一致
	// TODO 不能包含任何字符两次以上
	// 至少包含1个小写字母(逻辑替换第一个为小写)

	//a-z 26位，~!@#$%^&*()_+  13位， 总长26+26+10+13 = 75

	password, _ := utils.GenaratePassword()

	fmt.Printf("密码长度%v, 密码为%v", len(password), password)

	time.Sleep(1000000000000000)
}
