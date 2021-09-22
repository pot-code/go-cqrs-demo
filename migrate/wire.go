//+build wireinject

package migrate

import (
	"github.com/google/wire"
	"github.com/pot-code/go-cqrs-demo/migrate/config"
	pdb "github.com/pot-code/go-cqrs-demo/pkg/db"
	"github.com/pot-code/go-cqrs-demo/pkg/mq"
	"github.com/pot-code/gobit/pkg/db"
	"github.com/pot-code/gobit/pkg/util"
)

func InitMigrate(lm *util.LifecycleManager) *Migration {
	wire.Build(
		NewMigration, pdb.NewEntClient, db.NewSqlxProvider, config.ConfigSet, mq.NewKafkaAdmin,
	)
	return &Migration{}
}
