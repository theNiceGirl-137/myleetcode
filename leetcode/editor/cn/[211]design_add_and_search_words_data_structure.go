//请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。 
//
// 实现词典类 WordDictionary ： 
//
// 
// WordDictionary() 初始化词典对象 
// void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配 
// bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回 false 。word 中可能包含一些 
//'.' ，每个 . 都可以表示任何一个字母。 
// 
//
// 
//
// 示例： 
//
// 
//输入：
//["WordDictionary","addWord","addWord","addWord","search","search","search",
//"search"]
//[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
//输出：
//[null,null,null,null,false,true,true,true]
//
//解释：
//WordDictionary wordDictionary = new WordDictionary();
//wordDictionary.addWord("bad");
//wordDictionary.addWord("dad");
//wordDictionary.addWord("mad");
//wordDictionary.search("pad"); // return False
//wordDictionary.search("bad"); // return True
//wordDictionary.search(".ad"); // return True
//wordDictionary.search("b.."); // return True
// 
//
// 
//
// 提示： 
//
// 
// 1 <= word.length <= 500 
// addWord 中的 word 由小写英文字母组成 
// search 中的 word 由 '.' 或小写英文字母组成 
// 最多调用 50000 次 addWord 和 search 
// 
// Related Topics 深度优先搜索 设计 字典树 字符串 👍 361 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
type WordDictionary struct {
	tireRoot *tireNode
}


func Constructor211() WordDictionary {
	return WordDictionary{&tireNode{}}
}


func (this *WordDictionary) AddWord(word string)  {
	this.tireRoot.Insert(word)
}


func (this *WordDictionary) Search(word string) bool {
	return this.Dfs(0, word, this.tireRoot)
}

func (this *WordDictionary) Dfs(index int, word string, node *tireNode) bool {
	if index == len(word) {
		return node.isEnd
	}
	ch := word[index]
	if ch != '.' {
		child := node.children[ch-'a']
		if child != nil && this.Dfs(index+1, word, child) {
			return true
		}
	} else {
		for i :=  0; i < len(node.children); i++ {
			child := node.children[i]
			if child != nil && this.Dfs(index+1, word, child) {
				return true
			}
		}
	}
	return false
}

type tireNode struct {
	children [26]*tireNode
	isEnd bool
}

func (this *tireNode) Insert(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &tireNode{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
//leetcode submit region end(Prohibit modification and deletion)
