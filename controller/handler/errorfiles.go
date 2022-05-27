// Copyright 2019 HAProxy Technologies LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handler

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/haproxytech/models"

	config "github.com/haproxytech/kubernetes-ingress/controller/configuration"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/api"
	"github.com/haproxytech/kubernetes-ingress/controller/store"
)

type ErrorFiles struct {
	files files
}

func (h *ErrorFiles) Update(k store.K8s, cfg *config.ControllerCfg, api api.HAProxyClient) (reload bool, err error) {
	h.files.dir = cfg.Env.ErrFileDir
	if k.ConfigMaps.Errorfiles == nil {
		return false, nil
	}
	// Update Files
	for code, content := range k.ConfigMaps.Errorfiles.Annotations {
		logger.Error(h.writeFile(code, content))
	}
	var apiInput []*models.Errorfile
	apiInput, reload = h.refresh()
	// Update API
	defaults, err := api.DefaultsGetConfiguration()
	if err != nil {
		return false, err
	}
	defaults.ErrorFiles = apiInput
	if err = api.DefaultsPushConfiguration(*defaults); err != nil {
		return false, err
	}
	return reload, nil
}

func (h *ErrorFiles) writeFile(code, content string) (err error) {
	// Update file
	if _, ok := h.files.data[code]; !ok {
		err = checkCode(code)
		if err != nil {
			return
		}
	}
	err = h.files.writeFile(code, content)
	if err != nil {
		err = fmt.Errorf("failed writing errorfile for code '%s': %w", code, err)
	}
	return
}

func (h *ErrorFiles) refresh() (result []*models.Errorfile, reload bool) {
	for code, f := range h.files.data {
		if !f.inUse {
			reload = true
			err := h.files.deleteFile(code)
			if err != nil {
				logger.Errorf("failed deleting errorfile for code '%s': %s", code, err)
			}
			continue
		}
		if f.updated {
			logger.Debugf("updating errorfile for code '%s': reload required", code)
			reload = true
		}
		c, _ := strconv.Atoi(code) // code already checked in newCode
		result = append(result, &models.Errorfile{
			Code: int64(c),
			File: filepath.Join(h.files.dir, code),
		})
		f.inUse = false
		f.updated = false
	}
	return
}

func checkCode(code string) error {
	var codes = [15]string{"200", "400", "401", "403", "404", "405", "407", "408", "410", "425", "429", "500", "502", "503", "504"}
	var c string
	for _, c = range codes {
		if code == c {
			break
		}
	}
	if c != code {
		return fmt.Errorf("HTTP error code '%s' not supported", code)
	}
	return nil
}
