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

var Contents = newContentsTable("public", "contents", "")

type contentsTable struct {
	postgres.Table

	//Columns
	ID        postgres.ColumnInteger
	Content   postgres.ColumnString
	Type      postgres.ColumnString
	Processed postgres.ColumnBool

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ContentsTable struct {
	contentsTable

	EXCLUDED contentsTable
}

// AS creates new ContentsTable with assigned alias
func (a ContentsTable) AS(alias string) *ContentsTable {
	return newContentsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ContentsTable with assigned schema name
func (a ContentsTable) FromSchema(schemaName string) *ContentsTable {
	return newContentsTable(schemaName, a.TableName(), a.Alias())
}

func newContentsTable(schemaName, tableName, alias string) *ContentsTable {
	return &ContentsTable{
		contentsTable: newContentsTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newContentsTableImpl("", "excluded", ""),
	}
}

func newContentsTableImpl(schemaName, tableName, alias string) contentsTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		ContentColumn   = postgres.StringColumn("content")
		TypeColumn      = postgres.StringColumn("type")
		ProcessedColumn = postgres.BoolColumn("processed")
		allColumns      = postgres.ColumnList{IDColumn, ContentColumn, TypeColumn, ProcessedColumn}
		mutableColumns  = postgres.ColumnList{ContentColumn, TypeColumn, ProcessedColumn}
	)

	return contentsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		Content:   ContentColumn,
		Type:      TypeColumn,
		Processed: ProcessedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}