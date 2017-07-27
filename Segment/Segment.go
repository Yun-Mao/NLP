package Seg

import (
	"大作业/Segment/DictLoading"
	"大作业/Segment/DictLoading2"
)
//input接口
type input interface {
	Segment(sentence string) []string //分词
}
//常规分词方法
func DeafaultSegment() input {
	return DictLoading.MapTrieSeg
}
//逆序最大分词方法
func DeafaultSegment2() input {
	return DictLoading2.MapTrieSeg
}