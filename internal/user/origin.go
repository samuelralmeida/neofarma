package user

type Origin string

var (
	Internal Origin = "internal"
	External Origin = "external"
)

var origins = map[string]Origin{
	string(Internal): Internal,
	string(External): External,
}

func GetOriginByName(originName string) (Origin, bool) {
	o, exists := origins[originName]
	return o, exists
}
