package e2e_test

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func TestPipelines(t *testing.T) {
	// Pipelines
	var pipelineList []*models.Pipeline
	t.Run("PipelineListOk", func(t *testing.T) {
		listOut, listErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"pipelines", "list",
		})
		if listErr != nil {
			t.Errorf("List org pipeline should not err, out: %s\nerr: %s", listOut, listErr)
		}

		err := json.Unmarshal([]byte(listOut), &pipelineList)
		if err != nil {
			t.Errorf("failed to unmarshall list test output, out: %s\nerr: %s", listOut, err)
		}

		if len(pipelineList) == 0 {
			t.Errorf("There should be at least one pipeline in this org:\n%s", listOut)
		}
	})

	firstPipeline := *pipelineList[0]
	t.Run("GetPipelineOk", func(t *testing.T) {
		getOut, getErr := executeCommand([]string{
			"--output", "json",
			"pipelines", "get",
			"--project", *firstPipeline.Project.Canonical,
			"--env", *firstPipeline.Environment.Canonical,
			"--component", *firstPipeline.Component.Canonical,
			"--pipeline", *firstPipeline.Name,
		})
		if getErr != nil {
			t.Errorf("get pipeline should not err, out: %s\nerr: %s", getOut, getErr)
		}

		var getPipeline models.Pipeline
		err := json.Unmarshal([]byte(getOut), &getPipeline)
		if err != nil {
			t.Errorf("Failed to parse json output of pipelines get cmd, out: %s\nerr: %s", getOut, err)
		}

		if assert.ObjectsAreEqualValues(firstPipeline, getPipeline) {
			t.Errorf("both pipelines should be equal:\nexpect: %v\ngot: %v", firstPipeline, getPipeline)
		}
	})

	t.Run("PipelinePauseOk", func(t *testing.T) {
		_, pauseErr := executeCommand([]string{
			"--output", "json",
			"pipelines", "pause",
			"--project", *firstPipeline.Project.Canonical,
			"--env", *firstPipeline.Environment.Canonical,
			"--component", *firstPipeline.Component.Canonical,
			"--pipeline", *firstPipeline.Name,
		})
		if pauseErr != nil {
			t.Errorf("failed to pause pipeline '%s': %s", *firstPipeline.Name, pauseErr)
		}
	})

	t.Run("PipelineUnpauseOk", func(t *testing.T) {
		_, unpauseErr := executeCommand([]string{
			"--output", "json",
			"pipelines", "unpause",
			"--project", *firstPipeline.Project.Canonical,
			"--env", *firstPipeline.Environment.Canonical,
			"--component", *firstPipeline.Component.Canonical,
			"--pipeline", *firstPipeline.Name,
		})
		if unpauseErr != nil {
			t.Errorf("failed to unpause pipeline '%s': %s", *firstPipeline.Name, unpauseErr)
		}
	})

	t.Run("PipelineSynced", func(t *testing.T) {
		_, syncedErr := executeCommand([]string{
			"--output", "json",
			"pipelines", "synced",
			"--project", *firstPipeline.Project.Canonical,
			"--env", *firstPipeline.Environment.Canonical,
			"--component", *firstPipeline.Component.Canonical,
			"--pipeline", *firstPipeline.Name,
		})
		if syncedErr != nil {
			t.Errorf("failed to pause pipeline '%s': %s", *firstPipeline.Name, syncedErr)
		}
	})

	t.Run("PipelineLastUsedOk", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"pipelines", "last-used",
			"--since-days", "99",
		})
		if cmdErr != nil {
			t.Errorf("failed to list last-used pipelines, out: %s\nerr: %s", cmdOut, cmdErr)
		}
	})

	// Jobs
	var jobList []*models.Job
	t.Run("ListJobsOk", func(t *testing.T) {
		if len(pipelineList) == 0 {
			t.Fatal("Test setup error: pipeline list length should not be 0.")
		}

		firstPipeline := pipelineList[0]
		listOut, listErr := executeCommand([]string{
			"--output", "json",
			"pipeline", "jobs", "list",
			"--project", *firstPipeline.Project.Canonical,
			"--env", *firstPipeline.Environment.Canonical,
			"--component", *firstPipeline.Component.Canonical,
			"--pipeline", *firstPipeline.Name,
		})
		if listErr != nil {
			t.Errorf("List job in pipeline %q should not err, out: %s\nerr: %s", *firstPipeline.Name, listOut, listErr)
		}

		err := json.Unmarshal([]byte(listOut), &jobList)
		if err != nil {
			t.Errorf("failed to marshal output of cy pp job list, out: %s\nerr: %s", listOut, err)
		}

		if len(jobList) == 0 {
			t.Errorf("job list should not be empty:\n%s", litter.Sdump(jobList))
		}

		firstJob := jobList[0]
		t.Run("JobGetOk", func(t *testing.T) {
			getOut, getErr := executeCommand([]string{
				"--output", "json",
				"pipeline", "job", "get",
				"--project", *firstPipeline.Project.Canonical,
				"--env", *firstPipeline.Environment.Canonical,
				"--component", *firstPipeline.Component.Canonical,
				"--pipeline", *firstPipeline.Name,
				"--job", *firstJob.Name,
			})
			if getErr != nil {
				t.Errorf("cy get job in pipeline '%s' should not fail, out: %s\nerr: %s", *firstJob.Name, getOut, getErr)
			}

			var getJob *models.Job
			err := json.Unmarshal([]byte(getOut), &getJob)
			if err != nil {
				t.Errorf("failed to unmarshall get job cmd output, out: %s\nerr: %s", getOut, err)
			}

			assert.Equal(t, *firstJob.ID, *getJob.ID)
		})

		t.Run("PauseJobOk", func(t *testing.T) {
			pauseOut, pauseErr := executeCommand([]string{
				"--output", "json",
				"pipeline", "job", "pause",
				"--project", *firstPipeline.Project.Canonical,
				"--env", *firstPipeline.Environment.Canonical,
				"--component", *firstPipeline.Component.Canonical,
				"--pipeline", *firstPipeline.Name,
				"--job", *firstJob.Name,
			})
			if pauseErr != nil {
				t.Errorf("cmd cy pp job pause failed for pipeline '%s', out: %s, err: %s", *firstPipeline.Name, pauseOut, pauseErr)
			}

			t.Run("UnpauseJobOk", func(t *testing.T) {
				unpauseOut, unpauseErr := executeCommand([]string{
					"--output", "json",
					"pipeline", "job", "unpause",
					"--project", *firstPipeline.Project.Canonical,
					"--env", *firstPipeline.Environment.Canonical,
					"--component", *firstPipeline.Component.Canonical,
					"--pipeline", *firstPipeline.Name,
					"--job", *firstJob.Name,
				})
				if unpauseErr != nil {
					t.Errorf("cmd cy pp job unpause failed for pipeline '%s', out: %s, err: %s", *firstPipeline.Name, unpauseOut, unpauseErr)
				}
			})
		})

		var triggeredBuild *models.Build
		// Builds
		t.Run("CreateBuildOk", func(t *testing.T) {
			triggerOut, triggerErr := executeCommand([]string{
				"--output", "json",
				"pipeline", "build", "create",
				"--project", *firstPipeline.Project.Canonical,
				"--env", *firstPipeline.Environment.Canonical,
				"--component", *firstPipeline.Component.Canonical,
				"--pipeline", *firstPipeline.Name,
				"--job", *firstJob.Name,
			})
			if triggerErr != nil {
				t.Errorf("cmd cy pp build create failed for job '%s' in pipeline '%s', out: %s, err: %s", *firstJob.Name, *firstPipeline.Name, triggerOut, triggerErr)
			}

			err := json.Unmarshal([]byte(triggerOut), &triggeredBuild)
			if err != nil {
				t.Errorf("cmd output is not a models.Build, out: %s\nerr: %s", triggerOut, triggerErr)
			}

			buildIDStr := strconv.Itoa(int(*triggeredBuild.ID))
			if err != nil {
				t.Errorf("invalid build id in:\n%v\n%s", triggeredBuild, err)
			}

			t.Run("GetBuildOk", func(t *testing.T) {
				getOut, getErr := executeCommand([]string{
					"--output", "json",
					"pipeline", "build", "get",
					"--project", *firstPipeline.Project.Canonical,
					"--env", *firstPipeline.Environment.Canonical,
					"--component", *firstPipeline.Component.Canonical,
					"--pipeline", *firstPipeline.Name,
					"--job", *firstJob.Name,
					"--build-id", buildIDStr,
				})
				if getErr != nil {
					t.Errorf("cmd cy pp build get failed for job '%s' in pipeline '%s', out: %s, err: %s", *firstJob.Name, *firstPipeline.Name, getOut, getErr)
				}
			})

			t.Run("ListBuildOk", func(t *testing.T) {
				listOut, listErr := executeCommand([]string{
					"--output", "json",
					"pipeline", "build", "list",
					"--project", *firstPipeline.Project.Canonical,
					"--env", *firstPipeline.Environment.Canonical,
					"--component", *firstPipeline.Component.Canonical,
					"--pipeline", *firstPipeline.Name,
					"--job", *firstJob.Name,
				})
				if listErr != nil {
					t.Errorf("cmd cy pp build list failed for job '%s' in pipeline '%s', out: %s, err: %s", *firstJob.Name, *firstPipeline.Name, listOut, listErr)
				}
			})
		})
	})
}
