package es

//https://www.elastic.co/guide/en/elasticsearch/reference/7.17/common-options.html
type CommonOption struct {
	Pretty bool
	Human  bool
	//所有 REST API 都接受一个 filter_path 参数，该参数可用于减少 Elasticsearch 返回的响应。此参数采用逗号分隔的过滤器列表，用点符号表示
	// 它还支持 * 通配符匹配任何字段或字段名称的一部分：
	FilterPath []string
	// flat_settings 标志影响设置列表的呈现。当 flat_settings 标志为真时，设置以平面格式返回：
	// "index.number_of_replicas": "1",
	//FlatSettings bool

	//	默认情况下，当请求返回错误时，Elasticsearch 不包括错误的堆栈跟踪。您可以通过将 error_trace url 参数设置为 true 来启用该行为。
	ErrorTrace bool
}

type CatParam struct {
	CommonOption
	Local               bool
	MasterTimeout       string
	Sort                []string
	DisableTimestamping bool
	Columns             []string
}
