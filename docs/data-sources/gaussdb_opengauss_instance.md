---
subcategory: "GaussDB"
---

# hcs_gaussdb_opengauss_instance

Use this data source to get available HuaweiCloudStack gaussdb opengauss instance.

-> **NOTE:** If the endpoint is manually configured, **opengaussv31** should be configured.

## Example Usage

```hcl
data "hcs_gaussdb_opengauss_instance" "this" {
  name = "gaussdb-instance"
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the instance. If omitted, the provider-level region will
  be used.

* `name` - (Optional, String) Specifies the name of the instance.

* `vpc_id` - (Optional, String) Specifies the VPC ID.

* `subnet_id` - (Optional, String) Specifies the network ID of a subnet.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Indicates the ID of the instance.

* `status` - Indicates the DB instance status.

* `type` - Indicates the instance type.

* `flavor` - Indicates the instance specifications.

* `security_group_id` - Indicates the security group ID.

* `enterprise_project_id` - Indicates the enterprise project id.

* `db_user_name` - Indicates the default username.

* `time_zone` - Indicates the time zone.

* `availability_zone` - Indicates the instance availability zone.

* `port` - Indicates the database port.

* `switch_strategy` - Indicates the switch strategy.

* `maintenance_window` - Indicates the maintenance window.

* `coordinator_num` - Indicates the count of coordinator node.

* `sharding_num` - Indicates the sharding num.

* `replica_num` - Indicates the replica num.

* `private_ips` - Indicates the list of private IP address of the nodes.

* `public_ips` - Indicates the public IP address of the DB instance.

* `volume` - Indicates the volume information. The [volume](#opengauss_volume) object structure is documented below.

* `datastore` - Indicates the database information. The [datastore](#opengauss_datastore) object structure is
  documented below.

* `backup_strategy` - Indicates the advanced backup policy. The [backup_strategy](#opengauss_backup_strategy) object
  structure is documented below.

* `nodes` - Indicates the instance nodes information. The [nodes](#opengauss_nodes) object structure is
  documented below.

* `ha` - Indicates the instance ha information. The [ha](#opengauss_ha) object structure is documented below.

<a name="opengauss_volume"></a>
The `volume` block supports:

* `type` - Indicates the volume type.

* `size` - Indicates the volume size.

<a name="opengauss_datastore"></a>
The `datastore` block supports:

* `engine` - Indicates the database engine.

* `version` - Indicates the database version.

<a name="opengauss_backup_strategy"></a>
The `backup_strategy` block supports:

* `start_time` - Indicates the backup time window.

* `keep_days` - Indicates the number of days to retain the generated

<a name="opengauss_nodes"></a>
The `nodes` block contains:

* `id` - Indicates the node ID.

* `name` - Indicates the node name.

* `status` - Indicates the node status.

* `role` - Indicates whether the node support reduce.

* `availability_zone` - Indicates the availability zone where the node resides.

* `private_ip` - Indicates the internal IP address of the node. This parameter is valid only for CN nodes in the
  distributed edition. This parameter is valid for all nodes in the centralized edition.
  The parameter value exists after the ECS is created.

* `public_ip` - Indicates the bound external IP address of the node. This parameter is valid only for CN nodes in the
  distributed edition. This parameter is valid for all nodes in the centralized edition.
  The parameter value exists after the ECS is created and bound to an EIP.

* `data_ip` - Indicates the data ip of the node.

* `bms_hs_ip` - IP address of the high-speed NIC, which is a dedicated IP field of the BMS instance and is used for
  data synchronization.

* `management_ip` - Indicates the management ip of the node.

<a name="opengauss_ha"></a>
The `ha` block supports:

* `replication_mode` - Indicates the replication mode.
