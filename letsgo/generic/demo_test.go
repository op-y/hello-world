package generic

import (
	"fmt"
	"testing"
)

func TestDefaultKeyWordParams(t *testing.T) {
	category := []string{}
	realCategory := DefaultKeyWordParams[string]("default", category...)
	fmt.Println(realCategory)
}

func TestQuickSort(t *testing.T) {
	//nums := []int{9, 3, 1, 7, 4, 8, 6, 2, 5}
	nums := []int{6, 5, 6, 4, 5}
	fmt.Println("Unsorted:", nums)

	QuickSort[int](nums, 0, 4, func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	})
	fmt.Println("Sorted:  ", nums)

	strs := []string{"orange", "apple", "banana", "kiwi", "grape"}
	fmt.Println("Unsorted:", strs)

	QuickSort[string](strs, 0, 4, func(a, b string) int {
		if len(a) < len(b) {
			return -1
		} else if len(a) > len(b) {
			return 1
		} else {
			return 0
		}
	})
	fmt.Println("Sorted:  ", strs)
}

func TestRemoveDuplicate(t *testing.T) {
	fmt.Println(RemoveDuplicate[string]([]string{"a", "c", "a"}))
	fmt.Println(RemoveDuplicate[int]([]int{1, 2, 1, 1, 1}))
}

func TestRemoveDuplicateWithFilter(t *testing.T) {
	s := []*Student{
		NewStudent("a", 1),
		NewStudent("a", 1),
		NewStudent("b", 2),
		NewStudent("b", 2),
	}
	l := RemoveDuplicateWithFilter[*Student](s, DefaultFilter)
	for _, i := range l {
		fmt.Println(i.Name, i.Age)
	}
}

func TestNewUserModel(t *testing.T) {
	fmt.Println(NewUserModel[int, string](10, "hello"))
	fmt.Println(NewUserModel[string, string]("10", "hello"))
}

func TestPager(t *testing.T) {
	page := NewPager[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	list := page.Offset(1).Limit(3).Filter(
		func(item interface{}) KeepItem {
			if *item.(*int)%2 == 1 {
				return true
			}
			return false
		},
	).List()
	fmt.Println(list)
}

func TestNewModelObj(t *testing.T) {
	// new user model object
	user := NewModelObj[*User, User](32)
	fmt.Printf("%p \n", user)
	fmt.Printf("%T \n", user)
	fmt.Println(user.Uid)

	// new product model object
	prod := NewModelObj[*Product, Product](18)
	fmt.Printf("%p \n", prod)
	fmt.Printf("%T \n", prod)
	fmt.Println(prod.Pid)
}
