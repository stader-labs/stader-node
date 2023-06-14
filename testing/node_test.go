package testing

import (
	"context"
	"flag"
	"os"
	"testing"
	"time"

	// store "github.com/stader-labs/stader-node/stader-lib/contracts"
	"github.com/kurtosis-tech/kurtosis/api/golang/engine/lib/kurtosis_context"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli"
)

type StaderNodeSuite struct {
	done chan struct{}
	app  *cli.App
	suite.Suite
	kurtosisCtx *kurtosis_context.KurtosisContext
	enclaveId   string
}

func TestNodeSuite(t *testing.T) {
	suite.Run(t, new(StaderNodeSuite))
}

func (s *StaderNodeSuite) TestNodeDaemon() {
	// a := os.Args
	// run(time.Minute*1, s.done, func() {
	// 	err := s.app.Run([]string{
	// 		a[0],
	// 		"node",
	// 	})
	// 	require.Nil(s.T(), err)
	// })
	time.Sleep(time.Minute * 2)

	logrus.Printf("SOSSSSSSSSSSSSSSS")
}

// run once, before test suite methods
func (s *StaderNodeSuite) SetupSuite() {

	defer func() {
		r := recover()

		assert.Nil(s.T(), r, "------------ Recovered TESTS ---------------")
	}()

	ctx, cancelCtxFunc := context.WithCancel(context.Background())
	defer cancelCtxFunc()

	s.app = newApp()
	s.done = make(chan struct{}, 1)

	flagSet := flag.NewFlagSet("node_testing", flag.PanicOnError)
	var p string
	flagSet.StringVar(&p, "settings", UserSettingPath, "settings")
	var cp string
	flagSet.StringVar(&cp, "config-path", ConfigPath, "config-path")
	c := cli.NewContext(s.app, flagSet, nil)

	s.staderConfig(ctx, c)

	logrus.Println("Done SetupSuite()")

	go func() {

		a := os.Args
		// run(time.Minute*1, s.done, func() {
		err := s.app.Run([]string{
			a[0],
			"node",
		})
		require.Nil(s.T(), err)
	}()
}

// run once, after test suite methods
func (s *StaderNodeSuite) TearDownSuite() {
	// <-s.done
	logrus.Println("TearDown StaderNodeSuite")

	defer s.kurtosisCtx.Clean(context.Background(), true)

	path, err := homedir.Expand(ConfigPath)
	require.Nil(s.T(), err)
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
