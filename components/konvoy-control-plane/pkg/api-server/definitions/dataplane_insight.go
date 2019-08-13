package definitions

import (
	"github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/apis/mesh"
	"github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/model"
)

var DataplaneInsightWsDefinition = ResourceWsDefinition{
	Name: "Dataplane Insight",
	Path: "dataplane-insights",
	ResourceFactory: func() model.Resource {
		return &mesh.DataplaneInsightResource{}
	},
	ResourceListFactory: func() model.ResourceList {
		return &mesh.DataplaneInsightResourceList{}
	},
}