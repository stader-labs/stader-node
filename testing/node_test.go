package testing

import (
	"context"
	"flag"
	"fmt"
	"os"
	"testing"
	"time"

	// store "github.com/stader-labs/stader-node/stader-lib/contracts"
	"github.com/kurtosis-tech/kurtosis/api/golang/engine/lib/kurtosis_context"
	"github.com/mitchellh/go-homedir"

	//stader/register.go

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
	a := os.Args
	err := s.app.Run([]string{
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
	c := cli.NewContext(s.app, flagSet, nil)

	s.staderConfig(ctx, c)

	fmt.Println("Done SetupSuite()")

	go func() {
		a := os.Args
		err := s.app.Run([]string{
			a[0],
			"node",
		})
		assert.Nil(s.T(), err)
	}()

	go func() {
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
}

func removeTestFolder(t *testing.T) {
	path, err := homedir.Expand(ConfigPath)
	require.Nil(t, err)
	_ = os.RemoveAll(path)
}

func run(t time.Duration, done chan<- struct{}, f func()) {
	c1 := make(chan struct{}, 1)
	go func() {
		f()
		c1 <- struct{}{}
	}()

	select {
	case _ = <-c1:
	case <-time.After(t):
		done <- struct{}{}
		return
	}
}
