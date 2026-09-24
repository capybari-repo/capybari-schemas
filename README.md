# capybari-schemas

**Capybari Source Intelligence: language-neutral JSON Schemas.**

| Schema | Describes |
|---|---|
| [`schemas/finding.schema.json`](schemas/finding.schema.json) | One finding: dimension, category, severity, confidence, evidence, source analyzer, rule, remediation. |
| [`schemas/capability.schema.json`](schemas/capability.schema.json) | A capability's registry metadata (`capability.yaml`): inputs, evidence, outputs, cost, network/AI requirements, privacy, follow-ups. |
| [`schemas/report.schema.json`](schemas/report.schema.json) | The unified Software/Website X-Ray report produced by the CLI, GitHub Action and hosted service. |

The schemas use JSON Schema draft 2020-12, `$id` prefix `https://schemas.capybari.com/source-intelligence/0.1/`.

Any tool, in any language, can consume Capybari reports or contribute capability metadata by validating against these files. The Go package in this repository embeds them and provides `schemas.Validate(name, json)`. `capybari-core` runs its reports through it in CI, so the Go types and these schemas cannot drift apart.

```bash
capybari validate capybari-report/report.json
```

## Versioning

Schema versions follow the report's `schema_version`. Breaking changes need a new major version directory and a migration note in `capybari-docs`.

## License

Apache-2.0
