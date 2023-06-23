package testing

import (
	"context"
	"flag"
	"fmt"
	"log"
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

func TestNodeSuite(t *testing.T) {
	s := new(StaderNodeSuite)
	s.app = newApp()
	suite.Run(t, s)
}

func (s *StaderNodeSuite) TestNodeDaemon() {
	time.Sleep(time.Second * 10)
}

func (s *StaderNodeSuite) TestNodeDeposit() {
	// eth.EthToWei(100000).
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
			"9000000000000000000",
			"0",
			"1",
			"false",
		})
		assert.Nil(s.T(), err)
	}()

	time.Sleep(time.Minute)
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

	// flagSet.StringVar(&cp, "maxFee", "1000000000000000", "maxFee")
	// flagSet.StringVar(&cp, "maxPrioFee", "1000000000000000", "maxPrioFee")

	c := cli.NewContext(s.app, flagSet, nil)

	ePort := s.startAnvil(s.T())
	elUrl := fmt.Sprintf("http://127.0.0.1:%s", ePort)
	// cURL := "http://localhost:61143"
	s.staderConfig(ctx, c, nil, elUrl)

	fmt.Println("Done SetupSuite()")

	go func() {
		os.Setenv("PRE_SIGN_COOL_DOWN", "10s")
		a := os.Args
		err := s.app.Run([]string{
			a[0],
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
	fmt.Println("TearDown StaderNodeSuite")
	defer func() {
		if s.kurtosisCtx != nil {

			s.kurtosisCtx.Clean(context.Background(), true)
		}
	}()

	removeTestFolder(s.T())

	// You can't defer this because os.Exit doesn't care for defer
	if err := s.pool.Purge(s.anvil); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func removeTestFolder(t *testing.T) {
	path, err := homedir.Expand(ConfigPath)
	require.Nil(t, err)
	_ = os.RemoveAll(path)
}

func (s *StaderNodeSuite) startAnvil(t *testing.T) string {
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
			// PortBindings: map[docker.Port][]docker.PortBinding{
			// 	"8545/tcp": {{HostIP: "localhost", HostPort: "8545/tcp"}},
			// },
		}, func(hc *docker.HostConfig) {})

	require.Nil(s.T(), err)
	elPort := resource.GetPort("8545/tcp")

	err = pool.Retry(func() error {
		client, err := ethclient.Dial(fmt.Sprintf("http://127.0.0.1:%s", elPort))
		require.Nil(s.T(), err)

		s.client = client
		return err
	})

	time.Sleep(time.Second * 5)
	require.Nil(s.T(), err)

	s.pool = pool
	s.anvil = resource

	// Get resource's published port for our imposter.
	return elPort
}
