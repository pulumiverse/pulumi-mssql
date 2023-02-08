# Pulumi Provider for Microsoft SQL Server and Azure SQL

This repository will contain the Pulumi provider for Microsoft SQL Server and Azure SQL.

The [upstream Terraform Provider](https://github.com/PGSSoft/terraform-provider-mssql) uses the
new [Terraform Plugin Framework]( github.com/hashicorp/terraform-plugin-framework), which unfortunately
is not currently supported by the [Pulumi Terraform Bridge]( https://github.com/pulumi/pulumi-terraform-bridge).
Development on the provider will start shortly when the Terraform Pulumi Bridge supports the Terraform Plugin Framework.

In the meantime, if you want to push the development of the required changes to Pulumi Terraform Bridge, give the
respective [issue](https://github.com/pulumi/pulumi-terraform-bridge/issues/744) a thumbs up.
