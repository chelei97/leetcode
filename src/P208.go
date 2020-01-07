package main

import "fmt"

func main() {
	trie := Constructor1()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple") )  // 返回 true
	fmt.Println(trie.Search("app"))     // 返回 false
	fmt.Println(trie.StartsWith("app")) // 返回 true
	trie.Insert("app")
	fmt.Println(trie.Search("app"))     // 返回 true
}

type Trie struct {
	val string
	record map[string]*Trie
	isEnd bool
}


/** Initialize your data structure here. */
func Constructor1() Trie {
	return Trie{record: make(map[string]*Trie, 0)}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	for i := 0; i < len(word); i ++ {
		if v, ok := this.record[word[i : i + 1]]; ok {
			this = v
		}else {
			tmp := Constructor1()
			tmp.val = word[i : i + 1]
			this.record[word[i : i + 1]] = &tmp
			this = &tmp
		}
	}
	this.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for i := 0; i < len(word); i ++ {
		if v, ok := this.record[word[i : i + 1]]; ok {
			this = v
		}else {
			return false
		}
	}
	if this.isEnd {
		return true
	}
	return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i ++ {
		if v, ok := this.record[prefix[i : i + 1]]; ok {
			this = v
		}else {
			return false
		}
	}
	if !this.isEnd {
		return true
	}
	return false
}
