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

4.为什么要对键值进行约束：1.当哈希值一样是需要对键值做==或！=的判断。2.宽度（单个值所占用的字节数）越小越好，长度越小越好。

5.我们对一个值为nil的字典做任何操作都不会引起错误。当我们试图在一个值为nil的字典中添加键-元素对的时候。go语言运行时系统会立即抛出一个panic。

6.判断一个操作是否原子可以使用go run race做数据竞争。

# 十.通道的基本操作

1.关于通道的容量：当容量为0时该通道属于非缓冲通道。当容量大于0属于缓冲通道。

2.对通道的接收和发送操作都有那些基本的特性？（互斥，不可分割，阻塞）

    1.对于同一个通道，发送是互斥的，接收也是互斥的。
    
    2.发送和接收操作中对元素值的处理都是不可分割的。
    
    3.发送操作在完全完成之前会被阻塞。接收操作也是如此。
    
3.元素进入通道会被复制，也就是说进入通道的元素并不是接收操作符右边的那个元素，而是它的副本。当元素从通道进入外界时会被移动。这个移动包含了两步操作：1.复制通道中的元素值，并准备给到接收方2.删除在通道中的这个元素值。

4.缓冲通道的阻塞情况：

  1.当通道满的时候所有的发送操作都会被阻塞。
  
  2.当通道已空，那么它所有的接收操作都会被阻塞。
  
5.非缓冲通道的阻塞情况：
   
  1.无论发送还是接收操作，一开始执行就会被阻塞，直到配对的操作也开始执行，才会继续传递。由此可见非缓冲通道是在用同步的方式传递数据。也就是说只有双方对接上了，数据才会被传递。
  
  2.数据是直接送发送方复制到接收方的，中间并不会用通道做中转。
  
6.只声明而没有用make函数初始化的空通道，不管发送还是接收都会永久被阻塞。
  
7.发送和接收操作什么时候会引发panic.
  
  1.已初始化，但未关闭的通道，收发操作不一定会引发panic。但是通道一旦关闭，在对他进行发送操作，就会引发panic.
  
  2.试图关闭一个已经关闭的通道，也会引发panic。
  
  3.可以用接收表达式的第二个返回值检测通道是佛关闭。
  
8.通道长度代表当前通道内包含的元素，当通道满的时候长度等于容量。

9.通道的进出都是浅层复制。go没有深层复制，除非我们自己做。

# 十一.通道的高级用法。

1.单向通道。价值约束代码。

2.func getIntChan() <-chan int {

	num := 5
	
	ch := make(chan int, num)
	
	for i := 0; i < num; i++ {
	
		ch <- i
		
	}
	
	close(ch)
	
	return ch
	
}

3.与select连用

// 准备好几个通道。

intChannels := [3]chan int{

	make(chan int, 1),
	
	make(chan int, 1),
	
	make(chan int, 1),
}

// 随机选择一个通道，并向它发送元素值。

index := rand.Intn(3)

fmt.Printf("The index: %d\n", index)

intChannels[index] <- index

// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。

select {

case <-intChannels[0]:

	fmt.Println("The first candidate case is selected.")
	
case <-intChannels[1]:

	fmt.Println("The second candidate case is selected.")
	
case elem := <-intChannels[2]:

	fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
	
default:

	fmt.Println("No candidate case is selected!")
	
}

4.select语句的分支选择规则
  
   1.分支从上到下顺序执行，每个分支的表达式从左到右顺序执行。
   
   2.当所有case都不满足，就会执行default,如果没有default那么就会阻塞。
   
   3.如果同时发现多个条件满足，那么就会用一种伪随机算法在这些分支中选择一个执行。
   
   4.一个select只能有一个default。
  
5。如果在select语句中发现某个通道已关闭，那么应该改怎么样屏蔽掉他所在的分支？

答：直接将nil赋给那个通道变量。

# 十二.使用函数的正确姿势

1.只要两个函数的参数列表和结果列表里面的元素顺序及其类型是一样的，我们就可以说这两函数是一样的函数，也就是说他们实现了同一种函数类型。

2.高阶函数(函数作为参数传入，函数作为结果返回)

type operate func(x, y int) int

op := func(x, y int) int {

	return x + y
	
}

func calculate(x int, y int, op operate) (int, error) {

	if op == nil {
	
		return 0, errors.New("invalid operation")
		
	}
	
	return op(x, y), nil
	
}

3.闭包
func genCalculator(op operate) calculateFunc {

	return func(x int, y int) (int, error) {
	
		if op == nil {
		
			return 0, errors.New("invalid operation")
			
		}
		
		return op(x, y), nil
		
	}
	
}

4.传入函数的如果是值类型，那么修改它不会影响原值，如果是引用类型（切片，字典，通道）则会影响原值

5.函数返回指针不会发生拷贝，返回非指针并把结果赋给其他变量则必定拷贝（谨记go都是浅拷贝，起的区分值类型和引用类型）

# 十三.结构体及其方法的使用法门

1.// AnimalCategory 代表动物分类学中的基本分类法。

type AnimalCategory struct {

	kingdom string // 界。
	
	phylum string // 门。
	
	class  string // 纲。
	
	order  string // 目。
	
	family string // 科。
	
	genus  string // 属。
	
	species string // 种。
	
}

func (ac AnimalCategory) String() string {

	return fmt.Sprintf("%s%s%s%s%s%s%s",
	
		ac.kingdom, ac.phylum, ac.class, ac.order,
		
		ac.family, ac.genus, ac.species)
		
}

2.内嵌字段及使用

type Animal struct {

	scientificName string // 学名。
	
	AnimalCategory    // 动物基本分类。
	
}

func (a Animal) Category() string {

	return a.AnimalCategory.String()
	
}

3.如果外围字段和内嵌字段具有一种同名方法，则内嵌字段的方法会被屏蔽。

4.一个结构体嵌入了两个结构体，这两个结构体有同名方法则会引发编译错误。

5.值方法与指针方法的区别
  
  1.值方法的接收者是该方法所属的类型值的一个副本，我们在该方法内对副本的修改并不会影响原值，除非他是个引用类型的别名类型。
    
    而指针方法的接收者是方法所属的类型值的指针值的一个副本。我们在这样的方法内对副本指向的值进行修改，却一定会体现在原值上。
  
6.空的结构体可以当占位符用，占据内存空间小。

# 十四.接口类型的合理运用

1.tyoe Pet interface{

SetName(name string)

Name() string

Category() string

}

2.如何判定一个数据类型的某一个方法实现的就是某个接口类型中的某个方法？
   
   1.签名完全一致
   
   2.名称完全一致
   
3.简单来说我们给一个接口变量赋值的时候，该变量的动态类型与动态值会储存在一个专用的数据结构中，这变量实际上算是这个专用数据结构的一个实例。
这个专用的数据结构叫iface。

4.字面量nil才是真nil,因为它的类型和值都是nil。

5.关于静态动态值，类型
      接口变量  接口类型                 
 var  a       A=        &b          ；*B
      静值    静类型     动值           动类型

# 十五.关于指针的有限操作

1.go语言中不可寻址的值
  
  1.常量的值
  
  2.基本类型值的字面量
  
  3.算数操作的结果值
  
  4.对各种字面量的索引表达式和切片表达式的结果值。不过对切片字面量的索引结果是可寻址的。
  
  5.对字符串变量的索引切片表达式
  
  6.对字典变量的索引表达式的结果值
  
  7.函数方法字面量，以及对他们调用的表达式的结果值
  
  8.结构体字面的字段值，也就结构体.字段
  
  9.类型转化表达式的结果
  
  10.类型断言表达式的结果
  
  11.接收表达式的结果
  
他们的共同特征：1.不可改变。2.临时结果。3.不安全的

2.怎么样通过unsafe.Pointer操纵可寻址的值

dog := Dog{"little pig"}

dogP := &dog

dogPtr := uintptr(unsafe.Pointer(dogP))//起始存储地址

namePtr := dogPtr + unsafe.Offsetof(dogP.name)//加上偏移量等于字段起始地址

nameP := (*string)(unsafe.Pointer(namePtr))

# 十六.go语句以及其执行规则（上）

1.G(goroutine),p(processor)可承载若干个G且能使这G适时的与M对接,M(machine)系统级线程。

2.main函数就是一个go程序的主goroutine当他结束的时候就相当于go程序结束，其他未被执行的goroutine就再也没有机会执行了。
package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
	
		go func() {
		
			fmt.Println(i)
			
		}()
		
	}
	
}

# 十七.go语句以及其执行规则（下）

1.如何使主goroutine等待其他goroutine
  
  1.使用time.Sleep()
  
  2.使用atomic原子操作实现简单的自旋来保证顺序。
  
# 十八.if语句丶for语句和switch语句

1.numbers1 := []int{1, 2, 3, 4, 5, 6}

for i := range numbers1 {

	if i == 3 {
	
		numbers1[i] |= i
		
	}
	
}

fmt.Println(numbers1)

range语句迭代一个变量是索引值，迭代两个变量是索引值和元素值

2.  numbers2 := [...]int{1, 2, 3, 4, 5, 6}

maxIndex2 := len(numbers2) - 1

for i, e := range numbers2 {

	if i == maxIndex2 {
	
		numbers2[0] += e
		
	} else {
	
		numbers2[i+1] += e
		
	}
	
}

fmt.Println(numbers2)

这里要注意数组与切片的区别。

3.若switch表达式和case子表达式值类型不一样，会发生以switch表达式为基准的自动转换。

5.value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}

switch value5[4] {

case value5[0], value5[1], value5[2]:

	fmt.Println("0 or 1 or 2")
	
case value5[2], value5[3], value5[4]:

	fmt.Println("2 or 3 or 4")
	
case value5[4], value5[5], value5[6]:

	fmt.Println("4 or 5 or 6")
	
}

利用索引表达式这种间接的方式绕过switch对case唯一性的束缚。

4.value6 := interface{}(byte(127))

switch t := value6.(type) {

case uint8, uint16:

	fmt.Println("uint8 or uint16")
	
case byte:

	fmt.Printf("byte")
	
default:

	fmt.Printf("unsupported type: %T", t)
	
}

但是这种判断类型的switch语句无法被绕过，这里byte和uint8类型重复了，编译无法通过。

# 十九.错误处理（上）

1.package main

import (

	"errors"
	
	"fmt"
	
)

func echo(request string)(reponse string,err error){

	if request==""{
	
		err=errors.New("empty request")
		
		return
		
	}
	
	reponse=fmt.Sprintf("echo111:",request)
	
	return
	
}

func main()  {

	for  _,req:=range []string{"","shabi!"}{
	
		fmt.Printf("request",req)
		
		resp,err:=echo(req)
		
		if err!=nil{
		
			fmt.Println(err)
			
			continue
	
		}
		
		fmt.Println(resp)
		
	}
	
}

2.对于具体的错误判断，go语言中都有那些惯用法？

    1.对于类型在已知范围内的一系列错误值，一般使用类型断言表达式或类型switch语句判断。

    2.对于已有相应变量且类型相同的错误值，一般直接用判等操作来判断
    
    3.对于没有相应变量且类型未知的一系列错误值，只能使用错误信息的字符串表达形式来做判断。

# 二十.错误处理（下）

1.用类型建立起树形结构的错误体系，用统一字段建立起可追根溯源的链式错误关联。

2.由若干个名称不同但类型相同的错误值组成集合。如果他们是公开的，那就应该使成为常量而不是变量，或编写私有的错误值以及公开的获取和判等函数，避免恶意篡改。

# 二十一.panic函数，recover函数以及defer语句（上）

1.从panic被引发到程序终止运行的大致过程是什么？
 
  答：假如说某个函数的某行代码无意中引发了panic,这时初始化panic就会被建立起来，并且改程序的控制权就会从这行代码转移至调用其所属的函数的那行代码上，也就是调用栈中的上一级，但是控制权并不会在这里停留，而是会继续一级一级往上转移，直到最外层函数哪里，一般情况下就是主goroutine也就是main函数哪里，然后控制权被go语言运行时系统回收。随后程序崩溃并停止运行。在这个控制权不断转移的过程中，panic会不断地积累和完善，并在程序终止之前被打印出来。
  
# 二十二.panic函数，recover函数以及defer语句（下）
  
1.怎么样让panic包含一个值，以及应该让它包含什么样的值？
   
  答：直接调用panic函数就可以了，它的参数是一个空接口类型的，所以从语法上讲可以接受任何类型的值。最好是让他接收error类型的错误值。
  
2.怎样施加应对panic的保护措施，从而避免程序崩溃？
  
  答：使用内建函数recover。
  
package main

import (

 "fmt"
 
 "errors"
 
)

func main() {

 fmt.Println("Enter function main.")
 
 defer func(){
 
  fmt.Println("Enter defer function.")
  
  if p := recover(); p != nil {
  
   fmt.Printf("panic: %s\n", p)
   
  }
  
  fmt.Println("Exit defer function.")
  
 }()
 
 // 引发 panic。
 
 panic(errors.New("something wrong"))
 
 fmt.Println("Exit function main.")
 
}

3.多个defer函数调用的执行顺序

  1.从下到上一次调用
  
  2.存储defer函数以及参数的队列他是先进后出（FILO)的，相当于一个栈。

由上所述demo51的结果应该是：last ,2,1,0, frist

# 二十三.测试的基本规则和流程（上）

1.go语言对测试函数的名称和签名都有那些规定

  1.对于功能测试函数来说，名称必须以Test为前缀，并且参数列表中应有×testing.T类型的参数声明。
  
  2.对于性能测试函数来说，其名称必须以Benchmark为前缀，并且唯一参数的类型必须是×testing.B类型的。
  
  3.对于示例测试函数来说，其名称必须以Example为前缀，但对函数的参数列表没有强制规定
  
2.go test命令执的主要测试流程是什么？

  答：只有测试源码文件名对了，测试函数名和签名对了，我们运行go test的时候，其中的测试代码才有可能被运行。构建，执行测试函数，删除临时文件，打印结果。为了加快测试速度，他会对多个待测试的代码包并发的执行这个流程，但是结果还是按照我们给定顺序打印的。另一方面并发测试会让性能测试的结果发生偏差，所以性能测试一般都是串行执行。
 
3.testing.T 的部分功能有（判定失败接口，打印信息接口）
  testing.B 拥有testing.T 的全部接口，同时还可以统计内存消耗，指定并行数目和操作计时器

# 二十四.测试的基本规则和流程（下）

1.调用Fail()测试函数会执行下去，但结果会显示测试失败。
  调用FailNow()测试立即失败，后面的代码不会被执行。
  Log()和Logf()打印常规的测试日志。
  Error()和Errorf()在测试失败时打印失败测试日志。
  Fatal()和Fatalf()打印失败测试日志之后立即终止当前测试的执行并宣告测试失败。相当于最后都调用了FailNow()。

2.性能测试结果
  
  $ go test -bench=. -run=^$ puzzlers/article20/q3
  
  goos: darwin
  
  goarch: amd64
  
  pkg: puzzlers/article20/q3
  
  BenchmarkGetPrimes-8                     500000                    2314 ns/op
  
  被测试函数          最大P数（逻辑CPU最大数）  （最大测试次数）               平均时间
  
  PASS
  
  ok   puzzlers/article20/q3 1.192s
  
                             总时间
  
标记-bench才会进行性能测试 标记-run表明需要执行那些功能测试函数^$以为执行名称为空的功能测试函数，换句话说，不执行任何功能测试函数。

# 二十五.更多的测试手法

1.-cpu 1,2,4 相当于测试性能(CPU数)不同的计算机
  
  -count重复执行测试函数，它大于等于0，默认为1
  
  -parallel设置一个被侧代码包中的给你测试函数最大并发数、默认值是测试运行时的最大P数量。
  
  -cover标记。开启测试覆盖度分析。
  
2.性能测试的并发牵扯到b.RunParallel方法、b.SetParallelism方法和-cpu标记的联合运用。

3.性能测试函数中的计时器：StarTimer,StopTimer,ResetTimer。

# 二十六sync.Mutex与sync.RWMutex

1.使用互斥锁时的注意事项！
  
  1.不要重复锁定互斥锁；（对一个已经锁定了的互斥锁再次锁定，当前的goroutine就会立即阻塞，这个goroutine的执行流程，会一直停滞在互斥锁lock的那行代码上）
  
  2.不要忘记解锁互斥锁，必要时使用defer语句；
  
  3.不要对尚未锁定，或者已经解锁的互斥锁解锁；
  
  4.不要在多个函数之间传递互斥锁；
 
2.读写锁的规则

  与互斥基本相似，但是读写分离可以执行更细腻的控制。可以多读，但不可以多写，其中写锁的实现也是使用了互斥锁。
  
3.读写锁和互斥锁都实现Locker接口。

4.获取读锁func (rw RWMutex)RLocker() Locker

# 二十七.条件变量sync.Cond(上)

1.条件变量怎样与互斥锁配合使用
  
  答：条件变量的初始化离不开互斥锁，并且它的方法有的也是基于互斥锁的。
  
2.sysc.Cond类型的值一旦被使用就不应该再被传递了，应为传递往往意味着拷贝，拷贝一个已经使用过的sync.Cond值会引发panic。但是它的指针值是可以被拷贝的

# 二十八.条件变量sync.Cond(下)

1.条件变量的wait方法做了什么？
  
  1.把调用他的gouroutine也就是当前gouroutine加入到条件变量的通知队列中。
  
  2.解锁当前的条件变量基于的那个互斥锁。
  
  3.让当前的goroutine处于等待状态，等到通知到来时再做决定是否唤醒他，此时这个gouroutine会阻塞在调用这个wait的那行代码上。
  
  4.如果通知到来决定唤醒这个gouroutine,那么就在唤醒他后重新锁定当前条件变量基于的互斥锁。自此之后，当前的gourutine就会继续执行代码了。 

2.singal和broadcast方法的异同
  
  1.都会发送通知，但前者前者只会唤醒一个而后者会唤醒所有。
  
# 二十九.原子操作（上）

1.sync/atomic包中提供了几种原子操作？可操作的数据类型又有那些？

  1.加法（add）,比较并交换（cas）,加载（load）,存储（store）和交换（swap）。
  
  2.数据类型：int32,int64.uint32,uint64,uintptr,unsafe包中的pointer。还有一个value他可以用来保存所有类型。
  
# 三十.原子操作（下）
  
1.不能用原子值储存nil。

2.原子值里存储的第一个值，决定了它今后能且只能存储那一个类型。

3.对于给原子值储存引用类型的值引发并发问题的解决方案，对引用值进行深拷贝，把副本存入原子值，这样对对该值的任何操作都不会影响到原子值。

# 三十一.sync.WaitGroup和sync.Once

1.WaitGroup类型值中的计数器不能小于0

2.sync.Once类型值的Do方法是怎么保证只执行参数函数一次的？

答：1.参数函数必须是无参数声明无结果声明的。2.只执行首次调用时传入的参数函数，之后的就不执行了。所以多个函数需要多个Once。3.Once类型中有一个done字段，用来记录Do被执行的次数，该字段只能为1或0，一旦Do方法被调用他的值就会由0变为1。

# 三十二.context.Context类型

1.主要内容就是context树.以及他的四个函数一个可手动撤销的，两个根据时间撤销的，一个不可撤销的并且可携带数据的。background是他的根。value根据键值的方式存储数据。context值在传达撤销信号的时候是深度优先的。

# 三十三.临时对象池sync.Pool

1.sync.Pool临时对象池，可以被用来储存临时的对象。属于结构体类型，他的值在被真正使用后就不应该被复制了。

2.sync.Pool类型只有Put和Get方法。put接受一个interface{}类型的参数用于在当前的池中存放临时对象，而get则被用于从当前的池中获取临时对象，具体的说get方法会从池子里删除掉任何一个值，并把这个值作为结果返回。如果池子里没有值，那么方法就会用当前池子的new字段创建一个新值，并将其返回。该字段就是get方法最后的临时对象获取手段，这个对象不会被存入临时对象值，而是会直接返回给调用方。

3.sync包在被初始化的时候，会向go语言运行时系统注册一个函数，这个函数的功能就是清除所有以创建的临时对象池中的值。我们可以把它称为池清理函数。一旦被注册，后者在每次即将执行垃圾回收时就都会执行前者。

4.池总汇列表

5.临时对象池储存所用的数据结构是怎样的？
  
  答：顶层本地池列表，是一个数组，这个列表的长度总与go语言调度器中的p数量相同。在本地池列表中的每个本地池都包含了三个字段，他们是：储存私有临时对象的字段private,代表了共享临时对象列表的字段shared,以及一个sync.Mutex类型的嵌入字段。

6.临时对象池是怎样利用内部数据结构来存取值的？

  答：先存private,如果private有东西，就存在sahred。取的时候也是先取private后取shared。如果都没有就用new字段弄个新的返回去。
  
# 三十四.并发安全字典sync.Map(上)

1.并发字典和原生字典一样都不支持函数，切片，字典作为储存的键，值。

2.我们可以利用reflect.TypeOf函数获取一个键值的反射类型值，再用他的反射类型值调用Comparable方法，得到确切的判断结果。用来判断键，值是否违反1.。

3.我们可以从一开始（编码时期）就对存入的键，值作限定，这样编译器就可以帮助我们进行检查了。

# 三十五.并发安全字典sync.Map(下)

1.type ConcurrentMap struct {

 m         sync.Map
 
 keyType   reflect.Type
 
 valueType reflect.Type
 
}

多添加两个字段类型为反射类型，方便对键，值进行判断。不限制m的输入输出，增加了扩展性。但是反射或多或少会降低程序性能。

2.它有两个字典作为存储介质，一个read（原子值）,一个dirty（原生字典），如果进行查找会先去read,当read中没有时，会在所的保护下去dirty找，他们存储的都是unsafa.pointer的值（指针）。

3.当从脏字典查找键值的此书足够多的时候，脏字典会被转换成只读字典，脏字典变为nil,并根据已经改变的read字典重建脏字典。

# 三十六.unicode与字符编码

1.string值在底层是由一系列相对应的unicode代码点的utf-8编码值来表达的。
          
2.string->[]rune->[]byte

3.判断unicode字符是否为单字节字符，可以使用unicode/utf8代码包中的runelen,encoderune函数等。

# 三十七.string包与字符串操作

1.string.bulider的优势
  
  1.以存在的内容不可改变，但可以拼接更多内容。
  
  2.减少了内存分配和内容拷贝的次数。
  
  3.可将内容重置，可重用值。

2.builder值拥有的一系列指针方法write,writebyte,writerune,writestring。我们可把它们统称为拼接方法。

3.通过调用grow方法手动扩容builder

4.调用reset方法，我们可以让builder值重回零值状态。       

5.string.reader高效读取的关键就是已读计数。该计数代表着下一次读取的起始索引位置。调用seek方法可以直接设定该值中的已读计数器。

# 三十八.bytes包与字节串操作（上）

1.byte.buffer值的长度是未读内容的长度，而不是以存内容的长度。容量是容器容量

# 三十九.bytes包与字节串操作（下）

1.buffer在扩容的时候会先检测内容容器的剩余容量，如果够就进行长度扩充，如果不够就进行容量扩充（用新容器代替原容器新容器的容量 =2* 原有容量 + 所需字节数）。如果剩余容量的一半等于或大于现有长度再加上增加字节数的和。那么就会继续利用现有容器。并把未读内容拷贝到它的头部位置。这就意味着以读内容将会被覆盖。

2.处于灵值状态的buffer如果第一次扩容不超过64，那么该值就会基于一个预先定义好的长度为64的字节数组来创建内容容器。

3.Bytes方法和Next方法都有可能造成内容泄漏，原因是他们都把基于内容容器的切片直接返给了调用方。解决方法先做深拷贝再赋给别的变量。

# 四十.io包中的接口和工具（上）































  

  
  
  
  
  
  








  


 




      







  





























