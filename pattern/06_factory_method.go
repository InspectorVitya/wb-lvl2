package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

	Паттерн опредляет общий интерфейс создания объектов в суперклассе,
	позволяя классам потомкам переопределять создание, изменяя тип
	создаваемых объектов.
	Так мы создаем объекты не напрямую, а через фабричные методы, отделяя
	создание объектов от их использования. Остальной код будет работать
	с базовым интерфейсом, а мы сможем расширять количество объектов без
	изменения основного кода (мы не привязываемся к конкретным классам объектов).
	Плюсы:
	- Избавляет от привязки к использованию конкретного объекта,
	набор используемых объектов можно расширить, не изменяя остального кода.
	- Выделяет производство объектов в одно место, упрощяя навигацию
	по коду и его поддержку.
	Минусы:
	- Требуется реализация дополнительных объектов создания, что усложняет
	код
	- Может привести к созданию больших параллельных иерархий объектов,
	так как для создания каждого типа объекта нужен свой создатель.
	В Go реализовать классический вариант фабричного метода невозможно
	из-за отсутствия ООП в общем понимании: не создать базовый класс
	с абстрактным методом создания, который будет переопределяться в потомках.
	Но можно реализовать паттерн простая фабрика - класс с условным
	оператором, выбирающий какой объект требуется создать.
В примере реализации показано, как обеспечить хранилище данных с различными бэкэндами, такими как хранилище в памяти, дисковое хранилище.
*/

type Store interface {
	Save(string) error
}

// Различные реализации
type StorageType int

const (
	mongoStorage StorageType = 1 << iota
	memoryStorage
)

func NewStore(t StorageType) Store {
	switch t {
	case memoryStorage:
		return newMemoryStorage()
	case mongoStorage:
		return newMongoStorage()
	default:
		fmt.Println("unknown Storage type")
		return nil
	}
}

type MemoryStorage struct{}

func (ms *MemoryStorage) Save(s string) error {
	// сохранение записи в ОЗУ
	fmt.Printf("Запись '%s' успешно сохранена в Монго\n", s)
	return nil
}

func newMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

type MongoStorage struct{}

func (ms *MongoStorage) Save(s string) error {
	// сохранение записи в БД
	fmt.Printf("Запись '%s' успешно сохранена в Монго\n", s)
	return nil
}

func newMongoStorage() *MongoStorage {
	return &MongoStorage{}
}

// Использование
// С фабричным методом пользователь может определить тип хранилища, который он хочет.
func FactoryMethodPatternStart() {
	memStorage := NewStore(memoryStorage)
	memStorage.Save("memory")

	monStorage := NewStore(mongoStorage)
	monStorage.Save("Mongo")
}
