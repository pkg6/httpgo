package t

import (
	"strings"
)

// StringsReplaceMap
//您手机注册的验证码为：【变量1】，如有问题请拨打客服电话：【变量2】
//repl = map[string]string{"变量1"："111"，"变量2"::"222"}
//您手机注册的验证码为：111，如有问题请拨打客服电话：222
func StringsReplaceMap(s string, repl map[string]string) string {
	var r []string
	for k, v := range repl {
		r = append(r, k, v)
	}
	return strings.NewReplacer(r...).Replace(s)
}
