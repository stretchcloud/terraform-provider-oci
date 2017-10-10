# oci\_database\_data_guard_associations

Gets a list of data_guard_associations.

## Example Usage

```
data "oci_database_data_guard_associations" "testDataGuardAssociations" {
	database_id = "${var.database_id}"
}

```

## Argument Reference

The following arguments are supported:

* `database_id` - (Required) The database [OCID](/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `data_guard_associations` - The list of data_guard_associations.

## DataGuardAssociation Reference
* `apply_lag` - The lag time between updates to the primary database and application of the redo data on the standby database, as computed by the reporting database.  Example: `9 seconds` 
* `apply_rate` - The rate at which redo logs are synced between the associated databases.  Example: `180 Mb per second` 
* `database_id` - The [OCID](/Content/General/Concepts/identifiers.htm) of the reporting database.
* `id` - The [OCID](/Content/General/Concepts/identifiers.htm) of the Data Guard association.
* `lifecycle_details` - Additional information about the current lifecycleState, if available. 
* `lifecycle_state` - The current state of the Data Guard association.
* `peer_data_guard_association_id` - The [OCID](/Content/General/Concepts/identifiers.htm) of the peer database's Data Guard association.
* `peer_database_id` - The [OCID](/Content/General/Concepts/identifiers.htm) of the associated peer database.
* `peer_db_home_id` - The [OCID](/Content/General/Concepts/identifiers.htm) of the database home containing the associated peer database. 
* `peer_db_system_id` - The [OCID](/Content/General/Concepts/identifiers.htm) of the DB System containing the associated peer database. 
* `peer_role` - The role of the peer database in this Data Guard association.
* `protection_mode` - The protection mode of this Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
* `role` - The role of the reporting database in this Data Guard association.
* `time_created` - The date and time the Data Guard Association was created.
* `transport_type` - The redo transport type used by this Data Guard association.  For more information, see [Redo Transport Services](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400) in the Oracle Data Guard documentation. 

