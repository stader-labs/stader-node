package testing

import (
	"context"
	"flag"
	"fmt"
	"os"
	"testing"
	"time"

	// store "github.com/stader-labs/stader-node/stader-lib/contracts"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kurtosis-tech/kurtosis/api/golang/engine/lib/kurtosis_context"
	"github.com/mitchellh/go-homedir"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	//stader/register.go

	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	_ "github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/stader-labs/stader-node/testing/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli"
)

type StaderNodeSuite struct {
	app *cli.App
	suite.Suite
	kurtosisCtx *kurtosis_context.KurtosisContext
	enclaveId   string
	pool        *dockertest.Pool
	anvil       *dockertest.Resource
	client      *ethclient.Client
}

var (
	maxFee         = "100000"
	maxPriorityFee = "100000"
	gasLimit       = "100000000000"
)

func TestNodeSuite(t *testing.T) {
	s := new(StaderNodeSuite)
	s.app = newApp()
	suite.Run(t, s)
}

// func (s *StaderNodeSuite) TestNodeDaemon() {
// 	time.Sleep(time.Second * 10)
// }

func (s *StaderNodeSuite) TestNodeDeposit() {
	// eth.EthToWei(100000).
	time.Sleep(time.Second)
	go func() {
		a := os.Args

		err := s.app.Run([]string{
			a[0],
			"api",
			"node",
			"deposit-sd-approve-sd",
			eth.EthToWei(10000).String(),
		})
		assert.Nil(s.T(), err)
		err = s.app.Run([]string{
			a[0],
			"api",
			"node",
			"deposit-sd",
			eth.EthToWei(9000).String(),
		})
		assert.Nil(s.T(), err)

		err = s.app.Run([]string{
			a[0],
			"api",
			"validator",
			"deposit",
			eth.EthToWei(9).String(),
			"0",
			"1",
			"false",
		})
		assert.Nil(s.T(), err)
	}()

	time.Sleep(time.Second * 5)
}

// func (s *StaderNodeSuite) TestNode2() {
// 	time.Sleep(time.Second * 20)
// }

// run once, before test suite methods
func (s *StaderNodeSuite) SetupSuite() {
	removeTestFolder(s.T())
	defer func() {
		r := recover()
		assert.Nil(s.T(), r, "------------ Recovered TESTS ---------------")
	}()

	ctx, cancelCtxFunc := context.WithCancel(context.Background())
	defer cancelCtxFunc()

	flagSet := flag.NewFlagSet("node_testing", flag.PanicOnError)
	var p string
	flagSet.StringVar(&p, "settings", UserSettingPath, "settings")
	var cp string
	flagSet.StringVar(&cp, "config-path", ConfigPath, "config-path")

	flagSet.StringVar(&maxFee, "maxFee", maxFee, "Gas fee cap to use for the 1559 transaction execution")
	flagSet.StringVar(&maxPriorityFee, "maxPrioFee", maxPriorityFee, "Gas priority fee cap to use for the 1559 transaction execution")

	flagSet.StringVar(&gasLimit, "gasLimit", gasLimit, "Gas priority fee cap to use for the 1559 transaction execution")

	var localTestnet bool
	flagSet.BoolVar(&localTestnet, "local-testnet", true, "local-testnet")

	c := cli.NewContext(s.app, flagSet, nil)

	ePort := s.startAnvil(s.T())
	elUrl := fmt.Sprintf("http://127.0.0.1:%+v", ePort)
	// cURL := "http://127.0.0.1:59541"
	s.staderConfig(ctx, c, nil, elUrl)

	fmt.Println("Done SetupSuite()")

	go func() {
		a := os.Args
		err := s.app.Run([]string{
			a[0],
			"--local-testnet=true",
			"--presign-cooldown=1s",
			"node",
		})
		require.Nil(s.T(), err)
	}()

	go func() {
		defer func() {
			r := recover()
			fmt.Printf("RECOVER TEST SERVER %+v \n", r)
			require.Nil(s.T(), r)
		}()

		httptest.SererHttp(s.T())
	}()
}

// run once, after test suite methods
func (s *StaderNodeSuite) TearDownSuite() {
	if s.kurtosisCtx != nil {
		s.kurtosisCtx.Clean(context.Background(), true)
	}
	removeTestFolder(s.T())

	err := s.pool.Purge(s.anvil)
	require.Nil(s.T(), err)

	fmt.Println("TearDown StaderNodeSuite")
}

func removeTestFolder(t *testing.T) {
	path, err := homedir.Expand(ConfigPath)
	require.Nil(t, err)
	_ = os.RemoveAll(path)
}

func (s *StaderNodeSuite) startAnvil(t *testing.T) string {

	fmt.Println("------------ START ANVIL ---------------")

	pool, err := dockertest.NewPool("")
	require.Nil(s.T(), err)

	err = pool.Client.Ping()
	require.Nil(s.T(), err)

	resource, err := pool.RunWithOptions(
		&dockertest.RunOptions{
			Repository:   "ghcr.io/foundry-rs/foundry",
			Cmd:          []string{"anvil"},
			ExposedPorts: []string{"8545/tcp", "8545/udp"},
			Env: []string{
				"ANVIL_IP_ADDR=0.0.0.0",
			},
		}, func(hc *docker.HostConfig) {})

	require.Nil(s.T(), err)
	elPort := resource.GetPort("8545/tcp")

	err = pool.Retry(func() error {
		client, err := ethclient.Dial(fmt.Sprintf("http://127.0.0.1:%s", elPort))
		require.Nil(s.T(), err)

		s.client = client
		return err
	})

	require.Nil(s.T(), err)

	s.pool = pool
	s.anvil = resource

	// Get resource's published port for our imposter.
	fmt.Printf("------------ ANVIL STARTED AT: %+v --------------- \n", elPort)
	return elPort
}
