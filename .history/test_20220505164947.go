package main
 
import (
        "fmt"
        "strings"
)
 
//golang字符串操作
func main(){
        s := "->aa "
        //str := "wo"
 
        //以str为分隔符，将s切分成多个子切片，结果中**不包含**str本身。
        //如果str为空，则将s切分成Unicode字符列表。如果s中没有str子串，
        //则将整个s作为[]string的第一个元素返回
        index := strings.Split(s," ")
        fmt.Println(index[1]) //4
}