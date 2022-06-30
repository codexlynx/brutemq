package etcd

import (
	"context"
	etcd "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

type BruteEtcd struct {
	Endpoint string
	User     string
	client   *etcd.Client
}

func (brute *BruteEtcd) TryPassword(password string) (bool, error) {
	_, err := brute.client.Authenticate(context.Background(), brute.User, password)
	if err != nil {
		return false, nil
	}
	return true, err
}

func (brute *BruteEtcd) Close() {
	err := brute.client.Close()
	if err != nil {
		log.Println(err)
	}
}

func NewBruteEtcd(endpoint string, user string) (*BruteEtcd, error) {
	loggerCfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"/dev/null"},
		ErrorOutputPaths: []string{"/dev/null"},
	}

	logger, err := loggerCfg.Build()
	if err != nil {
		return nil, err
	}

	client, err := etcd.New(etcd.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: 5 * time.Second,
		Logger:      logger,
	})
	if err != nil {
		return nil, err
	}

	return &BruteEtcd{
		Endpoint: endpoint,
		User:     user,
		client:   client,
	}, nil
}
