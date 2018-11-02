package heap

type Object struct {
	class  *Class
	fields interface{}
}

func (object *Object) Class() *Class {
	return object.class
}

func (object *Object) Fields() Slots {
	return object.fields.(Slots)
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
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}
