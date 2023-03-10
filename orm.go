package goleap

import (
	"context"
	"fmt"
	"github.com/lab210-dev/dbkit/schema"
	"github.com/lab210-dev/dbkit/specs"
)

type Orm[T specs.Model] interface {
	Get(primaryKey any) (T, error)
	Delete(primaryKey any) error

	Create() (err error)
	Update() error

	Find() error
	FindAll() error

	Fields(field ...string) Orm[T]
	Where(condition ...any) Orm[T]
	Limit(limit int) Orm[T]
	Offset(offset int) Orm[T]
	OrderBy(fields ...string) Orm[T]

	Count() (total int64, err error)

	Payload() specs.Payload
}

type orm[T specs.Model] struct {
	context.Context
	specs.Connector

	model  *T
	schema specs.Schema
	fields []string

	focusedSchemaFields      []specs.SchemaField
	focusedDriverFields      []specs.DriverField
	requiredJoins            []specs.DriverJoin
	focusedSchemaValueFields []any

	payload specs.Payload
}

func (o *orm[T]) buildFieldsFromSchema() (err error) {
	for _, fieldName := range o.fields {
		field := o.schema.GetFieldByName(fieldName)

		if field == nil {
			// TODO (JG) : Use a real name of schema.
			err = &FieldNotFoundError{message: fmt.Sprintf("field `%s` not found in schema `%s`", fieldName, o.schema.TableName()), field: fieldName}
			return
		}

		o.focusedSchemaFields = append(o.focusedSchemaFields, field)
	}

	return
}

func (o *orm[T]) getFieldsTypeSchemaToDriver() (fields []any) {

	if len(o.focusedDriverFields) > 0 {
		return o.focusedSchemaValueFields
	}

	for _, field := range o.focusedSchemaFields {
		o.focusedSchemaValueFields = append(o.focusedSchemaValueFields, field.Copy())
	}

	return o.focusedSchemaValueFields
}

func (o *orm[T]) getFocusedDriverFields() []specs.DriverField {

	if len(o.focusedDriverFields) > 0 {
		return o.focusedDriverFields
	}

	for _, field := range o.focusedSchemaFields {
		o.focusedDriverFields = append(o.focusedDriverFields, field.Field())
	}

	return o.focusedDriverFields
}

func (o *orm[T]) getRequiredJoins() []specs.DriverJoin {

	if len(o.requiredJoins) > 0 {
		return o.requiredJoins
	}

	for _, field := range o.focusedSchemaFields {
		o.requiredJoins = append(o.requiredJoins, field.Join()...)
	}

	return o.requiredJoins
}

func (o *orm[T]) Payload() specs.Payload {
	if o.payload != nil {
		return o.payload
	}
	o.payload = newPayload[T](o)
	return o.payload
}

func (o *orm[T]) Get(primaryKeyValue any) (result T, err error) {
	/*	err = o.buildFieldsFromSchema()
		if err != nil {
			return
		}*/

	/*	primaryKeyField := o.schema.GetPrimaryKeyField()
		if primaryKeyField == nil {
			// TODO (JG) : Factory for error.
			err = &FieldNotFoundError{message: fmt.Sprintf("primary key not found in schema `%s`", o.schema.TableName())}
			return
		}*/

	// var where []driver.DriverWhere
	//	where = append(where, driver.NewWhere().SetFrom(primaryKeyField.Field()).SetOperator(driver.EqualOperator).SetTo(primaryKeyValue))
	err = o.buildFieldsFromSchema()
	if err != nil {
		return
	}

	pp := o.Payload()
	err = o.Connector.Select(o.Context, pp)
	if err != nil {
		return
	}

	result = pp.(*payload[T]).result[0]

	return
}

func (o *orm[T]) Delete(primaryKeyValue any) error {
	//TODO implement me
	panic("implement me")
}

func (o *orm[T]) Create() (err error) {
	//TODO implement me
	panic("implement me")
}

func (o *orm[T]) Update() error {
	//TODO implement me
	panic("implement me")
}

func (o *orm[T]) Find() error {
	//TODO implement me
	panic("implement me")
}

func (o *orm[T]) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (o *orm[T]) Fields(field ...string) Orm[T] {
	o.fields = field
	return o
}

func (o *orm[T]) Where(condition ...any) Orm[T] {
	//TODO implement me
	panic("implement me")
}

func (o *orm[T]) Limit(limit int) Orm[T] {
	//TODO implement me
	panic("implement me")
}

func (o *orm[T]) Offset(offset int) Orm[T] {
	//TODO implement me
	panic("implement me")
}

func (o *orm[T]) OrderBy(fields ...string) Orm[T] {
	//TODO implement me
	panic("implement me")
}

func (o *orm[T]) Count() (total int64, err error) {
	//TODO implement me
	panic("implement me")
}

func Use[T specs.Model](ctx context.Context, connector specs.Connector) Orm[T] {
	var model T

	return &orm[T]{
		Context:   ctx,
		Connector: connector,

		model:  &model,
		schema: schema.New(model).Parse(),
	}
}
