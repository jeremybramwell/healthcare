package config_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/healthcare/deploy/config"
	"github.com/ghodss/yaml"
)

func TestDefaultResource(t *testing.T) {
	resYAML := `
properties:
  name: foo-resource
`

	d := new(config.DefaultResource)
	if err := yaml.Unmarshal([]byte(resYAML), d); err != nil {
		t.Fatalf("yaml unmarshal: %v", err)
	}
	d.TmplPath = "foo"
	if err := d.Init(); err != nil {
		t.Fatalf("d.Init = %v", err)
	}
	if d.Name() != "foo-resource" {
		t.Fatalf("default resource name error: %v", d.Name())
	}
}
