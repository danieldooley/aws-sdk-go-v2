// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package glueiface provides an interface to enable mocking the AWS Glue service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package glueiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
)

// GlueAPI provides an interface to enable mocking the
// glue.Glue service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Glue.
//    func myFunc(svc glueiface.GlueAPI) bool {
//        // Make svc.BatchCreatePartition request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := glue.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockGlueClient struct {
//        glueiface.GlueAPI
//    }
//    func (m *mockGlueClient) BatchCreatePartition(input *glue.BatchCreatePartitionInput) (*glue.BatchCreatePartitionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockGlueClient{}
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
type GlueAPI interface {
	BatchCreatePartition(*glue.BatchCreatePartitionInput) (*glue.BatchCreatePartitionOutput, error)
	BatchCreatePartitionWithContext(aws.Context, *glue.BatchCreatePartitionInput, ...aws.Option) (*glue.BatchCreatePartitionOutput, error)
	BatchCreatePartitionRequest(*glue.BatchCreatePartitionInput) (*aws.Request, *glue.BatchCreatePartitionOutput)

	BatchDeleteConnection(*glue.BatchDeleteConnectionInput) (*glue.BatchDeleteConnectionOutput, error)
	BatchDeleteConnectionWithContext(aws.Context, *glue.BatchDeleteConnectionInput, ...aws.Option) (*glue.BatchDeleteConnectionOutput, error)
	BatchDeleteConnectionRequest(*glue.BatchDeleteConnectionInput) (*aws.Request, *glue.BatchDeleteConnectionOutput)

	BatchDeletePartition(*glue.BatchDeletePartitionInput) (*glue.BatchDeletePartitionOutput, error)
	BatchDeletePartitionWithContext(aws.Context, *glue.BatchDeletePartitionInput, ...aws.Option) (*glue.BatchDeletePartitionOutput, error)
	BatchDeletePartitionRequest(*glue.BatchDeletePartitionInput) (*aws.Request, *glue.BatchDeletePartitionOutput)

	BatchDeleteTable(*glue.BatchDeleteTableInput) (*glue.BatchDeleteTableOutput, error)
	BatchDeleteTableWithContext(aws.Context, *glue.BatchDeleteTableInput, ...aws.Option) (*glue.BatchDeleteTableOutput, error)
	BatchDeleteTableRequest(*glue.BatchDeleteTableInput) (*aws.Request, *glue.BatchDeleteTableOutput)

	BatchGetPartition(*glue.BatchGetPartitionInput) (*glue.BatchGetPartitionOutput, error)
	BatchGetPartitionWithContext(aws.Context, *glue.BatchGetPartitionInput, ...aws.Option) (*glue.BatchGetPartitionOutput, error)
	BatchGetPartitionRequest(*glue.BatchGetPartitionInput) (*aws.Request, *glue.BatchGetPartitionOutput)

	CreateClassifier(*glue.CreateClassifierInput) (*glue.CreateClassifierOutput, error)
	CreateClassifierWithContext(aws.Context, *glue.CreateClassifierInput, ...aws.Option) (*glue.CreateClassifierOutput, error)
	CreateClassifierRequest(*glue.CreateClassifierInput) (*aws.Request, *glue.CreateClassifierOutput)

	CreateConnection(*glue.CreateConnectionInput) (*glue.CreateConnectionOutput, error)
	CreateConnectionWithContext(aws.Context, *glue.CreateConnectionInput, ...aws.Option) (*glue.CreateConnectionOutput, error)
	CreateConnectionRequest(*glue.CreateConnectionInput) (*aws.Request, *glue.CreateConnectionOutput)

	CreateCrawler(*glue.CreateCrawlerInput) (*glue.CreateCrawlerOutput, error)
	CreateCrawlerWithContext(aws.Context, *glue.CreateCrawlerInput, ...aws.Option) (*glue.CreateCrawlerOutput, error)
	CreateCrawlerRequest(*glue.CreateCrawlerInput) (*aws.Request, *glue.CreateCrawlerOutput)

	CreateDatabase(*glue.CreateDatabaseInput) (*glue.CreateDatabaseOutput, error)
	CreateDatabaseWithContext(aws.Context, *glue.CreateDatabaseInput, ...aws.Option) (*glue.CreateDatabaseOutput, error)
	CreateDatabaseRequest(*glue.CreateDatabaseInput) (*aws.Request, *glue.CreateDatabaseOutput)

	CreateDevEndpoint(*glue.CreateDevEndpointInput) (*glue.CreateDevEndpointOutput, error)
	CreateDevEndpointWithContext(aws.Context, *glue.CreateDevEndpointInput, ...aws.Option) (*glue.CreateDevEndpointOutput, error)
	CreateDevEndpointRequest(*glue.CreateDevEndpointInput) (*aws.Request, *glue.CreateDevEndpointOutput)

	CreateJob(*glue.CreateJobInput) (*glue.CreateJobOutput, error)
	CreateJobWithContext(aws.Context, *glue.CreateJobInput, ...aws.Option) (*glue.CreateJobOutput, error)
	CreateJobRequest(*glue.CreateJobInput) (*aws.Request, *glue.CreateJobOutput)

	CreatePartition(*glue.CreatePartitionInput) (*glue.CreatePartitionOutput, error)
	CreatePartitionWithContext(aws.Context, *glue.CreatePartitionInput, ...aws.Option) (*glue.CreatePartitionOutput, error)
	CreatePartitionRequest(*glue.CreatePartitionInput) (*aws.Request, *glue.CreatePartitionOutput)

	CreateScript(*glue.CreateScriptInput) (*glue.CreateScriptOutput, error)
	CreateScriptWithContext(aws.Context, *glue.CreateScriptInput, ...aws.Option) (*glue.CreateScriptOutput, error)
	CreateScriptRequest(*glue.CreateScriptInput) (*aws.Request, *glue.CreateScriptOutput)

	CreateTable(*glue.CreateTableInput) (*glue.CreateTableOutput, error)
	CreateTableWithContext(aws.Context, *glue.CreateTableInput, ...aws.Option) (*glue.CreateTableOutput, error)
	CreateTableRequest(*glue.CreateTableInput) (*aws.Request, *glue.CreateTableOutput)

	CreateTrigger(*glue.CreateTriggerInput) (*glue.CreateTriggerOutput, error)
	CreateTriggerWithContext(aws.Context, *glue.CreateTriggerInput, ...aws.Option) (*glue.CreateTriggerOutput, error)
	CreateTriggerRequest(*glue.CreateTriggerInput) (*aws.Request, *glue.CreateTriggerOutput)

	CreateUserDefinedFunction(*glue.CreateUserDefinedFunctionInput) (*glue.CreateUserDefinedFunctionOutput, error)
	CreateUserDefinedFunctionWithContext(aws.Context, *glue.CreateUserDefinedFunctionInput, ...aws.Option) (*glue.CreateUserDefinedFunctionOutput, error)
	CreateUserDefinedFunctionRequest(*glue.CreateUserDefinedFunctionInput) (*aws.Request, *glue.CreateUserDefinedFunctionOutput)

	DeleteClassifier(*glue.DeleteClassifierInput) (*glue.DeleteClassifierOutput, error)
	DeleteClassifierWithContext(aws.Context, *glue.DeleteClassifierInput, ...aws.Option) (*glue.DeleteClassifierOutput, error)
	DeleteClassifierRequest(*glue.DeleteClassifierInput) (*aws.Request, *glue.DeleteClassifierOutput)

	DeleteConnection(*glue.DeleteConnectionInput) (*glue.DeleteConnectionOutput, error)
	DeleteConnectionWithContext(aws.Context, *glue.DeleteConnectionInput, ...aws.Option) (*glue.DeleteConnectionOutput, error)
	DeleteConnectionRequest(*glue.DeleteConnectionInput) (*aws.Request, *glue.DeleteConnectionOutput)

	DeleteCrawler(*glue.DeleteCrawlerInput) (*glue.DeleteCrawlerOutput, error)
	DeleteCrawlerWithContext(aws.Context, *glue.DeleteCrawlerInput, ...aws.Option) (*glue.DeleteCrawlerOutput, error)
	DeleteCrawlerRequest(*glue.DeleteCrawlerInput) (*aws.Request, *glue.DeleteCrawlerOutput)

	DeleteDatabase(*glue.DeleteDatabaseInput) (*glue.DeleteDatabaseOutput, error)
	DeleteDatabaseWithContext(aws.Context, *glue.DeleteDatabaseInput, ...aws.Option) (*glue.DeleteDatabaseOutput, error)
	DeleteDatabaseRequest(*glue.DeleteDatabaseInput) (*aws.Request, *glue.DeleteDatabaseOutput)

	DeleteDevEndpoint(*glue.DeleteDevEndpointInput) (*glue.DeleteDevEndpointOutput, error)
	DeleteDevEndpointWithContext(aws.Context, *glue.DeleteDevEndpointInput, ...aws.Option) (*glue.DeleteDevEndpointOutput, error)
	DeleteDevEndpointRequest(*glue.DeleteDevEndpointInput) (*aws.Request, *glue.DeleteDevEndpointOutput)

	DeleteJob(*glue.DeleteJobInput) (*glue.DeleteJobOutput, error)
	DeleteJobWithContext(aws.Context, *glue.DeleteJobInput, ...aws.Option) (*glue.DeleteJobOutput, error)
	DeleteJobRequest(*glue.DeleteJobInput) (*aws.Request, *glue.DeleteJobOutput)

	DeletePartition(*glue.DeletePartitionInput) (*glue.DeletePartitionOutput, error)
	DeletePartitionWithContext(aws.Context, *glue.DeletePartitionInput, ...aws.Option) (*glue.DeletePartitionOutput, error)
	DeletePartitionRequest(*glue.DeletePartitionInput) (*aws.Request, *glue.DeletePartitionOutput)

	DeleteTable(*glue.DeleteTableInput) (*glue.DeleteTableOutput, error)
	DeleteTableWithContext(aws.Context, *glue.DeleteTableInput, ...aws.Option) (*glue.DeleteTableOutput, error)
	DeleteTableRequest(*glue.DeleteTableInput) (*aws.Request, *glue.DeleteTableOutput)

	DeleteTrigger(*glue.DeleteTriggerInput) (*glue.DeleteTriggerOutput, error)
	DeleteTriggerWithContext(aws.Context, *glue.DeleteTriggerInput, ...aws.Option) (*glue.DeleteTriggerOutput, error)
	DeleteTriggerRequest(*glue.DeleteTriggerInput) (*aws.Request, *glue.DeleteTriggerOutput)

	DeleteUserDefinedFunction(*glue.DeleteUserDefinedFunctionInput) (*glue.DeleteUserDefinedFunctionOutput, error)
	DeleteUserDefinedFunctionWithContext(aws.Context, *glue.DeleteUserDefinedFunctionInput, ...aws.Option) (*glue.DeleteUserDefinedFunctionOutput, error)
	DeleteUserDefinedFunctionRequest(*glue.DeleteUserDefinedFunctionInput) (*aws.Request, *glue.DeleteUserDefinedFunctionOutput)

	GetCatalogImportStatus(*glue.GetCatalogImportStatusInput) (*glue.GetCatalogImportStatusOutput, error)
	GetCatalogImportStatusWithContext(aws.Context, *glue.GetCatalogImportStatusInput, ...aws.Option) (*glue.GetCatalogImportStatusOutput, error)
	GetCatalogImportStatusRequest(*glue.GetCatalogImportStatusInput) (*aws.Request, *glue.GetCatalogImportStatusOutput)

	GetClassifier(*glue.GetClassifierInput) (*glue.GetClassifierOutput, error)
	GetClassifierWithContext(aws.Context, *glue.GetClassifierInput, ...aws.Option) (*glue.GetClassifierOutput, error)
	GetClassifierRequest(*glue.GetClassifierInput) (*aws.Request, *glue.GetClassifierOutput)

	GetClassifiers(*glue.GetClassifiersInput) (*glue.GetClassifiersOutput, error)
	GetClassifiersWithContext(aws.Context, *glue.GetClassifiersInput, ...aws.Option) (*glue.GetClassifiersOutput, error)
	GetClassifiersRequest(*glue.GetClassifiersInput) (*aws.Request, *glue.GetClassifiersOutput)

	GetClassifiersPages(*glue.GetClassifiersInput, func(*glue.GetClassifiersOutput, bool) bool) error
	GetClassifiersPagesWithContext(aws.Context, *glue.GetClassifiersInput, func(*glue.GetClassifiersOutput, bool) bool, ...aws.Option) error

	GetConnection(*glue.GetConnectionInput) (*glue.GetConnectionOutput, error)
	GetConnectionWithContext(aws.Context, *glue.GetConnectionInput, ...aws.Option) (*glue.GetConnectionOutput, error)
	GetConnectionRequest(*glue.GetConnectionInput) (*aws.Request, *glue.GetConnectionOutput)

	GetConnections(*glue.GetConnectionsInput) (*glue.GetConnectionsOutput, error)
	GetConnectionsWithContext(aws.Context, *glue.GetConnectionsInput, ...aws.Option) (*glue.GetConnectionsOutput, error)
	GetConnectionsRequest(*glue.GetConnectionsInput) (*aws.Request, *glue.GetConnectionsOutput)

	GetConnectionsPages(*glue.GetConnectionsInput, func(*glue.GetConnectionsOutput, bool) bool) error
	GetConnectionsPagesWithContext(aws.Context, *glue.GetConnectionsInput, func(*glue.GetConnectionsOutput, bool) bool, ...aws.Option) error

	GetCrawler(*glue.GetCrawlerInput) (*glue.GetCrawlerOutput, error)
	GetCrawlerWithContext(aws.Context, *glue.GetCrawlerInput, ...aws.Option) (*glue.GetCrawlerOutput, error)
	GetCrawlerRequest(*glue.GetCrawlerInput) (*aws.Request, *glue.GetCrawlerOutput)

	GetCrawlerMetrics(*glue.GetCrawlerMetricsInput) (*glue.GetCrawlerMetricsOutput, error)
	GetCrawlerMetricsWithContext(aws.Context, *glue.GetCrawlerMetricsInput, ...aws.Option) (*glue.GetCrawlerMetricsOutput, error)
	GetCrawlerMetricsRequest(*glue.GetCrawlerMetricsInput) (*aws.Request, *glue.GetCrawlerMetricsOutput)

	GetCrawlerMetricsPages(*glue.GetCrawlerMetricsInput, func(*glue.GetCrawlerMetricsOutput, bool) bool) error
	GetCrawlerMetricsPagesWithContext(aws.Context, *glue.GetCrawlerMetricsInput, func(*glue.GetCrawlerMetricsOutput, bool) bool, ...aws.Option) error

	GetCrawlers(*glue.GetCrawlersInput) (*glue.GetCrawlersOutput, error)
	GetCrawlersWithContext(aws.Context, *glue.GetCrawlersInput, ...aws.Option) (*glue.GetCrawlersOutput, error)
	GetCrawlersRequest(*glue.GetCrawlersInput) (*aws.Request, *glue.GetCrawlersOutput)

	GetCrawlersPages(*glue.GetCrawlersInput, func(*glue.GetCrawlersOutput, bool) bool) error
	GetCrawlersPagesWithContext(aws.Context, *glue.GetCrawlersInput, func(*glue.GetCrawlersOutput, bool) bool, ...aws.Option) error

	GetDatabase(*glue.GetDatabaseInput) (*glue.GetDatabaseOutput, error)
	GetDatabaseWithContext(aws.Context, *glue.GetDatabaseInput, ...aws.Option) (*glue.GetDatabaseOutput, error)
	GetDatabaseRequest(*glue.GetDatabaseInput) (*aws.Request, *glue.GetDatabaseOutput)

	GetDatabases(*glue.GetDatabasesInput) (*glue.GetDatabasesOutput, error)
	GetDatabasesWithContext(aws.Context, *glue.GetDatabasesInput, ...aws.Option) (*glue.GetDatabasesOutput, error)
	GetDatabasesRequest(*glue.GetDatabasesInput) (*aws.Request, *glue.GetDatabasesOutput)

	GetDatabasesPages(*glue.GetDatabasesInput, func(*glue.GetDatabasesOutput, bool) bool) error
	GetDatabasesPagesWithContext(aws.Context, *glue.GetDatabasesInput, func(*glue.GetDatabasesOutput, bool) bool, ...aws.Option) error

	GetDataflowGraph(*glue.GetDataflowGraphInput) (*glue.GetDataflowGraphOutput, error)
	GetDataflowGraphWithContext(aws.Context, *glue.GetDataflowGraphInput, ...aws.Option) (*glue.GetDataflowGraphOutput, error)
	GetDataflowGraphRequest(*glue.GetDataflowGraphInput) (*aws.Request, *glue.GetDataflowGraphOutput)

	GetDevEndpoint(*glue.GetDevEndpointInput) (*glue.GetDevEndpointOutput, error)
	GetDevEndpointWithContext(aws.Context, *glue.GetDevEndpointInput, ...aws.Option) (*glue.GetDevEndpointOutput, error)
	GetDevEndpointRequest(*glue.GetDevEndpointInput) (*aws.Request, *glue.GetDevEndpointOutput)

	GetDevEndpoints(*glue.GetDevEndpointsInput) (*glue.GetDevEndpointsOutput, error)
	GetDevEndpointsWithContext(aws.Context, *glue.GetDevEndpointsInput, ...aws.Option) (*glue.GetDevEndpointsOutput, error)
	GetDevEndpointsRequest(*glue.GetDevEndpointsInput) (*aws.Request, *glue.GetDevEndpointsOutput)

	GetDevEndpointsPages(*glue.GetDevEndpointsInput, func(*glue.GetDevEndpointsOutput, bool) bool) error
	GetDevEndpointsPagesWithContext(aws.Context, *glue.GetDevEndpointsInput, func(*glue.GetDevEndpointsOutput, bool) bool, ...aws.Option) error

	GetJob(*glue.GetJobInput) (*glue.GetJobOutput, error)
	GetJobWithContext(aws.Context, *glue.GetJobInput, ...aws.Option) (*glue.GetJobOutput, error)
	GetJobRequest(*glue.GetJobInput) (*aws.Request, *glue.GetJobOutput)

	GetJobRun(*glue.GetJobRunInput) (*glue.GetJobRunOutput, error)
	GetJobRunWithContext(aws.Context, *glue.GetJobRunInput, ...aws.Option) (*glue.GetJobRunOutput, error)
	GetJobRunRequest(*glue.GetJobRunInput) (*aws.Request, *glue.GetJobRunOutput)

	GetJobRuns(*glue.GetJobRunsInput) (*glue.GetJobRunsOutput, error)
	GetJobRunsWithContext(aws.Context, *glue.GetJobRunsInput, ...aws.Option) (*glue.GetJobRunsOutput, error)
	GetJobRunsRequest(*glue.GetJobRunsInput) (*aws.Request, *glue.GetJobRunsOutput)

	GetJobRunsPages(*glue.GetJobRunsInput, func(*glue.GetJobRunsOutput, bool) bool) error
	GetJobRunsPagesWithContext(aws.Context, *glue.GetJobRunsInput, func(*glue.GetJobRunsOutput, bool) bool, ...aws.Option) error

	GetJobs(*glue.GetJobsInput) (*glue.GetJobsOutput, error)
	GetJobsWithContext(aws.Context, *glue.GetJobsInput, ...aws.Option) (*glue.GetJobsOutput, error)
	GetJobsRequest(*glue.GetJobsInput) (*aws.Request, *glue.GetJobsOutput)

	GetJobsPages(*glue.GetJobsInput, func(*glue.GetJobsOutput, bool) bool) error
	GetJobsPagesWithContext(aws.Context, *glue.GetJobsInput, func(*glue.GetJobsOutput, bool) bool, ...aws.Option) error

	GetMapping(*glue.GetMappingInput) (*glue.GetMappingOutput, error)
	GetMappingWithContext(aws.Context, *glue.GetMappingInput, ...aws.Option) (*glue.GetMappingOutput, error)
	GetMappingRequest(*glue.GetMappingInput) (*aws.Request, *glue.GetMappingOutput)

	GetPartition(*glue.GetPartitionInput) (*glue.GetPartitionOutput, error)
	GetPartitionWithContext(aws.Context, *glue.GetPartitionInput, ...aws.Option) (*glue.GetPartitionOutput, error)
	GetPartitionRequest(*glue.GetPartitionInput) (*aws.Request, *glue.GetPartitionOutput)

	GetPartitions(*glue.GetPartitionsInput) (*glue.GetPartitionsOutput, error)
	GetPartitionsWithContext(aws.Context, *glue.GetPartitionsInput, ...aws.Option) (*glue.GetPartitionsOutput, error)
	GetPartitionsRequest(*glue.GetPartitionsInput) (*aws.Request, *glue.GetPartitionsOutput)

	GetPartitionsPages(*glue.GetPartitionsInput, func(*glue.GetPartitionsOutput, bool) bool) error
	GetPartitionsPagesWithContext(aws.Context, *glue.GetPartitionsInput, func(*glue.GetPartitionsOutput, bool) bool, ...aws.Option) error

	GetPlan(*glue.GetPlanInput) (*glue.GetPlanOutput, error)
	GetPlanWithContext(aws.Context, *glue.GetPlanInput, ...aws.Option) (*glue.GetPlanOutput, error)
	GetPlanRequest(*glue.GetPlanInput) (*aws.Request, *glue.GetPlanOutput)

	GetTable(*glue.GetTableInput) (*glue.GetTableOutput, error)
	GetTableWithContext(aws.Context, *glue.GetTableInput, ...aws.Option) (*glue.GetTableOutput, error)
	GetTableRequest(*glue.GetTableInput) (*aws.Request, *glue.GetTableOutput)

	GetTableVersions(*glue.GetTableVersionsInput) (*glue.GetTableVersionsOutput, error)
	GetTableVersionsWithContext(aws.Context, *glue.GetTableVersionsInput, ...aws.Option) (*glue.GetTableVersionsOutput, error)
	GetTableVersionsRequest(*glue.GetTableVersionsInput) (*aws.Request, *glue.GetTableVersionsOutput)

	GetTableVersionsPages(*glue.GetTableVersionsInput, func(*glue.GetTableVersionsOutput, bool) bool) error
	GetTableVersionsPagesWithContext(aws.Context, *glue.GetTableVersionsInput, func(*glue.GetTableVersionsOutput, bool) bool, ...aws.Option) error

	GetTables(*glue.GetTablesInput) (*glue.GetTablesOutput, error)
	GetTablesWithContext(aws.Context, *glue.GetTablesInput, ...aws.Option) (*glue.GetTablesOutput, error)
	GetTablesRequest(*glue.GetTablesInput) (*aws.Request, *glue.GetTablesOutput)

	GetTablesPages(*glue.GetTablesInput, func(*glue.GetTablesOutput, bool) bool) error
	GetTablesPagesWithContext(aws.Context, *glue.GetTablesInput, func(*glue.GetTablesOutput, bool) bool, ...aws.Option) error

	GetTrigger(*glue.GetTriggerInput) (*glue.GetTriggerOutput, error)
	GetTriggerWithContext(aws.Context, *glue.GetTriggerInput, ...aws.Option) (*glue.GetTriggerOutput, error)
	GetTriggerRequest(*glue.GetTriggerInput) (*aws.Request, *glue.GetTriggerOutput)

	GetTriggers(*glue.GetTriggersInput) (*glue.GetTriggersOutput, error)
	GetTriggersWithContext(aws.Context, *glue.GetTriggersInput, ...aws.Option) (*glue.GetTriggersOutput, error)
	GetTriggersRequest(*glue.GetTriggersInput) (*aws.Request, *glue.GetTriggersOutput)

	GetTriggersPages(*glue.GetTriggersInput, func(*glue.GetTriggersOutput, bool) bool) error
	GetTriggersPagesWithContext(aws.Context, *glue.GetTriggersInput, func(*glue.GetTriggersOutput, bool) bool, ...aws.Option) error

	GetUserDefinedFunction(*glue.GetUserDefinedFunctionInput) (*glue.GetUserDefinedFunctionOutput, error)
	GetUserDefinedFunctionWithContext(aws.Context, *glue.GetUserDefinedFunctionInput, ...aws.Option) (*glue.GetUserDefinedFunctionOutput, error)
	GetUserDefinedFunctionRequest(*glue.GetUserDefinedFunctionInput) (*aws.Request, *glue.GetUserDefinedFunctionOutput)

	GetUserDefinedFunctions(*glue.GetUserDefinedFunctionsInput) (*glue.GetUserDefinedFunctionsOutput, error)
	GetUserDefinedFunctionsWithContext(aws.Context, *glue.GetUserDefinedFunctionsInput, ...aws.Option) (*glue.GetUserDefinedFunctionsOutput, error)
	GetUserDefinedFunctionsRequest(*glue.GetUserDefinedFunctionsInput) (*aws.Request, *glue.GetUserDefinedFunctionsOutput)

	GetUserDefinedFunctionsPages(*glue.GetUserDefinedFunctionsInput, func(*glue.GetUserDefinedFunctionsOutput, bool) bool) error
	GetUserDefinedFunctionsPagesWithContext(aws.Context, *glue.GetUserDefinedFunctionsInput, func(*glue.GetUserDefinedFunctionsOutput, bool) bool, ...aws.Option) error

	ImportCatalogToGlue(*glue.ImportCatalogToGlueInput) (*glue.ImportCatalogToGlueOutput, error)
	ImportCatalogToGlueWithContext(aws.Context, *glue.ImportCatalogToGlueInput, ...aws.Option) (*glue.ImportCatalogToGlueOutput, error)
	ImportCatalogToGlueRequest(*glue.ImportCatalogToGlueInput) (*aws.Request, *glue.ImportCatalogToGlueOutput)

	ResetJobBookmark(*glue.ResetJobBookmarkInput) (*glue.ResetJobBookmarkOutput, error)
	ResetJobBookmarkWithContext(aws.Context, *glue.ResetJobBookmarkInput, ...aws.Option) (*glue.ResetJobBookmarkOutput, error)
	ResetJobBookmarkRequest(*glue.ResetJobBookmarkInput) (*aws.Request, *glue.ResetJobBookmarkOutput)

	StartCrawler(*glue.StartCrawlerInput) (*glue.StartCrawlerOutput, error)
	StartCrawlerWithContext(aws.Context, *glue.StartCrawlerInput, ...aws.Option) (*glue.StartCrawlerOutput, error)
	StartCrawlerRequest(*glue.StartCrawlerInput) (*aws.Request, *glue.StartCrawlerOutput)

	StartCrawlerSchedule(*glue.StartCrawlerScheduleInput) (*glue.StartCrawlerScheduleOutput, error)
	StartCrawlerScheduleWithContext(aws.Context, *glue.StartCrawlerScheduleInput, ...aws.Option) (*glue.StartCrawlerScheduleOutput, error)
	StartCrawlerScheduleRequest(*glue.StartCrawlerScheduleInput) (*aws.Request, *glue.StartCrawlerScheduleOutput)

	StartJobRun(*glue.StartJobRunInput) (*glue.StartJobRunOutput, error)
	StartJobRunWithContext(aws.Context, *glue.StartJobRunInput, ...aws.Option) (*glue.StartJobRunOutput, error)
	StartJobRunRequest(*glue.StartJobRunInput) (*aws.Request, *glue.StartJobRunOutput)

	StartTrigger(*glue.StartTriggerInput) (*glue.StartTriggerOutput, error)
	StartTriggerWithContext(aws.Context, *glue.StartTriggerInput, ...aws.Option) (*glue.StartTriggerOutput, error)
	StartTriggerRequest(*glue.StartTriggerInput) (*aws.Request, *glue.StartTriggerOutput)

	StopCrawler(*glue.StopCrawlerInput) (*glue.StopCrawlerOutput, error)
	StopCrawlerWithContext(aws.Context, *glue.StopCrawlerInput, ...aws.Option) (*glue.StopCrawlerOutput, error)
	StopCrawlerRequest(*glue.StopCrawlerInput) (*aws.Request, *glue.StopCrawlerOutput)

	StopCrawlerSchedule(*glue.StopCrawlerScheduleInput) (*glue.StopCrawlerScheduleOutput, error)
	StopCrawlerScheduleWithContext(aws.Context, *glue.StopCrawlerScheduleInput, ...aws.Option) (*glue.StopCrawlerScheduleOutput, error)
	StopCrawlerScheduleRequest(*glue.StopCrawlerScheduleInput) (*aws.Request, *glue.StopCrawlerScheduleOutput)

	StopTrigger(*glue.StopTriggerInput) (*glue.StopTriggerOutput, error)
	StopTriggerWithContext(aws.Context, *glue.StopTriggerInput, ...aws.Option) (*glue.StopTriggerOutput, error)
	StopTriggerRequest(*glue.StopTriggerInput) (*aws.Request, *glue.StopTriggerOutput)

	UpdateClassifier(*glue.UpdateClassifierInput) (*glue.UpdateClassifierOutput, error)
	UpdateClassifierWithContext(aws.Context, *glue.UpdateClassifierInput, ...aws.Option) (*glue.UpdateClassifierOutput, error)
	UpdateClassifierRequest(*glue.UpdateClassifierInput) (*aws.Request, *glue.UpdateClassifierOutput)

	UpdateConnection(*glue.UpdateConnectionInput) (*glue.UpdateConnectionOutput, error)
	UpdateConnectionWithContext(aws.Context, *glue.UpdateConnectionInput, ...aws.Option) (*glue.UpdateConnectionOutput, error)
	UpdateConnectionRequest(*glue.UpdateConnectionInput) (*aws.Request, *glue.UpdateConnectionOutput)

	UpdateCrawler(*glue.UpdateCrawlerInput) (*glue.UpdateCrawlerOutput, error)
	UpdateCrawlerWithContext(aws.Context, *glue.UpdateCrawlerInput, ...aws.Option) (*glue.UpdateCrawlerOutput, error)
	UpdateCrawlerRequest(*glue.UpdateCrawlerInput) (*aws.Request, *glue.UpdateCrawlerOutput)

	UpdateCrawlerSchedule(*glue.UpdateCrawlerScheduleInput) (*glue.UpdateCrawlerScheduleOutput, error)
	UpdateCrawlerScheduleWithContext(aws.Context, *glue.UpdateCrawlerScheduleInput, ...aws.Option) (*glue.UpdateCrawlerScheduleOutput, error)
	UpdateCrawlerScheduleRequest(*glue.UpdateCrawlerScheduleInput) (*aws.Request, *glue.UpdateCrawlerScheduleOutput)

	UpdateDatabase(*glue.UpdateDatabaseInput) (*glue.UpdateDatabaseOutput, error)
	UpdateDatabaseWithContext(aws.Context, *glue.UpdateDatabaseInput, ...aws.Option) (*glue.UpdateDatabaseOutput, error)
	UpdateDatabaseRequest(*glue.UpdateDatabaseInput) (*aws.Request, *glue.UpdateDatabaseOutput)

	UpdateDevEndpoint(*glue.UpdateDevEndpointInput) (*glue.UpdateDevEndpointOutput, error)
	UpdateDevEndpointWithContext(aws.Context, *glue.UpdateDevEndpointInput, ...aws.Option) (*glue.UpdateDevEndpointOutput, error)
	UpdateDevEndpointRequest(*glue.UpdateDevEndpointInput) (*aws.Request, *glue.UpdateDevEndpointOutput)

	UpdateJob(*glue.UpdateJobInput) (*glue.UpdateJobOutput, error)
	UpdateJobWithContext(aws.Context, *glue.UpdateJobInput, ...aws.Option) (*glue.UpdateJobOutput, error)
	UpdateJobRequest(*glue.UpdateJobInput) (*aws.Request, *glue.UpdateJobOutput)

	UpdatePartition(*glue.UpdatePartitionInput) (*glue.UpdatePartitionOutput, error)
	UpdatePartitionWithContext(aws.Context, *glue.UpdatePartitionInput, ...aws.Option) (*glue.UpdatePartitionOutput, error)
	UpdatePartitionRequest(*glue.UpdatePartitionInput) (*aws.Request, *glue.UpdatePartitionOutput)

	UpdateTable(*glue.UpdateTableInput) (*glue.UpdateTableOutput, error)
	UpdateTableWithContext(aws.Context, *glue.UpdateTableInput, ...aws.Option) (*glue.UpdateTableOutput, error)
	UpdateTableRequest(*glue.UpdateTableInput) (*aws.Request, *glue.UpdateTableOutput)

	UpdateTrigger(*glue.UpdateTriggerInput) (*glue.UpdateTriggerOutput, error)
	UpdateTriggerWithContext(aws.Context, *glue.UpdateTriggerInput, ...aws.Option) (*glue.UpdateTriggerOutput, error)
	UpdateTriggerRequest(*glue.UpdateTriggerInput) (*aws.Request, *glue.UpdateTriggerOutput)

	UpdateUserDefinedFunction(*glue.UpdateUserDefinedFunctionInput) (*glue.UpdateUserDefinedFunctionOutput, error)
	UpdateUserDefinedFunctionWithContext(aws.Context, *glue.UpdateUserDefinedFunctionInput, ...aws.Option) (*glue.UpdateUserDefinedFunctionOutput, error)
	UpdateUserDefinedFunctionRequest(*glue.UpdateUserDefinedFunctionInput) (*aws.Request, *glue.UpdateUserDefinedFunctionOutput)
}

var _ GlueAPI = (*glue.Glue)(nil)