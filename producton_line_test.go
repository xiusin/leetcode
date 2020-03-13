package leetcode

import (
	"fmt"
	"testing"
)

/**
有一家生产奶酪的厂家，每天需要生产100000份奶酪卖给超市，通过一辆送货车发货，送货车辆每次送100份。
厂家有一个容量为1000份的冷库，用于奶酪保鲜，生产的奶酪需要先存放在冷库，运输车辆从冷库取货。
厂家有三条生产线，分别是牛奶供应生产线，发酵剂制作生产线，奶酪生产线。

生产每份奶酪需要2份牛奶和1份发酵剂。请设计生产系统。

发酵剂生产线 -------------|
						|
						| --------- 奶酪生产线 ----------  冷库 -------- 送货车
						|
牛奶供应生产线 -----------|
*/

var ferment = make(chan struct{}) //发酵剂
var milk = make(chan struct{})    // 牛奶

type cheese struct {
	ferment    chan struct{}
	milk       chan struct{}
	done       chan struct{}
	fermentNum int
	milkNum    int
}

const cheeseNum = 100000

var iceBox = make(chan cheese, 1000)

var truck = make(chan cheese, 100)

func TestProductionLine(t *testing.T) {
	var done = make(chan struct{})
	fermentNum := cheeseNum
	milkNum := cheeseNum
	go func() {
		for i := 0; i < fermentNum; i++ {
			ferment <- struct{}{}
		}
	}()

	go func() {
		for i := 0; i < milkNum; i++ {
			milk <- struct{}{}
		}
	}()

	go func() {
		for i := 0; i < cheeseNum; i++ {
			go func() {
				cheese := cheese{
					ferment: make(chan struct{}),
					milk:    make(chan struct{}),
					done:    make(chan struct{}),
				}
				go func() {
					cheese.milk <- <-milk
					milkNum++
					if cheese.fermentNum == 2 {
						cheese.done <- struct{}{}
					}
				}()
				go func() {
					for {
						cheese.ferment <- <-ferment
						cheese.fermentNum++
						if cheese.fermentNum == 2 {
							if cheese.milkNum == 1 {
								cheese.done <- struct{}{}
							}
							break
						}
					}
				}()
				go func() {
					<-cheese.done
					iceBox <- cheese
					fmt.Printf("第%d个奶酪已经制作完成\n", i+1)
				}()
			}()
		}
	}()

	//go func() {
	//	for cheese := range iceBox {
	//		car <- cheese
	//		fmt.Println("已经取出了奶酪放到运货车上")
	//	}
	//}()

	<-done
}
