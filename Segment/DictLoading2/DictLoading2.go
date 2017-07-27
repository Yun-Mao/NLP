package DictLoading2

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
	"log"
	"大作业/Config"
)

type Trie_Node struct {
	Children  map[rune]*Trie_Node //所有子节点
	Character string               //存储的值
	Num       int                //经过这个节点的单词数量
}


//map实现的trie树
type Trie struct {
	Root map[int32]*Trie_Node
}

func NewMapTrie() *Trie {
	ins := &Trie{
		Root: make(map[int32]*Trie_Node),
	}
	return ins
}
func reverse(str string) string {
	rs := []rune(str)
	len := len(rs)
	var tt []rune

	tt = make([]rune, 0)
	for i := 0; i < len; i++ {
		tt = append(tt, rs[len-i-1])
	}
	return string(tt[0:])
}

//插入一个单词
func (this *Trie) Insert(word string) {
	words := strings.Split(word, " ")
	//fmt.Println(words)
	word = words[0]
	character := words[2]
	//fmt.Println(character)
	runes := []rune(reverse(string(word)))
	//runes2 := []rune(string(character))
	var stopIdx = len(runes) - 1
	var current rune = runes[0]
	//fmt.Println(runes)
	var root *Trie_Node
	if len(runes)==1{
		root= this.getRoot(current,character)
	}else{
		root= this.getRoot(current,"nw")
	}
	//root.Num=1
	for i := 1; i <= stopIdx; i++ {
		current = runes[i]
		var seg *Trie_Node
		if i==stopIdx{
			seg = root.AddChild(current,character)
		}else{
			seg = root.AddChild(current,"nw")
		}

		root = seg
	}
	//seg := root.AddCharacter(runes2[0])
	root.addFrequency()
	//fmt.Println(this.Root)
}
func (this *Trie) getRoot(value int32,str string) *Trie_Node {
	node := this.Root[value]
	if node == nil {
		node = &Trie_Node{}
		node.Children = make(map[rune]*Trie_Node)
		this.Root[value] = node
		this.Root[value].Character=str;
	}
	return node
}
func (this *Trie_Node) addFrequency() {
	this.Num++
}

func (this *Trie_Node) AddChild(value int32,str string) *Trie_Node {
	node := this.Children[value]
	if node == nil {
		node = &Trie_Node{}
		node.Children = make(map[rune]*Trie_Node)
		this.Children[value] = node
		this.Children[value].Character=str;
	}
	return node
}
func (this *Trie_Node) getChild(value int32) *Trie_Node {
	return this.Children[value]
}

//分词
func (this *Trie) Segment(sentence string) []string {
	//sentence +="0"
	runes := []rune(reverse(sentence))
	words := make([]string, 0)
	length := len(runes)
	var current rune
	var root *Trie_Node = nil
	var word string = ""
	var mark int = 0
	var node2 string
	var flag bool
	flag = false
	for i := 0; i < length; i++ {
		//fmt.Println(length)
		current = runes[i]
		//fmt.Println(root)
		//fmt.Println(current)
		if root == nil {
			mark = i
			//fmt.Println(length)
			root = this.Root[current]
			//fmt.Println(root)
			if root != nil{
				node2 = root.Character
				if flag==true{
					word = string(current)
					flag=false
				}else{
					word += string(current)
				}
				if i==length-1{
					word = reverse(word)
					word += string("/")
					word += string(node2)
					words = append(words, word)
					word = ""
				}
			} else{
				//fmt.Println(length)
				word += string(current)
				node2 = "nw"
				word = reverse(word)
				//word += string(current)
				word += string("/")
				word += string(node2)

				words = append(words, word)
				word = ""
				node := this.Root[current]
				node = node
				flag = true
			}
		} else {
			//fmt.Println(length)
			node := root.getChild(current)

			//fmt.Println(node)
			//word += string(node2)
			if node != nil {
				node2 = node.Character
				word += string(current)
				if i==length-1{
					word = reverse(word)
					word += string("/")
					word += string(node2)
					words = append(words, word)
					word = ""
				}
				//fmt.Println(length)
				//fmt.Println(word)
				//node2 = node.Character
				//word += string("/")
				//word += string(node2)
				//fmt.Println(node2)
			} else {
				node2 = root.Character
				//fmt.Println(root.Num)
				if root.Num > 0 {
					word = reverse(word)
					word += string("/")
					word += string(node2)
					words = append(words, word)
					i--
					word = ""
				} else {
					word = string(runes[mark])
					word += string("/")
					word += string(node2)
					word += string(" ")
					//fmt.Println("hehe  ")
					//node2 := this.Root[current].Character
					//word += string(node2)
					i = mark
				}
			}
			root = node
		}
	}
	for i:=0;i<len(words)/2;i++{
		var temp string
		temp = words[i]
		words[i] = words[len(words)-i-1]
		words[len(words)-i-1]= temp
	}
	return words
}
var MapTrieSeg *Trie

//加载词典1
func (this *Trie) LoadDict(dictpath string,userpath string) {
	f, err := os.Open(dictpath)
	if err != nil {
		log.Println("open file:" + dictpath + "error !")
		return
	}
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		//fmt.Println(line)
		word := string(line)
		//fmt.Println(word)
		this.Insert(word)
	}
	f.Close()

	f, err = os.Open(userpath)
	if err != nil {
		log.Println("open file:" + userpath + "error !")
		return
	}
	reader = bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		//fmt.Println(line)
		word := string(line)
		//fmt.Println(word)
		this.Insert(word)
	}
	f.Close()
}

//加载词典2,加载dict2 用的 现在改了
func LoadDict2(dictpath string,userpath string) error{
	f, err := os.Open(dictpath)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		line = line[0:strings.IndexAny(line," ")]
		MapTrieSeg.Insert(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	f.Close()
	f, err = os.Open(userpath)
	if err != nil {
		return err
	}
	buf = bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		line = line[0:strings.IndexAny(line," ")]
		MapTrieSeg.Insert(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}


//init自动调用
func DictLoading2() {
	MapTrieSeg = NewMapTrie()
	MapTrieSeg.LoadDict(Config.DictPath,Config.UserPath)
	fmt.Println("字典加载完毕")
}