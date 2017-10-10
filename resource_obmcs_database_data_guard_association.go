// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"
	"github.com/oracle/terraform-provider-oci/crud"
)

func DataGuardAssociationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create: createDataGuardAssociation,
		Read:   readDataGuardAssociation,
		Schema: map[string]*schema.Schema{
			//Required
			"creation_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"database_admin_password": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"database_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"protection_mode": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"transport_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			//Optional

			//Computed
			"apply_lag": {
				Type: schema.TypeString,
				Computed: true,
			},
			"apply_rate": {
				Type: schema.TypeString,
				Computed: true,
			},
			"id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type: schema.TypeString,
				Computed: true,
			},
			"lifecycle_state": {
				Type: schema.TypeString,
				Computed: true,
			},
			"peer_data_guard_association_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"peer_database_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"peer_db_home_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"peer_db_system_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"peer_role": {
				Type: schema.TypeString,
				Computed: true,
			},
			"role": {
				Type: schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type: schema.TypeString,
				Computed: true,
			},

		},
	}
}

func createDataGuardAssociation(d *schema.ResourceData, m interface{}) (e error) {
	sync := &DataGuardAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.CreateResource(d, sync)
}

func readDataGuardAssociation(d *schema.ResourceData, m interface{}) (e error) {
	sync := &DataGuardAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.ReadResource(sync)
}

type DataGuardAssociationResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.DataGuardAssociation
}

func (s *DataGuardAssociationResourceCrud) ID() string {
	return s.Res.ID
}

func (s *DataGuardAssociationResourceCrud) Create() (e error) {
	creationType := s.D.Get("creation_type").(string)
	databaseAdminPassword := s.D.Get("database_admin_password").(string)
	databaseID := s.D.Get("database_id").(string)
	protectionMode := s.D.Get("protection_mode").(string)
	transportType := s.D.Get("transport_type").(string)

	opts := &baremetal.CreateDataGuardAssociationOptions{}

	s.Res, e = s.Client.CreateDataGuardAssociation(creationType, databaseAdminPassword, databaseID, protectionMode, transportType, opts)
	return
}

func (s *DataGuardAssociationResourceCrud) Get() (e error) {
    res, e := s.Client.GetDataGuardAssociation(
		s.D.Get("id").(string),
		s.D.Get("database_id").(string),
		nil,
	)
    if e == nil {
        s.Res = res
    }
    return
}

func (s *DataGuardAssociationResourceCrud) SetData() {
	s.D.Set("apply_lag", s.Res.ApplyLag)
	s.D.Set("apply_rate", s.Res.ApplyRate)
	s.D.Set("database_id", s.Res.DatabaseID)
	s.D.Set("id", s.Res.ID)
	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)
	s.D.Set("lifecycle_state", s.Res.LifecycleState)
	s.D.Set("peer_data_guard_association_id", s.Res.PeerDataGuardAssociationID)
	s.D.Set("peer_database_id", s.Res.PeerDatabaseID)
	s.D.Set("peer_db_home_id", s.Res.PeerDbHomeID)
	s.D.Set("peer_db_system_id", s.Res.PeerDbSystemID)
	s.D.Set("peer_role", s.Res.PeerRole)
	s.D.Set("protection_mode", s.Res.ProtectionMode)
	s.D.Set("role", s.Res.Role)
	s.D.Set("time_created", s.Res.TimeCreated.String())
	s.D.Set("transport_type", s.Res.TransportType)
}
