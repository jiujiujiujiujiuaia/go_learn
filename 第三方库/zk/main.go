package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/lexkong/log"

	"github.com/pkg/errors"
	"github.com/samuel/go-zookeeper/zk"
)

type ZkProvider struct {
	Conn  *zk.Conn
	Event <-chan zk.Event
	Once  sync.Once
	Cache sync.Map
}

type Node struct {
	Value    string
	State    *zk.Stat
	Old      string
	Listener []func(node *Node)
}

func (node *Node) AddListener(f func(node *Node)) {
	node.Listener = append(node.Listener, f)
}

var provider *ZkProvider

func init() {
	provider = &ZkProvider{}
}

func GetProvider() (*ZkProvider, error) {
	provider.Once.Do(
		func() {
			var err error
			zkHostsAddr := "129.204.8.88"
			provider.Conn, provider.Event, err = zk.Connect([]string{zkHostsAddr}, time.Second)
			if provider.Conn == nil || err != nil {
				go reconnect(provider, zkHostsAddr)
			}
		},
	)
	return provider, nil
}

//超时重试
func reconnect(provider *ZkProvider, addr string) (err error) {
	//10s超时
	<-time.After(10 * time.Second)
	provider.Conn, provider.Event, err = zk.Connect([]string{addr}, time.Second)
	if err != nil {
		go reconnect(provider, addr)
	}
	return nil
}

func Load(path string) (value string, err error) {
	node, err := getNode(path)
	if err != nil {
		log.Errorf(err, "getNode err:%v", err)
		return
	}
	value = node.Value
	return value, nil
}

func AddListener(path string, f func(node *Node)) (err error) {
	node, err := getNode(path)
	if err != nil {
		return errors.New("getNode err")
	}
	node.AddListener(f)
	return nil
}

func main() {
	err := AddListener("/wechat", change)
	if err != nil {
		log.Errorf(err, "err=[%+v]", err)
	}
	select {}
}

func change(node *Node) {
	fmt.Println("变动")
	fmt.Println(node.Value)
}

func getNode(path string) (node *Node, err error) {
	//获取ZkProvider,连接ZkHosts
	provider, err := GetProvider()
	//Load缓存中的数据，定时更新
	value, ok := provider.Cache.Load(path)
	if ok && value.(*Node).Value != "" {
		node = value.(*Node)
	} else {
		err = updateNode(provider, path)
		if err == nil {
			value, ok = provider.Cache.Load(path)
			if ok {
				node = value.(*Node)
			} else {
				err = errors.New("zk value load err")
			}
		}
	}
	return node, nil
}

func updateNode(provider *ZkProvider, path string) (err error) {
	if provider.Conn == nil {
		return errors.New("zk not connect")
	}
	// GetW returns the contents of a znode and sets a watch
	bytes, state, event, err := provider.Conn.GetW(path)
	if err != nil {
		return errors.New("zk conn GetW err")
	}
	value, ok := provider.Cache.Load(path)
	if ok {
		value.(*Node).Old = value.(*Node).Value
		value.(*Node).State = state
		value.(*Node).Value = string(bytes)
	} else {
		node := &Node{
			Value:    string(bytes),
			State:    state,
			Old:      "",
			Listener: make([]func(node *Node), 0),
		}
		provider.Cache.Store(path, node)
	}

	go func() {
		select {
		case e := <-event:
			fmt.Printf("zk event [%+v] [%+v]", e.Path, e.State.String())
			updateNode(provider, path)
			value, ok := provider.Cache.Load(path)
			if ok {
				node := value.(*Node)
				for _, fun := range node.Listener {
					go func() {
						if err := recover(); err != nil {
							log.Errorf(errors.New("qqq"), "recover from panic:%v", err)
						}
						fun(node)
					}()
				}
			}
		case <-time.After(time.Minute * 2):
			updateNode(provider, path)
		}
	}()
	return nil
}
