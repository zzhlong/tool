package httpclient

const File = 10
const Test = 20

type FromDataParams struct {
	ParamsType    int     //指参数类型 目前只有File、Test
	ParamsKey     string  //参数名
	ParamsContent *[]byte //参数内容
	FileName      string  //文件名称 如果ParamsType=File 需要传入文件名称
}

type ResponseBody struct {
	StateCode int
	Body      []byte
}
