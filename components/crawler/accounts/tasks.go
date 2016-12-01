package accounts

import "sync"

var tasks = struct {
	sync.RWMutex
	m []uint
}{}

func initTasks(length uint) {
	tasks.m = make([]uint, length)
}

func GetTaskByIndex(index uint) uint {
	tasks.RLock()
	task := tasks.m[index]
	tasks.RUnlock()
	return task
}

func UpdateTaskByIndex(index uint, session uint) {
	tasks.Lock()
	tasks.m[index] = session
	tasks.RUnlock()
}

