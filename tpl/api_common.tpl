syntax = "v1"

type (
    IdReq {
        Id int64 `path:"id"` // id
    }
    IdListReq {
        IdList []int64 `json:"id_list,optional" form:"id_list,optional"` // id列表
    }
    PageResult {
        Page     int64 `json:"page,optional" form:"page,optional"` // 页码
        PageSize int64 `json:"page_size,optional" form:"page_size,optional"` // 每页数量
        Total    int64 `json:"total,optional" form:"total,optional"` // 总数
    }
)