package database

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/alex123012/database-users-operator/api/v1alpha1"
	"github.com/alex123012/database-users-operator/controllers/database/connection"
	"github.com/alex123012/database-users-operator/controllers/database/postgresql"
	"github.com/alex123012/database-users-operator/controllers/internal"
)

type Interface interface {
	Close(cxt context.Context) error
	CreateUser(ctx context.Context, username, password string) (map[string]string, error)
	DeleteUser(ctx context.Context, username string) error
	ApplyPrivileges(ctx context.Context, username string, privileges []v1alpha1.PrivilegeSpec) error
	RevokePrivileges(ctx context.Context, username string, privileges []v1alpha1.PrivilegeSpec) error
}

type dbConnection interface {
	Close(ctx context.Context) error
	Copy() interface{}
	Connect(ctx context.Context, driver string, connString string) error
	Exec(ctx context.Context, disableLog connection.LogInfo, query string) error
}

func NewDatabase(ctx context.Context, s v1alpha1.DatabaseSpec, client client.Client, logger logr.Logger) (Interface, error) {
	conn := connection.NewDefaultConnector(logger)
	return newDatabase(ctx, conn, s, client, logger)
}

func newDatabase(ctx context.Context, conn dbConnection, s v1alpha1.DatabaseSpec, client client.Client, logger logr.Logger) (Interface, error) {
	var db Interface
	var err error
	switch s.Type {
	case v1alpha1.PostgreSQL:
		db, err = newPostgresql(ctx, conn, s.PostgreSQL, client, logger)
	default:
		err = fmt.Errorf("can't find supported DB type '%s'", s.Type)
	}
	return db, err
}

func newPostgresql(ctx context.Context, conn dbConnection, c v1alpha1.PostgreSQLConfig, client client.Client, logger logr.Logger) (*postgresql.Postgresql, error) {
	sslData := make(map[string]string, 0)
	var sslCAKey string
	if c.SSLMode == v1alpha1.SSLModeREQUIRE || c.SSLMode == v1alpha1.SSLModeVERIFYCA || c.SSLMode == v1alpha1.SSLModeVERIFYFULL {
		var err error
		sslData, err = internal.DecodeSecretData(ctx, types.NamespacedName(c.SSLCredentialsSecret), client)
		if err != nil {
			return nil, err
		}
		sslCAData, err := internal.DecodeSecretData(ctx, types.NamespacedName(c.SSLCAKey.Secret), client)
		if err != nil {
			return nil, err
		}
		sslCAKey = sslCAData[c.SSLCAKey.Key]
	}

	var password string
	if c.PasswordSecret.Key != "" && c.PasswordSecret.Secret.Name != "" && c.PasswordSecret.Secret.Namespace != "" {
		data, err := internal.DecodeSecretData(ctx, types.NamespacedName(c.PasswordSecret.Secret), client)
		if err != nil {
			return nil, err
		}
		password = data[c.PasswordSecret.Key]
	}

	cfg := postgresql.NewConfig(c.Host, c.Port, c.User, password, c.DatabaseName,
		c.SSLMode, sslData["ca.crt"], sslData["tls.crt"], sslData["tls.key"], sslCAKey)

	p := postgresql.NewPostgresql(conn, cfg, logger)
	return p, p.Connect(ctx)
}
