package service

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/HotHat/gin-admin/v10/internal/config"
	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/entity"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/repo"
	"github.com/HotHat/gin-admin/v10/internal/mods/rbac/schema"
	"github.com/HotHat/gin-admin/v10/pkg/cachex"
	"github.com/HotHat/gin-admin/v10/pkg/encoding/json"
	"github.com/HotHat/gin-admin/v10/pkg/encoding/yaml"
	"github.com/HotHat/gin-admin/v10/pkg/errors"
	"github.com/HotHat/gin-admin/v10/pkg/logging"
	"github.com/HotHat/gin-admin/v10/pkg/util"
	"go.uber.org/zap"
)

// MenuService management for RBAC
type MenuService struct {
	Cache            cachex.Cacher
	Trans            *util.Trans
	MenuRepo         *repo.MenuRepo
	MenuResourceRepo *repo.MenuResourceRepo
	RoleMenuRepo     *repo.RoleMenuRepo
}

func (a *MenuService) InitFromFile(ctx context.Context, menuFile string) error {
	f, err := os.ReadFile(menuFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			logging.Context(ctx).Warn("MenuService data file not found, skip init menu data from file", zap.String("file", menuFile))
			return nil
		}
		return err
	}

	var menus entity.Menus
	if ext := filepath.Ext(menuFile); ext == ".json" {
		if err := json.Unmarshal(f, &menus); err != nil {
			return errors.Wrapf(err, "Unmarshal JSON file '%s' failed", menuFile)
		}
	} else if ext == ".yaml" || ext == ".yml" {
		if err := yaml.Unmarshal(f, &menus); err != nil {
			return errors.Wrapf(err, "Unmarshal YAML file '%s' failed", menuFile)
		}
	} else {
		return errors.Errorf("Unsupported file type '%s'", ext)
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.createInBatchByParent(ctx, menus, nil)
	})
}

func (a *MenuService) createInBatchByParent(ctx context.Context, items entity.Menus, parent *entity.Menu) error {
	total := len(items)

	for i, item := range items {
		var parentID comm.ID
		if parent != nil {
			parentID = parent.ID
		}

		var (
			menuItem *entity.Menu
			err      error
		)

		if item.ID != 0 {
			menuItem, err = a.MenuRepo.Get(ctx, item.ID)
		} else if item.Code != "" {
			menuItem, err = a.MenuRepo.GetByCodeAndParentID(ctx, item.Code, parentID)
		} else if item.Name != "" {
			menuItem, err = a.MenuRepo.GetByNameAndParentID(ctx, item.Name, parentID)
		}

		if err != nil {
			return err
		}

		if item.Status == 0 {
			item.Status = entity.MenuStatusEnabled
		}

		if menuItem != nil {
			changed := false
			if menuItem.Name != item.Name {
				menuItem.Name = item.Name
				changed = true
			}
			if menuItem.Description != item.Description {
				menuItem.Description = item.Description
				changed = true
			}
			if menuItem.Path != item.Path {
				menuItem.Path = item.Path
				changed = true
			}
			if menuItem.Type != item.Type {
				menuItem.Type = item.Type
				changed = true
			}
			if menuItem.Sequence != item.Sequence {
				menuItem.Sequence = item.Sequence
				changed = true
			}
			if menuItem.Status != item.Status {
				menuItem.Status = item.Status
				changed = true
			}
			if changed {
				menuItem.UpdatedAt = time.Now()
				if err := a.MenuRepo.Update(ctx, menuItem); err != nil {
					return err
				}
			}
		} else {
			//if item.ID == 0 {
			//	item.ID = util.NewXID()
			//}
			if err := a.MenuRepo.Create(ctx, item); err != nil {
				return err
			}
			if item.Sequence == 0 {
				item.Sequence = total - i
			}
			item.ParentID = parentID
			parentIDStr := comm.IDToStr(parentID)
			if parent != nil {
				item.ParentPath = parent.ParentPath + parentIDStr + util.TreePathDelimiter
			}
			menuItem = item
		}

		for _, res := range item.Resources {
			if res.ID != 0 {
				exists, err := a.MenuResourceRepo.Exists(ctx, res.ID)
				if err != nil {
					return err
				} else if exists {
					continue
				}
			}

			if res.Path != "" {
				exists, err := a.MenuResourceRepo.ExistsMethodPathByMenuID(ctx, res.Method, res.Path, menuItem.ID)
				if err != nil {
					return err
				} else if exists {
					continue
				}
			}
			//if res.ID == "" {
			//	res.ID = util.NewXID()
			//}
			res.MenuID = menuItem.ID
			if err := a.MenuResourceRepo.Create(ctx, res); err != nil {
				return err
			}
		}

		if item.Children != nil {
			if err := a.createInBatchByParent(ctx, *item.Children, menuItem); err != nil {
				return err
			}
		}
	}
	return nil
}

// Query menus from the data access object based on the provided parameters and options.
func (a *MenuService) Query(ctx context.Context, params dto.MenuQueryParam) (*dto.MenuQueryResult, error) {
	params.Pagination = false

	if err := a.fillQueryParam(ctx, &params); err != nil {
		return nil, err
	}

	result, err := a.MenuRepo.Query(ctx, params, dto.MenuQueryOptions{
		QueryOptions: util.QueryOptions{
			OrderFields: schema.MenusOrderParams,
		},
	})
	if err != nil {
		return nil, err
	}

	if params.LikeName != "" || params.CodePath != "" {
		result.Data, err = a.appendChildren(ctx, result.Data)
		if err != nil {
			return nil, err
		}
	}

	if params.IncludeResources {
		for i, item := range result.Data {
			resResult, err := a.MenuResourceRepo.Query(ctx, dto.MenuResourceQueryParam{
				MenuID: comm.IDToStr(item.ID),
			})
			if err != nil {
				return nil, err
			}
			result.Data[i].Resources = resResult.Data
		}
	}

	result.Data = result.Data.ToTree()
	return result, nil
}

func (a *MenuService) fillQueryParam(ctx context.Context, params *dto.MenuQueryParam) error {
	if params.CodePath != "" {
		var (
			codes    []string
			lastMenu entity.Menu
		)
		for _, code := range strings.Split(params.CodePath, util.TreePathDelimiter) {
			if code == "" {
				continue
			}
			codes = append(codes, code)
			menu, err := a.MenuRepo.GetByCodeAndParentID(ctx, code, lastMenu.ParentID, dto.MenuQueryOptions{
				QueryOptions: util.QueryOptions{
					SelectFields: []string{"id", "parent_id", "parent_path"},
				},
			})
			if err != nil {
				return err
			} else if menu == nil {
				return errors.NotFound("", "MenuService not found by code '%s'", strings.Join(codes, util.TreePathDelimiter))
			}
			lastMenu = *menu
		}
		lastMenuIDStr := comm.IDToStr(lastMenu.ID)
		params.ParentPathPrefix = lastMenu.ParentPath + lastMenuIDStr + util.TreePathDelimiter
	}
	return nil
}

func (a *MenuService) appendChildren(ctx context.Context, data entity.Menus) (entity.Menus, error) {
	if len(data) == 0 {
		return data, nil
	}

	existsInData := func(id comm.ID) bool {
		for _, item := range data {
			if item.ID == id {
				return true
			}
		}
		return false
	}

	for _, item := range data {
		idStr := comm.IDToStr(item.ID)
		childResult, err := a.MenuRepo.Query(ctx, dto.MenuQueryParam{
			ParentPathPrefix: item.ParentPath + idStr + util.TreePathDelimiter,
		})
		if err != nil {
			return nil, err
		}
		for _, child := range childResult.Data {
			if existsInData(child.ID) {
				continue
			}
			data = append(data, child)
		}
	}

	if parentIDs := data.SplitParentIDs(); len(parentIDs) > 0 {
		parentResult, err := a.MenuRepo.Query(ctx, dto.MenuQueryParam{
			InIDs: comm.IDArrToStr(parentIDs),
		})
		if err != nil {
			return nil, err
		}
		for _, p := range parentResult.Data {
			if existsInData(p.ID) {
				continue
			}
			data = append(data, p)
		}
	}
	sort.Sort(data)

	return data, nil
}

// Get the specified menu from the data access object.
func (a *MenuService) Get(ctx context.Context, id comm.ID) (*entity.Menu, error) {
	menu, err := a.MenuRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if menu == nil {
		return nil, errors.NotFound("", "MenuService not found")
	}

	idStr := comm.IDToStr(menu.ID)
	menuResResult, err := a.MenuResourceRepo.Query(ctx, dto.MenuResourceQueryParam{
		MenuID: idStr,
	})
	if err != nil {
		return nil, err
	}
	menu.Resources = menuResResult.Data

	return menu, nil
}

// Create a new menu in the data access object.
func (a *MenuService) Create(ctx context.Context, formItem *dto.MenuForm) (*entity.Menu, error) {
	if config.C.General.DenyOperateMenu {
		return nil, errors.BadRequest("", "MenuService creation is not allowed")
	}

	menu := &entity.Menu{
		CreatedAt: time.Now(),
	}

	if parentID := formItem.ParentID; parentID != 0 {
		parent, err := a.MenuRepo.Get(ctx, parentID)
		if err != nil {
			return nil, err
		} else if parent == nil {
			return nil, errors.NotFound("", "Parent not found")
		}
		idStr := comm.IDToStr(parent.ID)
		menu.ParentPath = parent.ParentPath + idStr + util.TreePathDelimiter
	}

	if exists, err := a.MenuRepo.ExistsCodeByParentID(ctx, formItem.Code, formItem.ParentID); err != nil {
		return nil, err
	} else if exists {
		return nil, errors.BadRequest("", "MenuService code already exists at the same level")
	}

	if err := formItem.FillTo(menu); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.MenuRepo.Create(ctx, menu); err != nil {
			return err
		}

		for _, res := range formItem.Resources {
			//res.ID = util.NewXID()
			res.MenuID = menu.ID
			res.CreatedAt = time.Now()
			if err := a.MenuResourceRepo.Create(ctx, res); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return menu, nil
}

// Update the specified menu in the data access object.
func (a *MenuService) Update(ctx context.Context, id comm.ID, formItem *dto.MenuForm) error {
	if config.C.General.DenyOperateMenu {
		return errors.BadRequest("", "MenuService update is not allowed")
	}

	menu, err := a.MenuRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if menu == nil {
		return errors.NotFound("", "MenuService not found")
	}

	oldParentPath := menu.ParentPath
	oldStatus := menu.Status
	var childData entity.Menus

	if menu.ParentID != formItem.ParentID {
		if parentID := formItem.ParentID; parentID != 0 {
			parent, err := a.MenuRepo.Get(ctx, parentID)
			if err != nil {
				return err
			} else if parent == nil {
				return errors.NotFound("", "Parent not found")
			}
			idStr := comm.IDToStr(parent.ID)
			menu.ParentPath = parent.ParentPath + idStr + util.TreePathDelimiter
		} else {
			menu.ParentPath = ""
		}

		idStr := comm.IDToStr(menu.ID)
		childResult, err := a.MenuRepo.Query(ctx, dto.MenuQueryParam{
			ParentPathPrefix: oldParentPath + idStr + util.TreePathDelimiter,
		}, dto.MenuQueryOptions{
			QueryOptions: util.QueryOptions{
				SelectFields: []string{"id", "parent_path"},
			},
		})
		if err != nil {
			return err
		}
		childData = childResult.Data
	}

	if menu.Code != formItem.Code {
		if exists, err := a.MenuRepo.ExistsCodeByParentID(ctx, formItem.Code, formItem.ParentID); err != nil {
			return err
		} else if exists {
			return errors.BadRequest("", "MenuService code already exists at the same level")
		}
	}

	if err := formItem.FillTo(menu); err != nil {
		return err
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if oldStatus != formItem.Status {
			idStr := comm.IDToStr(menu.ID)
			oldPath := oldParentPath + idStr + util.TreePathDelimiter
			if err := a.MenuRepo.UpdateStatusByParentPath(ctx, oldPath, formItem.Status); err != nil {
				return err
			}
		}

		for _, child := range childData {
			idStr := comm.IDToStr(child.ID)
			oldPath := oldParentPath + idStr + util.TreePathDelimiter
			newPath := menu.ParentPath + idStr + util.TreePathDelimiter
			err := a.MenuRepo.UpdateParentPath(ctx, child.ID, strings.Replace(child.ParentPath, oldPath, newPath, 1))
			if err != nil {
				return err
			}
		}

		if err := a.MenuRepo.Update(ctx, menu); err != nil {
			return err
		}

		if err := a.MenuResourceRepo.DeleteByMenuID(ctx, id); err != nil {
			return err
		}
		for _, res := range formItem.Resources {
			//if res.ID == "" {
			//	res.ID = util.NewXID()
			//}
			res.MenuID = id
			if res.CreatedAt.IsZero() {
				res.CreatedAt = time.Now()
			}
			res.UpdatedAt = time.Now()
			if err := a.MenuResourceRepo.Create(ctx, res); err != nil {
				return err
			}
		}

		return a.syncToCasbin(ctx)
	})
}

// Delete the specified menu from the data access object.
func (a *MenuService) Delete(ctx context.Context, id comm.ID) error {
	if config.C.General.DenyOperateMenu {
		return errors.BadRequest("", "MenuService deletion is not allowed")
	}

	menu, err := a.MenuRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if menu == nil {
		return errors.NotFound("", "MenuService not found")
	}

	idStr := comm.IDToStr(menu.ID)
	childResult, err := a.MenuRepo.Query(ctx, dto.MenuQueryParam{
		ParentPathPrefix: menu.ParentPath + idStr + util.TreePathDelimiter,
	}, dto.MenuQueryOptions{
		QueryOptions: util.QueryOptions{
			SelectFields: []string{"id"},
		},
	})
	if err != nil {
		return err
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.delete(ctx, id); err != nil {
			return err
		}

		for _, child := range childResult.Data {
			if err := a.delete(ctx, child.ID); err != nil {
				return err
			}
		}

		return a.syncToCasbin(ctx)
	})
}

func (a *MenuService) delete(ctx context.Context, id comm.ID) error {
	if err := a.MenuRepo.Delete(ctx, id); err != nil {
		return err
	}
	if err := a.MenuResourceRepo.DeleteByMenuID(ctx, id); err != nil {
		return err
	}
	if err := a.RoleMenuRepo.DeleteByMenuID(ctx, id); err != nil {
		return err
	}
	return nil
}

func (a *MenuService) syncToCasbin(ctx context.Context) error {
	return a.Cache.Set(ctx, config.CacheNSForRole, config.CacheKeyForSyncToCasbin, fmt.Sprintf("%d", time.Now().Unix()))
}
