// Package schemas embeds the Capybari Source Intelligence JSON Schemas so Go
// programs can validate reports, findings and capability metadata. The
// .schema.json files are the language-neutral source of truth.
package schemas

import "embed"

// Version is the schema version implemented by these files.
const Version = "0.1"

// FS holds finding.schema.json, capability.schema.json and report.schema.json.
//
//go:embed schemas/*.schema.json
var FS embed.FS

// BaseURI is the $id prefix shared by all schemas.
const BaseURI = "https://schemas.capybari.com/source-intelligence/0.1/"

// Names lists the schema files.
var Names = []string{"finding.schema.json", "capability.schema.json", "report.schema.json"}
