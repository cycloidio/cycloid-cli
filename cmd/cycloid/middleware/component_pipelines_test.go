package middleware_test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestComponentPipeline(t *testing.T) {
	t.Parallel()
	config, err := GetTestConfig()
	if err != nil {
		t.Errorf("failed to get test config: %s", err)
		return
	}

	project, projectDelete, err := config.SetupTestProject(t.Name())
	defer projectDelete()
	if err != nil {
		t.Errorf("failed to create project for test '%s': %s", t.Name(), err)
	}

	env, envDelete, err := config.SetupTestEnv(t.Name(), *project.Canonical)
	defer envDelete()
	if err != nil {
		t.Errorf("failed to setup env for test '%s': %s", t.Name(), err)
		return
	}

	component, componentDelete, err := config.SetupTestComponent(
		*project.Canonical, *env.Canonical, t.Name(), pipelineTestStackRef, pipelineTestStackUseCase, &pipelineTestDefaultVars,
	)
	defer componentDelete()
	if err != nil {
		t.Errorf("failed to setup base component for test '%s': %s", t.Name(), err)
	}

	t.Run("GetPipeline", func(t *testing.T) {
		pipelineName := fmt.Sprintf("%s-%s-%s", *project.Canonical, *env.Canonical, *component.Canonical)
		got, err := m.GetPipeline(config.Org, *project.Canonical, *env.Canonical, *component.Canonical, pipelineName)
		if err != nil {
			t.Errorf("middleware.GetComponentPipelines() error = %v", err)
			return
		}

		if !reflect.DeepEqual(*got.Component, *component) {
			t.Fatalf("component in pipeline doesn't match, got:\n%v\nexpect: %v", *got.Component, *component)
		}
	})

	t.Run("PausePipeline", func(t *testing.T) {
		pipelineName := fmt.Sprintf("%s-%s-%s", *project.Canonical, *env.Canonical, *component.Canonical)
		err := m.PausePipeline(config.Org, *project.Canonical, *env.Canonical, *component.Canonical, pipelineName)
		if err != nil {
			t.Errorf("%s error = %v", t.Name(), err)
			return
		}
	})

	t.Run("UnpausePipeline", func(t *testing.T) {
		pipelineName := fmt.Sprintf("%s-%s-%s", *project.Canonical, *env.Canonical, *component.Canonical)
		err := m.UnpausePipeline(config.Org, *project.Canonical, *env.Canonical, *component.Canonical, pipelineName)
		if err != nil {
			t.Errorf("%s error = %v", t.Name(), err)
			return
		}
	})

	t.Run("SynchedPipelineOk", func(t *testing.T) {
		pipelineName := fmt.Sprintf("%s-%s-%s", *project.Canonical, *env.Canonical, *component.Canonical)
		got, err := m.SyncedPipeline(config.Org, *project.Canonical, *env.Canonical, *component.Canonical, pipelineName)
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
		pipelineName := fmt.Sprintf("%s-%s-%s", *project.Canonical, *env.Canonical, *component.Canonical)
		updatedPipeline, err := m.UpdatePipeline(config.Org, *project.Canonical, *env.Canonical, *component.Canonical, pipelineName, newPipeline, "---\nmessage: hello", false)
		if err != nil {
			t.Errorf("%s error = %v", t.Name(), err)
			return
		}

		pipelineJobs, err := m.GetJobs(config.Org, *updatedPipeline.Project.Canonical, *updatedPipeline.Environment.Canonical, *updatedPipeline.Component.Canonical, *updatedPipeline.Name)
		if err != nil {
			t.Errorf("failed to get jobs in pipeline '%s' in test '%s': %s", *updatedPipeline.Name, t.Name(), err)
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
			build, err := m.CreateBuild(config.Org, *project.Canonical, *env.Canonical, *component.Canonical, *updatedPipeline.Name, *pipelineJobs[0].Name)
			if err != nil {
				t.Errorf("failed to trigger build in job '%s': %s", *pipelineJobs[0].Name, err)
				return
			}

			buildIDStr := strconv.Itoa(int(*build.ID))
			getBuild, err := m.GetBuild(config.Org, *project.Canonical, *env.Canonical, *component.Canonical, *updatedPipeline.Name, *pipelineJobs[0].Name, buildIDStr)
			if err != nil {
				t.Errorf("failed to get build in job '%s': %s", *pipelineJobs[0].Name, err)
				return
			}

			if *build.ID != *getBuild.ID {
				t.Errorf("build are not matching:\ngot: %v\nexpected:\n%v", *getBuild, *build)
				return
			}

			err = m.AbortBuild(config.Org, *project.Canonical, *env.Canonical, *component.Canonical, *updatedPipeline.Name, *pipelineJobs[0].Name, buildIDStr)
			if err != nil {
				t.Fatalf("failed to abort build '%s': %s", buildIDStr, err)
			}

			_, err = m.RerunBuild(config.Org, *project.Canonical, *env.Canonical, *component.Canonical, *updatedPipeline.Name, *pipelineJobs[0].Name, buildIDStr)
			if err != nil {
				t.Fatalf("failed to re-run build '%s': %s", buildIDStr, err)
			}
		})
	})
}
