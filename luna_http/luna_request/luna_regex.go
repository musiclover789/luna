package luna_request

import (
	"fmt"
	"regexp"
	"strings"
)


func ParserOne(regex,str  string) string  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
		}
	}()
	r,err:=regexp.Compile(regex)
	if err !=nil{
		fmt.Print("正则表达式出错")
		return ""
	}
	return r.FindString(str)
}

// ISsMatch /****
func ISsMatch(regex,str  string) bool  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
		}
	}()
	if strings.Contains(str,regex){
		return true
	}
	r,err:=regexp.Compile(regex)
	if err !=nil{
		fmt.Print("正则表达式出错")
		return false
	}
	if len(r.FindString(str))>0{
		return true
	}
	return false
}

func ISABSMatch(regex,str  string) bool  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
		}
	}()
	r1:=regexp.MustCompile(regex)
	return r1.MatchString(str)
}

func RegexReplace(regex,str,new  string,group int) string  {
	defer func() {
		if err := recover(); err != nil {
			//
		}
	}()
	r,err:=regexp.Compile(regex)
	if err !=nil{
		fmt.Print("正则表达式出错")
		return ""
	}
	groupString := r.FindStringSubmatch(str)
	if len(groupString)==0{
		return ""
	}
	return strings.ReplaceAll(str,groupString[group],new)
}

func ParserOneByGroup(regex,str  string,group int) string  {
	defer func() {
		if err := recover(); err != nil {
			//
		}
	}()
	r,err:=regexp.Compile(regex)
	if err !=nil{
		fmt.Print("正则表达式出错")
		return ""
	}
	groupString := r.FindStringSubmatch(str)
	if len(groupString)==0{
		return ""
	}
	return groupString[group]
}

func CompressStr(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

func ParserMoreByGroup(regex,str  string) []string  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
		}
	}()
	r,err:=regexp.Compile(regex)
	if err !=nil{
		fmt.Print("正则表达式出错")
		return nil
	}
	groupString := r.FindAllString(str,-1)
	return groupString
}

func ParserMoreByGroupG(regex,str  string) [][]string  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
		}
	}()
	r,err:=regexp.Compile(regex)
	if err !=nil{
		fmt.Print("正则表达式出错")
		return nil
	}
	groupString := r.FindAllStringSubmatch(str,-1)
	return groupString
}

/***
 旧的字符串里面满足正则条件的 替换成新的字符串
 */
func ReplaceAll(regex,oldstr,newstr  string) string  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)

		}
	}()
	r,err:=regexp.Compile(regex)
	if err !=nil{
		fmt.Print("正则表达式出错")
		return ""
	}
	return r.ReplaceAllString(oldstr,newstr)
}

