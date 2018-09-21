# paas-terraform-provider

This project is a terraform provider to interact with the
[GOV.UK PaaS](https://www.cloud.service.gov.uk/).

The canonical way to use Cloud Foundry is to use the CF CLI and operations such
as:

- `cf push`
- `cf create-domain`
- `cf map-route`
- `cf marketplace`

This workflow is do-able but ideally would be more declarative, hence a
terraform provider.

# Resources

This project currently does not implement any resources.

# Data sources

## Org

```
data "paas_org" "myorg" {
  name = "my-org-name"
}
```

### Required attributes:

- `name`

### Computed:

- `guid`

## Space

```
data "paas_space" "space" {
  name     = "my-space-name"
  org_guid = "${data.paas_org.myorg.guid}"
}
```

### Required attributes:

- `name`
- `org_guid`

### Computed:

- `guid`

## User

User currently requires `org_guid` otherwise you need global user read
permissions in order to get information about them.

```
data "paas_user" "toby" {
  name     = "toby.lornewelch-richards@digital.cabinet-office.gov.uk"
  org_guid = "${data.paas_org.myorg.guid}"
}
```

### Required attributes:

- `name`
- `org_guid`

### Computed:

- `guid`

- `guid`

## Domain

Domains are for owned domains, NOT shared domains.

Shared domains are domains like:

- `cloudapps.digital`
- `apps.internal`

```
data "paas_domain" "mydomain" {
  name = "my-domain.com"
}
```

### Required attributes:

- `name`

### Computed:

- `guid`

## Shared domain

Shared domains are domains like:

- `cloudapps.digital`
- `apps.internal`

```
data "paas_shared_domain" "cloudapps" {
  name = "cloudapps.digital"
}
```

### Required attributes:

- `name`

### Computed:

- `guid`

## App
```
data "paas_app" "myapp" {
  name       = "myapp"
  org_guid   = "${data.paas_org.myorg.guid}"
  space_guid = "${data.paas_space.myspace.guid}"
}
```

### Required attributes:

- `name`
- `org_guid`
- `space_guid`

### Computed:

- `guid`
