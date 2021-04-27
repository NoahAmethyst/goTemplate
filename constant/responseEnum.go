package constant

const (
	RESPONSE   = "response"
	REQUEST_ID = "requestId"

	SUCCESS           = 200
	PERMISSION_DENIED = 403
	SYSTEM_ERROR      = 500
	INVALID_PARAMETER = 50001
)

func GetResMsg(code int) string {
	resMsgMap := map[int]string{
		SUCCESS:           "",
		PERMISSION_DENIED: "permission denied",
		INVALID_PARAMETER: "invalid parameters",
	}
	return resMsgMap[code]
}
