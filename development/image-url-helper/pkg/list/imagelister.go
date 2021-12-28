package list

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/kyma-project/test-infra/development/image-url-helper/pkg/common"
	"gopkg.in/yaml.v2"
)

func GetWalkFunc(resourcesDirectory string, images, testImages common.ComponentImageMap) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		//pass the error further, this shouldn't ever happen
		if err != nil {
			return err
		}

		// skip directory entries, we just want files
		if info.IsDir() {
			return nil
		}

		// we only want to check values.yaml files
		if info.Name() != "values.yaml" {
			return nil
		}

		var parsedFile common.ValueFile

		yamlFile, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		err = yaml.Unmarshal(yamlFile, &parsedFile)
		if err != nil {
			return err
		}

		component := strings.Replace(path, resourcesDirectory+"/", "", -1)
		component = strings.Replace(component, "/values.yaml", "", -1)

		common.AppendImagesToMap(parsedFile, images, testImages, component)

		return nil
	}
}
