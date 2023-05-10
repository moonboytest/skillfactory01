package postgressql

import (
	"context"
	//"vendor/golang.org/x/net/idna"

	"github.com/jackc/pgx/v4/pgxpool"
)

//Хранилище данных
type Storage struct {
	db *pgxpool.Pool
}

//Конструктор
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}

	s := Storage{
		db: db,
	}

	return &s, nil
}

//Задача
type Task struct {
	Id         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

//Возвращает список задач из БД
func (s *Storage) Tasks(taskID, authorID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			opened,
			close,
			author_id,
			assigned_id,
			title,
			content
		FROM tasks
		WHERE
		($1 = 0 OR id = $1 ) AND
		($2 = 0 OR author_id = $2)
		ORDER BY id;
		`,
		taskID,
		authorID)
	if err != nil {
		return nil, err
	}

	var tasks []Task

	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.Id,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}
	return tasks, rows.Err()

}

// Создает новую задачу и возвращает ее id
func (s *Storage) NewTask(t Task) (int, error) {
	var id int

	err := s.db.QueryRow(context.Background(), `
	
	INSERT INTO tasks(title, content) 
	VALUES ($1, $2) RETURNING id;
	`,
		t.Title,
		t.Content,
	).Scan(&id)

	return id, err
}

//Возвращает список задач из БД по автору
func (s *Storage) TasksByAuthor(authorID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			opened,
			close,
			author_id,
			assigned_id,
			title,
			content
		FROM tasks
		WHERE
		id = $1 AND
		ORDER BY id;
		`,
		authorID)
	if err != nil {
		return nil, err
	}

	var tasks []Task

	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.Id,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}
	return tasks, rows.Err()

}

//Получает список задач по метке
func (s *Storage) TasksByLable(lableID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
	SELECT 
	id,
	opened,
	close,
	author_id,
	assigned_id,
	title,
	content
FROM tasks
WHERE id = 
	(SELECT task_id
	FROM tasks_lables 
	WHERE label_id = $1 )
ORDER BY id;
`,
		lableID)
	if err != nil {
		return nil, err
	}

	var tasks []Task

	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.Id,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	return tasks, rows.Err()

}

//Обновляет название задачи по id
func (s *Storage) UpdateTaskTitle(taskID int, name string) error {
	_, err := s.db.Exec(context.Background(),`
	UPDATE tasks
	SET title = $2
	WHERE id = $1
	`,
		taskID, 
		name,
	)
	if err != nil{
		return err
	}

	return nil
}

//Обновляет текст задачи по id
func (s *Storage) UpdateTaskContent(taskID int, content string) error {
	_, err := s.db.Exec(context.Background(),`
	UPDATE tasks
	SET content = $2
	WHERE id = $1
	`,
		taskID, 
		content,
)
	if err != nil{
		return err
	}
	return nil
}

//Обновляет исполнителя задачи по id
func (s *Storage) UpdateTaskAssigned(taskID int, assigned string) error {
	_, err := s.db.Exec(context.Background(),`
	UPDATE tasks
	SET assigned = $2
	WHERE id = $1
	`,
	taskID,
	assigned,
	)
	if err != nil{
		return err
	}
	return nil
}

//Удаляет задачу по id
func (s *Storage) DeleteTask(taskID int) error {
	_, err := s.db.Exec(context.Background(), `
	DELETE FROM tasks
	WHERE id =$
	`,
		taskID)
	return err
}
