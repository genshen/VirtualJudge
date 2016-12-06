package status

import (
	"time"
	"log"
	"container/ring"
	"gensh.me/VirtualJudge/components/crawler/accounts"
)

type Task struct {
	StatusCode        int8
	RunId             string
	OjType            int8
	AccountIndex      uint
	QueryCount        uint
	LocalSubmissionId int
}

type TaskResult struct {
	entity      *ring.Ring
	LocalRunId  int
	RunId       string
	StatusCode  int8
	Memory      string
	ExecuteTime string
}

var tasks  chan Task
var tasksResult  chan TaskResult
var timer *time.Timer
var tasksRing *ring.Ring
var taskCount = 0
var onStatusChangeCallback  func(*TaskResult)

func InitializeStatusTasks(callback func(*TaskResult)) {
	tasks = make(chan Task, 10)  //init chan
	tasksResult = make(chan TaskResult, 10)
	timer = time.NewTimer(2 * time.Second)
	onStatusChangeCallback = callback
	go runTasks()
}

func runTasks() {
	for {
		select {
		case task := <-tasks:
		// add a new task to taskRing
			if taskCount == 0 || tasksRing == nil {
				tasksRing = ring.New(1)
				tasksRing.Value = &task
				taskCount = 1
			} else {
				tasksRing.Link(&ring.Ring{Value:&task})
				taskCount++
			}
		case <-timer.C:
			log.Print("time over", tasksRing == nil)
			if tasksRing != nil {
				//use Do
				//todo select a task and fetch data from remote oj
				go LaunchTask(*tasksRing.Value.(*Task), tasksRing) //copy value
			}
			timer = time.NewTimer(4 * time.Second)
		case result := <-tasksResult:
			entity := result.entity
			if taskCount != 0 && entity != nil {
				taskCount--
				if taskCount == 0 {
					tasksRing = nil
				} else if (tasksRing != nil && entity == tasksRing) {
					tasksRing = tasksRing.Next()
				}
				entity.Prev().Unlink(1)
				entity = nil
				onStatusChangeCallback(&result)
			} //else, has unknown error
		}
	}
}

func LaunchTask(task Task, entity *ring.Ring) {
	ojType := int(task.OjType)
	accountInterface, err := accounts.GetInterfaceByOjType(ojType)
	if err != nil {
		//todo error time ++,query time ++
		return
	}

	if index := ojType - 1; index < len(statusInterfaces) && index >= 0 {
		si := statusInterfaces[index]
		if result, err := si.FetchStatus(accountInterface, si, task.AccountIndex, task.RunId); err != nil {
			log.Println(err.Error())
		} else {
			result.entity = entity
			result.LocalRunId = task.LocalSubmissionId
			result.RunId = task.RunId
			tasksResult <- *result //todo result type
		}
	}

}

func AddTaskToQueue(localSubmissionId int, ojType int8, accountIndex uint, statusCode int8, runId string) {
	tasks <- Task{StatusCode:statusCode, RunId:runId, LocalSubmissionId:localSubmissionId, QueryCount:0,
		OjType:ojType, AccountIndex:accountIndex }
}