package main

import (
	"大作业/Segment"
	"fmt"
	"os"
	"os/exec"
	"log"
	"bufio"
	"大作业/Config"
	"strings"
	"大作业/Segment/DictLoading"
	"大作业/Segment/DictLoading2"
)
func Input(){
	DictLoading.DictLoading()
	var inputReader *bufio.Reader
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("请输入句子：")
	sentence, err := inputReader.ReadString('\n')
	check(err)
	sentence = strings.Replace(sentence , " ", "|", -1)
	sentence = strings.Replace(sentence , "\n", "", -1)
	//fmt.Println(sentence)
	var list []string
	list=Seg.DeafaultSegment().Segment(sentence)
	fmt.Println(list)
	Goon()
}
func Input2(){
	DictLoading2.DictLoading2()
	var inputReader *bufio.Reader
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("请输入句子：")
	sentence, err := inputReader.ReadString('\n')
	check(err)
	sentence = strings.Replace(sentence , " ", "|", -1)
	sentence = strings.Replace(sentence , "\n", "", -1)
	//print(sentence)
	fmt.Println(Seg.DeafaultSegment2().Segment(sentence))
	Goon()
}
func menu()  {//菜单函数
	fmt.Println("1------中文分词（正序）")
	fmt.Println("2------中文分词（逆序）")
	fmt.Println("3------添加词典")
	fmt.Println("4------退出")
	var a int
	fmt.Scanln(&a)
	switch a{
	case 1: Input()
	case 2: Input2()
	case 3: Save_File()
	case 4: os.Exit(0)
	}
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func Clsscr() {
	cmd := exec.Command("cmd", "/c", "cls") // /c执行字符串指定的命令然后终止
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
func Goon() {
	fmt.Println("按回车键继续")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
func help(){

}
func Save_File(){
	f, err := os.OpenFile(Config.UserPath,os.O_APPEND,0644)
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Println("请输入词汇")
	var num string
	fmt.Scanln(&num)
	w.WriteString(num)
	w.Flush()
	w.WriteString(" ")
	w.Flush()
	w.WriteString("0")
	w.Flush()
	w.WriteString(" ")
	w.Flush()
	fmt.Println("请输入词性：")
	var cha string
	fmt.Scanln(&cha)
	w.WriteString(cha)
	w.Flush()
	w.WriteString("\n")
	w.Flush()
    fmt.Println("保存成功！！！")
	Goon()
}
func main() {
	defer func() {
		if err := recover();err != nil {
			fmt.Println("程序异常退出:",err)
			Goon()
			Clsscr()
			main()
		} else {
		}
	}()
	for{
		Clsscr()
		menu()
	}

}
