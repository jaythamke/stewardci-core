package runmgr

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	api "github.com/SAP/stewardci-core/pkg/apis/steward/v1alpha1"
	"github.com/ghodss/yaml"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"gotest.tools/v3/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	taskStartTime = `2019-05-14T08:24:08Z`
	stepStartTime = `2019-05-14T08:24:11Z`
	emptyBuild    = `{}`

	runningBuild = `{"status":
	                  {"steps": [
	                    {"name": "jenkinsfile-runner",
	                     "running": {"startedAt": "` + stepStartTime + `"}}]}}`

	completedSuccess = `{"status":
	                        {"conditions": [
	                            {"message": "message1",
	                             "reason": "Succeeded",
	                             "status": "True",
	                             "type": "Succeeded"}],
	                         "steps": [
	                            {"name": "jenkinsfile-runner",
	                             "terminated": {
	                                "reason": "Completed",
	                                "message": "ok",
	                                "exitCode": 0}}]}}`

	completedErrorInfra = `{"status":
	                           {"conditions": [
	                               {"message": "message1",
	                                "reason": "Failed",
	                                "status": "False",
	                                "type": "Succeeded"}],
	                            "steps": [
	                               {"name": "jenkinsfile-runner",
	                                "terminated": {
	                                   "reason": "Error",
	                                   "message": "ko",
	                                   "exitCode": 1}}]}}`

	completedErrorContent = `{"status":
	                            {"conditions": [
	                                {"message": "message1",
	                                 "reason": "Failed",
	                                 "status": "False",
	                                 "type": "Succeeded"}],
	                             "steps": [
	                                {"name": "jenkinsfile-runner",
	                                 "terminated": {
	                                    "reason": "Error",
	                                    "message": "ko",
	                                    "exitCode": 2}}]}}`

	completedErrorConfig = `{"status":
	                           {"conditions": [
	                               {"message": "message1",
	                                "reason": "Failed",
	                                "status": "False",
	                                "type": "Succeeded"}],
	                            "steps": [
	                               {"name": "jenkinsfile-runner",
	                                "terminated": {
	                                   "reason": "Error",
	                                   "message": "ko",
	                                   "exitCode": 3}}]}}`

	completedValidationFailed = `{"status":
	                                {"conditions": [
	                                    {"message": "message1",
	                                     "reason": "TaskRunValidationFailed",
	                                     "status": "False",
	                                     "type": "Succeeded"}]}}`

	//See issue https://github.com/SAP/stewardci-core/issues/? TODO: create public issue. internal: 21
	timeout = `{"status": {"conditions": [{"message": "TaskRun \"steward-jenkinsfile-runner\" failed to finish within \"10m0s\"", "reason": "TaskRunTimeout", "status": "False", "type": "Succeeded"}]}}`

	realStartedBuild = `status:
  conditions:
  - lastTransitionTime: ` + taskStartTime + `
    message: Not all Steps in the Task have finished executing
    reason: Running
    status: Unknown
    type: Succeeded
  podName: build-pod-38aa76
  startTime: 
  steps:
  - container: step-jenkinsfile-runner
    imageID: docker-pullable://alpine@sha256:acd3ca9941a85e8ed16515bfc5328e4e2f8c128caa72959a58a127b7801ee01f
    name: jenkinsfile-runner
    running:
      startedAt: "` + stepStartTime + `"
`

	realCompletedSuccess = `status:
  completionTime: "2019-05-14T08:24:49Z"
  conditions:
  - lastTransitionTime: "2019-10-04T13:57:28Z"
    message: All Steps have completed executing
    reason: Succeeded
    status: "True"
    type: Succeeded
  podName: build-pod-38aa76
  startTime: "2019-05-14T08:24:08Z"
  steps:
  - container: step-jenkinsfile-runner
    imageID: docker-pullable://alpine@sha256:acd3ca9941a85e8ed16515bfc5328e4e2f8c128caa72959a58a127b7801ee01f
    name: jenkinsfile-runner
    terminated:
      containerID: docker://2ee92b9e6971cd76f896c5c4dc403203754bd4aa6c5191541e5f7d8e04ce9326
      exitCode: 0
      finishedAt: "2019-05-14T08:24:49Z"
      reason: Completed
      startedAt: "2019-05-14T08:24:11Z"
`

	completedMessageSuccess = `status:
  completionTime: "2019-05-14T08:24:49Z"
  conditions:
  - lastTransitionTime: "2019-10-04T13:57:28Z"
    message: All Steps have completed executing
    reason: Succeeded
    status: "True"
    type: Succeeded
  podName: build-pod-38aa76
  startTime: "2019-05-14T08:24:08Z"
  steps:
  - container: step-jenkinsfile-runner
    imageID: docker-pullable://alpine@sha256:acd3ca9941a85e8ed16515bfc5328e4e2f8c128caa72959a58a127b7801ee01f
    name: jenkinsfile-runner
    terminated:
      containerID: docker://2ee92b9e6971cd76f896c5c4dc403203754bd4aa6c5191541e5f7d8e04ce9326
      exitCode: 0
      finishedAt: "2019-05-14T08:24:49Z"
      reason: Completed
      message: %q
      startedAt: "2019-05-14T08:24:11Z"
`
	completionTimeSet = `status:
  completionTime: 2019-05-14T08:24:49Z
  `
	completionTimeNotSet = `status: {}`

	conditionSuccessWithTransitionTime = `status:
  conditions:
  - lastTransitionTime: "2021-10-07T08:59:59Z"
    message: 'foo'
    reason: CouldntGetTask
    status: "False"
    type: Succeeded
  `
	conditionSuccessWithoutTransitionTime = `status:
  conditions:
  - message: 'bar'
    reason: CouldntGetTask
    status: "False"
    type: Succeeded
  `
	noSuccessCondition = `status:
  conditions:
  - lastTransitionTime: "2021-10-07T08:59:59Z"
    message: 'baz'
    reason: CouldntGetTask
    status: "False"
    type: Foo
  `
	imagePullFailedCondition = `status:
  conditions:
  - lastTransitionTime: "2022-12-02T12:30:01Z"
    message: 'failed to pull the image'
    reason: TaskRunImagePullFailed
    status: "False"
    type: Succeeded
`
)

func generateTime(timeRFC3339String string) *metav1.Time {
	t, _ := time.Parse(time.RFC3339, timeRFC3339String)
	mt := metav1.NewTime(t)
	return &mt
}

func fakeTektonTaskRun(s string) *tekton.TaskRun {
	var result tekton.TaskRun
	json.Unmarshal([]byte(s), &result)
	return &result
}

func fakeTektonTaskRunYaml(s string) *tekton.TaskRun {
	var result tekton.TaskRun
	yaml.Unmarshal([]byte(s), &result)
	return &result
}

func Test__GetStartTime_UnsetReturnsNil(t *testing.T) {
	run := NewRun(fakeTektonTaskRun(emptyBuild))
	startTime := run.GetStartTime()
	assert.Assert(t, startTime == nil)
}

func Test__GetStartTime_Set(t *testing.T) {
	expectedTime := generateTime(stepStartTime)
	run := NewRun(fakeTektonTaskRunYaml(realStartedBuild))
	startTime := run.GetStartTime()
	assert.Assert(t, expectedTime.Equal(startTime), fmt.Sprintf("Expected: %s, Is: %s", expectedTime, startTime))
}

func Test__IsFinished_RunningUpdatesContainer(t *testing.T) {
	run := NewRun(fakeTektonTaskRun(runningBuild))
	finished, _ := run.IsFinished()
	assert.Assert(t, run.GetContainerInfo().Running != nil)
	assert.Assert(t, finished == false)
}

func Test__IsFinished_CompletedSuccess(t *testing.T) {
	build := fakeTektonTaskRunYaml(realCompletedSuccess)
	run := NewRun(build)
	finished, result := run.IsFinished()
	assert.Assert(t, run.GetContainerInfo().Terminated != nil)
	assert.Assert(t, finished == true)
	assert.Equal(t, result, api.ResultSuccess)
}

func Test__IsFinished_CompletedFail(t *testing.T) {
	for _, test := range []struct {
		name           string
		trString       string
		expectedResult api.Result
	}{
		{
			name:           "infra_error",
			trString:       completedErrorInfra,
			expectedResult: api.ResultErrorInfra,
		}, {
			name:           "error_content",
			trString:       completedErrorContent,
			expectedResult: api.ResultErrorContent,
		}, {
			name:           "error_confix",
			trString:       completedErrorConfig,
			expectedResult: api.ResultErrorConfig,
		},
	} {
		t.Run(test.name, func(t *testing.T) {

			build := fakeTektonTaskRun(test.trString)
			run := NewRun(build)
			finished, result := run.IsFinished()
			assert.Assert(t, run.GetContainerInfo().Terminated != nil)
			assert.Assert(t, finished == true)
			assert.Equal(t, result, test.expectedResult)
		})
	}
}

func Test__IsFinished_CompletedValidationFail(t *testing.T) {
	build := fakeTektonTaskRun(completedValidationFailed)
	run := NewRun(build)
	finished, result := run.IsFinished()
	assert.Assert(t, finished == true)
	assert.Equal(t, result, api.ResultErrorInfra)
}

func Test__IsFinished_Timeout(t *testing.T) {
	run := NewRun(fakeTektonTaskRun(timeout))
	finished, result := run.IsFinished()
	assert.Assert(t, run.GetContainerInfo() == nil)
	assert.Assert(t, finished == true)
	assert.Equal(t, result, api.ResultTimeout)
}

func Test__IsRestartable_False(t *testing.T) {
	for id, taskrun := range []string{
		completedSuccess,
		completedErrorInfra,
		completedErrorConfig,
		completedErrorContent,
		completedValidationFailed,
		timeout,
	} {
		t.Run(fmt.Sprintf("%d", id), func(t *testing.T) {
			run := NewRun(fakeTektonTaskRun(taskrun))
			result := run.IsRestartable()
			assert.Assert(t, result == false)
		})
	}
}

func Test__IsRestartable_ImagePullFailed(t *testing.T) {
	run := NewRun(fakeTektonTaskRunYaml(imagePullFailedCondition))
	result := run.IsRestartable()
	assert.Assert(t, result)
}

func Test__GetCompletionTime(t *testing.T) {
	for id, taskrun := range []string{
		completionTimeSet,
		completionTimeNotSet,
		conditionSuccessWithTransitionTime,
		conditionSuccessWithoutTransitionTime,
		noSuccessCondition,
	} {
		t.Run(fmt.Sprintf("%d", id), func(t *testing.T) {
			run := NewRun(fakeTektonTaskRunYaml(taskrun))
			completionTime := run.GetCompletionTime()
			assert.Assert(t, completionTime != nil)
		})
	}
}

func Test__GetMessage(t *testing.T) {
	for _, test := range []struct {
		name            string
		inputMessage    string
		expectedMessage string
	}{
		{name: "message_ok",
			inputMessage:    `[{"key":"jfr-termination-log","value":"foo"}]`,
			expectedMessage: "foo",
		},
		{name: "wrong_key",
			inputMessage:    `[{"key":"termination-log","value":"foo"}]`,
			expectedMessage: "internal error",
		},
		{name: "empty message",
			inputMessage:    "",
			expectedMessage: "All Steps have completed executing",
		},
		{name: "multi_key",
			inputMessage:    `[{"key": "foo", "value": "bar"}, {"key":"jfr-termination-log","value":"foo"}]`,
			expectedMessage: "foo",
		},
		{name: "invalid_yaml_message",
			inputMessage:    "{no valid yaml",
			expectedMessage: "{no valid yaml",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			test := test
			t.Parallel()
			buildString := fmt.Sprintf(completedMessageSuccess, test.inputMessage)
			build := fakeTektonTaskRunYaml(buildString)
			run := NewRun(build)
			result := run.GetMessage()
			assert.Equal(t, test.expectedMessage, result)
		})
	}
}
