---
layout: "alicloud"
page_title: "Alicloud: alicloud_bgp_group"
sidebar_current: "docs-alicloud-resource-bgp-group"
description: |-
  Provides a Alicloud bgp group resource.
---

# alicloud\_bgp_group

Provides a BGP Group resource.

-> **NOTE:** Terraform will auto build bgp group instance  while it uses `alicloud\_bgp_group to build a bgp group resource.

## Example Usage

Basic Usage

```
resource "alicloud_vpn_gateway" "foo" {
    dst_cidr_block  =   "192.168.2.11"
    router_id       =   "vbr-xxxxxxxxxxxxxx"
}
```
## Argument Reference

The following arguments are supported:

* `dst_cidr_block` - (Required, ForceNew) The network segment of the VPC or switch that needs to be interconnected with the local IDC.
* `router_id` - (Required, ForceNew) The ID of the vbr instance.

## Attributes Reference

The following attributes are exported:

* `id` - The Combination ID of the BGP Network.

BGP Network can be imported using the id, e.g.

```
$ terraform import alicloud_bgp_network.example vbr-123456:192.168.2.11
```
