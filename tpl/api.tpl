syntax = "v1"

@server (
    prefix: #{prefix}
)
service #{tableName} {
    //查询#{tableComment}列表
    @handler Get#{TableName}List
    get /#{tableName}/list (#{TableName}Req) returns (#{TableName}ListRes)

    //查询#{tableComment}详情
    @handler Get#{TableName}Info
    get /#{tableName}/info/:id returns (#{TableName}InfoRes)

    //新增#{tableComment}
    @handler Add#{TableName}
    post /#{tableName}/add (#{TableName}Req) returns (string)

    //删除#{tableComment}
    @handler Del#{TableName}
    delete /#{tableName}/delete/:id returns (string)

    //批量删除#{tableComment}
    @handler BatchDel#{TableName}
    delete /#{tableName}/batchDelete (IdListReq) returns (string)

    //修改#{tableComment}
    @handler Update#{TableName}
    put /#{tableName}/update (#{TableName}Req) returns (string)
}
type #{TableName}Req {
#{TableInfo}
    Page         int64  `json:"page,optional" form:"page,optional"` // 页码
    PageSize     int64  `json:"page_size,optional" form:"page_size,optional"` // 每页数量
}

type #{TableName}InfoRes {
#{TableInfo}
}

type #{TableName}ListRes {
    List       []#{TableName}InfoRes `json:"list,optional" form:"list,optional"` // #{tableComment}列表
    PageResult PageResult `json:"page_result,optional" form:"page_result,optional"` // 分页结果
}

type IdListReq {
    IdList []int64 `json:"id_list,optional" form:"id_list,optional"` // id列表
}

type PageResult {
    Page     int64 `json:"page,optional" form:"page,optional"` // 页码
    PageSize int64 `json:"page_size,optional" form:"page_size,optional"` // 每页数量
    Total    int64 `json:"total,optional" form:"total,optional"` // 总数
}
