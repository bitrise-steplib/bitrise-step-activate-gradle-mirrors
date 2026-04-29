# Activate Gradle Mirrors

[![Step changelog](https://shields.io/github/v/release/bitrise-steplib/bitrise-step-activate-gradle-mirrors?include_prereleases&label=changelog&color=blueviolet)](https://github.com/bitrise-steplib/bitrise-step-activate-gradle-mirrors/releases)

Activates Bitrise repository mirrors for subsequent Gradle builds in the workflow

<details>
<summary>Description</summary>

This step installs a Gradle init script that redirects repository requests to
Bitrise-hosted mirrors for faster, more reliable dependency resolution in
subsequent Gradle invocations.

The step is gated by the `BITRISE_MAVENCENTRAL_PROXY_ENABLED` environment
variable and only activates when run in a Bitrise datacenter that has a
mirror deployment. In every other case it logs an info message and exits
successfully.

</details>

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://docs.bitrise.io/en/bitrise-ci/workflows-and-pipelines/steps/adding-steps-to-a-workflow.html).

You can also run this step directly with [Bitrise CLI](https://github.com/bitrise-io/bitrise).

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `verbose` | Enable logging additional information for troubleshooting | required | `false` |
</details>

<details>
<summary>Outputs</summary>
There are no outputs defined in this step
</details>

## 🙋 Contributing

We welcome [pull requests](https://github.com/bitrise-steplib/bitrise-step-activate-gradle-mirrors/pulls) and [issues](https://github.com/bitrise-steplib/bitrise-step-activate-gradle-mirrors/issues) against this repository.

For pull requests, work on your changes in a forked repository and use the Bitrise CLI to [run step tests locally](https://docs.bitrise.io/en/bitrise-ci/bitrise-cli/running-your-first-local-build-with-the-cli.html).

Learn more about developing steps:

- [Create your own step](https://docs.bitrise.io/en/bitrise-ci/workflows-and-pipelines/developing-your-own-bitrise-step/developing-a-new-step.html)
