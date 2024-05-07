module github.com/stader-labs/stader-node

go 1.16

require (
	github.com/a8m/envsubst v1.3.0
	github.com/alessio/shellescape v1.4.1
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/blang/semver/v4 v4.0.0
	github.com/btcsuite/btcd v0.24.0
	github.com/btcsuite/btcd/btcutil v1.1.5
	github.com/cespare/cp v1.1.1 // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/docker v23.0.3+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/dustin/go-humanize v1.0.0
	github.com/ethereum/go-ethereum v1.13.15
	github.com/fatih/color v1.13.0
	github.com/ferranbt/fastssz v0.1.2
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08 // indirect
	github.com/glendc/go-external-ip v0.1.0
	github.com/google/uuid v1.3.0
	github.com/hashicorp/go-version v1.6.0
	github.com/herumi/bls-eth-go-binary v1.28.1 // indirect
	github.com/imdario/mergo v0.3.13
	github.com/klauspost/cpuid v1.2.3 // indirect
	github.com/klauspost/cpuid/v2 v2.1.1
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/moby/term v0.0.0-20220808134915-39b0c02b01ae // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58
	github.com/prometheus/client_golang v1.13.0
	github.com/prysmaticlabs/go-bitfield v0.0.0-20210809151128-385d8c5e3fb7
	github.com/prysmaticlabs/gohashtree v0.0.2-alpha // indirect
	github.com/rivo/tview v0.0.0-20230621164836-6cc0565babaf // indirect
	github.com/sethvargo/go-password v0.2.0
	github.com/shirou/gopsutil/v3 v3.23.1
	github.com/stader-labs/ethcli-ui/configuration v0.0.0-20240401032301-ba44ad7f2388
	github.com/stader-labs/ethcli-ui/wizard v0.0.0-20240401032301-ba44ad7f2388
	github.com/syndtr/goleveldb v1.0.1-0.20220614013038-64ee5596c38a // indirect
	github.com/tyler-smith/go-bip39 v1.1.0
	github.com/urfave/cli v1.22.10
	github.com/wealdtech/go-eth2-types/v2 v2.7.0
	github.com/wealdtech/go-eth2-util v1.7.0
	github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4 v1.3.0
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/goleak v1.1.12 // indirect
	golang.org/x/crypto v0.17.0
	golang.org/x/sync v0.5.0
	golang.org/x/term v0.15.0
	gopkg.in/yaml.v2 v2.4.0
	gotest.tools/v3 v3.3.0 // indirect
)

replace github.com/rivo/tview => github.com/hamidraza/tview v1.0.0
