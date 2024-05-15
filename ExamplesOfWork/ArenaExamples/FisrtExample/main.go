//go:build goexperiment.arenas
// +build goexperiment.arenas

package main

// Данная программа помогает наглядно увидеть процесс работы arena через trace.out файл
// В программе создается 1000 массивов длиной в 2000 элементов, в конце в память арены
// добавляем инт значение равное 42 (Кузбасский Регион), и выводим это значение.
// После запуска программы заходим в trace.out файл испольщуя команду go tool trace trace.out
// Видим что когда мы подошли GC не вызывался, а размер выделенной памяти в какой-то момент удвоился.
// Для наглядности под этим кодом закоменчем обычная ситуация с выделением памяти в которой
// будет вызываться GC.
import (
	"arena"
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

const (
	numberOfArrays = 1000
	arrayLen       = 2000
)

func main() {
	// Запись в trace файл
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()

	mem := arena.NewArena()
	for j := 0; j < numberOfArrays; j++ {
		o1 := arena.MakeSlice[float64](mem, arrayLen, arrayLen)
		for i := 0; i < arrayLen; i++ {
			o1[i] = float64(i)
		}
		time.Sleep(10 * time.Millisecond)
	}
	t := arena.New[int](mem)
	value := 42
	t = &value
	fmt.Println(*t)
	mem.Free()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("KUZBASS")
}

/*
package main


import (
	"fmt"
	"os"
	"runtime/debug"
	"runtime/trace"
	"sync"
	"time"
)

const (
	NumWorkers    = 4     // Количество воркеров
	NumTasks      = 500   // Количество задач
	MemoryIntense = 10000 // Размер память затратной задачи (число элементов)
)

func main() {
	// Запись в trace файл
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()

	// Установка целевого процента сборщика мусора GOGC=1000.
	debug.SetGCPercent(1000)

	// Очередь задач и очередь результата
	taskQueue := make(chan int, NumTasks)
	resultQueue := make(chan int, NumTasks)

	// Запуск воркеров
	var wg sync.WaitGroup
	wg.Add(NumWorkers)
	for i := 0; i < NumWorkers; i++ {
		go worker(taskQueue, resultQueue, &wg)
	}

	// Отправка задач в очередь
	for i := 0; i < NumTasks; i++ {
		taskQueue <- i
	}
	close(taskQueue)

	// Получение результатов из очереди
	go func() {
		wg.Wait()
		close(resultQueue)
	}()

	// Обработка результатов
	for result := range resultQueue {
		fmt.Println("Результат:", result)
	}

	fmt.Println("Готово!")
}

// Функция воркера
func worker(tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		result := performMemoryIntensiveTask(task)
		results <- result
	}
}

// performMemoryIntensiveTask функция требующая много памяти
func performMemoryIntensiveTask(task int) int {
	// Создание среза большого размера
	data := make([]int, MemoryIntense)
	for i := 0; i < MemoryIntense; i++ {
		data[i] = i + task
	}

	// Имитация временной задержки
	time.Sleep(10 * time.Millisecond)

	// Вычисление результата
	result := 0
	for _, value := range data {
		result += value
	}
	return result
}
*/
