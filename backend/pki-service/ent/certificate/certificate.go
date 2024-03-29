// Code generated by entc, DO NOT EDIT.

package certificate

import (
	"fmt"
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the certificate type in the database.
	Label = "certificate"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldSslId holds the string denoting the sslid field in the database.
	FieldSslId = "ssl_id"
	// FieldSerial holds the string denoting the serial field in the database.
	FieldSerial = "serial"
	// FieldCommonName holds the string denoting the commonname field in the database.
	FieldCommonName = "common_name"
	// FieldNotBefore holds the string denoting the notbefore field in the database.
	FieldNotBefore = "not_before"
	// FieldNotAfter holds the string denoting the notafter field in the database.
	FieldNotAfter = "not_after"
	// FieldIssuedBy holds the string denoting the issuedby field in the database.
	FieldIssuedBy = "issued_by"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeDomains holds the string denoting the domains edge name in mutations.
	EdgeDomains = "domains"
	// Table holds the table name of the certificate in the database.
	Table = "certificates"
	// DomainsTable is the table that holds the domains relation/edge. The primary key declared below.
	DomainsTable = "certificate_domains"
	// DomainsInverseTable is the table name for the Domain entity.
	// It exists in this package in order to avoid circular dependency with the "domain" package.
	DomainsInverseTable = "domains"
)

// Columns holds all SQL columns for certificate fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldSslId,
	FieldSerial,
	FieldCommonName,
	FieldNotBefore,
	FieldNotAfter,
	FieldIssuedBy,
	FieldSource,
	FieldCreated,
	FieldStatus,
}

var (
	// DomainsPrimaryKey and DomainsColumn2 are the table columns denoting the
	// primary key for the domains relation (M2M).
	DomainsPrimaryKey = []string{"certificate_id", "domain_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/hm-edu/pki-service/ent/runtime"
//
var (
	Hooks [1]ent.Hook
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// CommonNameValidator is a validator for the "commonName" field. It is called by the builders before save.
	CommonNameValidator func(string) error
)

// Status defines the type for the "status" enum field.
type Status string

// StatusInvalid is the default value of the Status enum.
const DefaultStatus = StatusInvalid

// Status values.
const (
	StatusInvalid    Status = "Invalid"
	StatusRequested  Status = "Requested"
	StatusApproved   Status = "Approved"
	StatusDeclined   Status = "Declined"
	StatusApplied    Status = "Applied"
	StatusIssued     Status = "Issued"
	StatusRevoked    Status = "Revoked"
	StatusExpired    Status = "Expired"
	StatusReplaced   Status = "Replaced"
	StatusRejected   Status = "Rejected"
	StatusUnmanaged  Status = "Unmanaged"
	StatusSAApproved Status = "SAApproved"
	StatusInit       Status = "Init"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusInvalid, StatusRequested, StatusApproved, StatusDeclined, StatusApplied, StatusIssued, StatusRevoked, StatusExpired, StatusReplaced, StatusRejected, StatusUnmanaged, StatusSAApproved, StatusInit:
		return nil
	default:
		return fmt.Errorf("certificate: invalid enum value for status field: %q", s)
	}
}
