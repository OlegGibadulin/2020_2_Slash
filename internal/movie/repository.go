package movie

import (
	"github.com/go-park-mail-ru/2020_2_Slash/internal/models"
)

type MovieRepository interface {
	Insert(movie *models.Movie) error
	Update(movie *models.Movie) error
	DeleteByID(movieID uint64) error
	SelectByID(movieID uint64) (*models.Movie, error)
	SelectFullByID(movieID uint64, curUserID uint64) (*models.Movie, error)
	SelectByContentID(contentID uint64) (*models.Movie, error)
	SelectByParams(params *models.ContentFilter, pgnt *models.Pagination,
		curUserID uint64) ([]*models.Movie, error)
	SelectLatest(pgnt *models.Pagination, curUserID uint64) ([]*models.Movie, error)
	SelectByRating(pgnt *models.Pagination, curUserID uint64) ([]*models.Movie, error)
	SelectWhereNameLike(curUserID uint64, name string, limit, offset uint64) ([]*models.Movie, error)
}
