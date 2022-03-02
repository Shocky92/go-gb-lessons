module lesson8

go 1.17

require gopkg.in/yaml.v2 v2.4.0
require config v1.0.0

replace (
	config v1.0.0 => ./config
)