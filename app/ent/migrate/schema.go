// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppQueriesColumns holds the columns for the "app_queries" table.
	AppQueriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "ip_queries", Type: field.TypeUUID, Nullable: true},
	}
	// AppQueriesTable holds the schema information for the "app_queries" table.
	AppQueriesTable = &schema.Table{
		Name:       "app_queries",
		Columns:    AppQueriesColumns,
		PrimaryKey: []*schema.Column{AppQueriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_queries_ips_queries",
				Columns:    []*schema.Column{AppQueriesColumns[3]},
				RefColumns: []*schema.Column{IpsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "appquery_updated_at_created_at",
				Unique:  false,
				Columns: []*schema.Column{AppQueriesColumns[2], AppQueriesColumns[1]},
			},
		},
	}
	// AppResponsesColumns holds the columns for the "app_responses" table.
	AppResponsesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "responsecode", Type: field.TypeString},
		{Name: "codedescription", Type: field.TypeString},
		{Name: "app_query_responses", Type: field.TypeUUID, Nullable: true},
	}
	// AppResponsesTable holds the schema information for the "app_responses" table.
	AppResponsesTable = &schema.Table{
		Name:       "app_responses",
		Columns:    AppResponsesColumns,
		PrimaryKey: []*schema.Column{AppResponsesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_responses_app_queries_responses",
				Columns:    []*schema.Column{AppResponsesColumns[5]},
				RefColumns: []*schema.Column{AppQueriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "appresponse_updated_at_created_at",
				Unique:  false,
				Columns: []*schema.Column{AppResponsesColumns[2], AppResponsesColumns[1]},
			},
		},
	}
	// IpsColumns holds the columns for the "ips" table.
	IpsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "ip_address", Type: field.TypeString},
	}
	// IpsTable holds the schema information for the "ips" table.
	IpsTable = &schema.Table{
		Name:       "ips",
		Columns:    IpsColumns,
		PrimaryKey: []*schema.Column{IpsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "ip_updated_at_created_at",
				Unique:  false,
				Columns: []*schema.Column{IpsColumns[2], IpsColumns[1]},
			},
		},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"IPDNSBL"}},
		{Name: "ipaddress", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"WAITING", "IN_PROGRESS", "DONE", "ERROR"}},
		{Name: "error", Type: field.TypeString, Nullable: true},
		{Name: "completed_by", Type: field.TypeTime, Nullable: true},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "task_updated_at_created_at",
				Unique:  false,
				Columns: []*schema.Column{TasksColumns[2], TasksColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeBytes},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_updated_at_created_at",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[2], UsersColumns[1]},
			},
			{
				Name:    "user_username",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppQueriesTable,
		AppResponsesTable,
		IpsTable,
		TasksTable,
		UsersTable,
	}
)

func init() {
	AppQueriesTable.ForeignKeys[0].RefTable = IpsTable
	AppResponsesTable.ForeignKeys[0].RefTable = AppQueriesTable
}
