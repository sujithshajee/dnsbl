// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/sujithshajee/dnsbl/app/ent/ip"
)

// IP is the model entity for the IP schema.
type IP struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// IPAddress holds the value of the "ip_address" field.
	IPAddress string `json:"ip_address,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IPQuery when eager-loading is set.
	Edges IPEdges `json:"edges"`
}

// IPEdges holds the relations/edges for other nodes in the graph.
type IPEdges struct {
	// Queries holds the value of the queries edge.
	Queries []*AppQuery `json:"queries,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// QueriesOrErr returns the Queries value or an error if the edge
// was not loaded in eager-loading.
func (e IPEdges) QueriesOrErr() ([]*AppQuery, error) {
	if e.loadedTypes[0] {
		return e.Queries, nil
	}
	return nil, &NotLoadedError{edge: "queries"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IP) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case ip.FieldIPAddress:
			values[i] = new(sql.NullString)
		case ip.FieldCreatedAt, ip.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case ip.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type IP", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IP fields.
func (i *IP) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case ip.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case ip.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case ip.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case ip.FieldIPAddress:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip_address", values[j])
			} else if value.Valid {
				i.IPAddress = value.String
			}
		}
	}
	return nil
}

// QueryQueries queries the "queries" edge of the IP entity.
func (i *IP) QueryQueries() *AppQueryQuery {
	return (&IPClient{config: i.config}).QueryQueries(i)
}

// Update returns a builder for updating this IP.
// Note that you need to call IP.Unwrap() before calling this method if this IP
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *IP) Update() *IPUpdateOne {
	return (&IPClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the IP entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *IP) Unwrap() *IP {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: IP is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *IP) String() string {
	var builder strings.Builder
	builder.WriteString("IP(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ip_address=")
	builder.WriteString(i.IPAddress)
	builder.WriteByte(')')
	return builder.String()
}

// IPs is a parsable slice of IP.
type IPs []*IP

func (i IPs) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}