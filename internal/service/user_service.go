package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go-user-age-api/db/sqlc"
	"go-user-age-api/internal/models"
	"go-user-age-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return models.UserResponse{}, errors.New("dob must be in YYYY-MM-DD format")
	}

	result, err := s.repo.CreateUser(ctx, sqlc.CreateUserParams{
		Name: req.Name,
		Dob:  dob,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	id, _ := result.LastInsertId()

	return models.UserResponse{
		ID:   int32(id),
		Name: req.Name,
		DOB:  req.DOB,
	}, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int32) (models.UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.UserResponse{}, errors.New("user not found")
		}
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  models.CalculateAge(user.Dob),
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]models.UserResponse, error) {
	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]models.UserResponse, 0)

	for _, user := range users {
		responses = append(responses, models.UserResponse{
			ID:   user.ID,
			Name: user.Name,
			DOB:  user.Dob.Format("2006-01-02"),
			Age:  models.CalculateAge(user.Dob),
		})
	}

	return responses, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return models.UserResponse{}, errors.New("dob must be in YYYY-MM-DD format")
	}

	result, err := s.repo.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: req.Name,
		Dob:  dob,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return models.UserResponse{}, errors.New("user not found")
	}

	return models.UserResponse{
		ID:   id,
		Name: req.Name,
		DOB:  req.DOB,
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	existingUser, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found")
		}
		return err
	}

	_ = existingUser

	return s.repo.DeleteUser(ctx, id)
}

func (s *UserService) ListUsersPaginated(ctx context.Context, page int32, limit int32) ([]models.UserResponse, error) {
	offset := (page - 1) * limit

	users, err := s.repo.ListUsersPaginated(ctx, sqlc.ListUsersPaginatedParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	responses := make([]models.UserResponse, 0)

	for _, user := range users {
		responses = append(responses, models.UserResponse{
			ID:   user.ID,
			Name: user.Name,
			DOB:  user.Dob.Format("2006-01-02"),
			Age:  models.CalculateAge(user.Dob),
		})
	}

	return responses, nil
}