package kubesecrets

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"
	ctrl "sigs.k8s.io/controller-runtime"
)

// backend wraps the backend framework
type secretsReaderBackend struct {
	*framework.Backend

	KubeSecretReader KubernetesSecretsReader
}

var _ logical.Factory = Factory

// Factory configures and returns the plugin backends
func Factory(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	b, err := newBackend()
	if err != nil {
		return nil, err
	}

	if conf == nil {
		return nil, fmt.Errorf("configuration passed into backend is nil")
	}

	if err := b.Setup(ctx, conf); err != nil {
		return nil, err
	}

	return b, nil
}

func newBackend(b_optional ...*testing.T) (*secretsReaderBackend, error) {
	var t *testing.T
	if len(b_optional) > 0 {
		t = b_optional[0]
	}
	t.Logf("Testing Fooffff")
	scheme := runtime.NewScheme()
	t.Logf("Testing rrr")
	err := clientgoscheme.AddToScheme(scheme)
	if err != nil {
		return nil, err
	}
	t.Logf("Testing Fooffeeff")

	// TODO: support configuration where Vault installed out of cluster
	client, err := kclient.New(ctrl.GetConfigOrDie(), kclient.Options{Scheme: scheme})
	if err != nil {
		return nil, err
	}
	t.Logf("Testing Foofeeeeefeeff")
	b := &secretsReaderBackend{
		KubeSecretReader: KubernetesSecretsReader{
			client: client,
		},
	}
	t.Logf("Testing Foofee444444eeefeeff")
	b.Backend = &framework.Backend{
		Help: strings.TrimSpace(backendHelp),
		// TypeLogical indicates that the backend (plugin) is a secret provider.
		BackendType: logical.TypeLogical,
		// Define the path for which this backend will respond.
		Paths: []*framework.Path{
			pathSecrets(b),
		},
	}
	t.Logf("Testing Foofeeeeerrerereeeefeeff")

	return b, nil
}
