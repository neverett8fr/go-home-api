package home

import "time"

func AfterMidday() bool {
	return time.Now().Hour() >= 12
}

func BeforeMidday() bool {
	return time.Now().Hour() < 12
}

func Always() bool {
	return true
}

func GetCondition(name string) func() bool {
	conditions := map[string]func() bool{
		"after-midday":  AfterMidday,
		"before-midday": BeforeMidday,
		"always":        Always,
	}

	return conditions[name]
}
