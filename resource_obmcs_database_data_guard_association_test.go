// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
)

type ResourceDataGuardAssociationTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.DataGuardAssociation
	DeletedRes   *baremetal.DataGuardAssociation
}

func (s *ResourceDataGuardAssociationTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders

	s.TimeCreated = baremetal.Time{Time: time.Now()}

	s.Config = `
resource "oci_database_data_guard_association" "testDataGuardAssociation" {
	creation_type = "${var.creation_type}"
	database_admin_password = "${var.database_admin_password}"
	database_id = "${var.database_id}"
	protection_mode = "${var.protection_mode}"
	transport_type = "${var.transport_type}"
}

	`

	s.Config += testProviderConfig()

	s.ResourceName = "oci_database_data_guard_association.t"

}

func (s *ResourceDataGuardAssociationTestSuite) TestCreateResourceDataGuardAssociation() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
				),
			},
		},
	})
}

func (s ResourceDataGuardAssociationTestSuite) TestUpdateCreationTypeForcesNewDataGuardAssociation() {

	config := `
resource "oci_database_data_guard_association" "testDataGuardAssociation" {
	creation_type = "${var.creation_type}"
	database_admin_password = "${var.database_admin_password}"
	database_id = "${var.database_id}"
	protection_mode = "${var.protection_mode}"
	transport_type = "${var.transport_type}"
}
`

	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "creation_type", "newValue"),
				),
			},
		},
	})
}

func (s ResourceDataGuardAssociationTestSuite) TestUpdateDatabaseAdminPasswordForcesNewDataGuardAssociation() {

	config := `
resource "oci_database_data_guard_association" "testDataGuardAssociation" {
	creation_type = "${var.creation_type}"
	database_admin_password = "${var.database_admin_password}"
	database_id = "${var.database_id}"
	protection_mode = "${var.protection_mode}"
	transport_type = "${var.transport_type}"
}
`

	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "database_admin_password", "newValue"),
				),
			},
		},
	})
}

func (s ResourceDataGuardAssociationTestSuite) TestUpdateDatabaseIDForcesNewDataGuardAssociation() {

	config := `
resource "oci_database_data_guard_association" "testDataGuardAssociation" {
	creation_type = "${var.creation_type}"
	database_admin_password = "${var.database_admin_password}"
	database_id = "${var.database_id}"
	protection_mode = "${var.protection_mode}"
	transport_type = "${var.transport_type}"
}
`

	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "database_id", "newValue"),
				),
			},
		},
	})
}

func (s ResourceDataGuardAssociationTestSuite) TestUpdateProtectionModeForcesNewDataGuardAssociation() {

	config := `
resource "oci_database_data_guard_association" "testDataGuardAssociation" {
	creation_type = "${var.creation_type}"
	database_admin_password = "${var.database_admin_password}"
	database_id = "${var.database_id}"
	protection_mode = "${var.protection_mode}"
	transport_type = "${var.transport_type}"
}
`

	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "protection_mode", "newValue"),
				),
			},
		},
	})
}

func (s ResourceDataGuardAssociationTestSuite) TestUpdateTransportTypeForcesNewDataGuardAssociation() {

	config := `
resource "oci_database_data_guard_association" "testDataGuardAssociation" {
	creation_type = "${var.creation_type}"
	database_admin_password = "${var.database_admin_password}"
	database_id = "${var.database_id}"
	protection_mode = "${var.protection_mode}"
	transport_type = "${var.transport_type}"
}
`

	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "transport_type", "newValue"),
				),
			},
		},
	})
}

func TestResourceDataGuardAssociationTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceDataGuardAssociationTestSuite))
}
