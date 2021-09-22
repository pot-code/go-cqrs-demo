package db

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pot-code/go-cqrs-demo/ent"
	gobit "github.com/pot-code/gobit/pkg"
	"github.com/pot-code/gobit/pkg/config"
	"github.com/pot-code/gobit/pkg/util"
)

func NewEntClient(bc *config.BaseConfig, sc *sqlx.DB, lm *util.LifecycleManager) *ent.Client {
	var drv dialect.Driver

	drv = entsql.OpenDB(dialect.Postgres, sc.DB)
	if bc.Env == gobit.EnvDevelop {
		drv = dialect.Debug(drv)
	}
	client := ent.NewClient(ent.Driver(drv))

	lm.OnExit(func(ctx context.Context) {
		log.Println("[ent.Client] close ent client")
		client.Close()
	})

	return client
}
