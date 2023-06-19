module github.com/stader-labs/stader-node

go 1.13

require (
	github.com/Masterminds/semver/v3 v3.2.1 // indirect
	github.com/Microsoft/go-winio v0.6.0 // indirect
	github.com/a8m/envsubst v1.3.0
	github.com/alessio/shellescape v1.4.1
	github.com/blang/semver/v4 v4.0.0
	github.com/btcsuite/btcd v0.23.1
	github.com/btcsuite/btcd/btcutil v1.1.2
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/docker v20.10.18+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/dustin/go-humanize v1.0.0
	github.com/ethereum/go-ethereum v1.10.26
	github.com/fatih/color v1.13.0
	github.com/ferranbt/fastssz v0.1.2
	github.com/glendc/go-external-ip v0.1.0
	github.com/google/uuid v1.3.0
	github.com/herumi/bls-eth-go-binary v1.28.1 // indirect
	github.com/imdario/mergo v0.3.13
	github.com/klauspost/cpuid/v2 v2.1.1
	github.com/kurtosis-tech/kurtosis-portal/api/golang v0.0.0-20230411133558-b983d9bebe4f // indirect
	github.com/kurtosis-tech/kurtosis/api/golang v0.78.0
	github.com/kurtosis-tech/kurtosis/grpc-file-transfer/golang v0.0.0-20230609162710-10b6b91dc9e8 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/moby/term v0.0.0-20220808134915-39b0c02b01ae // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58
	github.com/prometheus/client_golang v1.13.0
	github.com/prysmaticlabs/go-bitfield v0.0.0-20210809151128-385d8c5e3fb7
	github.com/prysmaticlabs/prysm/v3 v3.1.1
	github.com/rivo/tview v0.0.0-20230530133550-8bd761dda819 // indirect
	github.com/sethvargo/go-password v0.2.0
	github.com/shirou/gopsutil/v3 v3.23.1
	github.com/sirupsen/logrus v1.9.3
	github.com/stader-labs/ethcli-ui/configuration v0.0.0-20230602142021-378099c5eca8
	github.com/stader-labs/ethcli-ui/wizard v0.0.0-20230602142021-378099c5eca8
	github.com/stretchr/testify v1.8.4
	github.com/test-go/testify v1.1.4 // indirect
	github.com/tyler-smith/go-bip39 v1.1.0
	github.com/ulikunitz/xz v0.5.11 // indirect
	github.com/urfave/cli v1.22.10
	github.com/wealdtech/go-eth2-types/v2 v2.7.0
	github.com/wealdtech/go-eth2-util v1.7.0
	github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4 v1.3.0
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/crypto v0.7.0
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sync v0.1.0
	golang.org/x/term v0.8.0
	google.golang.org/genproto v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	gopkg.in/yaml.v2 v2.4.0
	gotest.tools/v3 v3.3.0 // indirect
)

replace github.com/rivo/tview => github.com/hamidraza/tview v1.0.0
