package constants

const (
	//HeaderOriginKey header parameter
	HeaderOriginKey string = "Access-Control-Allow-Origin"

	//HeaderOriginValue header parameter
	HeaderOriginValue string = "*"

	//HeaderMethodsKey header parameter
	HeaderMethodsKey string = "Access-Control-Allow-Methods"

	//HeaderMethodsValue header parameter
	HeaderMethodsValue string = "POST, GET, OPTIONS, PUT, DELETE"

	//HeaderNamesKey header parameter
	HeaderNamesKey string = "Access-Control-Allow-Headers"

	//HeaderNamesValue header parameter
	HeaderNamesValue string = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-user-agent, x-grpc-web, grpc-status, grpc-message"

	//HeaderExposeKey header parameter
	HeaderExposeKey string = "Access-Control-Expose-Headers"

	//HeaderExposeValue header parameter
	HeaderExposeValue string = "grpc-status, grpc-message"

	//RMethodOptions name
	RMethodOptions string = "OPTIONS"
)
