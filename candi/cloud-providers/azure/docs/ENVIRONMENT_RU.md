---
title: "Cloud provider — Azure: подготовка окружения"
---

Для управления облаком Microsoft Azure, необходимо иметь в нём учётную запись и хотя бы одну [Subscription, привязанную к ней](https://docs.microsoft.com/en-us/azure/cost-management-billing/manage/create-subscription).

Для управления ресурсами в облаке Microsoft Azure средствами Deckhouse, необходимо создать сервисный аккаунт:
- Установите [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli), залогиньтесь и получите Subscription ID:
```shell
export SUBSCRIPTION_ID=$(az login | jq -r '.[0].id')
```
- Создайте service account:
```shell
az ad sp create-for-rbac --role="Contributor" --scopes="/subscriptions/$SUBSCRIPTION_ID" --name "DeckhouseCANDI"
```

Для дальнейшей работы с утилитой az необходимо быть залогиненным. Если сервисный аккаунт уже создан и известен его логин/пароль/тенант, то можно использовать эти данные:
```shell
az login --service-principal -u <username> -p <password> --tenant <tenant>
```
