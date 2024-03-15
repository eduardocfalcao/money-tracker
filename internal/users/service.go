package users

import (
	"context"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"

	"github.com/eduardocfalcao/money-tracker/database/queries"
	"github.com/eduardocfalcao/money-tracker/internal/users/models"
)

const (
	saltSize = 16
)

type UsersService interface {
	CreateUser(ctx context.Context, userRequest models.CreateUserRequest) error
}

type service struct {
	Repository queries.QuerierTx
}

func NewService(repository queries.QuerierTx) *service {
	return &service{
		Repository: repository,
	}
}

func (s *service) CreateUser(ctx context.Context, userRequest models.CreateUserRequest) error {
	salt := generateRandomSalt(saltSize)
	params := queries.CreateUserParams{
		Name:         userRequest.Name,
		Email:        userRequest.Email,
		Passwordhash: hashPassword(userRequest.Password, salt),
		Salt:         string(salt),
	}

	return s.Repository.CreateUser(ctx, params)
}

func generateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

func hashPassword(password string, salt []byte) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a hex string
	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}

func checkPassword(hashedPassword, currPassword string,
	salt []byte) bool {
	var currPasswordHash = hashPassword(currPassword, salt)

	return hashedPassword == currPasswordHash
}
