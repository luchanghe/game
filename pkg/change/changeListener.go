package change

import "reflect"

type ChangeListener interface {
	Update()
	GetChanges()
}

type Watcher struct {
	originalValue reflect.Value
	currentValue  reflect.Value
	changes       []FieldChange
}

type FieldChange struct {
	FieldName string
	OldValue  interface{}
	NewValue  interface{}
}

func NewWatcher(v interface{}) *Watcher {
	originalValue := reflect.ValueOf(v).Elem()
	currentValue := reflect.New(originalValue.Type()).Elem()
	currentValue.Set(originalValue)

	return &Watcher{
		originalValue: originalValue,
		currentValue:  currentValue,
		changes:       []FieldChange{},
	}
}
func (w *Watcher) Update() {
	for i := 0; i < w.originalValue.NumField(); i++ {
		field := w.originalValue.Type().Field(i)
		originalValue := w.originalValue.Field(i).Interface()
		currentValue := w.currentValue.Field(i).Interface()

		if !reflect.DeepEqual(originalValue, currentValue) {
			w.changes = append(w.changes, FieldChange{
				FieldName: field.Name,
				OldValue:  originalValue,
				NewValue:  currentValue,
			})
			// 更新 originalValue 为 currentValue
			w.originalValue.Field(i).Set(w.currentValue.Field(i))
		}
	}
}

func (w *Watcher) GetChanges() []FieldChange {
	return w.changes
}
