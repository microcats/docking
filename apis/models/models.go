package models

import (
    "fmt"
    "log"
    "time"
    "golang.org/x/net/context"
    "github.com/coreos/etcd/client"
)

var dataSource EtcdInterface

func init() {
    dataSource = newEtcdClient()
}

type EtcdInterface interface {
    Get(key string, opts *client.GetOptions) (*client.Response, error)
    Set(key, value string, opts *client.SetOptions) (*client.Response, error)
    Delete(key string, opts *client.DeleteOptions) (*client.Response, error)
    Update(key, value string) (*client.Response, error)
    IsKey(key string) (bool)
}

type EtcdClient struct {
    etcdClient client.Client
    prefix string
}

func newEtcdClient() EtcdInterface {
    cfg := client.Config{
       Endpoints:               []string{"http://127.0.0.1:2379"},
       Transport:               client.DefaultTransport,
       // set timeout per request to fail fast when the target endpoint is unavailable
       HeaderTimeoutPerRequest: time.Second,
    }

    etcd, err := client.New(cfg)
    if err != nil {
       log.Fatal(err)
    }

    return &EtcdClient{
        etcdClient: etcd,
        prefix:     "/docking%s",
    }
}

func (e *EtcdClient) getKeyApi() client.KeysAPI {
    return client.NewKeysAPI(e.etcdClient)
}

func (e *EtcdClient) getKeyPrefix(key string) string {
    return fmt.Sprintf(e.prefix, key)
}

func (e *EtcdClient) Get(key string, opts *client.GetOptions) (*client.Response, error)  {
    return e.getKeyApi().Get(context.Background(), e.getKeyPrefix(key), opts)
}

func (e *EtcdClient) Set(key, value string, opts *client.SetOptions) (*client.Response, error) {
    return e.getKeyApi().Set(context.Background(), e.getKeyPrefix(key), value, opts)
}

func (e *EtcdClient) Delete(key string, opts *client.DeleteOptions) (*client.Response, error) {
    return e.getKeyApi().Delete(context.Background(), e.getKeyPrefix(key), opts)
}

func (e *EtcdClient) Update(key, value string) (*client.Response, error) {
    return e.getKeyApi().Update(context.Background(), e.getKeyPrefix(key), value)
}

func (e *EtcdClient) IsKey(key string) bool {
    if _, err := e.Get(key, nil); err != nil {
        return false
    }

    return true
}
