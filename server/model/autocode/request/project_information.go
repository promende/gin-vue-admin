package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ProjectInformationSearch struct{
    autocode.ProjectInformation
    request.PageInfo
}