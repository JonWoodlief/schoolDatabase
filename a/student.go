package databaseTypes

type Student struct {
    id int
    classes []Class
}

func (s Student) GetID() int {
    return s.id
}
func (s Student) SetID(newID int) int {
    s.id = newID
    if s.id == newID {
        return 0
    }
    return 1 //failed to set id
}
func (s Student) AddClass(newClass Class) int {
    for _, class := range s.classes {
        if class.id == newClass.id {
            return 1 //student already has class
        }
    }

    s.classes = append(s.classes, newClass)
    return 0
}
func (s Student) DelClass(remClass Class) int {
    for i, class := range s.classes {
        if class.id == remClass.id {
            s.classes = append(s.classes[:i], s.classes[i+1:]...)
            return 0
        }
    }

    return 1 //class not found
}
