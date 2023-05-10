package memdb


import "tasks_app/pkg/storage/postgressql"


// Заглушка
type DB []postgressql.Task

//Выполнение контракта интерфейса storage.Interface

func (db DB) Tasks(int, int) ([]postgressql.Tasks, error){
	return db, nil
}
func (db DB) NewTask(postgressql.Task) (int, error){
	return 0, nil
}

