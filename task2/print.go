package main

import "fmt"

func start(n int, sign chan bool) {
	go func() {
		var count int
		for i := 0; i <= n; i += 2 {
			fmt.Println(i)
			count++
			if count == 5 {
				sign <- true
				select {
				case <-sign:
					count = 0
				}
			}
		}
		sign <- true
	}()
	go func() {
		count, i := 0, -1
		for {
			if i > n {
				sign <- true
				return
			}
			select {
			case <-sign:
				for i += 2; i <= n; i += 2 {
					fmt.Println(i)
					count++
					if count == 5 {
						count = 0
						sign <- true
						break
					}
				}
			}
		}
	}()
}

/*func main() {
	start(100,make(chan bool))
	for {}
}
*/
