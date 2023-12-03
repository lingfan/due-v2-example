module due-v2-example

go 1.21.4

require (
	github.com/dobyte/due/eventbus/nats/v2 v2.0.0-20231113044848-500fc2dd4021
	github.com/dobyte/due/eventbus/redis/v2 v2.0.0-20231113044848-500fc2dd4021
	github.com/dobyte/due/locate/redis/v2 v2.0.0-20231113044848-500fc2dd4021
	github.com/dobyte/due/network/ws/v2 v2.0.0-20231113044848-500fc2dd4021
	github.com/dobyte/due/registry/etcd/v2 v2.0.0-20231113044848-500fc2dd4021
	github.com/dobyte/due/transport/grpc/v2 v2.0.0-20231113044848-500fc2dd4021
	github.com/dobyte/due/v2 v2.0.1-0.20231113044848-500fc2dd4021
	github.com/dobyte/jwt v0.1.3
	github.com/golang/protobuf v1.5.3
	github.com/redis/go-redis/v9 v9.3.0
	go.mongodb.org/mongo-driver v1.13.0
	golang.org/x/crypto v0.15.0
	golang.org/x/net v0.18.0
	google.golang.org/api v0.152.0
)

require (
	cloud.google.com/go/compute v1.23.3 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/BurntSushi/toml v1.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/jinzhu/copier v0.3.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/nats-io/nats.go v1.31.0 // indirect
	github.com/nats-io/nkeys v0.4.5 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/panjf2000/ants/v2 v2.7.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shamaton/msgpack/v2 v2.1.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.etcd.io/etcd/api/v3 v3.5.9 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.9 // indirect
	go.etcd.io/etcd/client/v3 v3.5.9 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/oauth2 v0.14.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231120223509-83a465c0220f // indirect
	google.golang.org/grpc v1.59.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)


// 替换规则
//replace (
//    github.com/dobyte/due/v2 => ../due // 绝对路径 或 相对路径 都可以
//    github.com/dobyte/due/locate/redis/v2 => ../due/locate/redis
//    github.com/dobyte/due/registry/etcd/v2 => ../due/registry/etcd
//)