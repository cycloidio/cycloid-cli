package e2e_test

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"sync"
	"testing"
)

const (
	helloWorldSource = "cycloid/plugin-hello-world:latest"
	// baseLocalImage is the base path in the host-reachable registry.
	baseLocalImage = "localhost:5000/plugin-hello-world"
	// baseAPIImage is the base path as seen from inside the docker network.
	// Must match the registry URL configured in the plugin-registry service.
	baseAPIImage = "docker-registry:5000/plugin-hello-world"
)

var imageOnce sync.Once

// ensureHelloWorldImage ensures the plugin-hello-world source image is pulled once.
// Unique tags are pushed on demand by uniqueHelloWorldVersion.
func ensureHelloWorldImage(t *testing.T) {
	t.Helper()
	imageOnce.Do(func() {
		run := func(name string, args ...string) {
			cmd := exec.Command(name, args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				t.Fatalf("command %q %v failed: %v", name, args, err)
			}
		}
		run("docker", "pull", helloWorldSource)
		run("docker", "login", "localhost:5000", "-u", "cycloid", "-p", "cycloid123")
	})
}

// uniqueHelloWorldVersion generates a unique semver-compatible image tag, pushes it
// to the local registry, and returns the docker-registry:5000 URL for use with the API.
// The plugin-registry enforces a global unique constraint on version URLs, so each
// call to CreatePluginVersion must use a different image tag.
func uniqueHelloWorldVersion(t *testing.T) string {
	t.Helper()
	ensureHelloWorldImage(t)

	tag := fmt.Sprintf("1.0.%d", rand.Intn(999999))
	local := fmt.Sprintf("%s:%s", baseLocalImage, tag)
	api := fmt.Sprintf("%s:%s", baseAPIImage, tag)

	run := func(name string, args ...string) {
		t.Helper()
		cmd := exec.Command(name, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			t.Fatalf("uniqueHelloWorldVersion: %q %v failed: %v", name, args, err)
		}
	}

	run("docker", "tag", helloWorldSource, local)
	run("docker", "push", local)
	return api
}
