// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"
	"github.com/oracle/terraform-provider-oci/options"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DataGuardAssociationDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDataGuardAssociations,
		Schema: map[string]*schema.Schema{
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"data_guard_associations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DataGuardAssociationResource(),
			},
		},
	}
}

func readDataGuardAssociations(d *schema.ResourceData, m interface{}) (e error) {
	sync := &DataGuardAssociationDatasourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.ReadResource(sync)
}

type DataGuardAssociationDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListDataGuardAssociations
}

func (s *DataGuardAssociationDatasourceCrud) Get() (e error) {
	databaseID := s.D.Get("database_id").(string)

	opts := &baremetal.ListDataGuardAssociationsOptions{}
	options.SetListOptions(s.D, &opts.ListOptions)

	s.Res = &baremetal.ListDataGuardAssociations{DataGuardAssociations: []baremetal.DataGuardAssociation{}}

	for {
		var list *baremetal.ListDataGuardAssociations
		if list, e = s.Client.ListDataGuardAssociations(databaseID, opts); e != nil {
			break
		}

		s.Res.DataGuardAssociations = append(s.Res.DataGuardAssociations, list.DataGuardAssociations...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNextPage {
			break
		}
	}
	return
}

func (s *DataGuardAssociationDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, r := range s.Res.DataGuardAssociations {
			res := map[string]interface{}{
				"apply_lag":                      r.ApplyLag,
				"apply_rate":                     r.ApplyRate,
				"database_id":                    r.DatabaseID,
				"id":                             r.ID,
				"lifecycle_details":              r.LifecycleDetails,
				"lifecycle_state":                r.LifecycleState,
				"peer_data_guard_association_id": r.PeerDataGuardAssociationID,
				"peer_database_id":               r.PeerDatabaseID,
				"peer_db_home_id":                r.PeerDbHomeID,
				"peer_db_system_id":              r.PeerDbSystemID,
				"peer_role":                      r.PeerRole,
				"protection_mode":                r.ProtectionMode,
				"role":                           r.Role,
				"time_created":                   r.TimeCreated.String(),
				"transport_type":                 r.TransportType,
			}
			resources = append(resources, res)
		}
		s.D.Set("data_guard_associations", resources)
	}
	return
}
