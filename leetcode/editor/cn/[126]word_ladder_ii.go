//按字典 wordList 完成从单词 beginWord 到单词 endWord 转化，一个表示此过程的 转换序列 是形式上像 beginWord -> s
//1 -> s2 -> ... -> sk 这样的单词序列，并满足： 
//
// 
// 
// 
// 每对相邻的单词之间仅有单个字母不同。 
// 转换过程中的每个单词 si（1 <= i <= k）必须是字典 wordList 中的单词。注意，beginWord 不必是字典 wordList 中的单
//词。 
// sk == endWord 
// 
//
// 给你两个单词 beginWord 和 endWord ，以及一个字典 wordList 。请你找出并返回所有从 beginWord 到 endWord 的
// 最短转换序列 ，如果不存在这样的转换序列，返回一个空列表。每个序列都应该以单词列表 [beginWord, s1, s2, ..., sk] 的形式返回。 
//
// 
//
// 示例 1： 
//
// 
//输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","lo
//g","cog"]
//输出：[["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
//解释：存在 2 种最短的转换序列：
//"hit" -> "hot" -> "dot" -> "dog" -> "cog"
//"hit" -> "hot" -> "lot" -> "log" -> "cog"
// 
//
// 示例 2： 
//
// 
//输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","lo
//g"]
//输出：[]
//解释：endWord "cog" 不在字典 wordList 中，所以不存在符合要求的转换序列。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= beginWord.length <= 7 
// endWord.length == beginWord.length 
// 1 <= wordList.length <= 5000 
// wordList[i].length == beginWord.length 
// beginWord、endWord 和 wordList[i] 由小写英文字母组成 
// beginWord != endWord 
// wordList 中的所有单词 互不相同 
// 
// 
// 
// Related Topics 广度优先搜索 数组 字符串 回溯算法 
// 👍 428 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// 将起始字符串、终止字符串、单词表里的所有字符串看作节点
	// 若两个字符串只有一个字符不同，那么将他们相连
	var ans [][]string
	dict := make(map[string]bool)
	for _, v := range wordList {
		dict[v] = true
	}
	if !dict[endWord] {
		return ans
	}
	delete(dict, beginWord)
	delete(dict, endWord)
	q1 := map[string]bool{beginWord: true}
	q2 := map[string]bool{endWord: true}
	next := make(map[string][]string)
	var reversed, found bool
	for len(q1) > 0 {
		q := make(map[string]bool)
		for key := range q1 {
			s := []byte(key)
			for i := 0; i < len(s); i++ {
				ch := s[i]
				for j := 0; j < 26; j++ {
					s[i] = byte(j + 'a')
					s1 := string(s)
					if q2[s1] {
						if reversed {
							next[s1] = append(next[s1], key)
						} else {
							next[key] = append(next[key], s1)
						}
						found = true
					}
					if dict[s1] {
						if reversed {
							next[s1] = append(next[s1], key)
						} else {
							next[key] = append(next[key], s1)
						}
						q[s1] = true
					}
				}
				s[i] = ch
			}
		}
		if found {
			break
		}
		for key := range q {
			delete(dict, key)
		}
		if len(q) <= len(q2) {
			q1 = q
		} else {
			reversed = !reversed
			q1 = q2
			q2 = q
		}
	}
	if found {
		path := []string{beginWord}
		backTracking126(beginWord, endWord, next, path, &ans)
	}
	return ans
}

func backTracking126(beginWord, endWord string, next map[string][]string, path []string, ans *[][]string) {
	if beginWord == endWord {
		temp := make([]string, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}
	for _, v := range next[beginWord] {
		path = append(path, v)
		backTracking126(v, endWord, next, path, ans)
		path = path[:len(path)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
