package async

type TypeAsync int

const (
	RABBITMQ TypeAsync = iota
)

type Top []TopAddr

type TopAddr struct {
	Addr  string `json:"addr"`
	Count int    `json:"count"`
}
