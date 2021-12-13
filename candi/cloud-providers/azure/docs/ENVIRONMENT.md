---
title: "Cloud provider — Azure: Preparing environment"
---

To rule the Microsoft Azure cloud, you need an account and at least a single [Subscription connected to id](https://docs.microsoft.com/en-us/azure/cost-management-billing/manage/create-subscription).

You have to create a service account with Microsoft Azure so that Deckhouse can manage cloud resources:
- Install the [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli), login and get Subscription ID:
```shell
export SUBSCRIPTION_ID=$(az login | jq -r '.[0].id')
```
- Create the service account:
```shell
az ad sp create-for-rbac --role="Contributor" --scopes="/subscriptions/$SUBSCRIPTION_ID" --name "DeckhouseCANDI"
```

For further work with the az tool, you have to be logged in. If you have a service account username, password, and tenant, you're able to use them:
```shell
az login --service-principal -u <username> -p <password> --tenant <tenant>
```
