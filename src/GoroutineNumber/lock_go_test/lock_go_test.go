package lock_go_test // 协程并发数限制库

import (
	"Golang_study/src/GoroutineNumber/lock_go_test"
	"log"
	"testing"
	"time"
)

func TestNewGoLimit(t *testing.T) {
	log.Println("开始测试...")
	g := lock_go.NewGoLimit(2) //max_num(最大允许并发数)设置为2

	for i := 0; i < 10; i++ {
		//并发计数加1.若 计数>=max_num, 则阻塞,直到 计数<max_num
		g.Add()

		//运行过程中可以随时修改最大可并发数据
		//g.SetMax(3)

		go func(g *lock_go.GoLimit, i int) {
			defer g.Done() //并发计数减1

			time.Sleep(time.Second * 2)
			log.Println(i, "done")
		}(g, i)
	}

	log.Println("循环结束")

	g.WaitZero() //阻塞, 直到所有并发都完成
	log.Println("测试结束")

}
