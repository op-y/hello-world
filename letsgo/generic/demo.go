package generic

import (
	"fmt"
	"math"
)

func DefaultKeyWordParams[D any](defVal D, params ...D) D {
	if len(params) == 0 {
		return defVal
	}
	return params[0]
}

// 快速排序
func Partition[T any](arr []T, lo, hi int, compareFn func(a, b T) int) int {
    v := lo
    i := lo
    j := hi + 1
    for {
        for i++; compareFn(arr[i], arr[v]) < 0; i++ {
            if i == hi {
                break
            }
        }
        for j--; compareFn(arr[j], arr[v]) > 0; j-- {
            if j == lo {
                break
            }
        }
        if i >= j {
            break
        }
		arr[i], arr[j] = arr[j], arr[i]
    }
	arr[lo], arr[j] = arr[j], arr[lo]
    return j
}

func QuickSort[T any](arr []T, lo, hi int, compareFn func(a, b T) int) {
	if len(arr) < 2 {
		return
	}
    if hi <= lo {
        return
    }

    j := Partition(arr, lo, hi, compareFn)
    QuickSort(arr, lo, j-1, compareFn)
    QuickSort(arr, j+1, hi, compareFn)
}

func RemoveDuplicate[T string | int | float64](duplicateSlice []T) []T {
	set := map[T]interface{}{}
	res := []T{}
	for _, item := range duplicateSlice {
		_, ok := set[item]
		if !ok {
			res = append(res, item)
			set[item] = nil
		}
	}
	return res
}

type Student struct {
	Name string
	Age  int
}

func NewStudent(name string, age int) *Student {
	return &Student{Name: name, Age: age}
}

func DefaultFilter(item interface{}) (uniqueKey interface{}) {
	return item.(*Student).Name
}

func RemoveDuplicateWithFilter[T comparable](compareSlice []T, filterFunc func(item interface{}) (key interface{})) []T {
	set := map[interface{}]interface{}{}
	res := []T{}
	for _, item := range compareSlice {
		i := filterFunc(item)
		_, ok := set[i]
		if !ok {
			res = append(res, item)
			set[i] = nil
		}
	}
	return res
}

// 联合约束类型
type ID interface {
	int | string
}

// 写法  [T ID, D string] == [T int | string, D string]
type UserModel[T ID, D string] struct {
	Id   T
	Name D
}

func NewUserModel[A ID, D string](id A, name D) *UserModel[A, D] {
	return &UserModel[A, D]{Id: id, Name: name}
}

// 分页的实现
type KeepItem bool

// 若需要保留的item 则返回true 即可
type FilterFunc func(item interface{}) KeepItem

type PageList[T any] struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Size  int `json:"size"`
	List  []T `json:"list"`
}

type Pager[T any] struct {
	limit   int
	offset  int
	total   int
	pageCnt int
	list    []T
}

func NewPager[T any](list []T) *Pager[T] {
	return &Pager[T]{
		limit:  10,
		offset: 1,
		total:  len(list),
		list:   list,
	}
}

func (this *Pager[T]) Filter(filterFn FilterFunc) *Pager[T] {
	tmpList := []T{}
	for _, item := range this.list {
		if filterFn(&item) {
			tmpList = append(tmpList, item)
		}
	}
	this.list = tmpList
	this.total = len(tmpList)
	return this
}

func (this *Pager[T]) Offset(c int) *Pager[T] {
	this.offset = c
	return this
}

func (this *Pager[T]) Limit(c int) *Pager[T] {
	this.limit = c
	return this
}

func (this *Pager[T]) List() []T {
	// 页码
	if this.offset <= 0 {
		this.offset = 1
	}
	// size
	if this.limit > this.total {
		this.limit = this.total
	}
	// 总页数
	this.pageCnt = int(math.Ceil(float64(this.total) / float64(this.limit)))
	if this.offset > this.pageCnt {
		return []T{}
	}
	startIdx := (this.offset - 1) * this.limit
	endIdx := startIdx + this.limit

	if endIdx > this.total {
		endIdx = this.total
	}

	return this.list[startIdx:endIdx]
}

func (this *Pager[T]) Output() *PageList[T] {

	return &PageList[T]{
		Total: this.total,
		Page:  this.offset,
		Size:  this.limit,
		List:  this.list,
	}
}

// 通用初始化类型
type ModelObj interface {
	User | Product
}

type User struct {
	Uid int
}

func (this *User) SetId(id int) {
	this.Uid = id
}

type Product struct {
	Pid int
}

func (this *Product) SetId(id int) {
	this.Pid = id
}

// TrimModelObj 是一个动态类型的 Interface, 由M决定当前Interface的最终类型
type TrimModelObj[M ModelObj] interface {
	*M
	SetId(id int)
}

// TrimModelObj[Model] 由第二个参数决定当前的动态类型；
// NewModelObj[*User, User](32) 如 Model 是 User 类型, 最终 TrimModelObj == *User，所以我们需要为 Trim 传递 *User
func NewModelObj[Trim TrimModelObj[Model], Model ModelObj](id int) Trim {
	m := new(Model)
	t := Trim(m)
	fmt.Printf("%p \n", m)
	// 类型转换成指定的*Model
	t.SetId(id)
	return t
}
