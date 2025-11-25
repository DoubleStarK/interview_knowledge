package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID int
}

type TaskQueue struct {
	Tasks         chan Task
	ConcurrentNum int
	wg            *sync.WaitGroup
}

func (q *TaskQueue) Init(concurrentNum int) error {
	if concurrentNum == 0 {
		return errors.New("invalid concurrentNum 0")
	}
	q.ConcurrentNum = concurrentNum
	q.Tasks = make(chan Task, concurrentNum)
	q.wg = &sync.WaitGroup{}
	return nil
}

func (q *TaskQueue) AddTask(t Task) error {
	if q.ConcurrentNum == 0 {
		return errors.New("0")
	}
	q.Tasks <- t
	return nil
}

func (q *TaskQueue) Wait() {
	q.wg.Wait()
	close(q.Tasks)
}

func (q *TaskQueue) Execute(f func(Task) error) {
	// 启动并发 worker，worker 从 q.Tasks 中读取任务并执行
	for i := 0; i < q.ConcurrentNum; i++ {
		go func() {
			for t := range q.Tasks {
				q.wg.Add(1)
				// 在当前 worker goroutine 中同步执行任务，执行完毕后 Done
				func(task Task) {
					defer q.wg.Done()
					if err := f(task); err != nil {
						fmt.Printf("task %d error: %v\n", task.ID, err)
					}
				}(t)
			}
		}()
	}
}

func main() {
	var tasks []Task
	for i := 0; i < 100; i++ {
		tasks = append(tasks, Task{
			ID: i,
		})
	}
	queue := TaskQueue{}
	queue.Init(10)

	go queue.Execute(func(t Task) error {
		time.Sleep(time.Millisecond * 1000)
		fmt.Println(t.ID)
		return nil
	})
	for _, t := range tasks {
		queue.AddTask(t)
	}
	queue.Wait()
}
