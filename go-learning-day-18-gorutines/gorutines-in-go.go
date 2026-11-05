package main
import (
	"fmt"
	"time"
	"sync"
	//  "sync/atomic"
)



// func main(){
// 	go test()
// 	go test()
// 	go test()
// }
// func main(){
// 	go test()	
// 	time.Sleep(1 * time.Second)
//     fmt.Println("Hello World from main")
// }


// func main(){
// 	var wg sync.WaitGroup
// 	for i:=0; i<10;i++{
// 		wg.Add(1)
// 		go test(i,&wg)
// 	}
// 	wg.Wait()
// }




// func test(id int, wg *sync.WaitGroup, count *int){
// 	defer wg.Done()
// 	fmt.Println("Hello World from test", id)
// 	for i:=0; i<1000000;i++ {
// 		(*count)++
// 	}
// 	fmt.Println("Count from test", id, ":", *count)
// }



// func main(){
// 	var wg sync.WaitGroup
// 	count := 0
// 	wg.Add(2)
// 	go test(1,&wg,&count)
// 	go test(2,&wg,&count)
// 	wg.Wait()
// }






// func test(id int, wg *sync.WaitGroup, count *int, mu *sync.Mutex){
// 	defer wg.Done()
// 	fmt.Println("Hello World from test", id)
// 	for i:=0; i<100000000;i++ {
// 		mu.Lock()
// 		(*count)++
// 		mu.Unlock()
// 	}
// 	fmt.Println("Count from test", id, ":", *count)
// }


// func main(){
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	var mu sync.Mutex
// 	count := 0
// 	test(1,&wg,&count,&mu)
// 	test(2,&wg,&count,&mu)
// 	wg.Wait()
// }






// func test(id int, wg *sync.WaitGroup, count *int){
// 	defer wg.Done()
// 	fmt.Println("Hello World from test", id)
// 	for i:=0; i<100000000;i++ {
// 		atomic.AddInt64((*int64)(count),1)
// 	}
// 	fmt.Println("Count from test", id, ":", *count)
// }


// func main(){
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	count := 0
// 	test(1,&wg,&count)
// 	test(2,&wg,&count)
// 	wg.Wait()
// }


// func main(){
// 	ch:=make(chan int)
// 	go func(){
// 		ch <- 1
// 	}()
// 	value:= <-ch
// 	fmt.Println(value)
// }




// func main(){
// 	ch:=make(chan int)
// 	go func(){
// 		ch <- 1
// 		ch <- 2
// 	}()
// 	value:= <-ch
// 	againValue:= <-ch

// 	fmt.Println(value)
// 	fmt.Println(againValue)
// }

//suppose if u try to read from channel before writing to it, it will block the execution of the program until the value is written to the channel. This is because channels in Go are blocking by default. When you try to read from a channel that has no values, the goroutine will wait until a value is sent to that channel. Similarly, if you try to send a value to a channel that has no receiver, the sending goroutine will block until another goroutine reads from that channel.
//and if u does not write any thing to the channel and try to read from it, it will block the execution of the program indefinitely, leading to a deadlock situation. This is because there is no value being sent to the channel, and the reading goroutine will wait forever for a value that will never come. it will result in a deadlock error
//same with the writing to the channel, if there is no receiver for the value being sent, the sending goroutine will block indefinitely, leading to a deadlock situation. This is because there is no goroutine waiting to receive the value being sent, and the sending goroutine will wait forever for a receiver that will never come. it will result in a deadlock error
//above are the examples of unbuffered channels, which means that the channel has no capacity to hold values. In contrast, buffered channels have a capacity to hold a certain number of values, and they allow sending and receiving values without blocking until the buffer is full or empty.


// func main(){
// 	ch:=make(chan int, 3)
// 	go func(){
// 		fmt.Println("Hello World from goroutine")
// 		ch <- 1
// 		fmt.Println("Hello World from goroutine")
// 		ch <- 2
// 		fmt.Println("Hello World from goroutine")
// 		ch <- 3
// 		fmt.Println("Hello World from goroutine")
// 		ch <- 4
// 		fmt.Println("Hello World from goroutine")
// 	}()
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Hello World from main")
// 	// fmt.Println(<-ch)
// 	// fmt.Println(<-ch)
// 	// fmt.Println(<-ch)
// }

//idea is simple like in go func out of 4 fmt.Println() statements, only 3 will be executed because the channel has a buffer size of 3. The fourth fmt.Println() statement will block the execution of the goroutine until there is space in the channel buffer to send the value. Since the main function is sleeping for 2 seconds, it will not read from the channel during that time, and the fifth fmt.Println() statement will not be executed. After 2 seconds, the main function will read from the channel and unblock the goroutine, allowing it to execute the fifth fmt.Println() statement.
//like in case o f 5th fmt.Println() statement, it will block the execution of the goroutine until there is space in the channel buffer to send the value. Since the main function is sleeping for 2 seconds, it will not read from the channel during that time, and the fifth fmt.Println() statement will not be executed. After 2 seconds, the main function will read from the channel and unblock the goroutine, allowing it to execute the fifth fmt.Println() statement. because the buffer is full to place 4th digit as it has capacity of 3, so it will block the execution of the goroutine until there is space in the channel buffer to send the value. Since the main function is sleeping for 2 seconds, it will not read from the channel during that time, and the fifth fmt.Println() statement will not be executed. After 2 seconds, the main function will read from the channel and unblock the goroutine, allowing it to execute the fifth fmt.Println() statement.	


func consumer(wg *sync.WaitGroup,id int,jobs <- chan int){
	defer wg.Done()	
	for job := range jobs {
		fmt.Printf("Consumer %d processing job %d\n", id, job)
		time.Sleep(1 * time.Second)
	}
}


func main(){
	var jobs sync.WaitGroup;
	ch:=make(chan int, 5)

	for i:=0; i<=3;i++{
		jobs.Add(1)
		go consumer(&jobs, i, ch)
	}
	for i:=0; i<5;i++{
		ch <- i
	}

	close(ch)
	jobs.Wait()
	fmt.Println("All jobs processed")
}



//idea of above programme is simples of the things exists on this entire