package position_role

import "errors"

type PositionRole struct {
	value int
	name  string
}

var (
	Employee = PositionRole{value: 0, name: "EMPLOYEE"}
	Manager  = PositionRole{value: 1, name: "MANAGER"}
	Director = PositionRole{value: 2, name: "DIRECTOR"}

	roles = []PositionRole{Employee, Manager, Director}
)

func (r PositionRole) GetWeight() int {
	switch r {
	case Manager:
		return 100
	case Director:
		return 1000
	default:
		return 0
	}
}

func FromValue(value int) (PositionRole, error) {
	for _, role := range roles {
		if role.value == value {
			return role, nil
		}
	}
	return PositionRole{}, errors.New("invalid role value")
}

func (r PositionRole) String() string {
	return r.name
}

func (r PositionRole) Value() int {
	return r.value
}

func (r PositionRole) Name() string {
	return r.name
}
