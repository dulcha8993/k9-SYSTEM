package utils 

import ( 
 "reflect"
 dto "k9-system/dto/activity_log" 
) 

func GetChanges(oldObj interface{}, newObj interface{}) []dto.FieldChange { 
	var changes []dto.FieldChange
	oldValue := reflect.ValueOf(oldObj) 
	newValue := reflect.ValueOf(newObj) 
	
	// Dereference pointers 
	
	if oldValue.Kind() == reflect.Ptr { 
		oldValue = oldValue.Elem() 
	} 
	if newValue.Kind() == reflect.Ptr { 
		newValue = newValue.Elem() 
	} 
	oldType := oldValue.Type() 
	
	for i := 0; i < oldValue.NumField(); i++ { 
		field := oldType.Field(i) 
		oldField := oldValue.Field(i).Interface() 
		newField := newValue.Field(i).Interface() 
		if !reflect.DeepEqual(oldField, newField) { 
			changes = append(changes, dto.FieldChange{ Field: field.Name, Old: oldField, New: newField, }) 
			} 
		} 
	return changes 
}