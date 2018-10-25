package heap

type Object struct {
	class  *Class
	fields Slots
}

func (o *Object) Class() *Class {
	return o.class
}

func (o *Object) Fields() Slots {
	return o.fields
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
