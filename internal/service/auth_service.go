package service

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	gocache "github.com/patrickmn/go-cache"
	"github.com/xryar/golang-grpc-ecommerce/internal/entity"
	jwtentity "github.com/xryar/golang-grpc-ecommerce/internal/entity/jwt"
	"github.com/xryar/golang-grpc-ecommerce/internal/repository"
	"github.com/xryar/golang-grpc-ecommerce/internal/utils"
	"github.com/xryar/golang-grpc-ecommerce/pb/auth"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IAuthService interface {
	Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error)
	Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error)
	Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error)
	ChangePassword(ctx context.Context, request *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error)
	GetProfile(ctx context.Context, request *auth.GetProfileRequest) (*auth.GetProfileResponse, error)
}

type authService struct {
	authRepository repository.IAuthRepository
	cacheService   *gocache.Cache
}

func (as *authService) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	if request.Password != request.PasswordConfirmation {
		return &auth.RegisterResponse{
			Base: utils.BadRequestResponse("Password is not matched"),
		}, nil
	}

	userEmail, err := as.authRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if userEmail != nil {
		return &auth.RegisterResponse{
			Base: utils.BadRequestResponse("User already Exist"),
		}, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return nil, err
	}

	newUser := entity.User{
		Id:        uuid.NewString(),
		Fullname:  request.Fullname,
		Email:     request.Email,
		Password:  string(hashedPassword),
		RoleCode:  entity.UserRoleCustomer,
		CreatedAt: time.Now(),
		CreatedBy: &request.Fullname,
	}
	err = as.authRepository.InsertUser(ctx, &newUser)
	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{
		Base: utils.SuccessResponse("User created"),
	}, nil
}

func (as *authService) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	user, err := as.authRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return &auth.LoginResponse{
			Base: utils.BadRequestResponse("User is not registered"),
		}, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated")
		}

		return nil, err
	}

	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtentity.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Id,
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Email:    user.Email,
		Fullname: user.Fullname,
		Role:     user.RoleCode,
	})
	secretKey := os.Getenv("JWT_SECRET")
	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{
		Base:        utils.SuccessResponse("Login Successfully"),
		AccessToken: accessToken,
	}, nil
}

func (as *authService) Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	jwtToken, err := jwtentity.ParseTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	tokenClaims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	as.cacheService.Set(jwtToken, "", time.Duration(tokenClaims.ExpiresAt.Time.Unix()-time.Now().Unix())*time.Second)

	return &auth.LogoutResponse{
		Base: utils.SuccessResponse("Logout Success"),
	}, nil
}

func (as *authService) ChangePassword(ctx context.Context, request *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error) {
	if request.NewPassword != request.NewPasswordConfirmation {
		return &auth.ChangePasswordResponse{
			Base: utils.BadRequestResponse("New Password is not matched"),
		}, nil
	}

	jwtToken, err := jwtentity.ParseTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}
	claims, err := jwtentity.GetClaimsFromToken(jwtToken)
	if err != nil {
		return nil, err
	}

	user, err := as.authRepository.GetUserByEmail(ctx, claims.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return &auth.ChangePasswordResponse{
			Base: utils.BadRequestResponse("User doesnt't exist"),
		}, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return &auth.ChangePasswordResponse{
				Base: utils.BadRequestResponse("Old password is not matched"),
			}, nil
		}

		return nil, err
	}

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), 10)
	if err != nil {
		return nil, err
	}
	err = as.authRepository.UpdateUserPassword(ctx, user.Id, string(hashedNewPassword), user.Fullname)
	if err != nil {
		return nil, err
	}

	return &auth.ChangePasswordResponse{
		Base: utils.SuccessResponse("Change Password success"),
	}, nil
}

func (as *authService) GetProfile(ctx context.Context, request *auth.GetProfileRequest) (*auth.GetProfileResponse, error) {
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := as.authRepository.GetUserByEmail(ctx, claims.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return &auth.GetProfileResponse{
			Base: utils.BadRequestResponse("User doesn't exist"),
		}, nil
	}

	return &auth.GetProfileResponse{
		Base:        utils.SuccessResponse("Success Get Profile User"),
		UserId:      claims.Subject,
		Fullname:    claims.Fullname,
		Email:       claims.Email,
		RoleCode:    claims.Role,
		MemberSince: timestamppb.New(user.CreatedAt),
	}, nil
}

func NewAuthService(authRepository repository.IAuthRepository, cacheService *gocache.Cache) IAuthService {
	return &authService{
		authRepository: authRepository,
		cacheService:   cacheService,
	}
}
