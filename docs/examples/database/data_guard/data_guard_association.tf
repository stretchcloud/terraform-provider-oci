variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "creation_type" {}
variable "database_admin_password" {}
variable "database_id" {}
variable "protection_mode" {}
variable "transport_type" {}


provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}

resource "oci_database_data_guard_association" "testDataGuardAssociation" {
	creation_type = "${var.creation_type}"
	database_admin_password = "${var.database_admin_password}"
	database_id = "${var.database_id}"
	protection_mode = "${var.protection_mode}"
	transport_type = "${var.transport_type}"
}


data "oci_database_data_guard_associations" "testDataGuardAssociations" {
	database_id = "${var.database_id}"
}
