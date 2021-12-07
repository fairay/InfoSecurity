package main

type statMap struct {
	st map[byte]int
}

func NewStatMap() (obj *statMap) {
	obj = &statMap{
		st: make(map[byte]int),
	}
	return obj
}
func (this *statMap) Add(b byte) {
	if _, ok := this.st[b]; ok {
		this.st[b]++
	} else {
		this.st[b] = 1
	}
}