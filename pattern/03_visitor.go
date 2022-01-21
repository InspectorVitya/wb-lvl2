package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

	Visitor -  это поведенческий паттерн, который позволяет добавлять в программу новые операции, не изменяя классы объектов,
	над которыми эти операции могут выполняться.
	Используется, когда:
	- Когда над объектами сложной структуры объектов надо выполнять некоторые не связанные между собой операции,
	но вы не хотите «засорять» классы такими операциями.
	- Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии
	Преимущества:
	Упрощает добавление операций, работающих со сложными структурами объектов.
	Объединяет родственные операции в одном классе.
	Посетитель может накапливать состояние при обходе структуры элементов
	Недостатки:
	усложняет расширение иерархии классов, поскольку новые классы обычно требуют добавления нового метода visit для каждого посетителя
*/

// Shape Базовый интерфейс Фигура - корень иерархии, требующий метод принятия визитора
type Shape interface {
	Accept(Visitor)
}

// Square Объект квадрат с полем - стороной
type Square struct {
	side int
}

// Accept Реализация метода принятия визитора
func (s *Square) Accept(v Visitor) {
	v.VisitSquare(*s)
}

// Circle Объект круг с полем - радиусом
type Circle struct {
	radius int
}

// Accept Реализация метода принятия визитора
func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(*c)
}

// Rectangle Объект прямоугольник с полями - сторонами
type Rectangle struct {
	a, b int
}

// Accept Реализация метода принятия визитора
func (r *Rectangle) Accept(v Visitor) {
	v.VisitRectangle(*r)
}

// Visitor Интерфейс визитор, требующий реализации методов посещения для каждого
//объекта из иерархии
type Visitor interface {
	VisitSquare(Square)
	VisitCircle(Circle)
	VisitRectangle(Rectangle)
}

// XMLPrinter Реализация интерфейса визитора для печати фигур в XML формате
type XMLPrinter struct{}

func (x *XMLPrinter) VisitSquare(Square) {
	fmt.Printf("<xml>Square</xml>\n")
}

func (x *XMLPrinter) VisitCircle(Circle) {
	fmt.Printf("<xml>Circle</xml>\n")
}

func (x *XMLPrinter) VisitRectangle(Rectangle) {
	fmt.Printf("<xml>Rectangle</xml>\n")
}

// JSONPrinter Реализация интерфейса визитора для печати фигур в JSON формате
type JSONPrinter struct{}

func (j *JSONPrinter) VisitSquare(Square) {
	fmt.Println(`{"Square"}`)
}

func (j *JSONPrinter) VisitCircle(Circle) {
	fmt.Println(`{"Circle"}`)
}

func (j *JSONPrinter) VisitRectangle(Rectangle) {
	fmt.Println(`{"Rectangle"}`)
}

func main() {
	//создаем объекты иерархии
	r := Rectangle{}
	c := Circle{}
	s := Square{}

	xml := XMLPrinter{} //создаем объект реализующий визитор
	//вызываем метод приема визитора для каждого объета иерархии
	r.Accept(&xml)
	c.Accept(&xml)
	s.Accept(&xml)

	json := JSONPrinter{} //создаем объект реализующий визитор
	//вызываем метод приема визитора для каждого объета иерархии
	r.Accept(&json)
	c.Accept(&json)
	s.Accept(&json)
}
