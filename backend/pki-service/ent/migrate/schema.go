// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CertificatesColumns holds the columns for the "certificates" table.
	CertificatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "ssl_id", Type: field.TypeInt, Nullable: true},
		{Name: "serial", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "common_name", Type: field.TypeString},
		{Name: "not_before", Type: field.TypeTime, Nullable: true},
		{Name: "not_after", Type: field.TypeTime, Nullable: true},
		{Name: "issued_by", Type: field.TypeString, Nullable: true},
		{Name: "source", Type: field.TypeString, Nullable: true},
		{Name: "created", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"Invalid", "Requested", "Approved", "Declined", "Applied", "Issued", "Revoked", "Expired", "Replaced", "Rejected", "Unmanaged", "SAApproved", "Init"}, Default: "Invalid"},
	}
	// CertificatesTable holds the schema information for the "certificates" table.
	CertificatesTable = &schema.Table{
		Name:       "certificates",
		Columns:    CertificatesColumns,
		PrimaryKey: []*schema.Column{CertificatesColumns[0]},
	}
	// DomainsColumns holds the columns for the "domains" table.
	DomainsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "fqdn", Type: field.TypeString, Unique: true},
	}
	// DomainsTable holds the schema information for the "domains" table.
	DomainsTable = &schema.Table{
		Name:       "domains",
		Columns:    DomainsColumns,
		PrimaryKey: []*schema.Column{DomainsColumns[0]},
	}
	// CertificateDomainsColumns holds the columns for the "certificate_domains" table.
	CertificateDomainsColumns = []*schema.Column{
		{Name: "certificate_id", Type: field.TypeInt},
		{Name: "domain_id", Type: field.TypeInt},
	}
	// CertificateDomainsTable holds the schema information for the "certificate_domains" table.
	CertificateDomainsTable = &schema.Table{
		Name:       "certificate_domains",
		Columns:    CertificateDomainsColumns,
		PrimaryKey: []*schema.Column{CertificateDomainsColumns[0], CertificateDomainsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "certificate_domains_certificate_id",
				Columns:    []*schema.Column{CertificateDomainsColumns[0]},
				RefColumns: []*schema.Column{CertificatesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "certificate_domains_domain_id",
				Columns:    []*schema.Column{CertificateDomainsColumns[1]},
				RefColumns: []*schema.Column{DomainsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CertificatesTable,
		DomainsTable,
		CertificateDomainsTable,
	}
)

func init() {
	CertificateDomainsTable.ForeignKeys[0].RefTable = CertificatesTable
	CertificateDomainsTable.ForeignKeys[1].RefTable = DomainsTable
}
