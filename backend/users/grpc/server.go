package grpc

import (
	"context"
	"errors"
	"main/internal/jwt"
	"main/internal/users"

	"github.com/jackc/pgx/v5"
	accounts_proto "gitlab.com/volgaIt/grpc-proto/accounts"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	accounts_proto.UnimplementedAccountsServer
	validator jwt.TokenManager
	repo      *users.Repo
}

func New(
	validator jwt.TokenManager,
	repo *users.Repo,
) *Server {

	return &Server{
		validator: validator,
		repo:      repo,
	}
}

func (s *Server) ValidateAccessToken(
	ctx context.Context,
	in *accounts_proto.ValidateAccessTokenRequest,
) (*accounts_proto.TokenPayload, error) {
	payload, err := s.validator.ParseToken(in.Token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	roles := make([]accounts_proto.Role, 0, len(payload.Roles))
	for _, role := range payload.Roles {
		roles = append(roles, userRoleToGrpcRole(users.Role(role)))
	}

	return &accounts_proto.TokenPayload{
		Uid:   int64(payload.Uid),
		Roles: roles,
	}, nil
}

func (s *Server) GetUserById(ctx context.Context, in *accounts_proto.GetUserByIdRequest) (*accounts_proto.User, error) {
	user, err := s.repo.GetUserById(ctx, int(in.Id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "user is not exist")
		}

		return nil, err
	}

	roles := make([]accounts_proto.Role, 0, len(user.Roles))
	for _, role := range user.Roles {
		roles = append(roles, userRoleToGrpcRole(users.Role(role)))
	}

	return &accounts_proto.User{
		Id:    int64(user.Id),
		Roles: roles,
	}, nil
}

func userRoleToGrpcRole(role users.Role) accounts_proto.Role {
	switch role {
	case users.Admin:
		return accounts_proto.Role_admin
	case users.Teacher:
		return accounts_proto.Role_manager
	case users.Student:
		return accounts_proto.Role_doctor
	default:
		return accounts_proto.Role_user
	}
}
