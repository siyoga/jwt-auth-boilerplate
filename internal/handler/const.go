package handler

type CtxKey string

const (
	IpHeader = "X-Forwarded-For"

	HeaderContentType = "Content-Type"
	JsonContentType   = "application/json"
	AuthHeader        = "Authorization"

	TokenStart    = "Bearer "       // Префикс значения заголовка с авторизацией
	TokenStartInd = len(TokenStart) // Индекс, с которого в заголовке авторизации должен начинаться jwt токен

	AccessTokenCookie  = "access_token"
	RefreshTokenCookie = "refresh_token"

	AccountCtxKey     = CtxKey("acc")
	IpCtxKey          = CtxKey("ip")
	TokenNumberCtxKey = CtxKey("token")
)
