# go学习笔记

# 一.工作区和go path
1.go version可以检验是否安装成功。

2.goroot go语言的安装目录，也就是go语言的安装路径

3.gopath 若干动作区目录。是我们自己定义的工作区

（我们这些工作区去放置go语言的源码文件，以及安装后的归档文件和可执行文件）

4.gobin   go语言程序生成的可执行文件的路径

5.go语言源码的组织方式

以代码包为基本组成单位，在文件系统中，这些代码包与目录一一对应。由于目录可以有子目录，所以代码包可以有子包。

一个代码包中可以含任意个以.go为扩展名的源码文件，这些文件都需要被声明属于同一个代码包。

代码包的名称一般会与源码文件所在的目录同名，如果不同名，那么在构建，安装的过程中会以代码包名称为准

每个代码包都会有导入路径。代码包的导入路径是其他代码需要使用该包的实体程序时需要引入的路径。在实际使用程序实体之前，我们必须先导入其所在的代码包。

例如 import "github.com/labstack/echo"

在工作区中，一个代码包的导入路径实际上就是从src子目录到该包的实际储存位置的相对路径。

所以说go语言源码的组织方式就是以环境变量gopath，工作区，src目录和代码包为主线的。

一般情况下，go语言的源码文件都需要被存放在环境变量gopath包含的某个工作区（目录）中的src目录的某个代码包中

6.了解源码安装后的结果

安装后产生的归档文件（以.a为扩展名的文件），就会放进该工作区的pkg子目录；如果产生了可执行文件，就可能会放进该工作区的bin子目录

归档文件存放的具体位置和规则

源码文件会以代码包的形式组织起来，一个代码包其实就对应一个目录，安装某个代码包而产生的归档文件是与这个代码包同名的。

放置它的相对目录就是代码包的导入路径的直接父级

如导入目录为github.com/labstack/echo，

那么执行命令go install github.com/labstack/echo

生成的归档文件的相对目录就是 github.com/labstack，文件名就是echo.a

归档文件的相对目录与pkg目录之间还有一级目录叫做平台相关目录

7.理解构建和安装go程序的过程

构建命令go build,安装命令go install。构建和安装代码包的时候都会执行编译，打包等操作，并且，这些操作生成的任何文件都会先被保存到某个临时的目录中。

如果构建的是库源码文件，那么操作后产生的结果文件只会存在与临时目录中。这里的构建主要意义在于检查和验证。

如果构建的是命令源码文件，那么操作的结果文件也会被搬运到源码文件所在的目录。

安装操作会先执行构建，然后还会链接操作，并且把结果文件搬运到指定目录。

进一步说，如果安装的库源码文件，那么结果文件会被搬运到它所在工作区的okg目录下的某个子目录中。

如果安装的是命令源码文件，那么结果文件会被搬运到它所在工作区的bin目录中，或者环境变量gobin指向的目录中。

# 二.命令源码文件

1.什么是命令源码文件

答：是程序的入口，是可以通过安装和构建，生成与其对应的可执行文件，后者一般会与命令源码文件的直接父级目录同名。

如果一个源码文件声明属于main包，并且包含一个无参数的无声明的main函数，那么他就是命令源码文件。

2.flag.StringVar(&name, "name", "everyone", "The greeting object.")

这个函数的四个参数：1.参数地址，2.命令参数，3.默认值，4.对参数简单的说明

或者这样var name = flag.String("name", "everyone", "The greeting object.")

运行：go run demo2.go -name="Robert"

参数说明：$ go run demo2.go --help

go run 命令会临时生成可执行文件

3.自定义参数使用说明

对flag.usage重新赋值

在main函数开始处

flag.Usage = func() {

 fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
 
 flag.PrintDefaults()
 
}

或者对flag.CommandLine重新赋值

在init()函数开始处

flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)

flag.CommandLine.Usage = func() {

	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
  
	flag.PrintDefaults()
  
}

或者创建一个自己的私有容器

var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

把对flag.stringvar的调用替换成cmdline.parse(os.Args[:1])

os.Args[1:]就是我们给定的那些命令参数

















