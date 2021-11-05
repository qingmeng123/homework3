/*******
* @Author:qingmeng
* @Description:
* @File:channelTest
* @Date2021/10/30
 */

package main

import "fmt"

var x int64


func add(ch chan int) {
	for i := 0; i < 50000; i++ {
		x = x + 1
	}
	<-ch
}

func main() {
	ch1:=make(chan int)
	ch2:=make(chan int)
	go add(ch1)
	ch1<-1
	go add(ch2)
	ch2<-2
	fmt.Println(x)
}
