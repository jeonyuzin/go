package main

import (
	"fmt"
	"math/rand"
	"time" //seed 생성용 패키지
)

//난수 추출된 수의 소수 판정 프로그램 ver0.6
func main() {
	seed := time.Now().Unix() //seed설정
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2 //2~151사이에수
	fmt.Println("임의 로 추출된 수 ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break //약수 발견시 바로 탈출
		}
		//fmt.Print(i, " ")
	}
	if isPrime { //자체가 bool이라 이렇게쓸수있음 비교연산자 제거
		fmt.Println(number, "는 소수 입니다.")
	} else {
		fmt.Println(number, "는 소수가 아닙니다.")
	}
	/*for i := 1; i < 6; i++ {
		dice := rand.Intn(6) + 1 //0부터 입력한 수 앞까지 3이면 0~2
		fmt.Println(dice)        //근데 5만 나옴 난수의 시드 값을 조정해줘야함
	}*/

}
