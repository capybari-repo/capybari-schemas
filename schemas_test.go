package schemas

import (
	"os"
	"testing"
)

func TestSchemasCompileAndExamplesValidate(t *testing.T) {
	b, err := os.ReadFile("examples/capability.example.json")
	if err != nil {
		t.Fatal(err)
	}
	if err := Validate("capability.schema.json", b); err != nil {
		t.Fatal(err)
	}
	bad := []byte(`{"id":"Bad Id"}`)
	if err := Validate("capability.schema.json", bad); err == nil {
		t.Fatal("expected validation error")
	}
	f := []byte(`{"id":"CSI-0123456789abcdef","dimension":"security","category":"secret","title":"t","severity":"high","confidence":"high","source":{"capability":"secrets","version":"0.1.0"},"detected_at":"2026-01-01T00:00:00Z"}`)
	if err := Validate("finding.schema.json", f); err != nil {
		t.Fatal(err)
	}
	if err := Validate("finding.schema.json", []byte(`{"severity":"urgent"}`)); err == nil {
		t.Fatal("expected finding validation error")
	}
}
