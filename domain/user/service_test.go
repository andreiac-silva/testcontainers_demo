package user

import (
	"context"
	"database/sql"

	"github.com/andreiac-silva/testcontainers_demo/domain/model"
)

func (s *IntegrationTestSuite) TestCreate() {
	userService := NewService(NewRepository(s.Database))

	s.Run("service should create the user successfully", func() {
		// Given
		user := model.User{
			Name:    "John",
			Surname: "Smith",
			Age:     45,
		}

		// When
		id, err := userService.Create(context.Background(), user)

		// Then
		s.NoError(err)
		s.NotNil(id)
	})
}

func (s *IntegrationTestSuite) TestFind() {
	userService := NewService(NewRepository(s.Database))

	s.Run("service should return the user successfully", func() {
		// Given
		user := model.User{
			Name:    "Monica",
			Surname: "Sanchez",
			Age:     33,
		}
		id, err := userService.Create(context.Background(), user)

		// When
		result, err := userService.Get(context.Background(), *id)

		// Then
		s.NoError(err)
		s.Equal(user.Name, result.Name)
		s.Equal(user.Surname, result.Surname)
		s.Equal(user.Age, result.Age)
	})

	s.Run("service should return an error when the user does not exist", func() {
		// When
		_, err := userService.Get(context.Background(), 0)

		// Then
		s.Error(err)
		s.ErrorIs(err, sql.ErrNoRows)
	})
}

func (s *IntegrationTestSuite) TestFindAll() {
	userService := NewService(NewRepository(s.Database))

	s.Run("service should return all users successfully", func() {
		// Given
		user1 := model.User{
			Name:    "Katherine",
			Surname: "McDonald",
			Age:     29,
		}
		user2 := model.User{
			Name:    "George",
			Surname: "Parker",
			Age:     71,
		}
		id1, err := userService.Create(context.Background(), user1)
		s.NoError(err)
		s.NotNil(id1)

		id2, err := userService.Create(context.Background(), user2)
		s.NoError(err)
		s.NotNil(id2)

		// When
		result, err := userService.List(context.Background())

		// Then
		s.NoError(err)

		var resultIDS []int64
		for _, user := range result {
			resultIDS = append(resultIDS, *user.ID)
		}

		s.Contains(resultIDS, *id1)
		s.Contains(resultIDS, *id2)
	})
}
