package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

// worker1: Сигнал через канал done
func worker1(done chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("worker1 завершён")
			return
		default:
			fmt.Println("worker1 работает...")
			time.Sleep(1 * time.Second)
		}
	}
}

// worker2: Контекст с отменой
func worker2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker2 завершён")
			return
		default:
			fmt.Println("worker2 работает...")
			time.Sleep(1 * time.Second)
		}
	}
}

// worker3: Контекст с тайм-аутом
func worker3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker3 завершён (тайм-аут)")
			return
		default:
			fmt.Println("worker3 работает...")
			time.Sleep(1 * time.Second)
		}
	}
}

// worker4: Принудительное завершение программы
func worker4() {
	for {
		fmt.Println("worker4 работает...")
		time.Sleep(1 * time.Second)
	}
}

// worker5: Завершение через WaitGroup
func worker5(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println("worker5 работает...")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Способ 1: Канал done
	done := make(chan bool)
	go worker1(done)

	// Способ 2: Контекст с отменой
	ctx2, cancel2 := context.WithCancel(context.Background())
	go worker2(ctx2)

	// Способ 3: Контекст с тайм-аутом
	ctx3, cancel3 := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel3()
	go worker3(ctx3)

	// Способ 5: WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)
	go worker5(&wg)

	// Завершение воркеров
	time.Sleep(3 * time.Second)
	done <- true // Завершение worker1
	cancel2()    // Завершение worker2

	// Ожидание worker5
	wg.Wait()
	fmt.Println("Все воркеры (кроме worker4) завершены")

	// Способ 4: Принудительное завершение программы
	go worker4()
	time.Sleep(2 * time.Second)
	fmt.Println("Принудительное завершение программы")
	os.Exit(0)
}
