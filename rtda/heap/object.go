package heap

type Object struct {
	class *Class
	vars  interface{}
	extra interface{}
}

func (object *Object) Extra() interface{} {
	return object.extra
}

func (object *Object) Class() *Class {
	return object.class
}

func (object *Object) Vars() Slots {
	return object.vars.(Slots)
}
func (object *Object) IsInstanceOf(other *Class) bool {
	clz := object.class
	return other.isAssignableFrom(clz)
}

func (class *Class) NewObject() *Object {
	return newObject(class)
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		vars:  newSlots(class.instanceSlotCount),
	}
}

func (object *Object) SetRefVar(fieldName string, fieldDesc string, value *Object) {
	for _, field := range object.Class().fields {
		if !field.isStatic() && field.name == fieldName && field.descriptor == fieldDesc {
			object.Vars()[field.slotId].ref = value
			return
		}
	}
}

func (object *Object) GetRefVar(fieldName string, fieldDesc string) *Object {
	for _, field := range object.Class().fields {
		if !field.isStatic() && field.name == fieldName && field.descriptor == fieldDesc {
			return object.Vars()[field.slotId].ref
		}
	}
	return nil
}
