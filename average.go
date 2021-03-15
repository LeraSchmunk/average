// Среднее значение

package main

import (
	"fmt"
	"log"

	"github.com/LeraSchmunk/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0              //Переменная для хранения суммы 3х чисел
	for _, number := range numbers { //Перебор всех чисел в массиве
		sum += number //Текущее число прибавляется в сумме
	}
	fmt.Println(sum)

	sampleCount := float64(len(numbers))                    //Длина массива в виде int и преобразование в float64
	fmt.Printf("Average/Среднее: %0.2f\n", sum/sampleCount) //Вычисление среднего
}
