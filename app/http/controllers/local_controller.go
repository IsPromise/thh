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
	repoList, _ := service.CountGitReposWithUnpushedCommits(workspace)
	sort.Slice(repoList, func(i, j int) bool {
		if repoList[i].HasCommits != repoList[j].HasCommits {
			// 如果有未推送的提交，则优先级高于没有未推送的提交
			return cast.ToInt(repoList[i].HasCommits) > cast.ToInt(repoList[j].HasCommits)
		} else if repoList[i].HasChanges != repoList[j].HasChanges {
			// 如果有变更但没有未推送的提交，则优先级高于没有变更
			return cast.ToInt(repoList[i].HasChanges) > cast.ToInt(repoList[j].HasChanges)
		} else {
			// 如果都没有未推送的提交，也没有变更，则按路径名降序排序
			return repoList[i].Path < repoList[j].Path
		}
	})
	return component.SuccessResponse(repoList)
}
