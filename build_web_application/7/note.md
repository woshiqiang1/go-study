### 文本处理

#### JSON处理
##### 解析JSON
解析到结构体
JSON数据如下：
```json
{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}
```
假如有了上面的JSON串，那么我们如何来解析这个JSON串呢？Go的JSON包中有如下函数：
```go
func Unmarshal(data []byte, v interface{}) error
```
通过这个函数我们就可以实现解析的目的，详细的解析例子请看如下代码：
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
    ServerName string
    ServerIP string
}

type ServerSlice struct {
    Servers []Server
}

func main() {
    var s ServerSlice
    str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
    json.Unmarshal([]byte(str), &s)
    fmt.Println(s)
}
```
在上面的示例代码中，我们首先定义了与json数据对应的结构体，数组对应slice，字段名对应JSON里面的KEY，在解析的时候，如何将json数据与struct字段相匹配呢？例如JSON的key是 `Foo` ，那么怎么找对应的字段呢？
- 首先查找tag含有 `Foo` 的可导出的struct字段(**首字母大写**)
- 其次查找字段名是 `Foo` 的导出字段
- 最后查找类似 `FOO` 或者 `FoO` 这样的除了首字母之外其他大小写不敏感的导出字段

聪明的你一定注意到了这一点：能够被赋值的字段必须是可导出字段(即首字母大写）。同时JSON解析的时候只会解析能找得到的字段，找不到的字段会被忽略，这样的一个好处是：当你接收到一个很大的JSON数据结构而你却只想获取其中的部分数据的时候，你只需将你想要的数据对应的字段名大写，即可轻松解决这个问题。

##### 解析到interface
上面那种解析方式是在我们知晓被解析的JSON数据的结构的前提下采取的方案，如果我们不知道被解析的数据的格式，又应该如何来解析呢？

我们知道interface{}可以用来存储任意数据类型的对象(类似TS的any类型)，这种数据结构正好用于存储解析的未知结构的json数据的结果。JSON包中采用map[string]interface{}和[]interface{}结构来存储任意的JSON对象和数组。Go类型和JSON类型的对应关系如下：
- bool    代表 JSON booleans
- float64 代表 JSON numbers
- string  代表 JSON strings
- nil     代表 JSON null
现在我们假设有如下的JSON数据：
```go
b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
```
如果在我们不知道他的结构的情况下，我们把他解析到interface{}里面：
```go
var f interface {}
err := json.Unmarshal(b, &f)
```
这个时候f里面存储了一个map类型，他们的key是string，值存储在空的interface{}里：
```go
f = map[string]interface{}{
    "Name": "Wednesday",
    "Age": 6,
    "Parents": []interface{}{
        "Gomez",
        "Morticia"
    }
}
```
那么如何来访问这些数据呢？通过断言的方式：
```go
m := f.(map[string]interface{})
```
通过断言之后，你就可以通过如下方式来访问里面的数据了：
```go
for k, v := range m {
    switch vv := v.(type) {
	case string:
		fmt.Println(k, "is string", vv)
	case int:
		fmt.Println(k, "is int", vv)
	case float64:
		fmt.Println(k,"is float64",vv)
	case []interface{}:
		fmt.Println(k, "is an array:")
		for i, u := range vv {
			fmt.Println(i, u)
		}
	default:
		fmt.Println(k, "is of a type I don't know how to handle")
	}
}
```
通过上面的示例可以看到，通过interface{}与type assert的配合，我们就可以解析未知结构的JSON数了。

上面这个是官方提供的解决方案，其实很多时候我们通过类型断言，操作起来不是很方便，目前bitly公司开源了一个叫做simplejson的包,在处理未知结构体的JSON时相当方便，详细例子如下所示：
```go
js, err := NewJson([]byte(`{
    "test": {
		"array": [1, "2", 3],
		"int": 10,
		"float": 5.150,
		"bignum": 9223372036854775807,
		"string": "simplejson",
		"bool": true
	}
}`))

arr, _ := js.Get("test").Get("array").Array()
i, _ := js.Get("test").Get("int").Int()
ms := js.Get("test").Get("string").MustString()
```
可以看到，使用这个库操作JSON比起官方包来说，简单的多,详细的请参考如下地址：https://github.com/bitly/go-simplejson

#### 生成JSON
我们开发很多应用的时候，最后都是要输出JSON数据串，那么如何来处理呢？JSON包里面通过 `Marshal` 函数来处理，函数定义如下：
```go
func Marshal(v interface{}) ([]byte, error)
```
假设我们还是需要生成上面的服务器列表信息，那么如何来处理呢？请看下面的例子：
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type ServerSlice struct {
	Servers []Server
}

func main() {
    var s ServerSlice
    s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
    b, error = json.Marshal(s)
    if err != nil {
        fmt.Println("json err:", err)
    }
    fmt.Println(string(b))
}
```
输出如下内容：
`{"Servers":[{"ServerName":"Shanghai_VPN","ServerIP":"127.0.0.1"},{"ServerName":"Beijing_VPN","ServerIP":"127.0.0.2"}]}`

我们看到上面的输出字段名的首字母都是大写的，如果你想用小写的首字母怎么办呢？把结构体的字段名改成首字母小写的？JSON输出的时候必须注意，只有**导出的字段(大写的)**才会被输出，如果修改字段名，那么就会发现什么都不会输出，所以必须通过struct tag定义来实现：
```go
type Server struct {
    ServerName string `json:"serverName"`
    serverIP string `json:"serverIP"`
}

type ServerSlice struct {
    Servers []Server `json:"servers"`
}
```
通过修改上面的结构体定义，输出的JSON串就和我们最开始定义的JSON串保持一致了。

针对JSON的输出，我们在定义struct tag的时候需要注意的几点是:
- 字段的tag是"-"，那么这个字段不会输出到JSON
- tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中，例如上面例子中serverName
- ag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
- 如果字段类型是bool, string, int, int64等，而tag中带有",string"选项，那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串
举例来说：
```go
type Server struct {
	// ID 不会导出到JSON中
	ID int `json:"-"`

	// ServerName2 的值会进行二次JSON编码
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`

	// 如果 ServerIP 为空，则不输出到JSON串中
	ServerIP   string `json:"serverIP,omitempty"`
}

s := Server {
	ID:         3,
	ServerName:  `Go "1.0" `,
	ServerName2: `Go "1.0" `,
	ServerIP:   ``,
}
b, _ := json.Marshal(s)
os.Stdout.Write(b)
```
会输出以下内容：
`{"serverName":"Go \"1.0\" ","serverName2":"\"Go \\\"1.0\\\" \""}`

Marshal函数只有在转换成功的时候才会返回数据，在转换的过程中我们需要注意几点：
- JSON对象只支持string作为key，所以要编码一个map，那么必须是map[string]T这种类型(T是Go语言中任意的类型)
- Channel, complex和function是不能被编码成JSON的
- 嵌套的数据是不能编码的，不然会让JSON编码进入死循环
- 指针在编码的时候会输出指针指向的内容，而空指针会输出null
本小节，我们介绍了如何使用Go语言的json标准包来编解码JSON数据，同时也简要介绍了如何使用第三方包go-simplejson来在一些情况下简化操作，学会并熟练运用它们将对我们接下来的Web开发相当重要。

#### 文件操作
