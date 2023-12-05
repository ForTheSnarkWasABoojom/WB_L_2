Что выведет программа? Объяснить вывод программы.

package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func asChan(vs ...int) <-chan int {
   c := make(chan int)
 
   go func() {
       for _, v := range vs {
           c <- v
           time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
      }
 
      close(c)
  }()
  return c
}
 
func merge(a, b <-chan int) <-chan int {
   c := make(chan int)
   go func() {
       for {
           select {
               case v := <-a:
                   c <- v
              case v := <-b:
                   c <- v
           }
      }
   }()
 return c
}
 
func main() {
 
   a := asChan(1, 3, 5, 7)
   b := asChan(2, 4 ,6, 8)
   c := merge(a, b )
   for v := range c {
       fmt.Println(v)
   }
}

Ответ

Данная программа будет выводить цифры от 1 до 8 в произвольном порядке, затем программа начнет выводить нули, пока не будет остановлена.

Проблема связана с тем, что канал c не закрывается после окончания данных из каналов a и b. 
Вследствие этого, когда оба канала a и b закрываются, но канал c остается открытым, 
операция чтения из c начинает возвращать значения по умолчанию для типа int, то есть нули.

Для решения этой проблемы, можно добавить закрытие канала c после завершения чтения из каналов a и b.