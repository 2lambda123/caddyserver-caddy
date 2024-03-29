package integration

import (
	"testing"

	"github.com/caddyserver/caddy/v2"
	_ "github.com/caddyserver/caddy/v2/modules/standard"
)

// Validates that Caddy's registered internal modules implement the necessary interfaces of their
// respective namespaces
func TestTypes(t *testing.T) {
	var i int
	for _, v := range caddy.Modules() {
		mod, _ := caddy.GetModule(v)
		if ok, err := caddy.ConformsToNamespace(mod.New(), mod.ID.Namespace()); !ok {
			t.Errorf("%s", err)
		}
		i++
	}
	t.Logf("Passed through %d modules", i)
}
