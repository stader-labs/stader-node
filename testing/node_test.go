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
	"github.com/kurtosis-tech/kurtosis/api/golang/engine/lib/kurtosis_context"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli"
)

type StaderNodeSuite struct {
	suite.Suite
	kurtosisCtx *kurtosis_context.KurtosisContext
	enclaveId   string
}

func TestNodeSuite(t *testing.T) {
	suite.Run(t, new(StaderNodeSuite))
}

func (s *StaderNodeSuite) TestExample1() {
	time.Sleep(time.Minute)
	s.Equal(true, true)
}

// run once, before test suite methods
func (s *StaderNodeSuite) SetupSuite() {
	ctx, cancelCtxFunc := context.WithCancel(context.Background())
	defer cancelCtxFunc()

	app := newApp()

	flagSet := flag.NewFlagSet("node_testing", flag.PanicOnError)
	var p string
	flagSet.StringVar(&p, "settings", ConfigPath, "settings")
	c := cli.NewContext(app, flagSet, nil)

	s.staderConfig(ctx, c)

	go func() {
		a := os.Args
		if err := app.Run([]string{
			a[0],
			"node",
		}); err != nil {
			fmt.Printf("ERROR RUN NODE %+v", err)
		}
	}()

	log.Println("Done SetupSuite()")
}

// run once, after test suite methods
func (s *StaderNodeSuite) TearDownSuite() {
	log.Println("TearDown StaderNodeSuite")

	defer s.kurtosisCtx.DestroyEnclave(context.Background(), s.enclaveId)
}
