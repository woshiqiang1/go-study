### 预防跨站脚本
Go的html/template里面带有下面几个函数可以帮你转义：
- func HTMLEscape(w io.Writer, b []byte) //把b进行转义之后写到w
- func HTMLEscapeString(s string) string //转义s之后返回结果字符串
- func HTMLEscaper(args ...interface{}) string //支持多个参数一起转义，返回结果字符串

### 处理文件上传
要使表单能够上传文件，首先第一步就是要添加form的enctype属性，enctype属性有如下三种情况:
> application/x-www-form-urlencoded   表示在发送前编码所有字符（默认）
> multipart/form-data	  不对字符编码。在使用包含文件上传控件的表单时，必须使用该值。
> text/plain	  空格转换为 "+" 加号，但不对特殊字符编码。

