package goroutine_channel

import (
	"fmt"
	"net/http"
	"time"
)

func Goroutine1() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}
	ch := make(chan string)
	for _, api := range apis {
		go func(api string, ch chan string) {
			_, err := http.Get(api)
			if err != nil {
				ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
				return
			}
			ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)

		}(api, ch)
	}
	// 需要了解这两个的不同
	//for range apis {
	//	fmt.Println(<-ch)
	//}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
