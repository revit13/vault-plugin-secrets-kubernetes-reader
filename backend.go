package kubesecrets

import (
	/*"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"
	kconfig "sigs.k8s.io/controller-runtime/pkg/client/config"*/

	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"
	kconfig "sigs.k8s.io/controller-runtime/pkg/client/config"
)

// backend wraps the backend framework
type secretsReaderBackend struct {
	*framework.Backend

	KubeSecretReader KubernetesSecretsReader
}

var _ logical.Factory = Factory

// Factory configures and returns the plugin backends
func Factory(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	scheme := runtime.NewScheme()
	err := clientgoscheme.AddToScheme(scheme)
	if err != nil {
		return nil, err
	}
	client, err := kclient.New(kconfig.GetConfigOrDie(), kclient.Options{Scheme: scheme})
	if err != nil {
		return nil, err
	}
	b, err := newBackend(&client)
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

func newBackend(k8sClient *kclient.Client, b_optional ...*testing.T) (*secretsReaderBackend, error) {
	var t *testing.T
	if len(b_optional) > 0 {
		t = b_optional[0]
	}
	t.Logf("Testing Fooffff")

	t.Logf("Testing Fooffeeff")

	// TODO: support configuration where Vault installed out of cluster

	t.Logf("Testing Foofeeeeefeeff")

	scheme := runtime.NewScheme()
	_ = clientgoscheme.AddToScheme(scheme)

	fc := fake.NewSimpleClientset(&exampleDeploy)

	// TODO: support configuration where Vault installed out of cluster
	client, err := kclient.New(kconfig.GetConfigOrDie(), kclient.Options{Scheme: scheme})
	if err != nil {
		return nil, err
	}
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
