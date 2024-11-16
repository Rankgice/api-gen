syntax = "v1"

import "../common.api"
@server (
    prefix: #{prefix}
    group: #{tableName}
)
service #{serviceName} {
    //查询#{tableComment}列表
    @handler Get#{TableName}List
    get /#{tableName}/list (#{TableName}Req) returns (#{TableName}ListRes)

    //查询#{tableComment}详情
    @handler Get#{TableName}Info
    get /#{tableName}/info/:id (IdReq) returns (#{TableName}InfoRes)

    //新增#{tableComment}
    @handler Add#{TableName}
    post /#{tableName}/add (#{TableName}InfoRes) returns (string)

    //删除#{tableComment}
    @handler Del#{TableName}
    delete /#{tableName}/delete/:id (IdReq) returns (string)

    //批量删除#{tableComment}
    @handler BatchDel#{TableName}
    delete /#{tableName}/batchDelete (IdListReq) returns (string)

    //修改#{tableComment}
    @handler Update#{TableName}
    put /#{tableName}/update (#{TableName}InfoRes) returns (string)
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