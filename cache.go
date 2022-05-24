package cache

import "time"

type Cache struct {
	grades map[string]Account
}
type Account struct {
	value    string
	deadline time.Time
}

func NewCache() Cache {
	newgrades := Cache{grades: make(map[string]Account)}
	return newgrades
}

func (from Cache) Put(key, value string) {
	from.grades[key] = Account{value: value}

}
func (from Cache) Keys() []string {
	var key []string
	for keys, k := range from.grades {
		if k.deadline.IsZero() || time.Now().Before(k.deadline) {
			key = append(key, keys)
		}
	}
	return key

}
func (from Cache) Get(key string) (string, bool) {
	if k, ok := from.grades[key]; ok {
		if k.deadline.IsZero() || time.Now().Before(k.deadline) {
			return k.value, true
		}
	}
	return "", false
}
func (from Cache) PutTill(key, value string, deadline time.Time) {
  from.grades[key] = Account{value: value,deadline: deadline }
}
