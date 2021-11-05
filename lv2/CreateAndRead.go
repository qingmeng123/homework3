/*******
* @Author:qingmeng
* @Description:
* @File:CreateAndRead
* @Date2021/11/6
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error)  {
	if e!=nil{
		panic(e)
	}
}

func main() {
	f,err:=os.Create("lv2/plan.txt")
	check(err)
	defer f.Close()

	str:="I’m not afraid of difficulties and insist on learning programming"
	n1,err:=f.WriteString(str)
	check(err)
	fmt.Println("wrote",n1,"bytes")

	f,err=os.Open("lv2/plan.txt")
	b:=make([]byte,n1)
	n2,err:=f.Read(b)
	fmt.Println(n2,"bytes:",string(b))

	dat,err:=ioutil.ReadFile("lv2/plan.txt")
	check(err)
	fmt.Println(string(dat))
}
