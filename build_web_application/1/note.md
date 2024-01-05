### GOROOT 与 GOPATH
GOROOT: GO语言包安装的位置(bin文件的位置)
GOPATH: `go get`的第三方安装包安装于此，GOPATH允许多个目录，当有多个目录时，请注意分隔符，多个目录的时候Windows是分号，Linux系统是冒号，当有多个GOPATH时，默认会将`go get`的内容放在第一个目录下。
以上 $GOPATH 目录约定有三个子目录：
- src 存放源代码（比如：.go .c .h .s等）
- pkg 编译后生成的文件（比如：.a）
- bin 编译后生成的可执行文件（为了方便，可以把此目录加入到 PATH 变量中，如果有多个gopath，那么使用{GOPATH//...//bin:}/bin添加所有的bin目录）

### GOPATH代码目录结构规划
GOPATH下的src目录就是接下来开发程序的主要目录，所有的源码都是放在这个目录下面，那么一般我们的做法就是一个目录一个项目，例如: $GOPATH/src/mymath 表示mymath这个应用包或者可执行应用，这个根据package是main还是其他来决定，main的话就是可执行应用，其他的话就是应用包，这个会在后续详细介绍package。

所以当新建应用或者一个代码包时都是在src目录下新建一个文件夹，文件夹名称一般是代码包名称，当然也允许多级目录，例如在src下面新建了目录$GOPATH/src/github.com/astaxie/beedb 那么这个包路径就是"github.com/astaxie/beedb"，包名称是最后一个目录beedb

下面我就以mymath为例来讲述如何编写应用包，执行如下代码：
> cd GOPATH/src // GOPATH 为自己电脑环境上的GOPATH路径
> mkdir mymath

新建文件sqrt.go，内容如下：
```go
/ $GOPATH/src/mymath/sqrt.go源码如下：
package mymath

func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}
```

### 编译应用
上面我们已经建立了自己的应用包，如何进行编译安装呢？有两种方式可以进行安装：
1. 只要进入对应的应用包目录，然后执行go install，就可以安装了
2. 在任意的目录执行如下代码go install mymath
安装完之后，我们可以进入如下目录：
> // C:\Users\zoucheng02\go\pkg\windows_amd64(window64的地址)
> 可以看到 mymath.a

这个.a文件是应用包，那么我们如何进行调用呢？

接下来我们新建一个应用程序来调用这个应用包

新建应用包mathapp
在src下新建mathapp文件夹，在mathapp文件夹下新建main.go文件：
```go
package main

import (
	"mymath"
	"fmt"
)

func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
}
```
可以看到这个的package是main，import里面调用的包是mymath,这个就是相对于GOPATH/src的路径，如果是多级目录，就在import里面引入多级目录，如果你有多个GOPATH，也是一样，Go会自动在多个GOPATH/src中寻找。

如何编译程序呢？进入该应用目录，然后执行go build，那么在该目录下面会生成一个mathapp的可执行文件：
> // 项目目录下执行，可执行文件会生成在当前目录下
> go build 

如何安装该应用，进入该目录执行go install,那么在$GOPATH/bin/下增加了一个可执行文件mathapp, 还记得前面我们把$GOPATH/bin加到我们的PATH里面了，这样可以在命令行输入如下命令就可以执行
> // 在src目录下执行，可执行文件会生成在bin目录下
> go install ./mathapp

在控制台执行：
> mathapp

也是输出如下内容：
> Hello, world.  Sqrt(2) = 1.414213562373095

### 获取远程包
go语言有一个获取远程包的工具就是go get，目前go get支持多数开源社区(例如：GitHub、googlecode、bitbucket、Launchpad)
`go get github.com/astaxie/beedb`

> go get -u 参数可以自动更新包，而且当go get的时候会自动获取该包依赖的其他第三方包

通过这个命令可以获取相应的源码，对应的开源平台采用不同的源码控制工具，例如GitHub采用git、googlecode采用hg，所以要想获取这些源码，必须先安装相应的源码控制工具

通过上面获取的代码在我们本地的源码相应的代码结构如下：
```
$GOPATH
  src
   |--github.com
		  |-astaxie
			  |-beedb
   pkg
	|--相应平台
		 |-github.com
			   |--astaxie
					|beedb.a
```
go get本质上可以理解为**首先第一步是通过源码工具clone代码到src下面，然后执行go install**

在代码中如何使用远程包，很简单的就是和使用本地包一样，只要在开头import相应的路径就可以
`import "github.com/astaxie/beedb"`

### 程序的整体结构
通过上面建立的我本地的mygo的目录结构如下所示：
```
bin/
	mathapp
pkg/
	平台名/ 如：darwin_amd64、linux_amd64
		 mymath.a
		 github.com/
			  astaxie/
				   beedb.a
src/
	mathapp
		  main.go
	mymath/
		  sqrt.go
	github.com/
		   astaxie/
				beedb/
					beedb.go
					util.go
```
从上面的结构我们可以很清晰的看到，bin目录下面存的是编译之后可执行的文件，pkg下面存放的是应用包，src下面保存的是应用源代码。

### Go命令
#### go build
这个命令主要用于编译代码。在包的编译过程中，若有必要，会同时编译与之相关联的包。
- 如果是普通包，就像我们在1.2节中编写的mymath包那样，当你执行go build之后，它不会产生任何文件。如果你需要在$GOPATH/pkg下生成相应的文件，那就得执行go install。
- 如果是main包，当你执行go build之后，它就会在当前目录下生成一个可执行文件。如果你需要在$GOPATH/bin下生成相应的文件，需要执行go install，或者使用go build -o 路径/a.exe。
- 如果某个项目文件夹下有多个文件，而你只想编译某个文件，就可在go build之后加上文件名，例如go build a.go；go build命令默认会编译当前目录下的所有go文件。
- 你也可以指定编译输出的文件名。例如1.2节中的mathapp应用，我们可以指定go build -o astaxie.exe，默认情况是你的package名(非main包)，或者是第一个源文件的文件名(main包)。
（注：实际上，package名在Go语言规范中指代码中“package”后使用的名称，此名称可以与文件夹名不同。默认生成的可执行文件名是文件夹名。）
- go build会忽略目录下以“_”或“.”开头的go文件。
- 如果你的源代码针对不同的操作系统需要不同的处理，那么你可以根据不同的操作系统后缀来命名文件。例如有一个读取数组的程序，它对于不同的操作系统可能有如下几个源文件：
`array_linux.go array_darwin.go array_windows.go array_freebsd.go`
go build的时候会选择性地编译以系统名结尾的文件（Linux、Darwin、Windows、Freebsd）。例如Linux系统下面编译只会选择array_linux.go文件，其它系统命名后缀文件全部忽略。

#### go clean
这个命令是用来移除当前源码包和关联源码包里面编译生成的文件。这些文件包括：
```
_obj/            旧的object目录，由Makefiles遗留
_test/           旧的test目录，由Makefiles遗留
_testmain.go     旧的gotest文件，由Makefiles遗留
test.out         旧的test记录，由Makefiles遗留
build.out        旧的test记录，由Makefiles遗留
*.[568ao]        object文件，由Makefiles遗留

DIR(.exe)        由go build产生
DIR.test(.exe)   由go test -c产生
MAINFILE(.exe)   由go build MAINFILE.go产生
*.so             由 SWIG 产生
```