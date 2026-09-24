package schemas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

var (
	once     sync.Once
	compiled map[string]*jsonschema.Schema
	initErr  error
)

func load() {
	c := jsonschema.NewCompiler()
	c.AssertFormat()
	for _, n := range Names {
		b, err := FS.ReadFile("schemas/" + n)
		if err != nil {
			initErr = err
			return
		}
		doc, err := jsonschema.UnmarshalJSON(bytes.NewReader(b))
		if err != nil {
			initErr = fmt.Errorf("%s: %w", n, err)
			return
		}
		if err := c.AddResource(BaseURI+n, doc); err != nil {
			initErr = err
			return
		}
	}
	compiled = map[string]*jsonschema.Schema{}
	for _, n := range Names {
		s, err := c.Compile(BaseURI + n)
		if err != nil {
			initErr = fmt.Errorf("compile %s: %w", n, err)
			return
		}
		compiled[n] = s
	}
}

// Validate checks a JSON document against the named schema
// (e.g. "report.schema.json").
func Validate(name string, doc []byte) error {
	once.Do(load)
	if initErr != nil {
		return initErr
	}
	s, ok := compiled[name]
	if !ok {
		return fmt.Errorf("unknown schema %q", name)
	}
	v, err := jsonschema.UnmarshalJSON(bytes.NewReader(doc))
	if err != nil {
		return err
	}
	return s.Validate(v)
}

// ValidateValue marshals v to JSON and validates it.
func ValidateValue(name string, v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return Validate(name, b)
}
