package packages

import (
	"fmt"
	"os/exec"
	"time"
)

var Packages = map[string]string{
	"dawn":            "github.com/dawnsgo/dawn",
	"cache/memcache":  "github.com/dawnsgo/dawn/cache/memcache",
	"cache/redis":     "github.com/dawnsgo/dawn/cache/redis",
	"component/http":  "github.com/dawnsgo/dawn/component/http",
	"config/consul":   "github.com/dawnsgo/dawn/config/consul",
	"config/etcd":     "github.com/dawnsgo/dawn/config/etcd",
	"config/nacos":    "github.com/dawnsgo/dawn/config/nacos",
	"crypto/ecc":      "github.com/dawnsgo/dawn/crypto/ecc",
	"crypto/rsa":      "github.com/dawnsgo/dawn/crypto/rsa",
	"eventbus/kafka":  "github.com/dawnsgo/dawn/eventbus/kafka",
	"eventbus/nats":   "github.com/dawnsgo/dawn/eventbus/nats",
	"eventbus/redis":  "github.com/dawnsgo/dawn/eventbus/redis",
	"locate/redis":    "github.com/dawnsgo/dawn/locate/redis",
	"lock/memcache":   "github.com/dawnsgo/dawn/lock/memcache",
	"lock/redis":      "github.com/dawnsgo/dawn/lock/redis",
	"log/aliyun":      "github.com/dawnsgo/dawn/log/aliyun",
	"log/tencent":     "github.com/dawnsgo/dawn/log/tencent",
	"network/kcp":     "github.com/dawnsgo/dawn/network/kcp",
	"network/tcp":     "github.com/dawnsgo/dawn/network/tcp",
	"network/ws":      "github.com/dawnsgo/dawn/network/ws",
	"registry/consul": "github.com/dawnsgo/dawn/registry/consul",
	"registry/etcd":   "github.com/dawnsgo/dawn/registry/etcd",
	"registry/nacos":  "github.com/dawnsgo/dawn/registry/nacos",
	"transport/grpc":  "github.com/dawnsgo/dawn/transport/grpc",
	"transport/rpcx":  "github.com/dawnsgo/dawn/transport/rpcx",
}

func CMD(pkg, major, version, dir string) *exec.Cmd {
	cmd := exec.Command("go", "get", fmt.Sprintf("%s/%s@%s", pkg, major, version))
	cmd.Dir = dir
	cmd.WaitDelay = 30 * time.Second

	return cmd
}
