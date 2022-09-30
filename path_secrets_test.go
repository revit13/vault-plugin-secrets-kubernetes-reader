package kubesecrets

import (
	"context"
	"fmt"
	"testing"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/hashicorp/vault/sdk/logical"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/hashicorp/go-hclog"
)

func getTestBackend(k8sclient *client.Client, t *testing.T) (logical.Backend, error) {
	t.Logf("Testing Fooxx")
	b, err := newBackend(k8sclient, t)
	if err != nil {
		t.Logf("Testing hererer")
		return nil, err
	}
	t.Logf("Testing Foo")
	c := &logical.BackendConfig{
		Logger: hclog.New(&hclog.LoggerOptions{}),
	}
	t.Logf("Testing Foo1")
	err = b.Setup(context.Background(), c)
	t.Logf("Testing Foo2")
	if err != nil {
		t.Fatalf("unable to create backend: %v", err)
	}
	t.Logf("Testing Foo3")
	return b, nil
}

func TestSecretNamespaceMissing(t *testing.T) {
	gomega.RegisterFailHandler(Fail)
	g := gomega.NewGomegaWithT(t)
	defer GinkgoRecover()
	t.Logf("Testing Foo7")
	/*scheme := runtime.NewScheme()
	err := clientgoscheme.AddToScheme(scheme)
	g.Expect(err).To(gomega.BeNil())
	// TODO: support configuration where Vault installed out of cluster
	k8sClient, err := kclient.New(kconfig.GetConfigOrDie(), kclient.Options{Scheme: scheme})
	g.Expect(err).To(gomega.BeNil())*/

	b, err := getTestBackend(nil, t)
	g.Expect(err).To(gomega.BeNil())

	request := &logical.Request{
		Operation: logical.ReadOperation,
		Path:      fmt.Sprintf("%s/", secretsPrefix),
		Data:      make(map[string]interface{}),
	}

	errMsg := "Missing secret namespace"
	resp, err := b.HandleRequest(context.Background(), request)
	g.Expect(err.Error()).Should(gomega.Equal(resp.Error().Error()), errMsg)
}
