package postgres

import (
	"github.com/CosmeticsShiraz/Backend/internal/domain/entity"
	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type ReportRepository interface {
	CreateReport(db database.Database, report *entity.Report) error
	GetReportsByObjectType(db database.Database, objectType string, statuses []enum.ReportStatus, opts ...QueryModifier) ([]*entity.Report, error)
	FindReportByID(db database.Database, id uint) (*entity.Report, error)
	UpdateReport(db database.Database, report *entity.Report) error
}
