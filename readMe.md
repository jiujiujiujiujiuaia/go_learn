


## golang 学习笔记

## 1命令源码文件和库源码文件

###命令源码文件
文件就是package为main并且有一个main函数的代码文件，命令源码文件可以接受参数
就像是linux下众多命令一样，因此可以很方便的写出命令工具。

一般标准的go工程目录结构是下图这样的，一个代码入口，其他都是库源码文件。
![](pic/20190629115035.png)

###库源码文件
库源码文件就是存在package非main的包中的代码，等待着被引用

###go命令和源码文件
最常见的go run 和go build 

- go build 

用于编译我们指定的源码文件或代码包以及它们的依赖包，但是之前遇到一种情况，和main.go同一个包下有库源码文件，然后main.go
有引用其代码，会报错**undefined**

![](pic/20190629120714.png)
![](pic/20190629120723.png)

为了规避这样的问题，go build 可以指定代码路径而不单独一个文件

## 2.go语言package和main意义

在golang项目中，只有package为main，且有main函数的代码是入口。
单个程序想要执行必须是package为main，才能执行其中的main函数。
多个这样的程序在一个目录下时，无法启动整个目录（为什么要启动整个目录，因为
程序之间有引用，不启动整个会报错，找不到某某）。
因此go语言文件名叫什么无所谓，引用其他代码时，只跟着package层次目录走（两个不同文件名
同一个package的程序，引用时路径一样，和文件名无关）


## 3.go中的nil
nil 是 interface、function、pointer、map、slice 和 channel 类型变量的默认初始值,
不像是java中null是给对象的默认值（习惯在java中检测字符串=="" && == null" go只需要 =="" 因为字符串没有nil"）

## 4.go中的string
golang中的string类型类似java的String类，都是不可变的，也就是你可以修改String对象的引用，但是无法直接修改底层的某个值
，但是你可以通过底层[]byte,[]rune获得并修改，byte是字节，rune是字符（一个字符可能由多个字节组成）

## 5.go语言中的反射i
sss1
11111test