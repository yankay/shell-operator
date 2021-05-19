package fake

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestRegisterCRD(t *testing.T) {
	f := NewFakeCluster("")
	f.RegisterCRD("deckhouse.io", "v1alpha2", "NodeGroup", false)

	_, err := f.KubeClient.Dynamic().Resource(schema.GroupVersionResource{
		Group:    "deckhouse.io",
		Version:  "v1alpha2",
		Resource: Pluralize("NodeGroup"),
	}).List(context.TODO(), v1.ListOptions{})
	require.NoError(t, err)
}
