module grpc-battle

go 1.13

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190121172915-509febef88a4
	golang.org/x/image => github.com/golang/image v0.0.0-20190802002840-cff245a6509b
	golang.org/x/lint => github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/net => github.com/golang/net v0.0.0-20190320064053-1272bf9dcd53
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sync => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190219203350-90b0e4468f99 // indirect
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20180221164845-07fd8470d635
	google.golang.org/grpc => github.com/grpc/grpc-go v1.26.0
	gopkg.in/yaml.v2 => github.com/go-yaml/yaml v0.0.0-20181115110504-51d6538a90f8
)

require (
	github.com/Shopify/sarama v1.26.1
	github.com/fsnotify/fsnotify v1.4.7
	github.com/golang/protobuf v1.4.0
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/schema v1.1.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.3.2
	golang.org/x/crypto v0.0.0-20200204104054-c9f3fb736b72
	golang.org/x/exp v0.0.0-20190121172915-509febef88a4
	golang.org/x/image v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	google.golang.org/genproto v0.0.0-20200331122359-1ee6d9798940
	google.golang.org/grpc v1.28.1
	google.golang.org/protobuf v1.21.0
	gopkg.in/gographics/imagick.v2 v2.6.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.2.8
)
