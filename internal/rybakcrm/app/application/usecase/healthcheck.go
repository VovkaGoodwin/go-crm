package usecase

import "crm-backend/internal/rybakcrm/app/application/dto"

type Healthcheck struct {
}

func NewHealthCheckUseCase() *Healthcheck {
	return &Healthcheck{}
}

func (u *Healthcheck) Handle() *dto.HealthcheckResponseDto {
	return &dto.HealthcheckResponseDto{
		Result: "ok",
	}
}
