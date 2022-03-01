// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kafkaconnectiface provides an interface to enable mocking the Managed Streaming for Kafka Connect service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kafkaconnectiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kafkaconnect"
)

// KafkaConnectAPI provides an interface to enable mocking the
// kafkaconnect.KafkaConnect service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Managed Streaming for Kafka Connect.
//    func myFunc(svc kafkaconnectiface.KafkaConnectAPI) bool {
//        // Make svc.CreateConnector request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := kafkaconnect.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKafkaConnectClient struct {
//        kafkaconnectiface.KafkaConnectAPI
//    }
//    func (m *mockKafkaConnectClient) CreateConnector(input *kafkaconnect.CreateConnectorInput) (*kafkaconnect.CreateConnectorOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKafkaConnectClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type KafkaConnectAPI interface {
	CreateConnector(*kafkaconnect.CreateConnectorInput) (*kafkaconnect.CreateConnectorOutput, error)
	CreateConnectorWithContext(aws.Context, *kafkaconnect.CreateConnectorInput, ...request.Option) (*kafkaconnect.CreateConnectorOutput, error)
	CreateConnectorRequest(*kafkaconnect.CreateConnectorInput) (*request.Request, *kafkaconnect.CreateConnectorOutput)

	CreateCustomPlugin(*kafkaconnect.CreateCustomPluginInput) (*kafkaconnect.CreateCustomPluginOutput, error)
	CreateCustomPluginWithContext(aws.Context, *kafkaconnect.CreateCustomPluginInput, ...request.Option) (*kafkaconnect.CreateCustomPluginOutput, error)
	CreateCustomPluginRequest(*kafkaconnect.CreateCustomPluginInput) (*request.Request, *kafkaconnect.CreateCustomPluginOutput)

	CreateWorkerConfiguration(*kafkaconnect.CreateWorkerConfigurationInput) (*kafkaconnect.CreateWorkerConfigurationOutput, error)
	CreateWorkerConfigurationWithContext(aws.Context, *kafkaconnect.CreateWorkerConfigurationInput, ...request.Option) (*kafkaconnect.CreateWorkerConfigurationOutput, error)
	CreateWorkerConfigurationRequest(*kafkaconnect.CreateWorkerConfigurationInput) (*request.Request, *kafkaconnect.CreateWorkerConfigurationOutput)

	DeleteConnector(*kafkaconnect.DeleteConnectorInput) (*kafkaconnect.DeleteConnectorOutput, error)
	DeleteConnectorWithContext(aws.Context, *kafkaconnect.DeleteConnectorInput, ...request.Option) (*kafkaconnect.DeleteConnectorOutput, error)
	DeleteConnectorRequest(*kafkaconnect.DeleteConnectorInput) (*request.Request, *kafkaconnect.DeleteConnectorOutput)

	DeleteCustomPlugin(*kafkaconnect.DeleteCustomPluginInput) (*kafkaconnect.DeleteCustomPluginOutput, error)
	DeleteCustomPluginWithContext(aws.Context, *kafkaconnect.DeleteCustomPluginInput, ...request.Option) (*kafkaconnect.DeleteCustomPluginOutput, error)
	DeleteCustomPluginRequest(*kafkaconnect.DeleteCustomPluginInput) (*request.Request, *kafkaconnect.DeleteCustomPluginOutput)

	DescribeConnector(*kafkaconnect.DescribeConnectorInput) (*kafkaconnect.DescribeConnectorOutput, error)
	DescribeConnectorWithContext(aws.Context, *kafkaconnect.DescribeConnectorInput, ...request.Option) (*kafkaconnect.DescribeConnectorOutput, error)
	DescribeConnectorRequest(*kafkaconnect.DescribeConnectorInput) (*request.Request, *kafkaconnect.DescribeConnectorOutput)

	DescribeCustomPlugin(*kafkaconnect.DescribeCustomPluginInput) (*kafkaconnect.DescribeCustomPluginOutput, error)
	DescribeCustomPluginWithContext(aws.Context, *kafkaconnect.DescribeCustomPluginInput, ...request.Option) (*kafkaconnect.DescribeCustomPluginOutput, error)
	DescribeCustomPluginRequest(*kafkaconnect.DescribeCustomPluginInput) (*request.Request, *kafkaconnect.DescribeCustomPluginOutput)

	DescribeWorkerConfiguration(*kafkaconnect.DescribeWorkerConfigurationInput) (*kafkaconnect.DescribeWorkerConfigurationOutput, error)
	DescribeWorkerConfigurationWithContext(aws.Context, *kafkaconnect.DescribeWorkerConfigurationInput, ...request.Option) (*kafkaconnect.DescribeWorkerConfigurationOutput, error)
	DescribeWorkerConfigurationRequest(*kafkaconnect.DescribeWorkerConfigurationInput) (*request.Request, *kafkaconnect.DescribeWorkerConfigurationOutput)

	ListConnectors(*kafkaconnect.ListConnectorsInput) (*kafkaconnect.ListConnectorsOutput, error)
	ListConnectorsWithContext(aws.Context, *kafkaconnect.ListConnectorsInput, ...request.Option) (*kafkaconnect.ListConnectorsOutput, error)
	ListConnectorsRequest(*kafkaconnect.ListConnectorsInput) (*request.Request, *kafkaconnect.ListConnectorsOutput)

	ListConnectorsPages(*kafkaconnect.ListConnectorsInput, func(*kafkaconnect.ListConnectorsOutput, bool) bool) error
	ListConnectorsPagesWithContext(aws.Context, *kafkaconnect.ListConnectorsInput, func(*kafkaconnect.ListConnectorsOutput, bool) bool, ...request.Option) error

	ListCustomPlugins(*kafkaconnect.ListCustomPluginsInput) (*kafkaconnect.ListCustomPluginsOutput, error)
	ListCustomPluginsWithContext(aws.Context, *kafkaconnect.ListCustomPluginsInput, ...request.Option) (*kafkaconnect.ListCustomPluginsOutput, error)
	ListCustomPluginsRequest(*kafkaconnect.ListCustomPluginsInput) (*request.Request, *kafkaconnect.ListCustomPluginsOutput)

	ListCustomPluginsPages(*kafkaconnect.ListCustomPluginsInput, func(*kafkaconnect.ListCustomPluginsOutput, bool) bool) error
	ListCustomPluginsPagesWithContext(aws.Context, *kafkaconnect.ListCustomPluginsInput, func(*kafkaconnect.ListCustomPluginsOutput, bool) bool, ...request.Option) error

	ListWorkerConfigurations(*kafkaconnect.ListWorkerConfigurationsInput) (*kafkaconnect.ListWorkerConfigurationsOutput, error)
	ListWorkerConfigurationsWithContext(aws.Context, *kafkaconnect.ListWorkerConfigurationsInput, ...request.Option) (*kafkaconnect.ListWorkerConfigurationsOutput, error)
	ListWorkerConfigurationsRequest(*kafkaconnect.ListWorkerConfigurationsInput) (*request.Request, *kafkaconnect.ListWorkerConfigurationsOutput)

	ListWorkerConfigurationsPages(*kafkaconnect.ListWorkerConfigurationsInput, func(*kafkaconnect.ListWorkerConfigurationsOutput, bool) bool) error
	ListWorkerConfigurationsPagesWithContext(aws.Context, *kafkaconnect.ListWorkerConfigurationsInput, func(*kafkaconnect.ListWorkerConfigurationsOutput, bool) bool, ...request.Option) error

	UpdateConnector(*kafkaconnect.UpdateConnectorInput) (*kafkaconnect.UpdateConnectorOutput, error)
	UpdateConnectorWithContext(aws.Context, *kafkaconnect.UpdateConnectorInput, ...request.Option) (*kafkaconnect.UpdateConnectorOutput, error)
	UpdateConnectorRequest(*kafkaconnect.UpdateConnectorInput) (*request.Request, *kafkaconnect.UpdateConnectorOutput)
}

var _ KafkaConnectAPI = (*kafkaconnect.KafkaConnect)(nil)
