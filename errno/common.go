package errno

import "sync"

var (
	errorMap            = new(sync.Map)
	OK                  = add(0, "SUCCESS")
	CookieErr           = add(302, "Cookie 失效")
	RequestErr          = add(400, "请求错误")
	Unauthorized        = add(401, "未授权")
	NothingFound        = add(404, "啥都木有")
	CodeExpired         = add(420, "Code过期")
	InvalidCode         = add(421, "无效的Code")
	InvalidClient       = add(422, "无效的Client")
	InvalidRefreshToken = add(422, "无效的RefreshToken")
	ServerErr           = add(500, "服务器错误")
	ServerBusy          = add(502, "服务器忙，请稍微尝试")

	// todo some error code and msg to add
)
