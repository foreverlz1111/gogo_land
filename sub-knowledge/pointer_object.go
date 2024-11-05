package main

type People interface {
	Speak(string) string
}

type Student struct {
}

func (student Student) Speak(think string) (talk string) {
	println(think)
	talk = "hi"
	return talk
}
func main() {
	var peo People = Student{}
	println(peo.Speak("hello"))
}
