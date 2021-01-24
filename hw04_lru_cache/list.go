package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(i *listItem)
	MoveToFront(i *listItem)
}

type listItem struct {
	Value interface{}
	Next  *listItem
	Prev  *listItem
}

type list struct {
	head, tail *listItem
	length     int
}

func (l list) Len() int { return l.length }

func (l list) Front() *listItem { return l.head }

func (l list) Back() *listItem { return l.tail }

func (l *list) initList(item *listItem) {
	l.head = item
	l.tail = item
	l.length++
}

func (l *list) PushFront(v interface{}) *listItem {
	item := &listItem{Value: v}
	if l.length == 0 {
		l.initList(item)
		return item
	}
	item.Next = l.head
	item.Next.Prev = item
	l.head = item
	l.length++
	return item
}

func (l *list) PushBack(v interface{}) *listItem {
	item := &listItem{Value: v}
	if l.length == 0 {
		l.initList(item)
		return item
	}
	item.Prev = l.tail
	item.Prev.Next = item
	l.tail = item
	l.length++
	return item
}

func (l *list) Remove(i *listItem) {
	// is this the last value?
	if i.Next == nil {
		l.tail = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	// is this the first value?
	if i.Prev == nil {
		l.head = i.Next
	} else {
		i.Prev.Next = i.Next
	}
	l.length--
}

func (l *list) MoveToFront(i *listItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return &list{}
}
