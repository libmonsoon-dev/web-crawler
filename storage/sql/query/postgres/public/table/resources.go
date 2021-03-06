//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Resources = newResourcesTable("public", "resources", "")

type resourcesTable struct {
	postgres.Table

	//Columns
	ID        postgres.ColumnInteger
	WebsiteID postgres.ColumnInteger
	Path      postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ResourcesTable struct {
	resourcesTable

	EXCLUDED resourcesTable
}

// AS creates new ResourcesTable with assigned alias
func (a ResourcesTable) AS(alias string) *ResourcesTable {
	return newResourcesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ResourcesTable with assigned schema name
func (a ResourcesTable) FromSchema(schemaName string) *ResourcesTable {
	return newResourcesTable(schemaName, a.TableName(), a.Alias())
}

func newResourcesTable(schemaName, tableName, alias string) *ResourcesTable {
	return &ResourcesTable{
		resourcesTable: newResourcesTableImpl(schemaName, tableName, alias),
		EXCLUDED:       newResourcesTableImpl("", "excluded", ""),
	}
}

func newResourcesTableImpl(schemaName, tableName, alias string) resourcesTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		WebsiteIDColumn = postgres.IntegerColumn("website_id")
		PathColumn      = postgres.StringColumn("path")
		allColumns      = postgres.ColumnList{IDColumn, WebsiteIDColumn, PathColumn}
		mutableColumns  = postgres.ColumnList{WebsiteIDColumn, PathColumn}
	)

	return resourcesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		WebsiteID: WebsiteIDColumn,
		Path:      PathColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
