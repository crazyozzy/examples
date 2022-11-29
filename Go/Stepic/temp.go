package main

type Rider struct {
	On bool
	Ammo int
	Power int
}

func (r *Rider) Shoot () bool {
	if r.On == false {
		return false
	}

	if r.Ammo > 0 {
		r.Ammo--
		return true
	} else {
		return false
	}
}

func (r *Rider) RideBike () bool {
	if r.On == false {
		return false
	}

	if r.Power > 0 {
		r.Power--
		return true
	} else {
		return false
	}
}

func main() {

	testStruct := &Rider{true, 5, 3}
	/*
	 * Экземпляр созданной вами структуры необходимо передать в качестве
	 * аргумента функции testStruct, которая выполнит проверку соблюдения
	 * всех условий задания/
	 */

// }
}