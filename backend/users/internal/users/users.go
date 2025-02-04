package users

import (
	"context"
	"errors"
	"fmt"
	"main/internal/jwt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/volgaIt/packages/errorx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int    `json:"id" db:"id"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	Roles     []Role `json:"roles" db:"roles"`
}

type Role string

const (
	Admin   Role = "admin"
	Teacher Role = "teacher"
	Student Role = "student"
)

func StringToRole(s string) Role {
	switch s {
	case "admin":
		return Admin
	case "teacher":
		return Teacher
	default:
		return Student
	}
}

type Service struct {
	repo  *Repo
	cost  int
	token jwt.TokenManager
	db    *pgxpool.Pool
}

func NewService(
	repo *Repo,
	cost int,
	tokenGenerator jwt.TokenManager,
	db *pgxpool.Pool,
) *Service {
	return &Service{repo, cost, tokenGenerator, db}
}

func (s *Service) SignUp(
	ctx context.Context,
	username string,
	password string,
	firstName string,
	lastName string,
	roles []Role,
) (int, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return 0, err
	}
	if user.Id != 0 {
		return 0, errorx.BadRequest(fmt.Sprintf(`username '%s' already taken`, username))
	}

	pwdHash, err := s.hashPassword(password)
	if err != nil {
		return 0, err
	}

	return s.repo.InsertUser(ctx, username, pwdHash, roles, firstName, lastName)
}

func (s *Service) hashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), s.cost)
	if errors.Is(err, bcrypt.ErrPasswordTooLong) {
		return "", errorx.BadRequest("password is too long")
	}

	return string(hash), err
}

func (s *Service) SignIn(
	ctx context.Context,
	username string,
	password string,
) (jwt.TokenPair, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return jwt.TokenPair{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return jwt.TokenPair{}, errorx.BadRequest("invalid username or password")
	}

	roles := make([]string, 0, len(user.Roles))
	for _, role := range user.Roles {
		roles = append(roles, string(role))
	}

	pair, err := s.token.GenerateTokens(user.Id, roles)
	if err != nil {
		return jwt.TokenPair{}, err
	}

	err = s.repo.InsertUserRefreshToken(ctx, user.Id, pair.RefreshToken, pair.RefreshTokenExpiresIn)
	if err != nil {
		return jwt.TokenPair{}, err
	}

	return pair, nil
}

func (s *Service) RefreshTokens(
	ctx context.Context,
	refreshToken string,
) (jwt.TokenPair, error) {

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return jwt.TokenPair{}, err
	}
	defer tx.Rollback(ctx) //nolint:errcheck

	// craete tx repo
	repo := NewRepo(tx)

	token, err := repo.GetUserToken(ctx, refreshToken)
	if err != nil {
		return jwt.TokenPair{}, err
	}

	if time.Now().After(token.ExpiresIn) {
		err = repo.DeleteUserToken(ctx, refreshToken)
		if err != nil {
			return jwt.TokenPair{}, err
		}

		if err := tx.Commit(ctx); err != nil {
			return jwt.TokenPair{}, err
		}

		return jwt.TokenPair{}, errorx.Unauthorized
	}

	user, err := repo.GetUserById(ctx, token.UserId)
	if err != nil {
		return jwt.TokenPair{}, err
	}

	// если не стух, то генерируем новую пачку токенов
	roles := make([]string, 0, len(user.Roles))
	for _, role := range user.Roles {
		roles = append(roles, string(role))
	}

	pair, err := s.token.GenerateTokens(user.Id, roles)
	if err != nil {
		return jwt.TokenPair{}, err
	}

	err = repo.DeleteUserToken(ctx, refreshToken)
	if err != nil {
		return jwt.TokenPair{}, err
	}

	err = repo.InsertUserRefreshToken(ctx, user.Id, pair.RefreshToken, pair.RefreshTokenExpiresIn)
	if err != nil {
		return jwt.TokenPair{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return jwt.TokenPair{}, err
	}

	return pair, nil
}

func (s *Service) UpdateProfile(
	ctx context.Context,
	uid int,
	password string,
	firstName string,
	lastName string,
) error {
	user, err := s.repo.GetUserById(ctx, uid)
	if err != nil {
		return err
	}

	hash, err := s.hashPassword(password)
	if err != nil {
		return err
	}

	return s.repo.UpdateUser(ctx, uid, user.Username, hash, firstName, lastName, user.Roles)
}

func (s *Service) UpdateUser(ctx context.Context, user User) error {
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadUncommitted})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx) //nolint:errcheck

	repo := NewRepo(tx)

	_, err = repo.GetUserById(ctx, user.Id)
	if err != nil {
		return err
	}

	existedUser, err := repo.GetUserByUsername(ctx, user.Username)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}
	if err == nil && user.Id != existedUser.Id {
		return errorx.BadRequest(fmt.Sprintf("username '%s' already taken", user.Username))
	}

	hash, err := s.hashPassword(user.Password)
	if err != nil {
		return err
	}

	err = repo.UpdateUser(
		ctx,
		user.Id,
		user.Username,
		hash,
		user.FirstName,
		user.LastName,
		user.Roles,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
