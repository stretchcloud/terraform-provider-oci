# oci\_database\_data_guard_association

## Example Usage

```
resource "oci_database_data_guard_association" "testDataGuardAssociation" {
	creation_type = "${var.creation_type}"
	database_admin_password = "${var.database_admin_password}"
	database_id = "${var.database_id}"
	protection_mode = "${var.protection_mode}"
	transport_type = "${var.transport_type}"
}

```

## Argument Reference

The following arguments are supported:

* `creation_type` - (Required) Specifies where to create the associated database. "ExistingDbSystem" is the only supported `creationType` value. 
* `database_admin_password` - (Required) A strong password for the `SYS`, `SYSTEM`, and `PDB Admin` users to apply during standby creation.  The password must contain no fewer than nine characters and include:  * At least two uppercase characters.  * At least two lowercase characters.  * At least two numeric characters.  * At least two special characters. Valid special characters include "_", "#", and "-" only.  **The password MUST be the same as the primary admin password.** 
* `database_id` - (Required) The database [OCID](/Content/General/Concepts/identifiers.htm).
* `protection_mode` - (Required) The protection mode to set up between the primary and standby databases. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation.  **IMPORTANT** - The only protection mode currently supported by the Database Service is MAXIMUM_PERFORMANCE. 
* `transport_type` - (Required) The redo transport type to use for this Data Guard association.  Valid values depend on the specified `protectionMode`:  * MAXIMUM_AVAILABILITY - SYNC or FASTSYNC * MAXIMUM_PERFORMANCE - ASYNC * MAXIMUM_PROTECTION - SYNC  For more information, see [Redo Transport Services](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400) in the Oracle Data Guard documentation.  **IMPORTANT** - The only transport type currently supported by the Database Service is ASYNC. 


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

