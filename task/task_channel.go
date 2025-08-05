package task
import "fmt"

func ChannelTest1(ch chan int){
	go func ()  {
		for i:=1;i<=10;i++{
			ch <- i
			fmt.Println("打印发送数据=",i)
		}
	}()
	
	go func() {
		for value:=range ch{
			fmt.Println("打印接受数据=",value)
		}
		close(ch)
	}()
}


func ChannelTest2(ch chan int){
	go func ()  {
		for i:=1;i<=100;i++{
			ch <- i
			fmt.Println("打印发送数据=",i)
		}
	}()
	
	go func() {
		for value:=range ch{
			fmt.Println("打印接受数据=",value)
		}
		close(ch)
	}()
}