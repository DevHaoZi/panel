package job

import (
	"github.com/TheTNB/panel/internal/biz"
	"github.com/google/wire"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"log/slog"
)

var ProviderSet = wire.NewSet(NewJobs)

type Jobs struct {
	db      *gorm.DB
	log     *slog.Logger
	setting biz.SettingRepo
	cert    biz.CertRepo
	backup  biz.BackupRepo
	cache   biz.CacheRepo
}

func NewJobs(db *gorm.DB, log *slog.Logger, setting biz.SettingRepo, cert biz.CertRepo, backup biz.BackupRepo, cache biz.CacheRepo) *Jobs {
	return &Jobs{
		db:      db,
		log:     log,
		setting: setting,
		cert:    cert,
		backup:  backup,
		cache:   cache,
	}
}

func (r *Jobs) Register(c *cron.Cron) error {
	if _, err := c.AddJob("* * * * *", NewMonitoring(r.db, r.log, r.setting)); err != nil {
		return err
	}
	if _, err := c.AddJob("0 4 * * *", NewCertRenew(r.db, r.log, r.cert)); err != nil {
		return err
	}

	if _, err := c.AddJob("0 2 * * *", NewPanelTask(r.db, r.log, r.backup, r.cache, r.setting)); err != nil {
		return err
	}

	return nil
}
