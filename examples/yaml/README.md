# Pulumi MSSQL

### Prerequisites

1. [Install Pulumi](https://www.pulumi.com/docs/get-started/install/)

### Steps

1. Pulumi login

    ```bash
    pulumi login
    ```

1. Create a new stack:

    ```bash
    pulumi stack init mssql-dev
    pulumi stack select mssql-dev
    ```

1. Login to Azure CLI (you will be prompted to do this during deployment if you forget this step):

    ```
    $ az login
    ```

    > ***Note:***  
    > you can use Service Principal Login as well
    > https://www.pulumi.com/registry/packages/azure/installation-configuration/#authenticate-using-a-service-principal

1. Configure the required settings:

    ```bash
    pulumi config set local-ip <your-external-ip>
    pulumi config set resourcegroup-location <azure-region>
    pulumi config set resourcegroup-name <name-of-resource-group>
    ```

1. Run `pulumi up` to preview and deploy changes:

    ```bash
    pulumi up
    Previewing changes:
    ...
