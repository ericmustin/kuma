package mesh

import (
	"context"
	"fmt"

	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	"github.com/kumahq/kuma/pkg/core/resources/manager"
	"github.com/kumahq/kuma/pkg/core/resources/store"
)

var defaultTrafficPermission = mesh_proto.TrafficPermission{
	Sources: []*mesh_proto.Selector{
		{
			Match: map[string]string{
				mesh_proto.ServiceTag: "*",
			},
		},
	},
	Destinations: []*mesh_proto.Selector{
		{
			Match: map[string]string{
				mesh_proto.ServiceTag: "*",
			},
		},
	},
}

// TrafficPermission needs to contain mesh name inside it. Otherwise if the name is the same (ex. "allow-all") creating new mesh would fail because there is already resource of name "allow-all" which is unique key on K8S
func defaultTrafficPermissionName(meshName string) string {
	return fmt.Sprintf("allow-all-%s", meshName)
}

func createDefaultTrafficPermission(resManager manager.ResourceManager, meshName string) error {
	tp := &core_mesh.TrafficPermissionResource{
		Spec: defaultTrafficPermission,
	}
	return resManager.Create(context.Background(), tp, store.CreateByKey(defaultTrafficPermissionName(meshName), meshName))
}