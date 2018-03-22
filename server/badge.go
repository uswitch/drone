// Copyright 2018 Drone.IO Inc.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//      http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"

	"github.com/drone/drone/model"
	"github.com/drone/drone/shared/httputil"
	"github.com/drone/drone/store"
)

var (
	badgeSuccessUrl = `https://img.shields.io/badge/drone.io-success-brightgreen.svg`
	badgeFailureUrl = `https://img.shields.io/badge/drone.io-failure-red.svg`
	badgeStartedUrl = `https://img.shields.io/badge/drone.io-started-yellow.svg`
	badgeErrorUrl   = `https://img.shields.io/badge/drone.io-error-lightgrey.svg`
	badgeNoneUrl    = `https://img.shields.io/badge/drone.io-none-lightgrey.svg`
)

func GetBadge(c *gin.Context) {
	repo, err := store.GetRepoOwnerName(c,
		c.Param("owner"),
		c.Param("name"),
	)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	// if no commit was found then display
	// the 'none' badge, instead of throwing
	// an error response
	branch := c.Query("branch")
	if len(branch) == 0 {
		branch = repo.Branch
	}

	repoUrl := fmt.Sprintf("%s/%s", httputil.GetURL(c.Request), repo.FullName)
	queryString := fmt.Sprintf("link=%s&link=%s&%s", repoUrl, repoUrl, c.Request.URL.RawQuery)
	build, err := store.GetBuildLast(c, repo, branch)
	if err != nil {
		log.Warning(err)
		c.Redirect(302, badgeNoneUrl+"?"+queryString)
		return
	}

	switch build.Status {
	case model.StatusSuccess:
		c.Redirect(302, badgeSuccessUrl+"?"+queryString)
	case model.StatusFailure:
		c.Redirect(302, badgeFailureUrl+"?"+queryString)
	case model.StatusError, model.StatusKilled:
		c.Redirect(302, badgeErrorUrl+"?"+queryString)
	case model.StatusPending, model.StatusRunning:
		c.Redirect(302, badgeStartedUrl+"?"+queryString)
	default:
		c.Redirect(302, badgeNoneUrl+"?"+queryString)
	}
}

func GetCC(c *gin.Context) {
	repo, err := store.GetRepoOwnerName(c,
		c.Param("owner"),
		c.Param("name"),
	)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	builds, err := store.GetBuildList(c, repo, 1)
	if err != nil || len(builds) == 0 {
		c.AbortWithStatus(404)
		return
	}

	url := fmt.Sprintf("%s/%s/%d", httputil.GetURL(c.Request), repo.FullName, builds[0].Number)
	cc := model.NewCC(repo, builds[0], url)
	c.XML(200, cc)
}
