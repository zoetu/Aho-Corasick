/*
 * @Description: 主程序
 * @Version: 1.0
 * @Author: zoetu 1753095374@qq.com
 * @Date: 2022-05-23 22:27:00
 * @LastEditTime: 2022-05-27 14:49:11
 */
//file split
// package main

// import (
// 	"PatternCourse/FileSplit"
// )

// func main() {
// 	fileToBeChunked := "./data/rule/url_data.txt"
// 	// FileSplit.Split_xMB(fileToBeChunked, 5)
// 	FileSplit.Split_xLINE(fileToBeChunked, 1000)
// }

// string match
package main

import (
	ac "PatternCourse/ac"
)

func main() {
	dictionary := []string{"he", "hello", "hers"}
	newac := ac.New(dictionary)
	newac.ACMatch("hershe", dictionary, 0)

}
