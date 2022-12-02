// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Experiments, M Lflow Artifacts, M Lflow Databricks, M Lflow Metrics, M Lflow Runs, Model Version Comments, Model Versions, Registered Models, Registry Webhooks, Transition Requests, etc.
package mlflow

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewExperiments(client *client.DatabricksClient) *ExperimentsAPI {
	return &ExperimentsAPI{
		impl: &experimentsImpl{
			client: client,
		},
	}
}

type ExperimentsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(ExperimentsService)
	impl ExperimentsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *ExperimentsAPI) WithImpl(impl ExperimentsService) *ExperimentsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Experiments API implementation
func (a *ExperimentsAPI) Impl() ExperimentsService {
	return a.impl
}

// Create experiment
//
// Creates an experiment with a name. Returns the ID of the newly created
// experiment. Validates that another experiment with the same name does not
// already exist and fails if another experiment with the same name already
// exists.
//
// Throws `RESOURCE_ALREADY_EXISTS` if a experiment with the given name exists.
func (a *ExperimentsAPI) Create(ctx context.Context, request CreateExperiment) (*CreateExperimentResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete an experiment
//
// Marks an experiment and associated metadata, runs, metrics, params, and tags
// for deletion. If the experiment uses FileStore, artifacts associated with
// experiment are also deleted.
func (a *ExperimentsAPI) Delete(ctx context.Context, request DeleteExperiment) error {
	return a.impl.Delete(ctx, request)
}

// Delete an experiment
//
// Marks an experiment and associated metadata, runs, metrics, params, and tags
// for deletion. If the experiment uses FileStore, artifacts associated with
// experiment are also deleted.
func (a *ExperimentsAPI) DeleteByExperimentId(ctx context.Context, experimentId string) error {
	return a.impl.Delete(ctx, DeleteExperiment{
		ExperimentId: experimentId,
	})
}

// Get an experiment
//
// Gets metadata for an experiment. This method works on deleted experiments.
func (a *ExperimentsAPI) Get(ctx context.Context, request GetExperimentRequest) (*Experiment, error) {
	return a.impl.Get(ctx, request)
}

// Get an experiment
//
// Gets metadata for an experiment. This method works on deleted experiments.
func (a *ExperimentsAPI) GetByExperimentId(ctx context.Context, experimentId string) (*Experiment, error) {
	return a.impl.Get(ctx, GetExperimentRequest{
		ExperimentId: experimentId,
	})
}

// Get metadata
//
// "Gets metadata for an experiment.
//
// This endpoint will return deleted experiments, but prefers the active
// experiment if an active and deleted experiment share the same name. If
// multiple deleted\nexperiments share the same name, the API will return one of
// them.
//
// Throws `RESOURCE_DOES_NOT_EXIST` if no experiment with the specified name
// exists.S
func (a *ExperimentsAPI) GetByName(ctx context.Context, request GetByNameRequest) (*GetExperimentByNameResponse, error) {
	return a.impl.GetByName(ctx, request)
}

// Get metadata
//
// "Gets metadata for an experiment.
//
// This endpoint will return deleted experiments, but prefers the active
// experiment if an active and deleted experiment share the same name. If
// multiple deleted\nexperiments share the same name, the API will return one of
// them.
//
// Throws `RESOURCE_DOES_NOT_EXIST` if no experiment with the specified name
// exists.S
func (a *ExperimentsAPI) GetByNameByExperimentName(ctx context.Context, experimentName string) (*GetExperimentByNameResponse, error) {
	return a.impl.GetByName(ctx, GetByNameRequest{
		ExperimentName: experimentName,
	})
}

// List experiments
//
// Gets a list of all experiments.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ExperimentsAPI) ListAll(ctx context.Context, request ListExperimentsRequest) ([]Experiment, error) {
	var results []Experiment
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.List(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Experiments) == 0 {
			break
		}
		for _, v := range response.Experiments {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// Restores an experiment
//
// "Restore an experiment marked for deletion. This also restores\nassociated
// metadata, runs, metrics, params, and tags. If experiment uses FileStore,
// underlying\nartifacts associated with experiment are also restored.\n\nThrows
// `RESOURCE_DOES_NOT_EXIST` if experiment was never created or was permanently
// deleted.",
func (a *ExperimentsAPI) Restore(ctx context.Context, request RestoreExperiment) error {
	return a.impl.Restore(ctx, request)
}

// Restores an experiment
//
// "Restore an experiment marked for deletion. This also restores\nassociated
// metadata, runs, metrics, params, and tags. If experiment uses FileStore,
// underlying\nartifacts associated with experiment are also restored.\n\nThrows
// `RESOURCE_DOES_NOT_EXIST` if experiment was never created or was permanently
// deleted.",
func (a *ExperimentsAPI) RestoreByExperimentId(ctx context.Context, experimentId string) error {
	return a.impl.Restore(ctx, RestoreExperiment{
		ExperimentId: experimentId,
	})
}

// Search experiments
//
// Searches for experiments that satisfy specified search criteria.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ExperimentsAPI) SearchAll(ctx context.Context, request SearchExperiments) ([]Experiment, error) {
	var results []Experiment
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.Search(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Experiments) == 0 {
			break
		}
		for _, v := range response.Experiments {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// Set a tag
//
// Sets a tag on an experiment. Experiment tags are metadata that can be
// updated.
func (a *ExperimentsAPI) SetExperimentTag(ctx context.Context, request SetExperimentTag) error {
	return a.impl.SetExperimentTag(ctx, request)
}

// Update an experiment
//
// Updates experiment metadata.
func (a *ExperimentsAPI) Update(ctx context.Context, request UpdateExperiment) error {
	return a.impl.Update(ctx, request)
}

func NewMLflowArtifacts(client *client.DatabricksClient) *MLflowArtifactsAPI {
	return &MLflowArtifactsAPI{
		impl: &mLflowArtifactsImpl{
			client: client,
		},
	}
}

type MLflowArtifactsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(MLflowArtifactsService)
	impl MLflowArtifactsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *MLflowArtifactsAPI) WithImpl(impl MLflowArtifactsService) *MLflowArtifactsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level MLflowArtifacts API implementation
func (a *MLflowArtifactsAPI) Impl() MLflowArtifactsService {
	return a.impl
}

// Get all artifacts
//
// List artifacts for a run. Takes an optional `artifact_path` prefix. If it is
// specified, the response contains only artifacts with the specified prefix.",
//
// This method is generated by Databricks SDK Code Generator.
func (a *MLflowArtifactsAPI) ListAll(ctx context.Context, request ListArtifactsRequest) ([]FileInfo, error) {
	var results []FileInfo
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.List(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Files) == 0 {
			break
		}
		for _, v := range response.Files {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

func NewMLflowDatabricks(client *client.DatabricksClient) *MLflowDatabricksAPI {
	return &MLflowDatabricksAPI{
		impl: &mLflowDatabricksImpl{
			client: client,
		},
	}
}

// These endpoints are modified versions of the MLflow API that accept
// additional input parameters or return additional information.
type MLflowDatabricksAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(MLflowDatabricksService)
	impl MLflowDatabricksService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *MLflowDatabricksAPI) WithImpl(impl MLflowDatabricksService) *MLflowDatabricksAPI {
	a.impl = impl
	return a
}

// Impl returns low-level MLflowDatabricks API implementation
func (a *MLflowDatabricksAPI) Impl() MLflowDatabricksService {
	return a.impl
}

// Get model
//
// Get the details of a model. This is a Databricks Workspace version of the
// [MLflow
// endpoint](https://www.mlflow.org/docs/latest/rest-api.html#get-registeredmodel)
// that also returns the model's Databricks Workspace ID and the permission
// level of the requesting user on the model.
func (a *MLflowDatabricksAPI) Get(ctx context.Context, request GetMLflowDatabrickRequest) (*GetResponse, error) {
	return a.impl.Get(ctx, request)
}

// Get model
//
// Get the details of a model. This is a Databricks Workspace version of the
// [MLflow
// endpoint](https://www.mlflow.org/docs/latest/rest-api.html#get-registeredmodel)
// that also returns the model's Databricks Workspace ID and the permission
// level of the requesting user on the model.
func (a *MLflowDatabricksAPI) GetByName(ctx context.Context, name string) (*GetResponse, error) {
	return a.impl.Get(ctx, GetMLflowDatabrickRequest{
		Name: name,
	})
}

// Transition a stage
//
// Transition a model version's stage. This is a Databricks Workspace version of
// the [MLflow
// endpoint](https://www.mlflow.org/docs/latest/rest-api.html#transition-modelversion-stage)
// that also accepts a comment associated with the transition to be recorded.",
func (a *MLflowDatabricksAPI) TransitionStage(ctx context.Context, request TransitionModelVersionStageDatabricks) (*TransitionStageResponse, error) {
	return a.impl.TransitionStage(ctx, request)
}

func NewMLflowMetrics(client *client.DatabricksClient) *MLflowMetricsAPI {
	return &MLflowMetricsAPI{
		impl: &mLflowMetricsImpl{
			client: client,
		},
	}
}

type MLflowMetricsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(MLflowMetricsService)
	impl MLflowMetricsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *MLflowMetricsAPI) WithImpl(impl MLflowMetricsService) *MLflowMetricsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level MLflowMetrics API implementation
func (a *MLflowMetricsAPI) Impl() MLflowMetricsService {
	return a.impl
}

// Get all history
//
// Gets a list of all values for the specified metric for a given run.
func (a *MLflowMetricsAPI) GetHistory(ctx context.Context, request GetHistoryRequest) (*GetMetricHistoryResponse, error) {
	return a.impl.GetHistory(ctx, request)
}

func NewMLflowRuns(client *client.DatabricksClient) *MLflowRunsAPI {
	return &MLflowRunsAPI{
		impl: &mLflowRunsImpl{
			client: client,
		},
	}
}

type MLflowRunsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(MLflowRunsService)
	impl MLflowRunsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *MLflowRunsAPI) WithImpl(impl MLflowRunsService) *MLflowRunsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level MLflowRuns API implementation
func (a *MLflowRunsAPI) Impl() MLflowRunsService {
	return a.impl
}

// Create a run
//
// Creates a new run within an experiment. A run is usually a single execution
// of a machine learning or data ETL pipeline. MLflow uses runs to track the
// `mlflowParam`, `mlflowMetric` and `mlflowRunTag` associated with a single
// execution.
func (a *MLflowRunsAPI) Create(ctx context.Context, request CreateRun) (*CreateRunResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete a run
//
// Marks a run for deletion.
func (a *MLflowRunsAPI) Delete(ctx context.Context, request DeleteRun) error {
	return a.impl.Delete(ctx, request)
}

// Delete a run
//
// Marks a run for deletion.
func (a *MLflowRunsAPI) DeleteByRunId(ctx context.Context, runId string) error {
	return a.impl.Delete(ctx, DeleteRun{
		RunId: runId,
	})
}

// Delete a tag
//
// Deletes a tag on a run. Tags are run metadata that can be updated during a
// run and after a run completes.
func (a *MLflowRunsAPI) DeleteTag(ctx context.Context, request DeleteTag) error {
	return a.impl.DeleteTag(ctx, request)
}

// Get a run
//
// "Gets the metadata, metrics, params, and tags for a run. In the case where
// multiple metrics with the same key are logged for a run, return only the
// value with the latest timestamp.
//
// If there are multiple values with the latest timestamp, return the maximum of
// these values.
func (a *MLflowRunsAPI) Get(ctx context.Context, request GetRunRequest) (*GetRunResponse, error) {
	return a.impl.Get(ctx, request)
}

// Log a batch
//
// Logs a batch of metrics, params, and tags for a run. If any data failed to be
// persisted, the server will respond with an error (non-200 status code).
//
// In case of error (due to internal server error or an invalid request),
// partial data may be written.
//
// You can write metrics, params, and tags in interleaving fashion, but within a
// given entity type are guaranteed to follow the order specified in the request
// body.
//
// The overwrite behavior for metrics, params, and tags is as follows:
//
// * Metrics: metric values are never overwritten. Logging a metric (key, value,
// timestamp) appends to the set of values for the metric with the provided key.
//
// * Tags: tag values can be overwritten by successive writes to the same tag
// key. That is, if multiple tag values with the same key are provided in the
// same API request, the last-provided tag value is written. Logging the same
// tag (key, value) is permitted. Specifically, logging a tag is idempotent.
//
// * Parameters: once written, param values cannot be changed (attempting to
// overwrite a param value will result in an error). However, logging the same
// param (key, value) is permitted. Specifically, logging a param is idempotent.
//
// Request Limits ------------------------------- A single JSON-serialized API
// request may be up to 1 MB in size and contain:
//
// * No more than 1000 metrics, params, and tags in total * Up to 1000
// metrics\n- Up to 100 params * Up to 100 tags
//
// For example, a valid request might contain 900 metrics, 50 params, and 50
// tags, but logging 900 metrics, 50 params, and 51 tags is invalid.
//
// The following limits also apply to metric, param, and tag keys and values:
//
// * Metric keyes, param keys, and tag keys can be up to 250 characters in
// length * Parameter and tag values can be up to 250 characters in length
func (a *MLflowRunsAPI) LogBatch(ctx context.Context, request LogBatch) error {
	return a.impl.LogBatch(ctx, request)
}

// Log a metric
//
// Logs a metric for a run. A metric is a key-value pair (string key, float
// value) with an associated timestamp. Examples include the various metrics
// that represent ML model accuracy. A metric can be logged multiple times.
func (a *MLflowRunsAPI) LogMetric(ctx context.Context, request LogMetric) error {
	return a.impl.LogMetric(ctx, request)
}

// Log a model
//
// **NOTE:** Experimental: This API may change or be removed in a future release
// without warning.
func (a *MLflowRunsAPI) LogModel(ctx context.Context, request LogModel) error {
	return a.impl.LogModel(ctx, request)
}

// Log a param
//
// Logs a param used for a run. A param is a key-value pair (string key, string
// value). Examples include hyperparameters used for ML model training and
// constant dates and values used in an ETL pipeline. A param can be logged only
// once for a run.
func (a *MLflowRunsAPI) LogParameter(ctx context.Context, request LogParam) error {
	return a.impl.LogParameter(ctx, request)
}

// Restore a run
//
// Restores a deleted run.
func (a *MLflowRunsAPI) Restore(ctx context.Context, request RestoreRun) error {
	return a.impl.Restore(ctx, request)
}

// Restore a run
//
// Restores a deleted run.
func (a *MLflowRunsAPI) RestoreByRunId(ctx context.Context, runId string) error {
	return a.impl.Restore(ctx, RestoreRun{
		RunId: runId,
	})
}

// Search for runs
//
// Searches for runs that satisfy expressions.
//
// Search expressions can use `mlflowMetric` and `mlflowParam` keys.",
//
// This method is generated by Databricks SDK Code Generator.
func (a *MLflowRunsAPI) SearchAll(ctx context.Context, request SearchRuns) ([]Run, error) {
	var results []Run
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.Search(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Runs) == 0 {
			break
		}
		for _, v := range response.Runs {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// Set a tag
//
// Sets a tag on a run. Tags are run metadata that can be updated during a run
// and after a run completes.
func (a *MLflowRunsAPI) SetTag(ctx context.Context, request SetTag) error {
	return a.impl.SetTag(ctx, request)
}

// Update a run
//
// Updates run metadata.
func (a *MLflowRunsAPI) Update(ctx context.Context, request UpdateRun) (*UpdateRunResponse, error) {
	return a.impl.Update(ctx, request)
}

func NewModelVersionComments(client *client.DatabricksClient) *ModelVersionCommentsAPI {
	return &ModelVersionCommentsAPI{
		impl: &modelVersionCommentsImpl{
			client: client,
		},
	}
}

type ModelVersionCommentsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(ModelVersionCommentsService)
	impl ModelVersionCommentsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *ModelVersionCommentsAPI) WithImpl(impl ModelVersionCommentsService) *ModelVersionCommentsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level ModelVersionComments API implementation
func (a *ModelVersionCommentsAPI) Impl() ModelVersionCommentsService {
	return a.impl
}

// Post a comment
//
// Posts a comment on a model version. A comment can be submitted either by a
// user or programmatically to display relevant information about the model. For
// example, test results or deployment errors.
func (a *ModelVersionCommentsAPI) Create(ctx context.Context, request CreateComment) (*CreateResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete a comment
//
// Deletes a comment on a model version.
func (a *ModelVersionCommentsAPI) Delete(ctx context.Context, request DeleteModelVersionCommentRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a comment
//
// Deletes a comment on a model version.
func (a *ModelVersionCommentsAPI) DeleteById(ctx context.Context, id string) error {
	return a.impl.Delete(ctx, DeleteModelVersionCommentRequest{
		Id: id,
	})
}

// Update a comment
//
// Post an edit to a comment on a model version.
func (a *ModelVersionCommentsAPI) Update(ctx context.Context, request UpdateComment) (*UpdateResponse, error) {
	return a.impl.Update(ctx, request)
}

func NewModelVersions(client *client.DatabricksClient) *ModelVersionsAPI {
	return &ModelVersionsAPI{
		impl: &modelVersionsImpl{
			client: client,
		},
	}
}

type ModelVersionsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(ModelVersionsService)
	impl ModelVersionsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *ModelVersionsAPI) WithImpl(impl ModelVersionsService) *ModelVersionsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level ModelVersions API implementation
func (a *ModelVersionsAPI) Impl() ModelVersionsService {
	return a.impl
}

// Create a model version
//
// Creates a model version.
func (a *ModelVersionsAPI) Create(ctx context.Context, request CreateModelVersionRequest) (*CreateModelVersionResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete a model version.
//
// Deletes a model version.
func (a *ModelVersionsAPI) Delete(ctx context.Context, request DeleteModelVersionRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a model version tag
//
// Deletes a model version tag.
func (a *ModelVersionsAPI) DeleteTag(ctx context.Context, request DeleteModelVersionTagRequest) error {
	return a.impl.DeleteTag(ctx, request)
}

// Get a model version
//
// Get a model version.
func (a *ModelVersionsAPI) Get(ctx context.Context, request GetModelVersionRequest) (*GetModelVersionResponse, error) {
	return a.impl.Get(ctx, request)
}

// Get a model version URI
//
// Gets a URI to download the model version.
func (a *ModelVersionsAPI) GetDownloadUri(ctx context.Context, request GetModelVersionDownloadUriRequest) (*GetModelVersionDownloadUriResponse, error) {
	return a.impl.GetDownloadUri(ctx, request)
}

// Searches model versions
//
// Searches for specific model versions based on the supplied __filter__.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ModelVersionsAPI) SearchAll(ctx context.Context, request SearchModelVersionsRequest) ([]ModelVersion, error) {
	var results []ModelVersion
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.Search(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.ModelVersions) == 0 {
			break
		}
		for _, v := range response.ModelVersions {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// Set a version tag
//
// Sets a model version tag.
func (a *ModelVersionsAPI) SetTag(ctx context.Context, request SetModelVersionTagRequest) error {
	return a.impl.SetTag(ctx, request)
}

// Transition a stage
//
// Transition to the next model stage.
func (a *ModelVersionsAPI) TransitionStage(ctx context.Context, request TransitionModelVersionStage) (*TransitionModelVersionStageResponse, error) {
	return a.impl.TransitionStage(ctx, request)
}

// Update model version
//
// Updates the model version.
func (a *ModelVersionsAPI) Update(ctx context.Context, request UpdateModelVersionRequest) error {
	return a.impl.Update(ctx, request)
}

func NewRegisteredModels(client *client.DatabricksClient) *RegisteredModelsAPI {
	return &RegisteredModelsAPI{
		impl: &registeredModelsImpl{
			client: client,
		},
	}
}

type RegisteredModelsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(RegisteredModelsService)
	impl RegisteredModelsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *RegisteredModelsAPI) WithImpl(impl RegisteredModelsService) *RegisteredModelsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level RegisteredModels API implementation
func (a *RegisteredModelsAPI) Impl() RegisteredModelsService {
	return a.impl
}

// Create a model
//
// Creates a new registered model with the name specified in the request body.
//
// Throws `RESOURCE_ALREADY_EXISTS` if a registered model with the given name
// exists.
func (a *RegisteredModelsAPI) Create(ctx context.Context, request CreateRegisteredModelRequest) (*CreateRegisteredModelResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete a model
//
// Deletes a registered model.
func (a *RegisteredModelsAPI) Delete(ctx context.Context, request DeleteRegisteredModelRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a model
//
// Deletes a registered model.
func (a *RegisteredModelsAPI) DeleteByName(ctx context.Context, name string) error {
	return a.impl.Delete(ctx, DeleteRegisteredModelRequest{
		Name: name,
	})
}

// Delete a model tag
//
// Deletes the tag for a registered model.
func (a *RegisteredModelsAPI) DeleteTag(ctx context.Context, request DeleteRegisteredModelTagRequest) error {
	return a.impl.DeleteTag(ctx, request)
}

// Get a model
//
// Gets the registered model that matches the specified ID.
func (a *RegisteredModelsAPI) Get(ctx context.Context, request GetRegisteredModelRequest) (*GetRegisteredModelResponse, error) {
	return a.impl.Get(ctx, request)
}

// Get a model
//
// Gets the registered model that matches the specified ID.
func (a *RegisteredModelsAPI) GetByName(ctx context.Context, name string) (*GetRegisteredModelResponse, error) {
	return a.impl.Get(ctx, GetRegisteredModelRequest{
		Name: name,
	})
}

// Get the latest version
//
// Gets the latest version of a registered model.
//
// This method is generated by Databricks SDK Code Generator.
func (a *RegisteredModelsAPI) GetLatestVersionsAll(ctx context.Context, request GetLatestVersionsRequest) ([]ModelVersion, error) {
	response, err := a.impl.GetLatestVersions(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.ModelVersions, nil
}

// List models
//
// Lists all available registered models, up to the limit specified in
// __max_results__.
//
// This method is generated by Databricks SDK Code Generator.
func (a *RegisteredModelsAPI) ListAll(ctx context.Context, request ListRegisteredModelsRequest) ([]RegisteredModel, error) {
	var results []RegisteredModel
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.List(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.RegisteredModels) == 0 {
			break
		}
		for _, v := range response.RegisteredModels {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// Rename a model
//
// Renames a registered model.
func (a *RegisteredModelsAPI) Rename(ctx context.Context, request RenameRegisteredModelRequest) (*RenameRegisteredModelResponse, error) {
	return a.impl.Rename(ctx, request)
}

// Search models
//
// Search for registered models based on the specified __filter__.
//
// This method is generated by Databricks SDK Code Generator.
func (a *RegisteredModelsAPI) SearchAll(ctx context.Context, request SearchRegisteredModelsRequest) ([]RegisteredModel, error) {
	var results []RegisteredModel
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.Search(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.RegisteredModels) == 0 {
			break
		}
		for _, v := range response.RegisteredModels {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// Set a tag
//
// Sets a tag on a registered model.
func (a *RegisteredModelsAPI) SetTag(ctx context.Context, request SetRegisteredModelTagRequest) error {
	return a.impl.SetTag(ctx, request)
}

// Update model
//
// Updates a registered model.
func (a *RegisteredModelsAPI) Update(ctx context.Context, request UpdateRegisteredModelRequest) error {
	return a.impl.Update(ctx, request)
}

func NewRegistryWebhooks(client *client.DatabricksClient) *RegistryWebhooksAPI {
	return &RegistryWebhooksAPI{
		impl: &registryWebhooksImpl{
			client: client,
		},
	}
}

type RegistryWebhooksAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(RegistryWebhooksService)
	impl RegistryWebhooksService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *RegistryWebhooksAPI) WithImpl(impl RegistryWebhooksService) *RegistryWebhooksAPI {
	a.impl = impl
	return a
}

// Impl returns low-level RegistryWebhooks API implementation
func (a *RegistryWebhooksAPI) Impl() RegistryWebhooksService {
	return a.impl
}

// Create a webhook
//
// **NOTE**: This endpoint is in Public Preview.
//
// Creates a registry webhook.
func (a *RegistryWebhooksAPI) Create(ctx context.Context, request CreateRegistryWebhook) (*CreateResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete a webhook
//
// **NOTE:** This endpoint is in Public Preview.
//
// Deletes a registry webhook.
func (a *RegistryWebhooksAPI) Delete(ctx context.Context, request DeleteRegistryWebhookRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a webhook
//
// **NOTE:** This endpoint is in Public Preview.
//
// Deletes a registry webhook.
func (a *RegistryWebhooksAPI) DeleteById(ctx context.Context, id string) error {
	return a.impl.Delete(ctx, DeleteRegistryWebhookRequest{
		Id: id,
	})
}

// List registry webhooks
//
// **NOTE:** This endpoint is in Public Preview.
//
// Lists all registry webhooks.
//
// This method is generated by Databricks SDK Code Generator.
func (a *RegistryWebhooksAPI) ListAll(ctx context.Context, request ListRegistryWebhooksRequest) ([]RegistryWebhook, error) {
	var results []RegistryWebhook
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.List(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Webhooks) == 0 {
			break
		}
		for _, v := range response.Webhooks {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// Test a webhook
//
// **NOTE:** This endpoint is in Public Preview.
//
// Tests a registry webhook.
func (a *RegistryWebhooksAPI) Test(ctx context.Context, request TestRegistryWebhookRequest) (*TestRegistryWebhookResponse, error) {
	return a.impl.Test(ctx, request)
}

// Update a webhook
//
// **NOTE:** This endpoint is in Public Preview.
//
// Updates a registry webhook.
func (a *RegistryWebhooksAPI) Update(ctx context.Context, request UpdateRegistryWebhook) error {
	return a.impl.Update(ctx, request)
}

func NewTransitionRequests(client *client.DatabricksClient) *TransitionRequestsAPI {
	return &TransitionRequestsAPI{
		impl: &transitionRequestsImpl{
			client: client,
		},
	}
}

type TransitionRequestsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(TransitionRequestsService)
	impl TransitionRequestsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *TransitionRequestsAPI) WithImpl(impl TransitionRequestsService) *TransitionRequestsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level TransitionRequests API implementation
func (a *TransitionRequestsAPI) Impl() TransitionRequestsService {
	return a.impl
}

// Approve transition requests
//
// Approves a model version stage transition request.
func (a *TransitionRequestsAPI) Approve(ctx context.Context, request ApproveTransitionRequest) (*ApproveResponse, error) {
	return a.impl.Approve(ctx, request)
}

// Make a transition request
//
// Creates a model version stage transition request.
func (a *TransitionRequestsAPI) Create(ctx context.Context, request CreateTransitionRequest) (*CreateResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete a ransition request
//
// Cancels a model version stage transition request.
func (a *TransitionRequestsAPI) Delete(ctx context.Context, request DeleteTransitionRequestRequest) error {
	return a.impl.Delete(ctx, request)
}

// List transition requests
//
// Gets a list of all open stage transition requests for the model version.
//
// This method is generated by Databricks SDK Code Generator.
func (a *TransitionRequestsAPI) ListAll(ctx context.Context, request ListTransitionRequestsRequest) ([]Activity, error) {
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Requests, nil
}

// Reject a transition request
//
// Rejects a model version stage transition request.
func (a *TransitionRequestsAPI) Reject(ctx context.Context, request RejectTransitionRequest) (*RejectResponse, error) {
	return a.impl.Reject(ctx, request)
}
