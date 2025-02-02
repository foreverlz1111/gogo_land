package main

import "log"

//49. 字母异位词分组
//
//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
//字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
//
//
//
//示例 1:
//
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//示例 2:
//
//输入: strs = [""]
//输出: [[""]]
//示例 3:
//
//输入: strs = ["a"]
//输出: [["a"]]

// 注意的问题：避免重复遍历

func groupAnagrams(strs []string) [][]string {
	mymap := map[rune]int{
		'a': 2, 'b': 3, 'c': 5, 'd': 7, 'e': 11, 'f': 13, 'g': 17, 'h': 19,
		'i': 23, 'j': 29, 'k': 31, 'l': 37, 'm': 41, 'n': 43, 'o': 47, 'p': 53,
		'q': 59, 'r': 61, 's': 67, 't': 71, 'u': 73, 'v': 79, 'w': 83, 'x': 89,
		'y': 97, 'z': 101,
	}

	arr := make([][]string, 0)
	if len(strs) == 0 {
		arr = append(arr, []string{})
		return arr
	} else if len(strs) == 1 {
		arr = append(arr, []string{strs[0]})
		return arr
	} else {
		mymap_number := map[int][]string{}
		for _, str := range strs {
			key := 1
			for _, s := range str {
				key *= mymap[s]
			}
			mymap_number[key] = append(mymap_number[key], str)
		}
		for _, v := range mymap_number {
			arr = append(arr, v)
		}
	}
	return arr
}
func main() {
	log.Println(groupAnagrams([]string{}))
	log.Println(groupAnagrams([]string{"one"}))
	log.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))

}
