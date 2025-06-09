package userservice_test


import (
    "auth-service/internal/model/user"
    "auth-service/internal/service"
    "errors"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type MockRepo struct {
    mock.Mock
}

func (m *MockRepo) GetUsers() []*user.User {
    args := m.Called()
    return args.Get(0).([]*user.User)
}

func (m *MockRepo) GetUserById(id int) *user.User {
    args := m.Called(id)
    if args.Get(0) == nil {
        return nil
    }
    return args.Get(0).(*user.User)
}

func (m *MockRepo) UpdateUser(id int, data *user.User) {
    m.Called(id, data)
}

func (m *MockRepo) DeleteUser(id int) {
    m.Called(id)
}

func (m *MockRepo) Save(u user.User) {
    m.Called(u)
}

func (m *MockRepo) SaveUserInFile() {}
func (m *MockRepo) Restore()        {}

func TestCreateUser(t *testing.T) {
    mockRepo := new(MockRepo)
    svc := userservice.New(mockRepo)

    valid := &user.User{Id: 1, Name: "TestUser"}
    mockRepo.On("Save", mock.AnythingOfType("user.User")).Return()

    t.Run("valid create", func(t *testing.T) {
        res, err := svc.CreateUser(valid)
        assert.NoError(t, err)
        assert.Equal(t, "TestUser", res.Name)
    })

    t.Run("invalid (empty name)", func(t *testing.T) {
        _, err := svc.CreateUser(&user.User{})
        assert.True(t, errors.Is(err, userservice.ErrInvalidUser))
    })
}

func TestGetUserByID(t *testing.T) {
    mockRepo := new(MockRepo)
    svc := userservice.New(mockRepo)

    u1 := &user.User{Id: 1, Name: "TestUser1"}
    mockRepo.On("GetUserById", 1).Return(u1)
    mockRepo.On("GetUserById", 99).Return(nil)

    t.Run("user found", func(t *testing.T) {
        u, err := svc.GetUserByID(1)
        assert.NoError(t, err)
        assert.Equal(t, "TestUser1", u.Name)
    })

    t.Run("user not found", func(t *testing.T) {
        _, err := svc.GetUserByID(99)
        assert.ErrorIs(t, err, userservice.ErrUserNotFound)
    })
}

func TestUpdateUser(t *testing.T) {
    mockRepo := new(MockRepo)
    svc := userservice.New(mockRepo)

    existing := &user.User{Id: 1, Name: "OldName"}
    updated := &user.User{Name: "NewName"}

    mockRepo.On("GetUserById", 1).Return(existing)
    mockRepo.On("UpdateUser", 1, updated).Return()
    mockRepo.On("GetUserById", 99).Return(nil)

    t.Run("update ok", func(t *testing.T) {
        err := svc.UpdateUser(1, updated)
        assert.NoError(t, err)
    })

    t.Run("update not found", func(t *testing.T) {
        err := svc.UpdateUser(99, updated)
        assert.ErrorIs(t, err, userservice.ErrUserNotFound)
    })
}

func TestDeleteUser(t *testing.T) {
    mockRepo := new(MockRepo)
    svc := userservice.New(mockRepo)

    mockRepo.On("GetUserById", 1).Return(&user.User{Id: 1})
    mockRepo.On("DeleteUser", 1).Return()
    mockRepo.On("GetUserById", 99).Return(nil)

    t.Run("delete ok", func(t *testing.T) {
        err := svc.DeleteUser(1)
        assert.NoError(t, err)
    })

    t.Run("delete not found", func(t *testing.T) {
        err := svc.DeleteUser(99)
        assert.ErrorIs(t, err, userservice.ErrUserNotFound)
    })
}

func TestGetUsers(t *testing.T) {
    mockRepo := new(MockRepo)
    svc := userservice.New(mockRepo)

    users := []*user.User{
        {Id: 1, Name: "User1"},
        {Id: 2, Name: "User2"},
    }

    mockRepo.On("GetUsers").Return(users)

    t.Run("get all users", func(t *testing.T) {
        res, err := svc.GetUsers()
        assert.NoError(t, err)
        assert.Len(t, res, 2)
    })
}