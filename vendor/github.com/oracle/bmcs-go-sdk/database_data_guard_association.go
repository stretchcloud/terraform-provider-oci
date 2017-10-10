// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

type DataGuardAssociation struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	ApplyLag                   string `json:"applyLag"`
	ApplyRate                  string `json:"applyRate"`
	DatabaseID                 string `json:"databaseId"`
	ID                         string `json:"id"`
	LifecycleDetails           string `json:"lifecycleDetails"`
	LifecycleState             string `json:"lifecycleState"`
	PeerDataGuardAssociationID string `json:"peerDataGuardAssociationId"`
	PeerDatabaseID             string `json:"peerDatabaseId"`
	PeerDbHomeID               string `json:"peerDbHomeId"`
	PeerDbSystemID             string `json:"peerDbSystemId"`
	PeerRole                   string `json:"peerRole"`
	ProtectionMode             string `json:"protectionMode"`
	Role                       string `json:"role"`
	TimeCreated                Time   `json:"timeCreated"`
	TransportType              string `json:"transportType"`
}

type CreateDataGuardAssociationOptions struct {
	CreateOptions
}

type GetDataGuardAssociationOptions struct {
}

type ListDataGuardAssociationsOptions struct {
	ListOptions
}

type ListDataGuardAssociations struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	DataGuardAssociations []DataGuardAssociation
}

func (l *ListDataGuardAssociations) GetList() interface{} {
	return &l.DataGuardAssociations
}

func (c *Client) CreateDataGuardAssociation(creationType string, databaseAdminPassword string, databaseID string, protectionMode string, transportType string, opts *CreateDataGuardAssociationOptions) (res *DataGuardAssociation, e error) {
	required := struct {
		CreationType          string `header:"-" json:"creationType" url:"-"`
		DatabaseAdminPassword string `header:"-" json:"databaseAdminPassword" url:"-"`
		ProtectionMode        string `header:"-" json:"protectionMode" url:"-"`
		TransportType         string `header:"-" json:"transportType" url:"-"`
	}{}
	required.CreationType = creationType
	required.DatabaseAdminPassword = databaseAdminPassword
	required.ProtectionMode = protectionMode
	required.TransportType = transportType

	details := &requestDetails{
		name: resourceDatabases,
		ids: urlParts{
			databaseID,
			resourceDataGuardAssociations,
		},
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.databaseApi.postRequest(details); e != nil {
		return
	}

	res = &DataGuardAssociation{}
	e = resp.unmarshal(res)
	return
}

func (c *Client) GetDataGuardAssociation(dataGuardAssociationID string, databaseID string, opts *GetDataGuardAssociationOptions) (res *DataGuardAssociation, e error) {
	required := struct {
	}{}

	details := &requestDetails{
		name: resourceDatabases,
		ids: urlParts{
			databaseID,
			resourceDataGuardAssociations,
			dataGuardAssociationID,
		},
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	res = &DataGuardAssociation{}
	e = resp.unmarshal(res)
	return
}

func (c *Client) ListDataGuardAssociations(databaseID string, opts *ListDataGuardAssociationsOptions) (res *ListDataGuardAssociations, e error) {
	required := struct {
	}{}

	details := &requestDetails{
		name: resourceDatabases,
		ids: urlParts{
			databaseID,
			resourceDataGuardAssociations,
		},
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	res = &ListDataGuardAssociations{}
	e = resp.unmarshal(res)
	return
}
