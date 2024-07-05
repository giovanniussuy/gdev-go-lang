package ports

import (
	"github.com/giovanniussuy/gdev-go-lang/hexagonal/core/domain"
)

type TodoRepository interface {
	Get(id string) (*domain.EntidadeImportante, error)
	List() ([]domain.EntidadeImportante, error)
	Create(todo *domain.EntidadeImportante) (*domain.EntidadeImportante, error)
}
