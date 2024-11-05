package main

func main() {
	const s = "golangtag"
	println(len(s[:]))
	var a byte = 1 << len(s) / 128    // 常量操作
	var b byte = 1 << len(s[:]) / 512 // 不是常量操作-越位报错
	println(a, b)
}
