// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/oracle/bmcs-go-sdk"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceDataGuardAssociationTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListDataGuardAssociations
}

func (s *DatasourceDataGuardAssociationTestSuite) SetupTest() {
	s.Client = testAccClient
    s.Provider = testAccProvider
    s.Providers = testAccProviders
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
	s.ResourceName = "data.oci_database_data_guard_associations.t"
}

func (s *DatasourceDataGuardAssociationTestSuite) TestReadDataGuardAssociations() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
data "oci_database_data_guard_associations" "testDataGuardAssociations" {
	database_id = "${var.database_id}"
}

				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "data_guard_associations.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "data_guard_associations.#"),
				),
			},
		},
	},
	)
}

func TestDatasourceDataGuardAssociationTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceDataGuardAssociationTestSuite))
}
