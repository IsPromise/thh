package controllers

import (
	"github.com/leancodebox/goose/preferences"
	"github.com/spf13/cast"
	"sort"
	"thh/app/http/controllers/component"
	"thh/app/service"
)

func GitStatusList() component.Response {
	workspace := preferences.Get("path.workspace")
	sL, _ := service.CountGitReposWithUnpushedCommits(workspace)
	sort.Slice(sL, func(i, j int) bool {
		return cast.ToInt(sL[i].HasCommits)*10+cast.ToInt(sL[i].HasChanges) > cast.ToInt(sL[j].HasCommits)*10+cast.ToInt(sL[j].HasChanges)
	})
	return component.SuccessResponse(sL)
}
