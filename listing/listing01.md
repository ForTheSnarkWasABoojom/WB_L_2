Что выведет программа? Объяснить вывод программы.


package main
 
import (
    "fmt"
)
 
func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}

Ответ:

[77 78 79]

Программа выведет срез b, который является подсрезом массива a с индексами от 1 до 3 (включительно)
