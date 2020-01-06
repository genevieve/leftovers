# Leftovers :turkey:

[![GoDoc](https://godoc.org/github.com/genevieve/leftovers?status.svg)](https://godoc.org/github.com/genevieve/leftovers)

Go cli & library for cleaning up **orphaned IaaS resources**.

* <a href='#why'>Why might you use this?</a>
* <a href='#what'>What's it look like?</a>
* <a href='#how'>Installation</a>
* <a href='#usage'>Usage</a>
* <a href='#maintainers'>Maintainers</a>



## <a name='why'></a> Why might you use this?
- You `terraform apply`'d way back when and lost your `terraform.tfstate`. Happens to the best of us.
- You used the console or cli for some IaaS to create some infrastructure and want to clean up!
- Your acceptance tests in CI failed, the container disappeared, and
infrastructure resources were tragically orphaned. :-(
- `terraform destroy` isn't working because the refresh step is failing due to a missing resource.



## <a name='what'></a>What's it look like?

It will **prompt you before deleting** any resource by default, ie:

```console
$ leftovers --filter banana

[Firewall: banana-http] Delete? (y/N)
```

It can be configured to **not** prompt, ie:

```console
$ leftovers --filter banana --no-confirm

[Firewall: banana-http] Deleting...
[Firewall: banana-http] Deleted!
```

Or maybe you want to **see all of the resources** in your IaaS, ie:
```console
$ leftovers --filter banana --dry-run

[Firewall: banana-http]
[Network: banana]
```


Finally, you might want to delete a single resource type::
```console
$ leftovers types
service-account

$ leftovers --filter banana --type service-account --no-confirm
[Service Account: banana@pivotal.io] Deleting...
[Service Account: banana@pivotal.io] Deleted!
```



## <a name='how'></a>Installation

### Option 1
[Install go.](https://golang.org/doc/install) Then:

```console
$ go get -u github.com/genevieve/leftovers/cmd/leftovers
```

### Option 2

```console
$ brew tap genevieve/tap
$ brew install leftovers
```

### Option 3

Linux binaries can be found on the [releases page](https://github.com/genevieve/leftovers/releases).



## <a name='how'></a>Usage

```console
$ leftovers -h

Usage:
  leftovers [OPTIONS]

Application Options:
  -v, --version                   Print version.
  -i, --iaas=                     The IaaS for clean up. [$BBL_IAAS]
  -n, --no-confirm                Destroy resources without prompting. This is dangerous, make good choices!
  -d, --dry-run                   List all resources without deleting any.
  -f, --filter=                   Filtering resources by an environment name.
  -t, --type=                     Type of resource to delete.
      --aws-access-key-id=        AWS access key id. [$BBL_AWS_ACCESS_KEY_ID]
      --aws-secret-access-key=    AWS secret access key. [$BBL_AWS_SECRET_ACCESS_KEY]
      --aws-session-token=        AWS session token. [$BBL_AWS_SESSION_TOKEN]
      --aws-region=               AWS region. [$BBL_AWS_REGION]
      --azure-client-id=          Azure client id. [$BBL_AZURE_CLIENT_ID]
      --azure-client-secret=      Azure client secret. [$BBL_AZURE_CLIENT_SECRET]
      --azure-tenant-id=          Azure tenant id. [$BBL_AZURE_TENANT_ID]
      --azure-subscription-id=    Azure subscription id. [$BBL_AZURE_SUBSCRIPTION_ID]
      --gcp-service-account-key=  GCP service account key path. [$BBL_GCP_SERVICE_ACCOUNT_KEY]
      --vsphere-vcenter-ip=       vSphere vCenter IP address. [$BBL_VSPHERE_VCENTER_IP]
      --vsphere-vcenter-password= vSphere vCenter password. [$BBL_VSPHERE_VCENTER_PASSWORD]
      --vsphere-vcenter-user=     vSphere vCenter username. [$BBL_VSPHERE_VCENTER_USER]
      --vsphere-vcenter-dc=       vSphere vCenter datacenter. [$BBL_VSPHERE_VCENTER_DC]
      --nsxt-manager-host=        NSX-T manager IP address or domain name. [$BBL_NSXT_MANAGER_HOST]
      --nsxt-username=            NSX-T manager username. [$BBL_NSXT_USERNAME]
      --nsxt-password=            NSX-T manager password. [$BBL_NSXT_PASSWORD]
      --openstack-auth-url=       Openstack auth URL. [$BBL_OPENSTACK_AUTH_URL]
      --openstack-username=       Openstack username. [$BBL_OPENSTACK_USERNAME]
      --openstack-password=       Openstack password. [$BBL_OPENSTACK_PASSWORD]
      --openstack-domain-name=    Openstack domain name. [$BBL_OPENSTACK_DOMAIN]
      --openstack-project-name=   Openstack project name. [$BBL_OPENSTACK_PROJECT]
      --openstack-region-name=    Openstack region name. [$BBL_OPENSTACK_REGION]

Help Options:
  -h, --help                      Show this help message
```

## <a name='maintainers'></a>Maintainers

- [Genevieve L'Esperance](https://twitter.com/genevieve_vl)
- [Rowan Jacobs](https://github.com/rowanjacobs)
