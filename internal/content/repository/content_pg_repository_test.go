package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_Slash/internal/content/mocks"
	"github.com/go-park-mail-ru/2020_2_Slash/internal/models"
	"github.com/stretchr/testify/assert"
)

var countries = []*models.Country{
	&models.Country{
		ID:   1,
		Name: "США",
	},
}

var genres = []*models.Genre{
	&models.Genre{
		Name: "Мультфильм",
	},
	&models.Genre{
		Name: "Комедия",
	},
}

var actors = []*models.Actor{
	&models.Actor{
		Name: "Майк Майерс",
	},
	&models.Actor{
		Name: "Эдди Мёрфи",
	},
}

var directors = []*models.Director{
	&models.Director{
		Name: "Эндрю Адамсон",
	},
	&models.Director{
		Name: "Вики Дженсон",
	},
}

var contentInst *models.Content = &models.Content{
	Name:             "Шрек",
	OriginalName:     "Shrek",
	Description:      "Полная сюрпризов сказка об ужасном болотном огре, который ненароком наводит порядок в Сказочной стране",
	ShortDescription: "Полная сюрпризов сказка об ужасном болотном огре",
	Year:             2001,
	Countries:        countries,
	Genres:           genres,
	Actors:           actors,
	Directors:        directors,
	Type:             "movie",
}

func TestContentPgRepository_DeleteByID_OK(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	contentPgRep := NewContentPgRepository(db)

	mocks.MockContentRepoDeleteReturnResultOk(mock, contentInst.ContentID)
	err = contentPgRep.DeleteByID(contentInst.ContentID)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestContentPgRepository_SelectByID_OK(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var contentInst *models.Content = &models.Content{
		Name:             "Шрек",
		OriginalName:     "Shrek",
		Description:      "Полная сюрпризов сказка об ужасном болотном огре, который ненароком наводит порядок в Сказочной стране",
		ShortDescription: "Полная сюрпризов сказка об ужасном болотном огре",
		Year:             2001,
		Type:             "movie",
	}

	contentPgRep := NewContentPgRepository(db)

	mocks.MockContentRepoSelectByIDReturnRows(mock, contentInst.ContentID, contentInst)
	dbContent, err := contentPgRep.SelectByID(contentInst.ContentID)
	assert.Equal(t, contentInst, dbContent)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestContentPgRepository_SelectById_NoContentWithThisID(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	contentPgRep := NewContentPgRepository(db)

	mocks.MockContentRepoSelectByIDReturnErrNoRows(mock, contentInst.ContentID)
	dbContent, err := contentPgRep.SelectByID(contentInst.ContentID)
	assert.Equal(t, dbContent, (*models.Content)(nil))
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestContentPgRepository_SelectCountries_OK(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	countriesID := []uint64{1}

	contentPgRep := NewContentPgRepository(db)

	mocks.MockContentRepoSelectCountriesReturnRows(mock, contentInst.ContentID, countriesID)
	dbCountires, err := contentPgRep.SelectCountriesByID(contentInst.ContentID)
	assert.Equal(t, countriesID, dbCountires)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestContentPgRepository_SelectDirectors_OK(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	directorsID := []uint64{1, 2}

	contentPgRep := NewContentPgRepository(db)

	mocks.MockContentRepoSelectDirectorsReturnRows(mock, contentInst.ContentID, directorsID)
	dbDirectors, err := contentPgRep.SelectDirectorsByID(contentInst.ContentID)
	assert.Equal(t, directorsID, dbDirectors)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestContentPgRepository_SelectActors_OK(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	actorsID := []uint64{1, 2}

	contentPgRep := NewContentPgRepository(db)

	mocks.MockContentRepoSelectActorsReturnRows(mock, contentInst.ContentID, actorsID)
	dbActors, err := contentPgRep.SelectActorsByID(contentInst.ContentID)
	assert.Equal(t, actorsID, dbActors)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestContentPgRepository_SelectGenres_OK(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	genresID := []uint64{1, 2}

	contentPgRep := NewContentPgRepository(db)

	mocks.MockContentRepoSelectGenresReturnRows(mock, contentInst.ContentID, genresID)
	dbGenres, err := contentPgRep.SelectGenresByID(contentInst.ContentID)
	assert.Equal(t, genresID, dbGenres)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestContentPgRepository_Insert_OK(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	contentPgRep := NewContentPgRepository(db)

	var countries = []*models.Country{
		&models.Country{
			ID:   1,
			Name: "США",
		},
	}
	var genres = []*models.Genre{
		&models.Genre{
			Name: "Мультфильм",
		},
	}
	var actors = []*models.Actor{
		&models.Actor{
			Name: "Майк Майерс",
		},
	}
	var directors = []*models.Director{
		&models.Director{
			Name: "Эндрю Адамсон",
		},
	}
	contentInst := &models.Content{
		Countries: countries,
		Genres:    genres,
		Actors:    actors,
		Directors: directors,
	}

	mocks.MockContentRepoInsertReturnResultOk(mock, contentInst, countries[0],
		genres[0], actors[0], directors[0])
	err = contentPgRep.Insert(contentInst)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestContentPgRepository_Update_OK(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	contentPgRep := NewContentPgRepository(db)

	var countries = []*models.Country{
		&models.Country{
			ID:   1,
			Name: "США",
		},
	}
	var genres = []*models.Genre{
		&models.Genre{
			Name: "Мультфильм",
		},
	}
	var actors = []*models.Actor{
		&models.Actor{
			Name: "Майк Майерс",
		},
	}
	var directors = []*models.Director{
		&models.Director{
			Name: "Эндрю Адамсон",
		},
	}
	contentInst := &models.Content{
		Countries: countries,
		Genres:    genres,
		Actors:    actors,
		Directors: directors,
	}

	mocks.MockContentRepoUpdateReturnResultOk(mock, contentInst, countries[0],
		genres[0], actors[0], directors[0])
	err = contentPgRep.Update(contentInst)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestContentPgRepository_UpdateImages_OK(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	contentPgRep := NewContentPgRepository(db)

	mocks.MockContentRepoUpdateImagesReturnResultOk(mock, contentInst)
	err = contentPgRep.UpdateImages(contentInst)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
