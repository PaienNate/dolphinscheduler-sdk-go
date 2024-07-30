package util

const (
	OfflineReleaseState = "OFFLINE" // 工作流下线状态
	OnlineReleaseState  = "ONLINE"  // 工作流上线状态
)

type Page struct {
	DefaultPage int
	DefaultSize int
}

var DefaultPage = Page{
	DefaultPage: 1,
	DefaultSize: 10,
}

type ExecuteType struct {
	ReRun string
	Stop  string
	Pause string
}

var DefaultExecuteType = ExecuteType{
	ReRun: "REPEAT_RUNNING",
	Stop:  "STOP",
	Pause: "PAUSE",
}

type HttpTask struct {
	HttpParameterTypeParameter string
	HttpParameterTypeBody      string
	HttpParameterTypeHeaders   string
	HttpMethodGet              string
	HttpMethodPost             string
	HttpMethodPut              string
}

var DefaultHttpTask = HttpTask{
	HttpParameterTypeParameter: "PARAMETER",
	HttpParameterTypeBody:      "BODY",
	HttpParameterTypeHeaders:   "HEADERS",
	HttpMethodGet:              "GET",
	HttpMethodPost:             "POST",
	HttpMethodPut:              "PUT",
}

type TaskType struct {
	Datax     string
	Http      string
	Condition string
	Shell     string
}

var DefaultTaskType = TaskType{
	Datax:     "DATAX",
	Http:      "HTTP",
	Condition: "CONDITIONS",
	Shell:     "SHELL",
}

type Resource struct {
	TypeFile          string
	TypeUdf           string
	DefaultPidFile    string
	DefaultCurrentDir string
}

var DefaultResource = Resource{
	TypeFile:          "FILE",
	TypeUdf:           "UDF",
	DefaultPidFile:    "-1",
	DefaultCurrentDir: "/",
}
