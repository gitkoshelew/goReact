package store_test

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	teardown()
	defer teardown()
	store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.User {
				testStore.Open()

				user := model.TestUser()
				return user
			},
			isValid: true,
		},
		{
			name: "email in use",
			model: func() *model.User {
				testStore.Open()

				user := model.TestUser()
				user.Email = "new@mail.org"
				testStore.User().Create(user)

				return user
			},
			isValid: false,
		},
		{
			name: "DB closed",
			model: func() *model.User {
				testStore.Close()
				return model.TestUser()
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.User().Create(tc.model())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.User().Create(tc.model())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestUserRepository_GetAll(t *testing.T) {
	teardown()
	defer teardown()
	store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		s       func() *store.Store
		isValid bool
	}{
		{
			name: "valid",
			s: func() *store.Store {
				testStore.Open()
				return testStore
			},
			isValid: true,
		},
		{
			name: "DB closed",
			s: func() *store.Store {
				st := testStore
				st.Close()
				return st
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := tc.s().User().GetAll()
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := tc.s().User().GetAll()
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestUserRepository_FindByID(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		id      func() int
		isValid bool
	}{
		{
			name: "valid",
			id: func() int {
				testStore.Open()
				return id.User
			},
			isValid: true,
		},
		{
			name: "invalid ID",
			id: func() int {
				testStore.Open()
				return 0
			},
			isValid: false,
		},
		{
			name: "DB closed",
			id: func() int {
				testStore.Close()
				return id.User
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.User().FindByID(tc.id())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.User().FindByID(tc.id())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestUserRepository_Delete(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		id      func() int
		isValid bool
	}{
		{
			name: "valid",
			id: func() int {
				testStore.Open()
				user := model.TestUser()
				id, _ := testStore.User().Create(user)
				return *id
			},
			isValid: true,
		},
		{
			name: "Invalid ID",
			id: func() int {
				testStore.Open()
				return 0
			},
			isValid: false,
		},
		{
			name: "DB closed",
			id: func() int {
				testStore.Close()
				return id.User
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.User().Delete(tc.id())
				assert.NoError(t, err)
				testStore.Close()
			} else {
				err := testStore.User().Delete(tc.id())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}

func TestUserRepository_Update(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.User {
				testStore.Open()

				user := model.TestUser()
				user.UserID = id.User

				return user
			},
			isValid: true,
		},
		{
			name: "invalid ID",
			model: func() *model.User {
				testStore.Open()

				user := model.TestUser()
				user.UserID = 0

				return user
			},
			isValid: false,
		},
		{
			name: "DB closed",
			model: func() *model.User {
				testStore.Close()

				user := model.TestUser()
				user.UserID = id.User

				return user
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.User().Update(tc.model())
				testStore.Close()
				assert.NoError(t, err)
			} else {
				err := testStore.User().Update(tc.model())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}

func TestUserRepository_ModelFromDTO(t *testing.T) {
	teardown()
	defer teardown()
	store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.UserDTO
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.UserDTO {
				testStore.Open()

				user := model.TestUserDTO()

				return user
			},
			isValid: true,
		},
		{
			name: "DB closed",
			model: func() *model.UserDTO {
				testStore.Close()

				user := model.TestUserDTO()

				return user
			},
			isValid: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.User().ModelFromDTO(tc.model())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.User().ModelFromDTO(tc.model())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestUserRepository_FindByEmail(t *testing.T) {
	teardown()
	defer teardown()
	store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		email   func() string
		isValid bool
	}{
		{
			name: "valid",
			email: func() string {
				testStore.Open()

				user := model.TestUser()
				user.Email = "searching@email.org"
				testStore.User().Create(user)

				return user.Email
			},
			isValid: true,
		},
		{
			name: "invalid Email",
			email: func() string {
				testStore.Open()

				user := model.TestUser()
				user.Email = "searching@email.org"
				testStore.User().Create(user)

				return "notThis@email.org"
			},
			isValid: false,
		},
		{
			name: "DB closed",
			email: func() string {
				testStore.Close()

				user := model.TestUser()
				user.Email = "searching@email.org"
				testStore.User().Create(user)

				return user.Email
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.User().FindByEmail(tc.email())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.User().FindByEmail(tc.email())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestUserRepository_VerifyEmail(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		id      func() int
		isValid bool
	}{
		{
			name: "valid",
			id: func() int {
				testStore.Open()

				return id.User
			},
			isValid: true,
		},
		{
			name: "invalid ID",
			id: func() int {
				testStore.Open()

				return -1
			},
			isValid: false,
		},
		{
			name: "DB closed",
			id: func() int {
				testStore.Close()

				return id.User
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.User().VerifyEmail(tc.id())
				testStore.Close()
				assert.NoError(t, err)
			} else {
				err := testStore.User().VerifyEmail(tc.id())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}

func TestUserRepository_EmailCheck(t *testing.T) {
	teardown()
	defer teardown()
	store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		email   func() string
		isValid bool
		isExist bool
	}{
		{
			name: "Email in use",
			email: func() string {
				testStore.Open()

				user := model.TestUser()
				user.Email = "inusing@email.org"
				testStore.User().Create(user)

				return user.Email
			},
			isValid: true,
			isExist: true,
		},
		{
			name: "Email is not in use",
			email: func() string {
				testStore.Open()

				return "unusing@email.org"
			},
			isValid: true,
			isExist: false,
		},
		{
			name: "DB closed",
			email: func() string {
				testStore.Close()

				return "unusing@email.org"
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.User().EmailCheck(tc.email())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
				if tc.isExist {
					assert.True(t, *result)
				} else {
					assert.False(t, *result)
				}

			} else {
				result, err := testStore.User().EmailCheck(tc.email())
				testStore.Close()
				assert.Error(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestUserRepository_PasswordChange(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.User {
				testStore.Open()

				user := model.TestUser()
				user.UserID = id.User
				user.Password = "NewPassword"

				return user
			},
			isValid: true,
		},
		{
			name: "invalid ID",
			model: func() *model.User {
				testStore.Open()

				user := model.TestUser()
				user.UserID = 0
				user.Password = "NewPassword"

				return user
			},
			isValid: false,
		},
		{
			name: "DB closed",
			model: func() *model.User {
				testStore.Close()

				user := model.TestUser()
				user.UserID = id.User
				user.Password = "NewPassword"

				return user
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.User().PasswordChange(tc.model())
				testStore.Close()
				assert.NoError(t, err)
			} else {
				err := testStore.User().PasswordChange(tc.model())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}
