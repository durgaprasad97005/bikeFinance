package services

import (
	"context"
	"errors"
	"time"

	"github.com/durgaprasad97005/bikeFinance/config"
	"github.com/durgaprasad97005/bikeFinance/dto/request"
	"github.com/durgaprasad97005/bikeFinance/dto/response"
	"github.com/durgaprasad97005/bikeFinance/models"
	"github.com/durgaprasad97005/bikeFinance/repository"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

// struct for AuthService
type AuthService struct {
	userRepository *repository.UserRepository
	cfg *config.Config
}

// Constructor like func to initialize AuthService
func NewAuthService(repo *repository.UserRepository, config *config.Config) *AuthService {
	return &AuthService{
		userRepository: repo,
		cfg: config,
	}
}

// Register service
func (s *AuthService) Register(ctx context.Context, req request.RegisterUser) (*response.User, error) {
	// Validate the req
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, err
	}
	
	// Create models.User model from the req
	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Role:      req.Role,
		Branch:    req.Branch,
		AuditFields: models.AuditFields{
			// CreatedBy: ,
			// LastModifiedBy: ,
			CreatedAt:      time.Now(),
			LastModifiedAt: time.Now(),
		},
	}

	// Check whether given email already exists or not
	existingUser, err := s.userRepository.Get(ctx, bson.M{"email": req.Email})
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("There exists another user with given email")
	}

	// Password hashing
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return nil, err
	}
	user.PasswordHash = string(passwordHash)

	// Calling repository to create new user
	if err := s.userRepository.Create(ctx, &user); err != nil {
		return nil, err
	}

	// Create responseUser
	responseUser := &response.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
		Branch:    user.Branch,
		AuditFields: models.AuditFields{
			CreatedBy:      user.CreatedBy,
			LastModifiedBy: user.LastModifiedBy,
			CreatedAt:      user.CreatedAt,
			LastModifiedAt: user.LastModifiedAt,
		},
	}

	return responseUser, nil
}

// Login service
func (s *AuthService) Login(ctx context.Context, req request.LoginUser) (string, error) {
	// Validate the req
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return "", err
	}

	// Check whether user exists
	user, err := s.userRepository.Get(ctx, bson.M{"email": req.Email})
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("There exists no user with given email")
	}

	// Compare Password and PasswordHash
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", err
	}

	// Create jwt access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.Hex(), 
		"exp": time.Now().Add(15 * time.Minute),
	})

	accessToken, err := token.SignedString([]byte(s.cfg.JwtSecret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
