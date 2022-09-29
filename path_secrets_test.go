package kubesecrets

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/vault/sdk/logical"

	"github.com/hashicorp/go-hclog"
)

func getTestBackend(t *testing.T) logical.Backend {
	t.Logf("Testing Fooxx")
	b, _ := newBackend()
	t.Logf("Testing Foo")
	c := &logical.BackendConfig{
		Logger: hclog.New(&hclog.LoggerOptions{}),
	}
	t.Logf("Testing Foo1")
	err := b.Setup(context.Background(), c)
	t.Logf("Testing Foo2")
	if err != nil {
		t.Fatalf("unable to create backend: %v", err)
	}
	t.Logf("Testing Foo3")
	return b
}

func TestSecretNamespaceMissing(t *testing.T) {
	t.Logf("Testing Foo7")
	b := getTestBackend(t)

	request := &logical.Request{
		Operation: logical.ReadOperation,
		Path:      fmt.Sprintf("%s/", secretsPrefix),
		Data:      make(map[string]interface{}),
	}

	errMsg := "Missing secret namespace"
	resp, _ := b.HandleRequest(context.Background(), request)
	if resp.Error().Error() != errMsg {
		t.Errorf("Error must be '%s', get '%s'", errMsg, resp.Error())
	}
}
