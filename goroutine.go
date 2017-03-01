package main
import ("fmt"
	// "math/rand"
	"time"
	//"sync"
	"sync"
)

func Walking(walker string, done chan bool){

	timeout := time.Now().Add(time.Duration(3)*time.Second)
	for time.Now().Before(timeout){
		// I don't want the screen is populated with messages
		time.Sleep(time.Duration(1)*time.Second)
		fmt.Printf("%v is still walking\n", walker)
	}

        // send value to chan
	done <- true
}

func Running(runner string){
	fmt.Printf("%s is running!\n",runner)
}


func WaitForOther(){

	go Running("Alex")
	go Running("Helen")
	go Running("Nova")
	// if you don't use time.Sleep, you will get no result
	// because you fire 3 go routines then come back and continue with
	// main routine it is the entrance and exit of the program, it exits immediately
	// this program use one cpu which is only capable of executing one task at a time
	time.Sleep(time.Duration(4)*time.Second)

}


func Communicate() {
	signal := make(chan bool)
	go Walking("Puppy", signal)
	// receive value from chan signal
	if <-signal {
		fmt.Printf("%v is done with walking\n", "Puppy")
	}
}


func WriteSharedCounter(){
	var counter int
	BuySauce := func (caller string, num *int){
					for i:=0;i<=5;i++{
						fmt.Printf("buy one Soy Sause for %v\n",caller)
						*num++
					}
				}

	go BuySauce("Mom",&counter)
	go BuySauce("Father",&counter)
	// this make sure writeSharedCounter does not return immediately
	time.Sleep(time.Duration(3)*time.Second)
	fmt.Printf("counter is %v\n", counter)

}

func WriteMutexCounter(){
	var mux sync.Mutex
	var counter int
	BuySauce := func (caller string){
					for i:=0;i<=5;i++{
						fmt.Printf("buy one Soy Sause for %v\n",caller)
						mux.Lock()
						counter++
						mux.Unlock()
					}
				}

	go BuySauce("Mom")
	go BuySauce("Father")
	// this make sure writeSharedCounter does not return immediately
	time.Sleep(time.Duration(3)*time.Second)
	fmt.Printf("counter is %v\n", counter)



}

func WaitGroup(){
        // the caller will wait for other goroutines to finish before it returns
	var wg sync.WaitGroup
	DoTask := func(){
			fmt.Printf("Doing task! Guranteed to be Printed!\n")
			wg.Done()
	}
	wg.Add(2)
	go DoTask()
	go DoTask()
	wg.Wait()


}

func main(){
       	Communicate()
	WaitForOther()
	WriteSharedCounter()
	WriteMutexCounter()
	WaitGroup()


}
