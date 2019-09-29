package databaseTypes

type Teacher struct {
	id      int
	classes []Class
}

func (t Teacher) GetID() int {
	return t.id
}
func (t Teacher) SetID(newID int) int {
	t.id = newID
	if t.id == newID {
		return 0
	}
	return 1 //failed to set id
}
func (t Teacher) getClasses() []Class {
    return t.classes
}
func (t Teacher) AddClass(newClass Class) int {
	for _, class := range t.classes {
		if class.id == newClass.id {
			return 1 //student already has class
		}
	}

	t.classes = append(t.classes, newClass)
	return 0
}
func (t Teacher) DelClass(remClass Class) int {
	for i, class := range t.classes {
		if class.id == remClass.id {
			t.classes = append(t.classes[:i], t.classes[i+1:]...)
			return 0
		}
	}

	return 1 //class not found
}
