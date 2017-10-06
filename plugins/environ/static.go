package environ

import (
	"strings"

	"github.com/drone/drone/model"
)

type static struct {
	envs []*model.Environ
}

func NewStatic(list []string) model.EnvironService {
	envs := make([]*model.Environ, len(list))

	for i, raw := range list {
		parts := strings.Split(raw, "=")
		envs[i] = &model.Environ{int64(i), parts[0], parts[1]}
	}

	return &static{envs}
}

func (b *static) EnvironList(_ *model.Repo) ([]*model.Environ, error) {
	return b.envs, nil
}
