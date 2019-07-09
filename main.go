package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var err error
var nums [1024]string
var indexNums int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readThisLine(theLine *bufio.Scanner) [5]string {
	var w [5]string
	/*標頭定義
	w[0] 自動 @#1/人工 @as
	w[1] 型態(type) @int
	w[2] 範圍(range) 1~99
	w[3] 個數(number)
	w[4] 顯示格式(format)補零、逗號...
	*/
	//讀到新的一行，要set w[5] all null
	for i := 0; i < 5; i++ {
		w[i] = ""
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = ""
	}
	indexNums = 0
	theLine.Split(bufio.ScanWords)
	for i := 0; theLine.Scan(); i++ {
		if i < 5 { //先讀標頭格式w[0]~w[5]
			w[i] = theLine.Text()
			//fmt.Println(w[i])
			if i == 0 {
				if w[0] == "@as" {
					//fmt.Println("人工測資")
					i = 4 //標頭結束
				}
			}
		} else {
			nums[indexNums] = theLine.Text()
			indexNums++
			//標頭讀完，把剩下nums[] 讀完
			//fmt.Print(theLine.Text() + ", ")

		}

	}
	return w

}

func showData(w [5]string) {
	/*產生測資*/
	//w[0] 自動 @#1/人工 @as
	//w[1] 型態(type) @int
	//w[2] 範圍(range) 1~99
	rangeBegin := 0
	rangeEnd := 0
	rangeSpecial := 0
	rangeFormat := ""

	if strings.ContainsAny(w[2], "~") {
		rangeBegin, err = strconv.Atoi(strings.Split(w[2], "~")[0])
		if err != nil {
			panic(err.Error())
		}

		if strings.ContainsAny(w[2], "&") {
			rangeFormat = "&"
			word2 := strings.Split(w[2], "~")[1]
			rangeEnd, err = strconv.Atoi(strings.Split(word2, rangeFormat)[0])
			if err != nil {
				panic(err.Error())
			}
			rangeSpecial, err = strconv.Atoi(strings.Split(word2, rangeFormat)[1])
			if err != nil {
				panic(err.Error())
			}
		} else if strings.ContainsAny(w[2], "*") {
			rangeFormat = "*"
			word2 := strings.Split(w[2], "~")[1]
			rangeEnd, err = strconv.Atoi(strings.Split(word2, rangeFormat)[0])
			if err != nil {
				panic(err.Error())
			}
			rangeSpecial, err = strconv.Atoi(strings.Split(word2, rangeFormat)[1])
			if err != nil {
				panic(err.Error())
			}
		} else if strings.ContainsAny(w[2], "^") {
			rangeFormat = "^"
			word2 := strings.Split(w[2], "~")[1]
			rangeEnd, err = strconv.Atoi(strings.Split(word2, rangeFormat)[0])
			if err != nil {
				panic(err.Error())
			}
			rangeSpecial, err = strconv.Atoi(strings.Split(word2, rangeFormat)[1])
			if err != nil {
				panic(err.Error())
			}
		} else {

			rangeEnd, err = strconv.Atoi(strings.Split(w[2], "~")[1])
			if err != nil {
				panic(err.Error())
			}
		}
	}
	//w[3] 個數(num)
	var num int
	if w[3] == "" {
		num = 1
		//fmt.Printf("從%d~%d中選%d個數字\n", rangeBegin, rangeEnd, num)
	} else if strings.ContainsAny(w[3], "~") {
		randBegin := 0
		randEnd := 0
		if strings.ContainsAny(w[3], "~") {
			randBegin, err = strconv.Atoi(strings.Split(w[3], "~")[0])
			if err != nil {
				panic(err.Error())
			}
			randEnd, err = strconv.Atoi(strings.Split(w[3], "~")[1])
			if err != nil {
				panic(err.Error())
			}
		}

		randNum := rand.Intn(randEnd - randBegin + 1)
		//fmt.Printf("randNum = %d\n", randNum)
		num = randNum + randBegin

		//fmt.Printf("從%d~%d中選%d~%d個數字, num = %d\n", rangeBegin, rangeEnd, randBegin, randEnd, num)
	} else {
		num, err = strconv.Atoi(w[3])
		if err != nil {
			panic(err.Error())
		}
		//fmt.Printf("從%d~%d中選%d個數字\n", rangeBegin, rangeEnd, num)
	}
	//w[4] 顯示格式(format)補零、逗號...

	//fmt.Println("最終測資為：")
	if w[2] == "0" {
		if w[4] == "g" { //從nums中挑選num個數字，可重複
			for i := 0; i < num; i++ {
				idx := rand.Intn(indexNums)
				fmt.Printf("%s ", nums[idx])
			}
		} else { //從nums中挑選num個數字，不可重複
			for i := 0; i < num; i++ {
				idx := rand.Intn(indexNums)
				for nums[idx] == "used" {
					idx++
					idx = idx % indexNums
				}
				fmt.Printf("%s ", nums[idx])
				nums[idx] = "used"
			}
		}
	} else {
		switch w[4] {
		case "":
			switch rangeFormat {
			case "":
				for i := 0; i < num; i++ {
					fmt.Printf("%d ", rand.Intn(rangeEnd-rangeBegin+1)+rangeBegin)
				}
				break
			case "&": //間隔rangeSpeical
				indexNums = 0
				for i := rangeBegin; i <= rangeEnd; i += rangeSpecial {
					nums[indexNums] = strconv.Itoa(i)
					indexNums++
				}
				for i := 0; i < num; i++ {
					idx := rand.Intn(indexNums)
					fmt.Printf("%s ", nums[idx])
				}
				break
			case "*":
				for i := 0; i < num; i++ {
					fmt.Printf("%d ", (rand.Intn(rangeEnd-rangeBegin+1)+rangeBegin)*rangeSpecial)
				}
				break
			case "^":
				for i := 0; i < num; i++ {
					outData := rand.Intn(rangeEnd-rangeBegin+1) + rangeBegin
					//fmt.Printf("begin=%d, end=%d, out=%d\n", rangeBegin, rangeEnd, outData)
					fmt.Printf("%d ", outData*outData)
				}
				break
			}
			break

		case "h": //前面補上個數
			fmt.Printf("%d ", num)
			switch rangeFormat {
			case "":
				for i := 0; i < num; i++ {
					fmt.Printf("%d ", rand.Intn(rangeEnd-rangeBegin+1)+rangeBegin)
				}
				break
			case "&":
				//fmt.Println("間隔晚點想")
				indexNums = 0
				for i := rangeBegin; i <= rangeEnd; i += rangeSpecial {
					nums[indexNums] = strconv.Itoa(i)
					indexNums++
				}
				for i := 0; i < num; i++ {
					idx := rand.Intn(indexNums)
					fmt.Printf("%s ", nums[idx])
				}
				break
			case "*":
				for i := 0; i < num; i++ {
					fmt.Printf("%d ", (rand.Intn(rangeEnd-rangeBegin+1)+rangeBegin)*rangeSpecial)
				}
				break
			case "^":
				for i := 0; i < num; i++ {
					outData := rand.Intn(rangeEnd-rangeBegin+1) + rangeBegin
					//fmt.Printf("begin=%d, end=%d, out=%d\n", rangeBegin, rangeEnd, outData)
					fmt.Printf("%d ", outData*outData)
				}
				break
			}
			break
		case "z": //尾巴補零
			switch rangeFormat {
			case "":
				for i := 0; i < num; i++ {
					fmt.Printf("%d ", rand.Intn(rangeEnd-rangeBegin+1)+rangeBegin)
				}
				break
			case "&":
				//fmt.Println("間隔晚點想")
				indexNums = 0
				for i := rangeBegin; i <= rangeEnd; i += rangeSpecial {
					nums[indexNums] = strconv.Itoa(i)
					indexNums++
				}
				for i := 0; i < num; i++ {
					idx := rand.Intn(indexNums)
					fmt.Printf("%s ", nums[idx])
				}
				break
			case "*":
				for i := 0; i < num; i++ {
					fmt.Printf("%d ", (rand.Intn(rangeEnd-rangeBegin+1)+rangeBegin)*rangeSpecial)
				}
				break
			case "^":
				for i := 0; i < num; i++ {
					outData := rand.Intn(rangeEnd-rangeBegin+1) + rangeBegin
					//fmt.Printf("begin=%d, end=%d, out=%d\n", rangeBegin, rangeEnd, outData)
					fmt.Printf("%d ", outData*outData)
				}
				break
			}
			fmt.Print("0")
			break
		case "e": //不留白
			switch rangeFormat {
			case "":
				for i := 0; i < num; i++ {
					fmt.Printf("%d", rand.Intn(rangeEnd-rangeBegin+1)+rangeBegin)
				}
				break
			case "&":
				//fmt.Println("間隔晚點想")
				indexNums = 0
				for i := rangeBegin; i <= rangeEnd; i += rangeSpecial {
					nums[indexNums] = strconv.Itoa(i)
					indexNums++
				}
				for i := 0; i < num; i++ {
					idx := rand.Intn(indexNums)
					fmt.Printf("%s", nums[idx])
				}
				break
			case "*":
				for i := 0; i < num; i++ {
					fmt.Printf("%d", (rand.Intn(rangeEnd-rangeBegin+1)+rangeBegin)*rangeSpecial)
				}
				break
			case "^":
				for i := 0; i < num; i++ {
					outData := rand.Intn(rangeEnd-rangeBegin+1) + rangeBegin
					//fmt.Printf("begin=%d, end=%d, out=%d\n", rangeBegin, rangeEnd, outData)
					fmt.Printf("%d", outData*outData)
				}
				break
			}
			break
		case "c": //加逗點
			switch rangeFormat {
			case "":
				for i := 0; i < num; i++ {
					if i != 0 {
						fmt.Print(",")
					}
					fmt.Printf("%d", rand.Intn(rangeEnd-rangeBegin+1)+rangeBegin)
				}
				break
			case "&":
				//fmt.Println("間隔晚點想")
				indexNums = 0
				for i := rangeBegin; i <= rangeEnd; i += rangeSpecial {
					nums[indexNums] = strconv.Itoa(i)
					indexNums++
				}
				for i := 0; i < num; i++ {
					idx := rand.Intn(indexNums)
					if i != 0 {
						fmt.Print(",")
					}
					fmt.Printf("%s", nums[idx])
				}
				break
			case "*":
				for i := 0; i < num; i++ {
					if i != 0 {
						fmt.Print(",")
					}
					fmt.Printf("%d", (rand.Intn(rangeEnd-rangeBegin+1)+rangeBegin)*rangeSpecial)
				}
				break
			case "^":
				for i := 0; i < num; i++ {
					if i != 0 {
						fmt.Print(",")
					}
					outData := rand.Intn(rangeEnd-rangeBegin+1) + rangeBegin
					//fmt.Printf("begin=%d, end=%d, out=%d\n", rangeBegin, rangeEnd, outData)
					fmt.Printf("%d", outData*outData)
				}
				break
			}
			break
		}
	}
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().Unix())
	cnt := 0
	file, err := os.Open("inputdata.txt")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close() //延後執行(func 結束時)

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {

		theLine := bufio.NewScanner(strings.NewReader(scanner.Text()))
		cnt++
		fmt.Printf("第 %d筆：", cnt)
		fmt.Println(scanner.Text())
		w := readThisLine(theLine)
		if w[0] != "@as" {

			showData(w)

		} else {
			//人工測資
			for i := 0; i < indexNums; i++ {
				fmt.Print(nums[i] + " ")
			}
		}
		fmt.Println()
	}

}
