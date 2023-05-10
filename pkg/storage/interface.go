package storage


import "tasks_app/pkg/storage/postgressql"


type Interface interface{
	Tasks(int, int) ([]postgressql.Task, error)
	NewTask(postgressql.Task) (int, error)
}

