syntax = "v1"

@server (
    prefix: #{prefix}
)
service #{tableName} {
    //查询业务列表
    @handler Get#{TableName}List
    get /#{tableName}/list (#{TableName}Req) returns (#{TableName}ListRes)

    //查询业务详情
    @handler Get#{TableName}Info
    get /#{tableName}/info/:id returns (#{TableName}InfoRes)

    //新增业务
    @handler Add#{TableName}
    post /#{tableName}/add (#{TableName}Req) returns (string)

    //删除业务
    @handler Del#{TableName}
    delete /#{tableName}/delete/:id returns (string)

    //批量删除业务
    @handler BatchDel#{TableName}
    delete /#{tableName}/batchDelete (IdListReq) returns (string)

    //修改业务
    @handler Update#{TableName}
    put /#{tableName}/update (#{TableName}Req) returns (string)
}
type #{TableName}Req {
#{TableInfo}
    Page         int64  `json:"page,optional" form:"page,optional"`
    PageSize     int64  `json:"page_size,optional" form:"page_size,optional"`
}

type #{TableName}InfoRes {
#{TableInfo}
}

type #{TableName}ListRes {
    List       []#{TableName}InfoRes `json:"list,optional" form:"list,optional"`
    PageResult PageResult `json:"page_result,optional" form:"page_result,optional"`
}

type IdListReq {
    IdList []int64 `json:"id_list,optional" form:"id_list,optional"`
}

type PageResult {
    Page     int64 `json:"page,optional" form:"page,optional"`
    PageSize int64 `json:"page_size,optional" form:"page_size,optional"`
    Total    int64 `json:"total,optional" form:"total,optional"`
}
