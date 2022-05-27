/*
 * @Description: ac测试
 * @Version: 1.0
 * @Author: zoetu 1753095374@qq.com
 * @Date: 2022-05-25 10:33:19
 * @LastEditTime: 2022-05-27 14:42:50
 */
package ac_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	ac "PatternCourse/ac"
)

var benchmarks = []struct {
	filetext  string //不同规则集
	textSizes []int  //文件大小
	textLines []int  //规则数
}{

	{
		"url",
		[]int{1, 5, 10, 50}, // 不足50MB那为文件最大内存
		[]int{1000, 5000, 10000, 100000, 1000000}, //规则数
	},
}

func Benchmark(b *testing.B) {
	for _, bb := range benchmarks {
		b.Run(bb.filetext, func(b *testing.B) {
			var dict []string
			if d, err := readDictionary(filepath.Join("../data/rule", fmt.Sprintf("%s_2w.txt", bb.filetext))); err != nil {
				b.Fatal(err)
			} else {
				dict = bytes2Strings(d) //读取文件为字符串数组
			}

			b.Run("bulding", func(b *testing.B) {
				b.ResetTimer() //计时器清零 build test
				for n := 0; n < b.N; n++ {
					ac.New(dict) //构造AC
				}
			})
			// 实验1：不同文件大小的搜索时间记录
			for _, size := range bb.textSizes {
				size := size
				b.Run(fmt.Sprintf("searching in %d MB text", size), func(b *testing.B) {
					trie := ac.New(dict) //大文本构造ac
					var txt string
					if t, err := readTextofSize(bb.filetext, size); err != nil {
						b.Fatal(err)
					} else {
						txt = string(t)
					}
					b.ResetTimer() //search test
					for n := 0; n < b.N; n++ {
						trie.Match(txt) //小文本搜索
					}
				})
			}
			// 实验2：不同规则数目的搜索时间
			for _, line := range bb.textLines {
				line := line
				b.Run(fmt.Sprintf("searching in %d lines text", line), func(b *testing.B) {
					trie := ac.New(dict) //大文本构造ac
					var txt string
					if t, err := readTextofLine(bb.filetext, line); err != nil {
						b.Fatal(err)
					} else {
						txt = string(t)
					}
					b.ResetTimer() //search test
					for n := 0; n < b.N; n++ {
						trie.Match(txt) //小文本搜索
					}
				})
			}

		})
	}
}

// 读取规则文件，并返回[][]byte
func readDictionary(filename string) ([][]byte, error) {
	var dict [][]byte

	f, err := os.OpenFile(filename, os.O_RDONLY, 0660)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		l, err := r.ReadBytes('\n')
		if err != nil || err == io.EOF {
			break
		}
		l = bytes.TrimSpace(l)
		dict = append(dict, l)
	}

	return dict, nil
}

// 读取不同文件大小的测试文件
func readTextofSize(filetext string, size int) ([]byte, error) {
	f := filepath.Join("../data/test_size", fmt.Sprintf("%s_data_%dMB.txt", filetext, size))
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("read text from %q: %s", f, err)
	}
	return b, nil
}

// 读取不同规则数的测试文件
func readTextofLine(filetext string, line int) ([]byte, error) {
	f := filepath.Join("../data/test_line", fmt.Sprintf("%s_data_%dline.txt", filetext, line))
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("read text from %q: %s", f, err)
	}
	return b, nil
}

// [][]byte转换成[]string切片
func bytes2Strings(b [][]byte) []string {
	var ss = make([]string, 0, len(b))
	for _, x := range b {
		ss = append(ss, string(x))
	}
	return ss
}
