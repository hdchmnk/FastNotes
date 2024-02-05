package user

import (
	"FastNotes/pkg"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

const secretKey = "secretkey"

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{repository, time.Duration(2) * time.Second}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := pkg.HashPassword(req.Password)
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
		ID:       strconv.Itoa(int(r.Id)),
		Username: r.Username,
		Email:    r.Email,
	}

	return res, nil
}

type MyJWTClaims struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserRes{}, err
	}

	err = pkg.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginUserRes{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID: strconv.Itoa(int(u.Id)), Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.Id)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		}})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginUserRes{}, err
	}

	return &LoginUserRes{accessToken: ss, Username: u.Username, ID: strconv.Itoa(int(u.Id))}, nil
}
