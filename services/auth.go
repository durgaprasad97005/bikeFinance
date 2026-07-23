package services

import (
	"errors"
	"time"

	"github.com/durgaprasad97005/bikeFinance/constants"
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
	jwtSecret      string
}

// Constructor like func to initialize AuthService
func NewAuthService(repo *repository.UserRepository, secret string) *AuthService {
	return &AuthService{
		userRepository: repo,
		jwtSecret:      secret,
	}
}

// Register service
func (s *AuthService) Register(req request.RegisterUser) (*response.User, error) {
	// Validate the req
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, err
	}
	if req.Role < 1 || req.Role >= constants.RoleCount {
		return nil, errors.New("Invalid role")
	}
	if req.Branch < 1 || req.Branch >= constants.BranchCount {
		return nil, errors.New("Invalid branch")
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
	existingUser, err := s.userRepository.Get(bson.M{"email": req.Email})
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("Duplicate user")
	}

	// Password hashing
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return nil, err
	}
	user.PasswordHash = string(passwordHash)

	// Calling repository to create new user
	if err := s.userRepository.Create(&user); err != nil {
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
func (s *AuthService) Login(req request.LoginUser) (string, error) {
	// Validate the req
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return "", err
	}

	// Check whether user exists
	user, err := s.userRepository.Get(bson.M{"email": req.Email})
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("User not found")
	}

	// Compare Password and PasswordHash
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", err
	}

	// Create jwt access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.Hex(),
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	})

	accessToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	// Return accessToken
	return accessToken, nil
}

// Profile service
func (s *AuthService) Profile(id string) (*response.User, error) {
	// Converting id to userId object
	userId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Get user
	user, err := s.userRepository.Get(bson.M{"_id": userId})
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

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
