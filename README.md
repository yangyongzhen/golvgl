# golvgl
golang language binding for littlevgl, what is littlevgl? a Open-source Embedded GUI Library 
( 开源UI库littlevgl的go语言绑定)
注：我本意就是在嵌入式linux上使用的封装。其他环境没考虑
使用方法是首先交叉编译littlevGL的源码生成动态库。
然后，main.go里为使用的demo

GOOS=linux GOARCH=arm GOARM=7 go build main.go

跑在板子上即可