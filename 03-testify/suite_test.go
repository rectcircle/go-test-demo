package testifydemo_test

// Basic imports
import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// 定义测试套件结构体，嵌入一个 suite.Suite，该结构体包含一个 T() 方法可以返回原生的 *testing.T
type ExampleTestSuite struct {
	suite.Suite
}

// 运行套件内所有测试函数前，执行且只执行一次该函数。
func (suite *ExampleTestSuite) SetupSuite() {
	fmt.Println("+++SetupSuite+++")
}

// 运行套件内所有测试函数后，执行且只执行一次该函数。
func (suite *ExampleTestSuite) TearDownSuite() {
	fmt.Println("+++TearDownSuite+++")
}

// 运行套件内的每个测试前，都会执行一次该函数。
func (suite *ExampleTestSuite) SetupTest() {
	fmt.Println("+++SetupTest+++")
}

// 运行套件内的每个测试后，都会执行一次该函数。
func (suite *ExampleTestSuite) TearDownTest() {
	fmt.Println("+++TearDownTest+++")
}

// 运行套件内的每个测试前，都会执行一次该函数。
func (suite *ExampleTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("+++BeforeTest(suiteName=%s, testName=%s)+++\n", suiteName, testName)
}

// 运行套件内的每个测试后，都会执行一次该函数。
func (suite *ExampleTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("+++AfterTest(suiteName=%s, testName=%s)+++\n", suiteName, testName)
}

// 运行套件内所有测试函数后，执行且只执行一次该函数，可以获取执行结果（起止时间、是否通过）相关信息。
func (suite *ExampleTestSuite) HandleStats(suiteName string, stats *suite.SuiteInformation) {
	fmt.Printf("+++HandleStats(suiteName=%s, stats=%+v)+++\n", suiteName, stats)
}

// 测试套件内，所有以 Test 开头的方法都会作为测试运行
func (suite *ExampleTestSuite) TestExample1() {
	fmt.Println("+++TestExample1+++")
	assert.True(suite.T(), true)
}

// 测试套件内，所有以 Test 开头的方法都会作为测试运行
// 注意：可以使用 suite.Suite 导出的断言函数，以方便测试
func (suite *ExampleTestSuite) TestExample2() {
	fmt.Println("+++TestExample1+++")
	suite.True(true)
}

// 为了让 go test 运行这个套件，我们需要创建一个正常的测试函数并将套件的指针传递给 suite.Run 函数
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
