package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type Worker struct {
	ch       *chan int // Канал для получения чисел, которые будут обрабатываться воркером
	cancelCh chan bool // Канал для сигнализации о завершении работы воркера
}

func (w Worker) work() {
	for {
		select {
		// Если приходит сигнал в канал cancelCh, воркер завершает работу
		case <-w.cancelCh:
			return
		// Получаем число из канала ch, вычисляем его квадрат и выводим результат
		case num := <-*w.ch:
			fmt.Printf("Квадрат числа %d равен %d\n", num, num*num)
		}
	}
}

func (w Worker) cancel() {
	// Отправляем сигнал завершения работы в канал cancelCh
	w.cancelCh <- true
	// Закрываем канал cancelCh, чтобы избежать утечек
	close(w.cancelCh)
	fmt.Println("Воркер завершил работу")
}

func main() {
	// Создаем сканер для считывания пользовательского ввода
	scanner := bufio.NewScanner(os.Stdin)
	// Настраиваем сканер на разделение текста по словам
	scanner.Split(bufio.ScanWords)

	n := 0

	fmt.Println("Введите число воркеров:")

	// Считываем количество воркеров
	if scanner.Scan() {
		text := scanner.Text()
		num, err := strconv.Atoi(text)
		n = num
		if err != nil {
			fmt.Println("Ошибка: введенное значение не является числом.")
		}
	}

	// Проверяем наличие ошибок ввода
	if err := scanner.Err(); err != nil {
		_, err2 := fmt.Fprintln(os.Stderr, "Ошибка чтения ввода:", err)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	// Создаем канал для передачи чисел воркерам
	ch := make(chan int)

	// Создаем и запускаем воркеров
	workers := make([]Worker, n)
	for i := 0; i < n; i++ {
		workers[i] = Worker{&ch, make(chan bool)}
		go workers[i].work()
	}

	// Отправляем числа в канал ch
	go func() {
		for i := 0; i < 1000000; i++ {
			ch <- i
		}
	}()

	// Создаем канал для отслеживания системных сигналов завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// Ожидаем сигнал завершения
	call := <-stop

	// Останавливаем всех воркеров
	for _, worker := range workers {
		worker.cancel()
	}
	close(ch)
	// Выводим сигнал завершения
	fmt.Println(call.String())
}
