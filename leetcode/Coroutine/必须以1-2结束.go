
// 控制最后一定是1，2结尾
// 方式1，最后要结束的时候发给协程一个false
func main2() {
	chSin <- true
	go PrintSin2()
	go PrintDou2()
	time.Sleep(2)
	chDou <- false
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(input.Text())
	}

}

func PrintSin2() {
	for {
		_, ok := <-chSin
		if !ok {
			fmt.Println("Single Out")
			return
		}
		fmt.Println(1)
		chDou <- true
	}
}

func PrintDou2() {
	for {
		ok := <-chDou
		fmt.Println(2)
		if !ok {
			close(chDou)
			close(chSin)
			fmt.Println("Double Out")
			return
		}
		chSin <- true
	}
}

// 方式2:原子变量
var NeedQuit int32

func PrintDou21() {
	for {
		<-chDou
		fmt.Println(2)
		if atomic.CompareAndSwapInt32(&NeedQuit, 1, 2) {
			close(chDou)
			close(chSin)
			fmt.Println("Double Out")
			return
		}
		chSin <- true
	}
}
