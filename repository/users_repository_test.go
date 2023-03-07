package repository

import (
	"errors"
	"kel1-stockbite-projects/models"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyUser =  models.Users{
	Id:       "24912488124",
	Name:     "John Doe",
	Email:    "john.doe@example.com",
	Password: "password123",
	Balance:  100.0,
}

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

func (suite *UsersRepositoryTestSuite) TestGetUserBalance_Success() {
	balance := sqlmock.NewRows([]string{"balance"})
	balance.AddRow(dummyUser.Balance)

	
	suite.mock.ExpectQuery("SELECT balance FROM users WHERE id").WillReturnRows(balance)

	repo := NewUsersRepository(suite.mockResource)

	actual, err := repo.GetUserBalance(dummyUser.Id)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), actual)

}

func (suite *UsersRepositoryTestSuite) TestGetUserBalance_Failed() {
	balance := sqlmock.NewRows([]string{"NotBalance"})
	balance.AddRow(dummyUser.Balance)

	
	suite.mock.ExpectQuery("SELECT balance FROM users WHERE id").WillReturnError(errors.New("failed"))
	repo := NewUsersRepository(suite.mockResource)



	actual, err := repo.GetUserBalance(dummyUser.Id)

	func ()  {

		defer func(){
			if r := recover(); r == nil{
				assert.Error(suite.T(), err)
			}
		}()
		repo.GetUserBalance(dummyUser.Id)
	}()

	assert.NotEqual(suite.T(), dummyUser.Balance, actual)
	assert.Error(suite.T(), err)

}

func (suite *UsersRepositoryTestSuite) TestUpdateUserBalance_Success() {

	suite.mock.ExpectExec("UPDATE users SET balance=\\$1 WHERE id=\\$2").WithArgs(int(dummyUser.Balance), dummyUser.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewUsersRepository(suite.mockResource)

	 err := repo.UpdateUserBalance(int(dummyUser.Balance), dummyUser.Id)

	assert.Nil(suite.T(), err)

}

func (suite *UsersRepositoryTestSuite) TestUpdateUserBalance_Failed() {

	suite.mock.ExpectExec("UPDATE users SET balance=\\$1 WHERE id=\\$2").WithArgs(int(dummyUser.Balance), dummyUser.Id).WillReturnError(errors.New("failed"))

	repo := NewUsersRepository(suite.mockResource)

	 err := repo.UpdateUserBalance(int(dummyUser.Balance), dummyUser.Id)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errors.New("failed"), err)

}



func (suite *UsersRepositoryTestSuite) TestUsersRepositoryInsert_Success() {
	// Prepare the user data for insertion.
	user :=	dummyUser

	// Configure mock to expect the insert query.
	suite.mock.ExpectExec("^INSERT INTO users").
		WithArgs(user.Id, user.Name, user.Email, user.Password, user.Balance).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the insert method.
	repo := NewUsersRepository(suite.mockResource)
	err := repo.Insert(&user)
	assert.Nil(suite.T(), err)
}
func (suite *UsersRepositoryTestSuite) TestUsersRepositoryInsert_Failed() {
	user :=	dummyUser

	suite.mock.ExpectExec("^INSERT INTO users").
	WithArgs(user.Id, user.Name, user.Email, user.Password, user.Balance).
	WillReturnError(errors.New("failedAddUser"))
	repo := NewUsersRepository(suite.mockResource)
	err := repo.Insert(&user)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errors.New("failedAddUser"), err)


}
func (suite *UsersRepositoryTestSuite) TestUsersRepositoryValidateUserLogin_Success() {
 valid := sqlmock.NewRows([]string{"EXISTS"})
valid.AddRow(true)

suite.mock.ExpectQuery(regexp.QuoteMeta( `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1 AND password=$2)`)).WillReturnRows(valid)

repo := NewUsersRepository(suite.mockResource)

err, actual := repo.ValidateUserLogin(dummyUser.Email, dummyUser.Password)

assert.Nil(suite.T(), err)
assert.NotNil(suite.T(), actual)

}

func (suite *UsersRepositoryTestSuite) TestUsersRepositoryValidateUserLogin_Failed() {
	valid := sqlmock.NewRows([]string{"EXISTS"})
   valid.AddRow(true)
   
   suite.mock.ExpectQuery(regexp.QuoteMeta( `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1 AND password=$2)`)).WillReturnError(errors.New("failedUserValidation"))
   
   repo := NewUsersRepository(suite.mockResource)
   
   err, actual := repo.ValidateUserLogin(dummyUser.Email, dummyUser.Password)

   func ()  {
	defer func(){

		if r := recover(); r == nil {
			assert.Error(suite.T(), err)
		}

	}()

	repo.ValidateUserLogin(dummyUser.Email, dummyUser.Password)
	
   }()
   
   assert.NotEqual(suite.T(), valid, actual)
   assert.Error(suite.T(), err)
   
   }

func TestStudentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UsersRepositoryTestSuite))
}
