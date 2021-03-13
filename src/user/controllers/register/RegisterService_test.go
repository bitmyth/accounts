package register

import (
	"errors"
	"github.com/bitmyth/accounts/src/user"
	"testing"
)

type Expect struct {
	Data   interface{}
	Return error
}

type Expects map[int]*Expect

type MockUserRepo struct {
	step    int
	expects Expects
	user    user.User
	users   []user.User
}

func NewMockUserRepo() *MockUserRepo {
	return &MockUserRepo{step: 0}

}

func (ur *MockUserRepo) ShouldReturn(expects Expects) *MockUserRepo {
	ur.expects = expects
	return ur
}

func (ur MockUserRepo) Find(out interface{}, query interface{}, args ...interface{}) error {
	expect := ur.expects[ur.step]
	out = expect.Data
	ur.step++
	return expect.Return
}

func (ur MockUserRepo) First(out interface{}, query interface{}, args ...interface{}) error {
	expect := ur.expects[ur.step]
	out = expect.Data
	ur.step++
	return expect.Return
}

func (ur MockUserRepo) Save(value interface{}) error {
	ur.expects[ur.step].Data = value
	ur.step++
	return nil
}

var mockUserRepo = NewMockUserRepo()

func TestServiceDo(t *testing.T) {
	u := user.User{
		Name:     "fake",
		Password: "password",
	}

	mockUserRepo.ShouldReturn(Expects{
		0: {
			Data:   nil,
			Return: errors.New("found"),
		},
		1: {
			Data:   user.User{Name: "fake"},
			Return: errors.New("found"),
		},
	})

	s := NewService(u, mockUserRepo)

	result, err := s.Do()

	t.Log("registered user:", result)

	if err != nil {
		t.Fatal(err)
	}
}
