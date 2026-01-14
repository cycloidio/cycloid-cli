package middleware_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestComponentPipeline(t *testing.T) {
	m := config.Middleware

	component, err := config.NewTestComponent(
		*config.Project.Canonical, *config.Environment.Canonical, t.Name(), config.Org+":"+pipelineTestStackCanonical, pipelineTestStackUseCase, "", "", *config.CatalogRepoVersionStacks.CommitHash, pipelineTestDefaultVars,
	)
	if err != nil {
		t.Errorf("failed to setup base component for test %q: %v", t.Name(), err)
	}

	t.Run("GetPipeline", func(t *testing.T) {
		pipelineName := fmt.Sprintf("%s-%s-%s", *config.Project.Canonical, *config.Environment.Canonical, *component.Canonical)
		got, err := m.GetPipeline(config.Org, *config.Project.Canonical, *config.Environment.Canonical, *component.Canonical, pipelineName)
		if err != nil {
			t.Errorf("middleware.GetComponentPipelines() error = %v", err)
			return
		}

		got.Component.Project.Owner = component.Project.Owner
		require.Equal(t, *got.Component, *component)
	})

	t.Run("PausePipeline", func(t *testing.T) {
		pipelineName := fmt.Sprintf("%s-%s-%s", *config.Project.Canonical, *config.Environment.Canonical, *component.Canonical)
		err := m.PausePipeline(config.Org, *config.Project.Canonical, *config.Environment.Canonical, *component.Canonical, pipelineName)
		if err != nil {
			t.Errorf("%s error = %v", t.Name(), err)
			return
		}
	})

	t.Run("UnpausePipeline", func(t *testing.T) {
		pipelineName := fmt.Sprintf("%s-%s-%s", *config.Project.Canonical, *config.Environment.Canonical, *component.Canonical)
		err := m.UnpausePipeline(config.Org, *config.Project.Canonical, *config.Environment.Canonical, *component.Canonical, pipelineName)
		if err != nil {
			t.Errorf("%s error = %v", t.Name(), err)
			return
		}
	})

	t.Run("SynchedPipelineOk", func(t *testing.T) {
		pipelineName := fmt.Sprintf("%s-%s-%s", *config.Project.Canonical, *config.Environment.Canonical, *component.Canonical)
		got, err := m.SyncedPipeline(config.Org, *config.Project.Canonical, *config.Environment.Canonical, *component.Canonical, pipelineName)
		if err != nil {
			t.Errorf("%s error = %v", t.Name(), err)
			return
		}

		if *got.Synced != "synced" {
			t.Errorf("pipeline must be in sync: %v", *got)
		}
	})

	t.Run("UpdatePipelineAndGetJobs", func(t *testing.T) {
		newPipeline := `---
jobs:
- name: job-hello-world
  build_logs_to_retain: 1
  plan:
  - task: hello-world
    config:
      platform: linux
      image_resource:
        type: registry-image
        source: {repository: busybox, tag: latest}
      run:
        path: /bin/sh
        args:
        - -ec
        - |
          echo ${MESSAGE}
      params:
        MESSAGE: ((message))
`
		pipelineName := fmt.Sprintf("%s-%s-%s", *config.Project.Canonical, *config.Environment.Canonical, *component.Canonical)
		updatedPipeline, err := m.UpdatePipeline(config.Org, *config.Project.Canonical, *config.Environment.Canonical, *component.Canonical, pipelineName, newPipeline, "---\nmessage: hello", false)
		if err != nil {
			t.Errorf("%s error = %v", t.Name(), err)
			return
		}

		pipelineJobs, err := m.GetJobs(config.Org, *updatedPipeline.Project.Canonical, *updatedPipeline.Environment.Canonical, *updatedPipeline.Component.Canonical, *updatedPipeline.Name)
		if err != nil {
			t.Logf("failed to get jobs in pipeline '%s' in test '%s': %s", *updatedPipeline.Name, t.Name(), err)
			return
		}

		if len(pipelineJobs) != 1 {
			t.Errorf("invalid number of jobs, must be only one, got: %d", len(pipelineJobs))
			return
		}

		if *pipelineJobs[0].Name != "job-hello-world" {
			t.Errorf("job name should be 'job-hello-world' as in the updated pipeline, got: %s", *pipelineJobs[0].Name)
			return
		}

		t.Run("TestBuilds", func(t *testing.T) {
			build, err := m.CreateBuild(config.Org, *config.Project.Canonical, *config.Environment.Canonical, *component.Canonical, *updatedPipeline.Name, *pipelineJobs[0].Name)
			if err != nil {
				t.Errorf("failed to trigger build in job '%s': %s", *pipelineJobs[0].Name, err)
				return
			}

			buildIDStr := strconv.Itoa(int(*build.ID))
			getBuild, err := m.GetBuild(config.Org, *config.Project.Canonical, *config.Environment.Canonical, *component.Canonical, *updatedPipeline.Name, *pipelineJobs[0].Name, buildIDStr)
			if err != nil {
				t.Errorf("failed to get build in job '%s': %s", *pipelineJobs[0].Name, err)
				return
			}

			if *build.ID != *getBuild.ID {
				t.Errorf("build are not matching:\ngot: %v\nexpected:\n%v", *getBuild, *build)
				return
			}

			time.Sleep(1 * time.Second)

			err = m.AbortBuild(config.Org, *config.Project.Canonical, *config.Environment.Canonical, *component.Canonical, *updatedPipeline.Name, *pipelineJobs[0].Name, buildIDStr)
			if err != nil {
				t.Errorf("failed to abort build '%s': %s", buildIDStr, err)
			}

			// // Add a bit of time, concourse seems to not like it
			// time.Sleep(3 * time.Second)
			//
			// _, err = m.RerunBuild(config.Org, *config.Project.Canonical, *config.Environment.Canonical, *component.Canonical, *updatedPipeline.Name, *pipelineJobs[0].Name, buildIDStr)
			// if err != nil {
			// 	t.Errorf("failed to re-run build '%s': %s", buildIDStr, err)
			// }
		})
	})
}
