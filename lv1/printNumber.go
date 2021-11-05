/*******
* @Author:qingmeng
* @Description:
* @File:printNumber
* @Date2021/10/30
 */

package main

import "fmt"

func printNumber(i int ,ch chan<- int ){
	fmt.Print(i,"\t")
	if i!=0&&i%10==0{
		fmt.Println()
	}
	ch<-i
}

func main() {
	var (
		ch1 =make(chan int)
		ch2=make(chan int)
	)

	for i := 0; i < 100; i++ {
		go printNumber(i,ch1)
		<-ch1
		go printNumber(i,ch2)
		<-ch2
	}

}
