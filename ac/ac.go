/*
 * @Description:ac算法
 * @Version: 1.0
 * @Author: zoetu 1753095374@qq.com
 * @Date: 2022-05-26 15:22:12
 * @LastEditTime: 2022-05-27 14:43:23
 */
package ac

import (
	"container/list"
	"fmt"
)

type trieNode struct {
	end   int                //模式串结尾标志
	fail  *trieNode          //fail跳转
	child map[rune]*trieNode //goto结点
	index int                //模式串序号
}

// 添加新的结点
func newTrieNode() *trieNode {
	return &trieNode{
		end:   0,
		fail:  nil,
		child: make(map[rune]*trieNode),
		index: -1,
	}
}

//自动机结点结构
type AhoCorasick struct {
	root  *trieNode //根节点
	size  int       //模式串数量
	mark  []bool    //已读标志
	count int       //节点数量
}

// 根据dic string构造AC
func New(dic []string) *AhoCorasick {
	ac := NewAhoCorasick()
	ac.Build(dic)
	return ac
}

//新建AC自动机根节点
func NewAhoCorasick() *AhoCorasick {
	return &AhoCorasick{
		root:  newTrieNode(),
		size:  0,
		mark:  make([]bool, 0),
		count: 0,
	}
}

// 1. AhoCorasick自动机初始化
func (ac *AhoCorasick) Build(dictionary []string) {

	for i, _ := range dictionary {
		ac.Insert(dictionary[i]) //goto
	}
	ac.Build_fail() //fail
	ac.mark = make([]bool, ac.size)
}

//2. goto
func (ac *AhoCorasick) Insert(s string) {
	curNode := ac.root
	for _, v := range s {
		if curNode.child[v] == nil { //自动机没有该字符边
			curNode.child[v] = newTrieNode()
			ac.count++ //自动机结点数+1
		}
		curNode = curNode.child[v] //当前节点往下走
	}
	curNode.end = len(s)    //字符串结尾标志设为串长度
	curNode.index = ac.size //记录当前字符串是第几串
	ac.size++
}

//3. failture
func (ac *AhoCorasick) Build_fail() { //fail构造 --BFS
	ll := list.New()
	ll.PushBack(ac.root)
	for ll.Len() > 0 {
		temp := ll.Remove(ll.Front()).(*trieNode)
		var p *trieNode = nil

		for i, v := range temp.child {
			if temp == ac.root {
				v.fail = ac.root
			} else {
				p = temp.fail
				for p != nil {
					if p.child[i] != nil {
						v.fail = p.child[i]
						break
					}
					p = p.fail
				}
				if p == nil {
					v.fail = ac.root
				}
			}
			ll.PushBack(v)
		}
	}
}

//4. Match  返回模式串序号及匹配文本末位置
func (ac *AhoCorasick) Match(s string) (ret []int, pos []int) {
	curNode := ac.root
	ac.resetMark()
	var p *trieNode = nil

	for key, v := range s {
		for curNode.child[v] == nil && curNode != ac.root {
			curNode = curNode.fail
		}
		curNode = curNode.child[v]
		if curNode == nil {
			curNode = ac.root
		}

		p = curNode
		for p != ac.root && p.end > 0 && !ac.mark[p.index] { //该结点不是根节点，标志为1即为字符串结尾，
			// ac.mark[p.index] = true  //标记已匹配，重复的不进入
			// for i := 0; i < p.end; i++ {//如果规则串重复则 end>1
			ret = append(ret, p.index)     //匹配到的模式串序号
			pos = append(pos, key-p.end+2) //文本匹配到的结尾字符
			// }
			p = p.fail
		}
	}
	return ret, pos
}

//resetMark 匹配标志位重置
func (ac *AhoCorasick) resetMark() { //标志位重置
	for i := 0; i < ac.size; i++ {
		ac.mark[i] = false
	}
}

//ACMatch   匹配输出 -- 匹配第几串模式串，共匹配多少串，在文本起始位置
var total int //总匹配数
func (ac *AhoCorasick) ACMatch(s string, dictionary []string, num int) {
	ret, pos := ac.Match(s) //查找字符串s
	for key, i := range ret {
		total++
		fmt.Println("匹配第", num, "个文件第", i+1, "串", dictionary[i], " 在文本的起始位置=", pos[key])
	}
}
