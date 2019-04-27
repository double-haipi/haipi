package utils

import (
	"math/rand"
	"time"
)

type task struct {
	id        int64
	countdown int64
	interval  int64
	callback  func()
}
type timers struct {
	taskList []task
	ticker   *time.Ticker
}

var store timers
var taskListToDelete []task
var isTimerRuning bool

func init() {
	store = timers{ticker: time.NewTicker(time.Second)}
	rand.Seed(time.Now().UnixNano())
}

func runTimers() {
	for {
		select {
		case <-store.ticker.C:
			for i := 0; i < len(store.taskList); i++ {
				pointer := &store.taskList[i]
				pointer.countdown = pointer.countdown - 1

				if pointer.countdown > 0 {
					continue
				}
				pointer.callback()
				if pointer.interval > 0 {
					pointer.countdown = pointer.interval
					continue
				}
				taskListToDelete = append(taskListToDelete, *pointer)
			}

			// if len(taskListToDelete) == 0 {
			// 	continue
			// }
			// var refreshTaskList []task
			// for _, pointer := range store.taskList {
			// 	for _, delValue := range taskListToDelete {
			// 		if value.id != delValue.id {
			// 			refreshTaskList = append(refreshTaskList, value)
			// 		}
			// 	}
			// }
			// store.taskList = refreshTaskList
			// taskListToDelete = taskListToDelete[0:0]
		}
	}
}

//AddTimer 添加timer
func AddTimer(countdown, interval int64, callback func()) (timerId int64) {
	timerId = rand.Int63()
	store.taskList = append(store.taskList, task{timerId, countdown, interval, callback})
	if isTimerRuning == false {
		isTimerRuning = true
		runTimers()
	}
	return timerId
}

//DelTimer 删除timer
func DelTimer(id int64) {
	index := -1
	for key, value := range store.taskList {
		if id == value.id {
			index = key
			break
		}
	}
	if index != -1 {
		store.taskList = append(store.taskList[:index], store.taskList[index+1:]...)
	}
}
