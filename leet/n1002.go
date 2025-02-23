package main

import "log"

//1002. 查找共用字符
//简单
//给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（包括重复字符），并以数组形式返回。你可以按 任意顺序 返回答案。
//
//
//示例 1：
//
//输入：words = ["bella","label","roller"]
//输出：["e","l","l"]
//示例 2：
//
//输入：words = ["cool","lock","cook"]
//输出：["c","o"]
//
//
//提示：
//
//1 <= words.length <= 100
//1 <= words[i].length <= 100
//words[i] 由小写英文字母组成

func commonChars(words []string) []string {
	first := [26]int{}
	for _, w := range words[0] {
		first[w-'a']++
	}

	for _, word := range words[1:] {
		other := [26]int{}
		for _, w := range word {
			other[w-'a']++
		}
		for i := 0; i < 26; i++ {
			first[i] = min(first[i], other[i])
		}
	}

	var hashdit []string
	for i := 0; i < 26; i++ {
		for first[i] > 0 {
			hashdit = append(hashdit, string(rune(i+'a')))
			first[i]--
		}
	}

	return hashdit
}
func main() {
	words := []string{"bella", "label", "roller"}
	log.Println(commonChars(words))
}
