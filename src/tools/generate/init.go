// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package generate

import (
	"github.com/hexya-erp/hexya/src/tools/logging"
)

const (
	// HexyaPath is the go import path of the base hexya package
	HexyaPath = "github.com/hexya-erp/hexya"
	// ModelsPath is the go import path of the hexya/models package
	ModelsPath = "github.com/hexya-erp/hexya/src/models"
	// DatesPath is the go import path of the hexya/models/types/dates package
	DatesPath = "github.com/hexya-erp/hexya/src/models/types/dates"
	// PoolPath is the go import path of the autogenerated pool package
	PoolPath = "github.com/hexya-erp/pool"
	// PoolModelPackage is the name of the pool package with model data
	PoolModelPackage = "h"
	// PoolQueryPackage is the name of the pool package with query dat
	PoolQueryPackage = "q"
)

var (
	log logging.Logger
	// ModelMixins are the names of the mixins declared in the models package
	ModelMixins = map[string]bool{
		"CommonMixin":    true,
		"BaseMixin":      true,
		"ModelMixin":     true,
		"TransientMixin": true,
	}
)

func init() {
	log = logging.GetLogger("tools/generate")
}
