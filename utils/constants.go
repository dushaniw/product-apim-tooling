package utils

const EnvConfigFileName string = "env_config.yaml"
const HeaderAuthorization string = "Authorization"
const HeaderContentType string = "Content-Type"
const HeaderAccept string = "Accept"
const HeaderProduces string = "Produces"
const HeaderConsumes string = "Consumes"
const HeaderValueApplicationZip = "application/zip"
const HeaderValueApplicationJSON string = "application/json"
const HeaderValueXWWWFormUrlEncoded string = "application/x-www-form-urlencoded"
const HeaderValueAuthBearerPrefix string = "Bearer"
const HeaderValueAuthBasicPrefix string = "Basic"
const HeaderValueMultiPartFormData string = "multipart/form-data"
const APICallTimeout = 5
const WSO2APIMUpdateTokenTimeout = 2
const DefaultTokenValidityPeriod string = "3600"
const LogPrefixInfo = "INFO: "
const LogPrefixWarning = "WARN: "
const LogPrefixError = "ERROR: "

const (
	BYTE     = 1.0
	KILOBYTE = 1024 * BYTE
	MEGABYTE = 1024 * KILOBYTE
	GIGABYTE = 1024 * MEGABYTE
)
