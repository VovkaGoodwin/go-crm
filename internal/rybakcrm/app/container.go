package app

import (
	"crm-backend/internal/rybakcrm/config"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"crm-backend/internal/rybakcrm/app/domain/repository"
	"crm-backend/internal/rybakcrm/app/domain/service"
	infra_repo "crm-backend/internal/rybakcrm/app/infrastructure/repository"
)

var (
	userRepository         repository.UserRepository
	accessTokenRepository  repository.AccessTokenRepository
	refreshTokenRepository repository.RefreshTokenRepository
	departmentRepository   repository.DepartmentRepository
)

func initContainer(cfg *config.Config, postgresDb *sqlx.DB, redis *redis.Client) {
	userRepository = infra_repo.NewUserRepository(postgresDb)
	accessTokenRepository = infra_repo.NewAccessTokenRepository(cfg, redis)
	refreshTokenRepository = infra_repo.NewRefreshTokenRepository(cfg, redis)
	departmentRepository = infra_repo.NewDepartmentRepositoryRepository(postgresDb)
}

func AuthService() *service.AuthService {
	return service.NewAuthService(
		userRepository,
		accessTokenRepository,
		refreshTokenRepository,
	)
}

func DepartmentService() *service.DepartmentService {
	return service.NewDepartmentService(departmentRepository)
}
