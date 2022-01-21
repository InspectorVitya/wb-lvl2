package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

	Фабричный метод — это порождающий паттерн проектирования, который решает проблему создания различных продуктов, без указания конкретных классов продуктов.
	В Golang можно реализовать лишь базовую версию паттерна - простая фабрика.

	Фабричный метод —это класс, в котором есть один метод с большим условным оператором, выбирающим создаваемый продукт.

	Применимость:
	- Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код.
	Фабричный метод отделяет код производства продуктов от остального кода, который эти продукты использует.

	- Когда вы хотите дать возможность пользователям расширять части вашего фреймворка или библиотеки.

	Преимущества:
	- Избавляет класс от привязки к конкретным классам продуктов.
	- Выделяет код производства продуктов в одно место, упрощая поддержку кода.
	- Упрощает добавление новых продуктов в программу.
	- Реализует принцип открытости/закрытости.

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
