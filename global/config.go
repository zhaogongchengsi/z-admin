package global

type Service struct {
	Port         int
	Mode         string
	ReadTimeout  int
	WriteTimeout int
}

type DataBase struct {
	DbName        string
	UserName      string
	PassWord      string
	Url           string
	Port          string
	MaxIdleConns  int
	MaxOpenConns  int
	TablePrefix   string
	SingularTable bool
	Charset       string
}

type Captcha struct {
	ImgWidth  int
	ImgHeight int
	KeyLong   int
	MaxSkew   float64
	DotCount  int
}

type JWT struct {
	SigningKey string
	ExpiresAt  int
	Issuer     string
}
