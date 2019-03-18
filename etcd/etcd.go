package etcd

import (
	"context"
	"github.com/Madq92/etcd-http-client/config"
	"github.com/Madq92/etcd-http-client/log"
	"github.com/sirupsen/logrus"

	"github.com/coreos/etcd/clientv3"
	"time"
)

var client *clientv3.Client

func init() {
	var err error
	etcdConfig := config.Conf.Etcd
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{etcdConfig.Hosts},
		DialTimeout: 10 * time.Second,
		Username:    etcdConfig.Username,
		Password:    etcdConfig.Password,
	})

	if err != nil {
		log.LOGGER.WithError(err).Error("could not init etcd client")
	}
	log.LOGGER.Info("init etcd client success")
}

func GetKey(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := client.Get(ctx, key)
	cancel()
	if err != nil {
		log.LOGGER.WithError(err).Panic(err)
	}
	for _, ev := range resp.Kvs {
		log.LOGGER.
			WithFields(logrus.Fields{
				"key":   string(ev.Key),
				"value": string(ev.Value)}).
			Debug("get kv ok")
	}

	if len(resp.Kvs) > 0 {
		return string(resp.Kvs[0].Value), nil
	} else {
		return "", nil
	}
}

func PutKey(key, value string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := client.Put(ctx, key, value)
	cancel()
	if err != nil {
		log.LOGGER.WithError(err).Panic(err)
		return 0, err
	}
	return resp.Header.Revision, nil
}
