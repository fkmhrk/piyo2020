package game

type Storage interface {
	Save(key string, data map[string]interface{}) error
	Load(key string) (map[string]interface{}, error)
}
