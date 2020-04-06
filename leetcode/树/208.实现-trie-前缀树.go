package main

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start

//前缀树的优点1.能够找到同一个前缀的所有单词
//2.前缀树节省空间，同一个前缀的单词公用一个空间

type Trie struct {
	Root *Node
}

type Node struct {
	IsEnd     bool
	Childrens map[byte]*Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	root := &Node{
		IsEnd:     false,
		Childrens: make(map[byte]*Node),
	}
	trie := Trie{Root: root}
	return trie
}

//遍历待插入的字符，如果存在该字符的节点，则跳过，如果不存在，则新建节点继续遍历
//最终给最后一个字符打上结束的标记
func (this *Trie) Insert(word string) {
	if word == "" {
		return
	}

	root := this.Root
	for i := 0; i < len(word); i++ {
		if _, ok := root.Childrens[word[i]]; ok {
			root = root.Childrens[word[i]]
		} else {
			root.Childrens[word[i]] = new(Node)
			root.Childrens[word[i]].Childrens = make(map[byte]*Node)
			root = root.Childrens[word[i]]
		}

	}
	root.IsEnd = true
}

//1.查找成功：word为空且isEnd=true
//2.查找失败：word不为空但是儿子节点为空了
//3.查找失败，但是是前缀：word为空但是isEnd=false
func (this *Trie) Search(word string) bool {
	root := this.Root
	for i := 0; i < len(word); i++ {
		if _, ok := root.Childrens[word[i]]; ok {
			root = root.Childrens[word[i]]
		} else {
			return false
		}
	}
	return root.IsEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this.Root
	for i := 0; i < len(prefix); i++ {
		if _, ok := root.Childrens[prefix[i]]; ok {
			root = root.Childrens[prefix[i]]
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
