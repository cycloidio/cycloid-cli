package uri

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

// RecurseFS will walk through each paths excluding paths that matches
// all the ignores patterns. This function is ment to process the paths
// for commands that uses --recurse and --ignore flags.
func RecurseFS(paths []string, ignores []string) ([]string, error) {
	var result []string

	var outErr error
	cleanPaths := slices.DeleteFunc(paths, func(path string) bool {
		ignored, err := IsIgnored(path, ignores)
		if err != nil {
			outErr = errors.Join(outErr, err)
			return false
		}

		return ignored
	})

	for _, path := range cleanPaths {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return nil, fmt.Errorf("invalid path %q: %w", path, err)
		}

		if !fileInfo.IsDir() {
			result = append(result, path)
			continue
		}

		if err := filepath.WalkDir(path, func(child string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			ignore, err := IsIgnored(d.Name(), ignores)
			if err != nil {
				return err
			}

			if ignore {
				if d.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}

			if !d.IsDir() {
				result = append(result, child)
			}

			return nil
		}); err != nil {
			return nil, fmt.Errorf("failed to crawl path %q: %w", path, err)
		}
	}

	return result, nil
}

func IsIgnored(path string, ignores []string) (bool, error) {
	pathIsAbs := filepath.IsAbs(path)
	for _, ignore := range ignores {
		if strings.HasPrefix(ignore, "/") && !pathIsAbs {
			pwd, err := os.Getwd()
			if err != nil {
				return false, fmt.Errorf("failed to get current directory: %w", err)
			}

			relPath, err := filepath.Rel(pwd, path)
			if err != nil {
				// In that case, the path and ignore pattern won't match
				return false, nil
			}

			path = relPath
		}

		re, err := regexp.Compile(ignore)
		if err != nil {
			return false, fmt.Errorf("ignore pattern %q is an invalid regex: %w", ignore, err)
		}

		if re.MatchString(path) {
			return true, nil
		}
	}

	return false, nil
}
