# go-demo

golang语言特性

1. 垃圾回收

   - 内存自动回收，再也不需要开发人员管理内存
   - 开发人员专注业务实现
   - 只需要new分配内存，不需要释放

2. 天然并发

   - 从语言层面支持并发

   - goro ute，轻量级线程，创建成千上万个gorute成为可能

   - 基于CSP(Communicating Sequential Process)模型实现

     ```go
      func main(){
         go fmt.Println("Hello")
     }
     ```

3. channel

   - 管道，类似unix/linux下的pipe
   - goroute之间通过channel进行通信
   - 支持任何类型

   ```go
   func main(){
       //chan：管道类型 int：管道内存放int型 3:容量 管道内只能放3个int数据
       pipe := make(chan int, 3)
       pipe <- 1 
       pipe <- 2
   }
   ```

   

4. 多返回值

   - 一个函数多个返回值

     ```go
     func calc(a int, b int)(int, int){
         sum := a+b
         avg := (a+b)/2
     	return sum, avg
     }
     ```



包的概念

1. 相同功能的代码放在同一目录下，该目录就被称为包
2. 包可以被其他包引用
3. main包可以用来生成可执行文件，每个程序只能有一个main包
4. 包的主要用途是提高代码的可用性

```shell
// 生成二进制文件 到bin目录下，名称为main
go build -o bin/main day_01/package_demo/main
```

   

文件名&关键字&标识符

1. 所有go文件以.go结尾

2. 标识符以字母或下划线开头，大小写敏感

3. _是特殊标识符，用来忽略结果

4. 保留关键字

   | break    | default    | func   | interface | select |
   | -------- | ---------- | ------ | --------- | ------ |
   | case     | defer      | go     | map       | struct |
   | chan     | else       | goto   | package   | switch |
   | const    | fallthough | if     | range     | type   |
   | continue | for        | import | return    | var    |



函数声明及注释

- 函数声明

  func 函数名 (参数列表) (返回值列表){}

  

- 注释
  1. 单行注释： //
  2. 多行注释：/* */

常量

1. 常量使用const修饰，代表永远只读，不能修改
2. const只能修饰boolean，number(int,浮点类型,complex)，string类型
3. 语法：const  identifier [type] = value，其中type可以省略

变量

1. 语法：var identifier type

值类型与引用类型

1. 值类型：变量直接存储值，内存通常在栈中分配。
   - 基础类型有int、float、bool、string以及数组和struct。
2. 引用类型：变量存储的是一个内存地址，该地址存储最终的值。内存通常在堆上分配。通过GC回收。
   - 指针、slice、map、chan等都是引用类型。

变量的作用域

1. 局部变量：在函数内部声明的变量，作用域仅限于函数内部。
2. 全局变量：在函数外部声明的变量，作用域仅限于整个包，如果是大写的话，作用于整个程序。

数据类型和操作符

1. bool类型，布尔类型：ture、false。
2. 数字类型，只要有int、int8、int16、int32、int64、int8、uint16、uint32、uint64、float32、float64。
3. byte类型，字符类型：var b byte = 'C'
4. string类型，字符串类型：var s string="hello"
   - "" ：双引号，转义特殊字符，如：\n
   - ``：反引号，原样输出
5. 类型转换，type(variable)，比如var a int=1；var b int32=int32(a)。
6. 短路操作符，!、&&、||。
7. 逻辑操作符，==、!=、<、<=、>、>=

指针类型

1. 普通类型，变量存储的是值

2. 指针类型，变量存储的是一个地址，改地址存储的才是值

3. 获取变量的地址，&，var a int 获取a的地址：&a

4. 获取指针类型的值，*，eg.：var *p int，*p即为p指针指向的值

   ```go
   var a int = 5
   var p *int = &a
   ```



流程控制

1. if/else if/else

2. switch

   ```go
   switch{
       case bool1:
       case bool2:
   	case bool3:
   	default :
   }
   
   switch value{
       case v1:
       case v2:
       case v3
       default
   }
   ```

3. for

   ```go
   //写法1
   for 初始化语句;条件判断;变量修改{
       
   }
   //写法2
   for 条件{
       
   }
   //写法3，用来遍历数组、slice、map、chan
   //for range 语句
   str:="Hello world"
   for i,v :=range str{
       
   }
   ```

4. continue和break

5. goto和label

   ```go
   lable1:
   for outer{
       label2:
       for inner{
           if xx{
               continue label1
           }else{
               goto label2
           }
       }
   }
   ```

   

   ​    

go 函数特点

1. 不支持重载，一个包不能有两个同名函数
2. 函数是一种类型，一个函数可以赋值给变量
3. 匿名函数
4. 多返回值

函数传参

1. 值传递
2. 引用传递

> 无论是值传递还是引用传递，传递给函数的都是变量的副本。值传递传递的是值副本，引用传递传的是引用的副本。一般来说，地址传递更为高效。因为值传递取决于值副本的大小，副本越大，则性能越低。

defer关键字

1. 当函数返回时，执行defer语句。可以用来做资源处理
2. 多个defer语句，按先进后出方式执行
3. defer语句中的变量，在defer声明时就决定了
4. 在return之前执行

内置函数、递归函数、闭包

1. close：主要来关闭chanel
2. len：用来求长度，比如string、array、slice、map、channel
3. new：用来分配内存。主要用来分配值类型，比如int、struct。返回的是指针
4. make：用来分配内存，主要用来分配引用类型，比如chan、map、slice
5. append：用来追加元素到数组、slice中
6. panic和recover：用来做错误处理
    > new声明一个对象并返回一个指针，make初始化并返回实际的对象
7. 闭包：一个函数及其相关的引用环境组合而成的实体,类似java中类的封装
   
   ```go
        func Adder() func(int) int{
            var x int
            return func (num int) int{
                x+=num
                return x
            }
        }
   ```

8. x
数组与分片
map数据结构
package介绍

NOTE:

- 格式化代码

  ```shell
  gofmt -w xxx.go
  ```