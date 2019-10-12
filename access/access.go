package access

import "fmt"

//可以被所有人访问 类似Java的public权限
func OutPrint(msg string) {
	innerPrint(msg)
}

//只能被包内访问 类似于Java的default权限
func innerPrint(msg string) {
	fmt.Print(msg)
}
