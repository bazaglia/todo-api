package models

import (
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"todo/storage"
)

// Task model
type Task struct {
	ID                 string          `json:"id" form:"id"`
	UserID             string          `json:"user_id" form:"-" db:"user_id"`
	Name               string          `json:"name" form:"name" binding:"required"`
	Description        *string         `json:"description,omitempty" form:"description"`
	Location           *string         `json:"location,omitempty" form:"location"`
	Date               *string         `json:"date,omitempty" form:"date"`
	Status             string          `json:"status" form:"-"`
	Labels             *pq.StringArray `json:"labels,omitempty" form:"labels"`
	Comments           *pq.StringArray `json:"comments,omitempty" form:"comments"`
	IsFavorite         bool            `json:"is_favorite,omitempty" form:"is_favorite" db:"is_favorite"`
	AttachmentProvider *string         `json:"-" form:"-" db:"attachment_provider"`
	AttachmentBucket   *string         `json:"-" form:"-" db:"attachment_bucket"`
	AttachmentObject   *string         `json:"-" form:"-" db:"attachment_object"`
	AttachmentFile     []byte          `json:"-" form:"-"`
	AttachmentURL      string          `json:"attachment_url,omitempty"`
}

// TaskFilter items for filtering task
type TaskFilter struct {
	UserID string
}

// TaskRepository operations for persistance layer
type TaskRepository struct {
	db      *sqlx.DB
	storage storage.Adapter
}

func (repository *TaskRepository) uploadAttachment(file *[]byte, data *Task) error {
	provider := repository.storage.GetName()
	bucketName := "attachment-photos"
	object := "attachment-" + strconv.FormatInt(time.Now().Unix(), 10)

	err := repository.storage.CreateObject(bucketName, object, *file)
	if err != nil {
		return err
	}

	data.AttachmentProvider = &provider
	data.AttachmentBucket = &bucketName
	data.AttachmentObject = &object

	data.AttachmentURL, err = repository.storage.GetSignedURL(bucketName, object)
	return err
}

func (repository *TaskRepository) FindAll(filter *TaskFilter) (*[]*Task, error) {
	sql := "SELECT * FROM tasks"
	parameters := map[string]interface{}{}

	if (*filter != TaskFilter{}) {
		sql += " WHERE"

		if filter.UserID != "" {
			sql += " user_id = :userID"
			parameters["userID"] = filter.UserID
		}
	}

	tasks := []*Task{}
	rows, err := repository.db.NamedQuery(sql, parameters)

	for rows.Next() {
		t := Task{}

		if err := rows.StructScan(&t); err != nil {
			return nil, err
		}

		if t.AttachmentObject != nil {
			t.AttachmentURL, err = repository.storage.GetSignedURL(*t.AttachmentBucket, *t.AttachmentObject)

			if err != nil {
				return nil, err
			}
		}

		tasks = append(tasks, &t)
	}

	return &tasks, err
}

func (repository *TaskRepository) Get(id string) (*Task, error) {
	t := &Task{}

	err := repository.db.Get(
		t,
		"SELECT * FROM tasks WHERE id = $1",
		id,
	)

	if err != nil {
		return nil, err
	}

	if t.AttachmentObject != nil {
		t.AttachmentURL, err = repository.storage.GetSignedURL(*t.AttachmentBucket, *t.AttachmentObject)
		if err != nil {
			return nil, err
		}
	}

	return t, err
}

func (repository *TaskRepository) Create(data *Task) (*Task, error) {
	if len(data.AttachmentFile) > 0 {
		if err := repository.uploadAttachment(&data.AttachmentFile, data); err != nil {
			return nil, err
		}
	}

	statement, err := repository.db.PrepareNamed(`
		INSERT INTO tasks (
			user_id, name, description, location, date, status, labels, comments, is_favorite,
			attachment_provider, attachment_bucket, attachment_object
		)
		VALUES (
			:user_id, :name, :description, :location, :date, :status, :labels, :comments, :is_favorite,
			:attachment_provider, :attachment_bucket, :attachment_object
		)
		RETURNING id
	`)

	if err != nil {
		return nil, err
	}

	err = statement.Get(&data.ID, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repository *TaskRepository) Update(data *Task) (*Task, error) {
	var statement = `
		UPDATE tasks SET
			name = :name, location = :location, amount_currency = :amount_currency,
			amount_value = :amount_value, date= :date, is_refundable = :is_refundable,
			payment_method = :payment_method, category = :category, description = :description
		WHERE id = :id
	`

	if len(data.AttachmentFile) > 0 {
		if err := repository.uploadAttachment(&data.AttachmentFile, data); err != nil {
			return nil, err
		}

		statement = `
			UPDATE tasks SET
				name = :name, location = :location, amount_currency = :amount_currency,
				amount_value = :amount_value, date= :date, is_refundable = :is_refundable,
				payment_method = :payment_method, category = :category, description = :description,
				attachment_provider = :attachment_provider, attachment_bucket = :attachment_bucket,
				attachment_object = :attachment_object
			WHERE id = :id
		`
	}

	_, err := repository.db.NamedExec(statement, data)

	return data, err
}

// NewTask instantiates a task setting default fields
func NewTask(userID string) *Task {
	return &Task{
		Status: "open",
		UserID: userID,
	}
}

func NewTaskRepository(db *sqlx.DB, storage storage.Adapter) *TaskRepository {
	return &TaskRepository{db: db, storage: storage}
}
