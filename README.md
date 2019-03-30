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

4.默认情况下命令源码文件可以接受那些类型的参数值：布尔，整数，浮点，字符串，以及time.Duration类型

可以用自定义的数据类型作为参数值的类型吗？可以，只要实现flag包中的value接口。

# 三.库源码文件

1.如果该目录下有一个命令源码文件，那么为了让同在一个目录下的文件都通过编译，其他源码文件也应该声明属于main包。

2.源码文件声明的包名可以与目录名不一致，只要这些源码文件声明的包名一致就可以。

3.构建时生成的可执行文件的主名称与其父目录名称一致。

4.如何跨包调用实体程序呢？很简单，首先导入包，然后包名.函数名就可以调用，这里有一点需要注意就是函数，名首字母必须大写，这样才可以跨包。

5.对于实体程序的另外一种访问方式就是，创建internal代码包让实体程序只能被当前模块中的其他代码引用（也就是模块级私有）。

6.包名冲突的两种解决方式：1.起别名-import 别名 "包名/相对目录"2.本地化方式导入代码包import."相对目录"。

# 四.程序实体那些事（上）

1.go语言中的实体程序包括变量，常量，函数，结构体，接口。

2.变量名：=值（短变量声明），短变量声明只能在函数体内部使用。

3.类型推断的好处体现在重构的时候。他可以提高程序灵活性，使重构变得简单，同时又不会给代码的维护带来额外的负担，更不会损失运行效率。

4.变量重声明的一些条件：1.因为变量的类型在初始化的时候已经确定，所以重声明的类型必须与原本的类型相同。2.变量的重声明只有在使用短变量声明时才会发生。3.声明并赋值的变量必须是多个，并且其中至少有一个是新的变量。这时我们才可以说对其中的旧变量进行了重声明。4.必须发生在同一个代码块里，如果跨代码块，那就是另外一种含义了（即可重名变量，并且这个变量会屏蔽外层变量）。

# 五.程序实体那些事（中）

1.作用于最大的作用就是对实体程序的访问权限的控制。

2.程序在查找可重名变量时，会由内向外逐层代码块查找，直到找到当前代码包那一层还没找到就会报错，并不会去导入的代码包里找。如果导入的包是用本地化的方式（import.xxx）那么程序是会去导入的包里面找变量的。

3.可重名变量类型可以不一致。

# 六.程序实体那些事（下）

1.如何判断一个变量的类型。答：使用类型断言表达式。

value, ok := interface{}(container).([]string)//先转化为一个空接口然后再判断他是不是一个切片类型。

断言表达式的语法类型是想x.(T).x代表要被判断的类型的值，这个值当下的类型必须是接口类型，不过具体是那个接口类型其实是无所谓的。

多类型判断可以配合switch使用

2.别名类型。例如： type MyString =string。 给string换个名。

byte实际就是uint8。rune是int32

3.类型再定义。例如： type MyString2 string

只有潜在类型相同的不同类型可以互相转化。但他们的值之间不能比较。

# 七.数组和切片

1.数组是定长的切片是变长的。

2.package main

import "fmt"

func main() {

	// 示例 1。
	
	s1 := make([]int, 5)
	
	fmt.Printf("The length of s1: %d\n", len(s1))
	
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	
	fmt.Printf("The value of s1: %d\n", s1)
	
	s2 := make([]int, 5, 8)
	
	fmt.Printf("The length of s2: %d\n", len(s2))
	
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	
	fmt.Printf("The value of s2: %d\n", s2)
	
}

3.s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}

s4 := s3[3:6]//容量3（4，5，6）

fmt.Printf("The length of s4: %d\n", len(s4))

fmt.Printf("The capacity of s4: %d\n", cap(s4))

fmt.Printf("The value of s4: %d\n", s4)

4.切片窗口向右扩展到最大s4[0:cap(s4)]

5.估算切片容量增长：1.一般情况下新切片的容量是原切片的两倍。2.当原切片的长度大于1024，则会以1.25倍扩容。直到结果长度不小于要追加元素之和。

6.确切的说切片的底层数组并不会被替换，因为扩容生成新的底层数组时也生成了新切片。

7.无需扩容时append返回的原底层数组的新切片，需要扩容时返回的是新底层数组的新切片。

# 八.container包中的那些容器

1.list以及其元素结构Element

首先是list的四种方法：MoveBefore和MoAfter,他们分别作用于把给定的元素移动到另一个元素的前面和后面。

MoveToFront和MoveToBack,分别作用于把给定元素移动到链表最前面和最后面。

func (l *List) MoveBefore(e, mark *Element)

func (l *List) MoveAfter(e, mark *Element)


func (l *List) MoveToFront(e *Element)

func (l *List) MoveToBack(e *Element)

然后是Front和Back，分别用于获取链表最前和最后的元素。

insertBefore和insertAfter，分别用于在指定的元素之前和之后插入新元素

PushFront和PushBack，分别用于在链表的最前和最后插入元素

func (l *List) Front() *Element

func (l *List) Back() *Element


func (l *List) InsertBefore(v interface{}, mark *Element) *Element

func (l *List) InsertAfter(v interface{}, mark *Element) *Element


func (l *List) PushFront(v interface{}) *Element

func (l *List) PushBack(v interface{}) *Element

2.关于ring

  1.ring仅由他自身就可以代表，而list需要自身和Element。
 
  2.ring只代表所属循环链表中的一个元素，而list代表整个链表
  
  3.通过var r ring.Ring得到的是一个长度为1的循环链表，而list类型的零值则代表一个长度为0的链表。list的根元素不会持有实际元素
  
  4.创建一个ring值时可以指定他包含的元素数量。循环链表一旦被创建长度是不可变的。
  
  5.ring的len方法时间复杂度o(n),list为o(1)。
  
  # 九.字典的操作和约束
  
  1.aMap := map[string]int{
  
	"one":    1,
	
	"two":    2,
	
	"three": 3,
	
}
k := "two"

v, ok := aMap[k]

if ok {

	fmt.Printf("The element of key %q: %d\n", k, v)
	
} else {

	fmt.Println("Not found!")
	
}

2.首先哈希表会把键值转化为哈希值，然后利用这个哈希值的低几位去定位到一个哈希桶，然后再去这个哈希桶中，查找这个键，由于键-元素是绑定的
找到这个键就一定能找到其对应的元素值，随后，哈希表就会把相对应的元素值作为结果返回。

3.字典的键值类型不可以是：函数，字典，切片。

4.

























