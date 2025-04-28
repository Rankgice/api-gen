syntax = "v1"

import "../common.api"
@server (
    prefix: #{prefix}/#{tableName}
    group: #{tableName}
)
service #{serviceName} {
    @doc "查询#{tableComment}列表"
    @handler Get#{TableName}List
    get / (#{TableName}Req) returns (#{TableName}ListRes)

    @doc "查询#{tableComment}详情"
    @handler Get#{TableName}Info
    get /:id (IdReq) returns (#{TableName}InfoRes)

    @doc "新增#{tableComment}"
    @handler Add#{TableName}
    post / (#{TableName}InfoRes) returns (string)

    @doc "删除#{tableComment}"
    @handler Del#{TableName}
    delete /:id (IdReq) returns (string)

    @doc "批量删除#{tableComment}"
    @handler BatchDel#{TableName}
    delete /batch (IdListReq) returns (string)

    @doc "修改#{tableComment}"
    @handler Update#{TableName}
    put / (#{TableName}InfoRes) returns (string)
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