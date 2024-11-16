syntax = "proto3";

package #{packageName};
option go_package = "./#{packageName}";

// Common
message IdReq {
    int64 Id = 1;   // id
}
message StringRes {
    string Resp = 1;    // 返回信息
}
message IdListReq {
    repeated int64 IdList = 1;  // id列表
}
message PageResult {
    int64 Page = 1;// 页码
    int64 PageSize  = 2;// 每页数量
    int64 Total = 3; // 总数
}

// #{tableName} #{tableComment} Copy to the main proto file
service #{tableName} {
    //查询#{tableComment}列表
    rpc Get#{TableName}List(#{TableName}Req) returns (#{TableName}ListRes);

    //查询#{tableComment}详情
    rpc Get#{TableName}Info (IdReq)returns (#{TableName}InfoRes);

    //新增#{tableComment}
    rpc Add#{TableName} (#{TableName}InfoRes) returns (StringRes);

    //删除#{tableComment}
    rpc Del#{TableName} (IdReq) returns (StringRes);

    //批量删除#{tableComment}
    rpc BatchDel#{TableName} (IdListReq) returns (StringRes);

    //修改#{tableComment}
    rpc Update#{TableName} (#{TableName}InfoRes) returns (StringRes);
}

message #{TableName}Req {
#{TableInfo}
#{TableInfoProtoPage}
}

message #{TableName}InfoRes {
#{TableInfo}
}

message #{TableName}ListRes {
    repeated #{TableName}InfoRes List = 1;  // #{tableComment}列表
    PageResult PageResult = 2;// 分页结果
}
