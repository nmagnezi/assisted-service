package hardware

import (
	"github.com/openshift/assisted-service/internal/common"
	models "github.com/openshift/assisted-service/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func GetBootDevice(log logrus.FieldLogger, hwValidator Validator, host *models.Host) (string, error) {
	path := hwValidator.GetHostInstallationPath(&common.Host{Host: *host})

	if path != "" {
		return path, nil
	}

	log.Errorf("Failed to determine installation path for host with id %s", host.ID)
	return "", errors.Errorf("host has no installation path %s", host.ID)
}
