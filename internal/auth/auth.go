package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	pb "taxiya/proto/auth"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID        string
	Email     string
	Password  string
	Name      string
	Phone     string
	CreatedAt time.Time
}

type AuthService struct {
	users map[string]*User
}

func NewAuthService() *AuthService {
	return &AuthService{
		users: make(map[string]*User),
	}
}

func (s *AuthService) hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func (s *AuthService) generateToken(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte("your-secret-key")) // En producción, usar una clave segura
}

func (s *AuthService) Register(email, password, name, phone string) (*pb.RegisterResponse, error) {
	if _, exists := s.users[email]; exists {
		return nil, errors.New("user already exists")
	}

	user := &User{
		ID:        generateID(), // Implementar esta función
		Email:     email,
		Password:  s.hashPassword(password),
		Name:      name,
		Phone:     phone,
		CreatedAt: time.Now(),
	}

	s.users[email] = user

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Token: token,
		User: &pb.User{
			Id:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			Phone:     user.Phone,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
		},
	}, nil
}

func (s *AuthService) Login(email, password string) (*pb.LoginResponse, error) {
	user, exists := s.users[email]
	if !exists {
		return nil, errors.New("user not found")
	}

	if user.Password != s.hashPassword(password) {
		return nil, errors.New("invalid password")
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token: token,
		User: &pb.User{
			Id:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			Phone:     user.Phone,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
		},
	}, nil
}

func generateID() string {
	// Implementar generación de ID único
	return "user-" + time.Now().Format("20060102150405")
}
