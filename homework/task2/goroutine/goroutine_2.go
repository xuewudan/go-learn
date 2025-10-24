package main

import (
	"fmt"
	"time"
)

// Task 定义任务类型，接收一个字符串参数并返回字符串结果
type Task func(string) string

// TaskResult 存储任务执行结果和统计信息
type TaskResult struct {
	TaskName      string
	Input         string
	Output        string
	ExecutionTime time.Duration
	Err           error
}

// Scheduler 任务调度器
type Scheduler struct {
	tasks []struct {
		task  Task
		name  string
		input string
	}
}

// NewScheduler 创建新的任务调度器
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make([]struct {
			task  Task
			name  string
			input string
		}, 0),
	}
}

// AddTask 向调度器添加任务
func (s *Scheduler) AddTask(name string, task Task, input string) {
	s.tasks = append(s.tasks, struct {
		task  Task
		name  string
		input string
	}{task, name, input})
}

// Run 并发执行所有任务并返回结果
func (s *Scheduler) Run() []TaskResult {
	results := make([]TaskResult, len(s.tasks))
	resultChan := make(chan TaskResult, len(s.tasks))

	// 启动所有任务的协程
	for _, t := range s.tasks {
		go func(task Task, name, input string) {
			startTime := time.Now()
			var result TaskResult

			defer func() {
				// 捕获可能的panic
				if r := recover(); r != nil {
					result.Err = fmt.Errorf("task panicked: %v", r)
				}
				resultChan <- result
			}()

			output := task(input)
			result = TaskResult{
				TaskName:      name,
				Input:         input,
				Output:        output,
				ExecutionTime: time.Since(startTime),
				Err:           nil,
			}
		}(t.task, t.name, t.input)
	}

	// 收集所有结果
	for i := 0; i < len(s.tasks); i++ {
		results[i] = <-resultChan
	}

	return results
}

func main() {
	// 创建调度器
	scheduler := NewScheduler()

	// 添加示例任务
	scheduler.AddTask("任务1", func(input string) string {
		time.Sleep(100 * time.Millisecond) // 模拟任务执行时间
		return fmt.Sprintf("处理完成: %s", input)
	}, "数据A")

	scheduler.AddTask("任务2", func(input string) string {
		time.Sleep(200 * time.Millisecond)
		return fmt.Sprintf("处理完成: %s", input)
	}, "数据B")

	scheduler.AddTask("任务3", func(input string) string {
		time.Sleep(150 * time.Millisecond)
		return fmt.Sprintf("处理完成: %s", input)
	}, "数据C")

	// 执行所有任务并获取结果
	start := time.Now()
	results := scheduler.Run()
	totalTime := time.Since(start)

	// 打印结果
	fmt.Printf("所有任务执行完成，总耗时: %v\n\n", totalTime)
	for _, res := range results {
		if res.Err != nil {
			fmt.Printf("任务 %s 执行失败: %v\n", res.TaskName, res.Err)
			continue
		}
		fmt.Printf("任务 %s:\n", res.TaskName)
		fmt.Printf("  输入: %s\n", res.Input)
		fmt.Printf("  输出: %s\n", res.Output)
		fmt.Printf("  执行时间: %v\n\n", res.ExecutionTime)
	}
}
