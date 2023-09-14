package environment

import (
	"nginx-ui/server/internal/analytic"
	"nginx-ui/server/query"
)

func RetrieveEnvironmentList() (envs []*analytic.Node, err error) {
	envQuery := query.Environment

	data, err := envQuery.Find()
	if err != nil {
		return
	}

	for _, v := range data {
		t := analytic.GetNode(v)

		envs = append(envs, t)
	}

	return
}
