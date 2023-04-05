---
title: Mssql Setup
meta_desc: Information on how to install the Mssql provider.
layout: installation
---

## Installation

The Pulumi Mssql provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/mssql`](https://www.npmjs.com/package/@pulumiverse/mssql)
* Python: [`pulumiverse_mssql`](https://pypi.org/project/pulumiverse_mssql/)
* Go: [`github.com/pulumiverse/pulumi-mssql/sdk/go/mssql`](https://github.com/pulumiverse/pulumi-mssql/sdk/go/mssql)
* .NET: [`Pulumiverse.Mssql`](https://www.nuget.org/packages/Pulumiverse.Mssql)

### Provider Binary

The Mssql provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource mssql v0.1.7
```

Replace the version string with your desired version.
