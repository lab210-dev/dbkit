package schema

import (
	"github.com/lab210-dev/dbkit/specs"
	"reflect"
	"sync"
)

type schema struct {
	specs.Model
	sync.Mutex

	index       int
	counter     int
	parsed      bool
	modelOrigin reflect.Value
	modelType   reflect.Type
	modelValue  reflect.Value

	fields      []specs.SchemaField
	fieldByName map[string]specs.SchemaField

	fromField specs.SchemaField
}

func (schema *schema) ModelOrigin() reflect.Value {
	return schema.modelOrigin
}

func (schema *schema) ModelValue() reflect.Value {
	return schema.modelValue
}

func (schema *schema) Copy() specs.Schema {
	tmp := reflect.New(reflect.TypeOf(schema)).Elem()
	tmp.Set(reflect.ValueOf(schema))

	return tmp.Interface().(specs.Schema)
}

func New(model specs.Model) specs.Schema {
	schema := new(schema)
	schema.Model = model

	schema.modelType = reflect.TypeOf(model)
	schema.modelValue = reflect.ValueOf(model)
	schema.modelOrigin = schema.modelValue

	if schema.modelType.Kind() == reflect.Struct {
		schema.modelValue = reflect.New(schema.modelType).Elem().Addr()
		schema.modelType = schema.modelValue.Type()
	}

	if schema.modelType.Kind() == reflect.Ptr {
		schema.modelType = schema.modelType.Elem()

		if schema.modelValue.IsNil() {
			schema.modelValue = reflect.New(schema.modelType)
			schema.Model = schema.modelValue.Interface().(specs.Model)
		}

		schema.modelValue = schema.modelValue.Elem()
	}

	schema.fieldByName = make(map[string]specs.SchemaField)

	return schema
}

func (schema *schema) Get() specs.Model {
	return schema.modelValue.Addr().Interface().(specs.Model)
}

func (schema *schema) SetIndex(index int) specs.Schema {
	schema.index = index
	schema.counter = index
	return schema
}

func (schema *schema) GetPrimaryKeyField() specs.SchemaField {
	for _, field := range schema.fields {

		if field.Schema() != schema {
			continue
		}

		if field.IsPrimaryKey() {
			return field
		}

	}
	return nil
}

func (schema *schema) Index() int {
	return schema.index
}

func (schema *schema) Counter() int {
	if schema.FromField() != nil {
		return schema.FromField().Schema().Counter()
	}

	schema.counter++
	return schema.counter
}

func (schema *schema) Fields() []specs.SchemaField {
	return schema.fields
}

func (schema *schema) FieldByName() map[string]specs.SchemaField {
	return schema.fieldByName
}

func (schema *schema) AddField(field specs.SchemaField) {
	defer schema.Unlock()
	schema.Lock()

	if field.HasEmbeddedSchema() {
		schema.fields = append(schema.fields, field.EmbeddedSchema().Fields()...)

		for key, value := range field.EmbeddedSchema().FieldByName() {
			schema.fieldByName[key] = value
		}
		return
	}

	schema.fields = append(schema.fields, field)
	schema.fieldByName[field.RecursiveFullName()] = field
}

func (schema *schema) FromField() specs.SchemaField {
	return schema.fromField
}

func (schema *schema) SetFromField(fromField specs.SchemaField) specs.Schema {
	schema.fromField = fromField
	return schema
}

func (schema *schema) GetFieldByName(name string) specs.SchemaField {
	return schema.fieldByName[name]
}

func (schema *schema) Parse() specs.Schema {
	// no need to parse again normally...
	if schema.parsed {
		return schema
	}

	for i := 0; i < schema.modelType.NumField(); i++ {
		fieldStruct := schema.modelType.Field(i)
		if !fieldStruct.IsExported() {
			continue
		}

		field := schema.parseField(i)
		if field == nil {
			continue
		}

		schema.AddField(field)
	}

	// TODO use a setter to set this value
	schema.parsed = true

	return schema
}
