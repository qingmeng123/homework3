/*******
* @Author:qingmeng
* @Description:
* @File:main
* @Date2021/11/6
 */

package main

import "fmt"

func main() {
	over := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			<-over		//在分协程里释放管道的值
		}()
			over <- true	//每次都给管道传输内容，使其堵塞

	}

	fmt.Println("over!!!")
}
