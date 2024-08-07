package users

import (
	"context"
	"crypto/rand"
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"errors"
	"io"

	"github.com/eduardocfalcao/money-tracker/internal/users/models"
	"github.com/eduardocfalcao/money-tracker/internal/users/repository"
)

const (
	saltSize = 16
)

var ErrNotFound = errors.New("not found in database")

type UsersService interface {
	CreateUser(ctx context.Context, userRequest models.CreateUserRequest) error
	GetUserByEmailAndPassword(ctx context.Context, userInfo models.LoginRequest) (models.User, error)
}

type service struct {
	Repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{
		Repository: repository,
	}
}

func (s *service) CreateUser(ctx context.Context, userRequest models.CreateUserRequest) error {
	salt := generateRandomSalt(saltSize)
	saltString := hex.EncodeToString(salt)
	params := models.User{
		Name:         userRequest.Name,
		Email:        userRequest.Email,
		Passwordhash: hashPassword(userRequest.Password, salt),
		Salt:         saltString,
	}

	return s.Repository.CreateUser(ctx, params)
}

func (s *service) GetUserByEmailAndPassword(ctx context.Context, userInfo models.LoginRequest) (models.User, error) {
	user, err := s.Repository.GetUserByEmail(ctx, userInfo.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, models.ErrUserNotFoundOrWrongPassword
		}
		return models.User{}, err
	}

	if checkPassword(user.Passwordhash, userInfo.Password, []byte(user.Salt)) {
		return user, nil
	}

	return models.User{}, models.ErrUserNotFoundOrWrongPassword
}

func generateRandomSalt(saltSize int) []byte {
	salt := make([]byte, saltSize)

	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		panic(err)
	}

	return salt
}

func hashPassword(password string, salt []byte) string {
	// Convert password string to byte slice
	passwordBytes := []byte(password)

	// Create sha-512 hasher
	sha512Hasher := sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	hashedPasswordBytes := sha512Hasher.Sum(nil)

	// Convert the hashed password to a hex string
	hashedPasswordHex := hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}

func checkPassword(hashedPassword, currPassword string,
	salt []byte,
) bool {
	currPasswordHash := hashPassword(currPassword, salt)

	return hashedPassword == currPasswordHash
}
