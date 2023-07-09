package user

import (
	"context"
	"strconv"
	"time"

	"github.com/ady243/go-project-chat/utils"
	"github.com/golang-jwt/jwt/v4"
)

const (
	secretKey = "secret"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		Repository: repository,
		timeout:    time.Duration(2) * time.Second,
	}
}
func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}

	return res, nil

}

type MyJwtClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = utils.CheckPassword(req.Password, u.Password)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, MyJwtClaims{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    strconv.Itoa(int(u.ID)),
		},
	})
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	res := &LoginUserRes{
		accesstoken: ss,
		ID:          strconv.Itoa(int(u.ID)),
		Username:    u.Username,
	}

	return res, nil
}

func (s *service) DeleteUser(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	id, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	err = s.Repository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
