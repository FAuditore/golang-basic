package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	/*
			字符串"",``
		    字符->编码值 A->65 a->97
			字节 byte->uint8
			len()返回字节个数
			字符串类型的零值是空串 ""
	*/
	str1 := "abc 哈哈"
	str2 := `\abc` + "`"
	fmt.Println(len(str1))                    //返回字节个数
	fmt.Println(len([]rune(str1)))            //字符个数
	fmt.Println(utf8.RuneCountInString(str1)) //字符个数
	fmt.Println(str2)

	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c ", str1[i])
	}

	fmt.Println()
	for i, v := range str1 {
		fmt.Printf("%d,%c ", i, v)
	}

	fmt.Println()

	//slice转字符串
	slice1 := []byte{65, 66, 67}
	str3 := string(slice1)
	fmt.Println(str3)
	//字符串转slice
	slice2 := []byte(str3)
	fmt.Println(slice2)

	//str3[1] = "D" 字符串不能修改

	//strings包
	strings1 := "hello world"
	fmt.Println(strings.Contains(strings1, "h"))       //是否包含
	fmt.Println(strings.ContainsAny(strings1, "abcd")) //是否包含任意一个字符
	fmt.Println(strings.Count(strings1, "o"))          //计数
	fmt.Println(strings.HasPrefix(strings1, "hel"))    //是否前缀
	fmt.Println(strings.LastIndex(strings1, "l"))      //最后一次位置

	slice3 := []string{"abc", "bcd", "cde"}
	strings2 := strings.Join(slice3, "-")
	fmt.Println(strings2)

	slice4 := strings.Split(strings2, "-")
	fmt.Println(slice4)

	strings3 := strings.Repeat("hello", 5)
	fmt.Println(strings3)

	//字符串替换,n代表个数 -1 全替换
	strings4 := strings.Replace(strings1, "l", "*", -1)
	fmt.Println(strings4)

	fmt.Println(strings.ToUpper(strings1))
	fmt.Println(strings.ToLower(strings1))

	//截取子串
	slice5 := strings1[0:6]
	fmt.Println(slice5)

	s1 := "true"
	//字符串转bool
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T %t\n", b1, b1)
	bs := strconv.FormatBool(b1)
	fmt.Printf("%T %s\n", bs, bs)

	s2 := "1100"
	//字符串转int(10进制)    (字符串，字符串数的进制，位数)
	i2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T %d\n", i2, i2)

	//int转字符串      (整型，整型的进制)
	is := strconv.FormatInt(i2, 10)
	fmt.Printf("%T %s\n", is, is)

	i3, err := strconv.Atoi("123")
	fmt.Printf("%T %d\n", i3, i3)

	is2 := strconv.Itoa(i3)
	fmt.Printf("%T %s\n", is2, is2)

}
