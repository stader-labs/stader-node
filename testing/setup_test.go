package testing

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/stader-labs/stader-node/stader/api"
	"github.com/stader-labs/stader-node/stader/node"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
// func (suite *ExampleTestSuite) SetupSuite() {
// 	// Initialise application

// }

type MyTestSuite struct {
	suite.Suite
}

// listen for 'go test' command --> run test methods
func TestMyTestSuite(t *testing.T) {
	suite.Run(t, new(MyTestSuite))
}

// run once, before test suite methods
func (s *MyTestSuite) SetupSuite() {
	log.Println("SetupSuite()")

	app := cli.NewApp()
	// Register commands
	api.RegisterCommands(app, "api", []string{"a"})
	node.RegisterCommands(app, "node", []string{"n"})

	// Run application
	go func() {
		a := os.Args
		if err := app.Run([]string{
			a[0],
			"--settings=/Users/batphonghan/.stader/user-settings.yml",
			"node",
		}); err != nil {
			panic(err)
		}
	}()

	log.Println("Done SetupSuite()")
	// connect the database, save to 's.db'
}

// run once, after test suite methods
func (s *MyTestSuite) TearDownSuite() {
	log.Println("TearDownSuite()")

	// delete the created database
}

// run before each test
func (s *MyTestSuite) SetupTest() {
	log.Println("SetupTest()")
}

// run after each test
func (s *MyTestSuite) TearDownTest() {
	log.Println("TearDownTest()")
}

// run before each test
func (s *MyTestSuite) BeforeTest(suiteName, testName string) {
	log.Println("BeforeTest()", suiteName, testName)
}

// run after each test
func (s *MyTestSuite) AfterTest(suiteName, testName string) {
	log.Println("AfterTest()", suiteName, testName)
}

func (s *MyTestSuite) TestExample1() {
	time.Sleep(time.Minute)
	s.Equal(true, true)
}

func (s *MyTestSuite) TestExample2() {
	s.Equal(true, true)
}
