// go luôn cần package, khai báo package main đầu
package main

// import thêm package con cần thiết
import (
	"fmt"
	"math"
	"os"
)

// hàm trong go có cấu trúc như sau
func addNum(a,b int) (int, rune) {
	return a + b, rune(a)
}

func main() {
	fmt.Println("Hello, World!")
	// package os dùng để xử lý các biến môi trường env, được khởi tạo trong file launch.json
	fmt.Print("In ten app ", os.Getenv("APP_NAME"))
	// mọi hàm, biến muốn export để sử dụng chung trong GO đều cần viết hoa chữ cái đầu
	fmt.Println(math.Pi)
	var a int = 1 // khai báo biến đồng thời khởi tạo giá trị, mặc định là 0
	b := 2 // Go tự động gán type dữ liệu theo giá trị khởi tạo
	fmt.Println(addNum(a,b))
	c := float64(a) // ép kiểu trong Go
	fmt.Println(c)
	const V = 123 // khai báo hằng V
	fmt.Println(V)

	// vòng lặp for
		sum := 0
		for i :=0; i<10;i++ {
			sum += i
		}
		fmt.Println(sum)

		for ;sum<50; sum ++ {
			fmt.Println(sum)
		}
	// for while và forever
		for sum <= 100 {
			fmt.Println(sum)
			sum += sum
		}
		// forever = while(true)
		/*
		for {
			
		}
		*/
	// if switch và defer
	if sum > 100 {
		fmt.Println("sum > 100")
	} else {
		fmt.Println("sum <= 100")
	}

	if count := 10; count > 0 {
		fmt.Println("count > 0")
	}
	// khong can break
	switch sum {
	case 1:
		fmt.Println("sum = 1")
	case 2:
		fmt.Println("sum = 2")
	default:
		fmt.Println("sum khong phai 1 hoac 2")
	}

	// defer => thực thi ngay trước khi hàm main kết thúc, ứng dụng dọn rác, handle error
	// defer có tính chất stacking tức last in fisrt out
	defer fmt.Println("sum = ", sum)

	for i := 0; i < 10; i++ {
		defer fmt.Println("i = ", i)
	}
}