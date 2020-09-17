package lookup

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/markbates/pkger"

	"github.com/spf13/viper"
)

func downloadFile(filepath, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func lookupEmbeddedPlugin(pluginName, pluginPath string) (bool, error) {

	info, _ := pkger.Stat(path.Join("/plugins", pluginName))
	if info == nil {
		return false, nil
	}

	fmt.Printf("Extracting plugin from embedded files\n")

	// Extract embedded plugin
	f, err := pkger.Open(path.Join("/plugins", pluginName))
	if err != nil {
		return true, err
	}
	defer f.Close()

	pluginContent, err := ioutil.ReadAll(f)
	if err != nil {
		return true, err
	}

	err = ioutil.WriteFile(pluginPath, pluginContent, 0640)
	if err != nil {
		return true, err
	}

	return true, nil
}

func lookupGithubPlugin(pluginName, pluginPath string) (bool, error) {
	fileUrl := fmt.Sprintf("https://github.com/cycloidio/%s", pluginName)
	fmt.Println("Downloading: " + fileUrl)

	err := downloadFile(pluginPath, fileUrl)
	if err != nil {
		return false, err
	}
	fmt.Println("Downloaded: " + fileUrl)
	return true, nil
}

func lookupLocalPlugin(pluginName, pluginPath string) (bool, error) {
	info, err := os.Stat(pluginPath)
	if err != nil && !os.IsNotExist(err) {
		return false, err
	}
	if info == nil {
		return false, nil
	}
	return true, nil
}

func LookupPlugin(version string) (string, error) {

	pluginDir := viper.GetString("cy-plugin-dir")
	pluginName := fmt.Sprintf("v%s.so", version)
	pluginPath := path.Join(pluginDir, pluginName)

	// Init, if pluginDir does not exist, try to create it
	info, _ := os.Stat(pluginDir)
	if info == nil {
		err := os.Mkdir(pluginDir, 0750)
		if err != nil {
			return "", err
		}
	}

	// check if we have the plugin locally
	found, err := lookupLocalPlugin(pluginName, pluginPath)
	if err != nil {
		return "", err
	}
	if found {
		return pluginPath, nil
	}

	// check if embedded so extract it
	found, err = lookupEmbeddedPlugin(pluginName, pluginPath)
	if err != nil {
		return "", err
	}
	if found {
		return pluginPath, nil
	}

	// TODO: if not download it from the API

	// if not download it from internet
	found, err = lookupGithubPlugin(pluginName, pluginPath)
	if err != nil {
		return "", err
	}
	if found {
		return pluginPath, nil
	}

	return "", errors.New("Unable to find any plugin for your API Version")
}
