# Let's golang

Practice tutorial by nico: https://nomadcoders.co/go-for-beginners

---

# 1. Theory

## 1.1 Packages and Imports

-   package func  
    대문자: public  
    소문자: private

## 1.2 Variables and Constants

-   go는 type language / type 명시 필요

## 1.3-4 Functions

-   compiler가 build 할 수 있도록 args type 명시 필요
-   여러 개의 return 값을 가질 수 있음
-   여러 개의 args 받으려면 ...string
-   'naked' return: return var 굳이 명시 필요 없음
-   'defer': finally절 같이 func 종료 후 추가 작업 명시

## 1.5 for, range

-   for 문법 1개뿐

## 1.6 If with a Twist

-   if else 조건문에서만 사용하는 변수 선언 가능

## 1.8 Pointers

-   low level programming
-   &a: 변수 a의 메모리상 주소값 (=address)
-   \*a: 변수 a가 가리키는 값 (=see through)

## 1.9 Arrays and Slices

-   arrays: static 크기
-   slices: dynamic 크기

## 1.10 Maps

-   key: value

## 1.11 Structs

-   go는 class와 object가 따로 없음
-   map과 달리 여러 type의 value 가지기 가능
-   python에 \_\_init\_\_ 과 같은 constructer method가 없어 직접 실행시켜야함

<br><br>

# 3. URL CHECKER & GO ROUTINES

## 3.2 Goroutines

-   앞에 'go'만 붙이면 됨
-   goroutines은 main func이 실행되는 동안만 유효
-   main func은 goroutines를 기다려주지 않음

## 3.3 Channels

-   goroutines과 main func 사이의 소통창구
-   '<-' 로 채널의 메세지 받기
