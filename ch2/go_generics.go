package ch2

import (
	"io"
	"strconv"
)

func Reverse1(s []int64) []int64 {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
	return s
}

//Reverse2 any是Go 1.18新增的预声明标识符，是interface{}的别名。
func Reverse2(s []any) []any {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
	return s
}

type Number interface {
	int | int32 | int64 // |表示取并集,Number这个interface可以作为类型限制，用于限定类型参数必须是int，int32和int64这3种类型。
}

/*
type Foo struct {}

//编译会报错， 类型参数列表不能用于方法，只能用于函数。
func (Foo) bar[T any](t T) {
}
*/

//AnyString AnyString这个interface可以作为类型限制，用于限定类型参数的底层类型必须是string。string本身以及下面的MyString都满足AnyString这个类型限制。
type AnyString interface {
	~string //~T ~ 是Go 1.18新增的符号，~T表示底层类型是T的所有类型。~的英文读作tilde
}
type MyString string

//类型限制customConstraint，用于限定底层类型为int并且实现了String() string方法的所有类型。下面的customInt就满足这个type constraint。
type customConstraint interface {
	~int
	String() string
}

type customInt int

func (i customInt) String() string {
	return strconv.Itoa(int(i))
}

//类型限制有2个作用：
//1.用于约定有效的类型实参，不满足类型限制的类型实参会被编译器报错。
//2.如果类型限制里的所有类型都支持某个操作，那在代码里，对应的类型参数就可以使用这个操作。
//https: //go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#type-inference

//什么请求适合适应泛型
//1.需要使用slice, map, channel类型，但是slice, map, channel里的元素类型可能有多种。
//通用的数据结构，比如链表，二叉树等。下面的代码实现了一个支持任意数据类型的二叉树。

type Tree[T any] struct {
	Cmp  func(T, T) int
	Root *Node[T]
}

type Node[T any] struct {
	Left, Right *Node[T]
	Data        T
}

func (bt *Tree[T]) Find(val T) **Node[T] {
	pl := &bt.Root
	for *pl != nil {
		switch cmp := bt.Cmp(val, (*pl).Data); {
		case cmp < 0:
			pl = &(*pl).Left
		case cmp > 0:
			pl = &(*pl).Right
		default:
			return pl
		}
	}
	return pl
}

func CmpInt(v, data int) int {
	if v > data {
		return 1
	}
	if v < data {
		return -1
	}
	return 0
}

//2. 当一个方法的实现对所有类型都一样。

type SliceFn[T any] struct {
	//comparable
	s   []T
	cmp func(T, T) bool
}

func (s SliceFn[T]) Len() int { return len(s.s) }

func (s SliceFn[T]) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}
func (s SliceFn[T]) Less(i, j int) bool {
	return s.cmp(s.s[i], s.s[j])
}

//什么时候不要使用泛型
//1.只是单纯调用实参的方法时，不要用泛型。
// good
func foo1(w io.Writer) {
	b := getBytes()
	_, _ = w.Write(b) //单纯是调用io.Writer的Write方法，把内容写到指定地方。使用interface作为参数更合适，可读性更强。
}

func getBytes() []byte {
	return []byte("test")
}

// bad
//func foo2[T io.Writer](w T) {
//	b := getBytes()
//	_, _ = w.Write(b)
//}

//2.当函数或者方法或者具体的实现逻辑，对于不同类型不一样时，不要用泛型。比如encoding/json这个包使用了reflect，如果用泛型反而不合适。
