package testing

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/repository"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UsersRepositoryTestSuite struct {
	suite.Suite
	mockResource *sqlx.DB
	mock         sqlmock.Sqlmock
}

func (suite *UsersRepositoryTestSuite) SetupTest() {
	mockdb, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	sqlxDB := sqlx.NewDb(mockdb, "sqlmock")
	suite.mockResource = sqlxDB
	suite.mock = mock
}

func (suite *UsersRepositoryTestSuite) TearDownTest() {
	suite.mockResource.Close()
}

func (suite *UsersRepositoryTestSuite) TestUsersRepository_Insert() {
	// Prepare the user data for insertion.
	user := models.Users{
		Id:       "24912488124",
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
		Balance:  100.0,
	}

	// Configure mock to expect the insert query.
	suite.mock.ExpectExec("^INSERT INTO users").
		WithArgs(user.Id, user.Name, user.Email, user.Password, user.Balance).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the insert method.
	repo := repository.NewUsersRepository(suite.mockResource)
	err := repo.Insert(&user)
	assert.Nil(suite.T(), err)
}
func TestUsersRepository_GetById(t *testing.T) {

}

func TestStudentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UsersRepositoryTestSuite))
}
