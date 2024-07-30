package task

// SqlTask 定义 SQL 任务结构体
// SqlTask defines the SQL task structure
type SqlTask struct {
	AbstractTask // 嵌入 AbstractTask embedding AbstractTask

	Type           string   `json:"type"`           // 数据源类型，例如MYSQL、POSTGRES、HIVE ... datasource type, e.g., MYSQL, POSTGRES, HIVE ...
	Datasource     int      `json:"datasource"`     // 数据源 ID datasource ID
	SQL            string   `json:"sql"`            // SQL 语句 SQL statement
	SqlType        string   `json:"sqlType"`        // SQL 类型，0 查询 1 非查询 SQL type, 0 query 1 NON_QUERY
	SendEmail      bool     `json:"sendEmail"`      // 是否发送邮件 send email
	DisplayRows    int      `json:"displayRows"`    // 显示行数 display rows
	Udfs           string   `json:"udfs"`           // UDF 列表 udf list
	ConnParams     string   `json:"connParams"`     // SQL 连接参数 SQL connection parameters
	PreStatements  []string `json:"preStatements"`  // 前置语句 pre-statements
	PostStatements []string `json:"postStatements"` // 后置语句 post-statements
	GroupId        int      `json:"groupId"`        // 分组 ID group ID
	Title          string   `json:"title"`          // 标题 title
	Limit          int      `json:"limit"`          // 限制 limit
}

// NewSqlTask 创建 SqlTask 实例
// NewSqlTask creates a new instance of SqlTask
func NewSqlTask() *SqlTask {
	return &SqlTask{}
}

// GetTaskType 实现 Task 接口方法
// GetTaskType implements Task interface method
func (s *SqlTask) GetTaskType() string {
	return "SQL"
}
