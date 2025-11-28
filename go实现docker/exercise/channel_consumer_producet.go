package main

import (
	"fmt"
	"time"
)

// 用channel实现一个生产者消费者模型,实现恒定消费者

var MAXConcurrent = 10

type Task struct {
	ID int
}

type Consumer struct {
	TaskChan      chan Task
	MaxConcurrent int
}

func NewConsumer(maxConcurrent int) *Consumer {
	taskChan := make(chan Task, maxConcurrent)
	for i := 0; i < maxConcurrent; i++ {
		go func(i int) {
			for {
				task := <-taskChan
				fmt.Printf("Worker %d, handling task %d\n", i, task.ID)
				time.Sleep(time.Millisecond * 100)
			}
		}(i)
	}

	return &Consumer{
		TaskChan:      taskChan,
		MaxConcurrent: maxConcurrent,
	}
}

func Producer(taskChan chan Task) {
	for i := 0; i < 10000; i++ {
		taskChan <- Task{
			ID: i,
		}
	}
}

func main() {
	consumer := NewConsumer(10)
	Producer(consumer.TaskChan)

	time.Sleep(time.Second)
}
