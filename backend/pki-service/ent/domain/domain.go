// Code generated by entc, DO NOT EDIT.

package domain

const (
	// Label holds the string label denoting the domain type in the database.
	Label = "domain"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFqdn holds the string denoting the fqdn field in the database.
	FieldFqdn = "fqdn"
	// EdgeCertificates holds the string denoting the certificates edge name in mutations.
	EdgeCertificates = "certificates"
	// Table holds the table name of the domain in the database.
	Table = "domains"
	// CertificatesTable is the table that holds the certificates relation/edge. The primary key declared below.
	CertificatesTable = "certificate_domains"
	// CertificatesInverseTable is the table name for the Certificate entity.
	// It exists in this package in order to avoid circular dependency with the "certificate" package.
	CertificatesInverseTable = "certificates"
)

// Columns holds all SQL columns for domain fields.
var Columns = []string{
	FieldID,
	FieldFqdn,
}

var (
	// CertificatesPrimaryKey and CertificatesColumn2 are the table columns denoting the
	// primary key for the certificates relation (M2M).
	CertificatesPrimaryKey = []string{"certificate_id", "domain_id"}
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

var (
	// FqdnValidator is a validator for the "fqdn" field. It is called by the builders before save.
	FqdnValidator func(string) error
)
