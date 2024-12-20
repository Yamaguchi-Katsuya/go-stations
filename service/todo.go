package service

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/TechBowl-japan/go-stations/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	newTodo, err := s.db.ExecContext(ctx, insert, subject, description)
	if err != nil {
		return nil, err
	}
	id, err := newTodo.LastInsertId()
	if err != nil {
		return nil, err
	}

	todo := &model.TODO{}
	err = s.db.QueryRowContext(ctx, confirm, id).Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)

	var rows *sql.Rows
	var err error
	if prevID == 0 {
		rows, err = s.db.QueryContext(ctx, read, size)
		if err != nil {
			return nil, err
		}
	} else {
		rows, err = s.db.QueryContext(ctx, readWithID, prevID, size)
		if err != nil {
			return nil, err
		}
	}

	defer rows.Close()

	todos := []*model.TODO{}
	for rows.Next() {
		todo := &model.TODO{}
		err := rows.Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	updated, err := s.db.ExecContext(ctx, update, subject, description, id)
	if err != nil {
		return nil, err
	}

	rows, err := updated.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return nil, &model.ErrNotFound{}
	}

	todo := &model.TODO{}
	err = s.db.QueryRowContext(ctx, confirm, id).Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	if len(ids) == 0 {
		return fmt.Errorf("ids is empty")
	}

	idList := []interface{}{}
	for _, id := range ids {
		idList = append(idList, id)
	}
	deleted, err := s.db.ExecContext(ctx, fmt.Sprintf(deleteFmt, strings.Repeat(", ?", len(ids)-1)), idList...)
	if err != nil {
		return err
	}

	rows, err := deleted.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return &model.ErrNotFound{}
	}

	return nil
}
