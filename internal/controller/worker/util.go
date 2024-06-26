package worker

import (
	"github.com/zncdatadev/trino-operator/internal/common"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	logger = ctrl.Log.WithName("worker")
)

func createWorkerConfigmapName(instanceName string, groupName string) string {
	return common.NewResourceNameGenerator(instanceName, string(common.Worker), groupName).GenerateResourceName("")
}

func createWorkerDeploymentName(instanceName string, groupName string) string {
	return common.NewResourceNameGenerator(instanceName, string(common.Worker), groupName).GenerateResourceName("")
}

func GetRole() common.Role {
	return common.Worker
}
