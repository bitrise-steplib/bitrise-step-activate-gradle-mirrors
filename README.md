# Activate Gradle Mirrors

Activates Bitrise repository mirrors for subsequent Gradle builds in the workflow.

This step installs a Gradle init script (`~/.gradle/init.d/bitrise-gradle-mirrors.init.gradle.kts`) that redirects repository requests to Bitrise-hosted mirrors for faster, more reliable dependency resolution. Subsequent Gradle invocations in the same workflow pick the init script up automatically.

## When the step does nothing

The step gracefully no-ops (logs an info message and exits successfully) in any of these cases:

- The `BITRISE_MAVENCENTRAL_PROXY_ENABLED` environment variable is not set to `"true"` (typically: not running on Bitrise CI).
- `BITRISE_DEN_VM_DATACENTER` is empty or points at a datacenter that does not have a Bitrise mirror deployment.
- No mirror flag is enabled and the `KnownMirrors` registry is empty (cannot happen in practice).

## Inputs

| Key | Default | Notes |
| --- | --- | --- |
| `mavencentral` | `""` | When `"true"`, route Maven Central through the Bitrise mirror. |
| `mavencentral_apache` | `""` | When `"true"`, route `repo.maven.apache.org` through the Bitrise mirror. |
| `google` | `""` | When `"true"`, route the Google Maven repository through the Bitrise mirror. |
| `verbose` | `"false"` | Enable debug logging on the CLI. |

When all three mirror flags are left empty, the CLI defaults to enabling **all known mirrors**.
