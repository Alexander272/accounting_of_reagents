package services

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
)

type PermissionService struct {
	enforcer casbin.IEnforcer
}

type Permission interface {
	// Register(confPath string, menu Menu) error
	Enforce(params ...interface{}) (bool, error)
}

// func NewPermissionService(confPath string, menu Menu) *PermissionService {
// 	permission := &PermissionService{}
// 	if err := permission.Register(confPath, menu); err != nil {
// 		log.Fatalf("failed to initialize permission service. error: %s", err.Error())
// 	}
// 	return permission
// }

// func (s *PermissionService) Register(path string, menu Menu) error {
// 	var err error
// 	adapter := NewPolicyAdapter(menu)

// 	s.enforcer, err = casbin.NewEnforcer(path, adapter)
// 	if err != nil {
// 		return fmt.Errorf("failed to create enforcer. error: %w", err)
// 	}

// 	if err = s.enforcer.LoadPolicy(); err != nil {
// 		return fmt.Errorf("failed to load policy. error: %w", err)
// 	}

// 	return nil
// }

// func (s *PermissionService) Enforce(params ...interface{}) (bool, error) {
// 	return s.enforcer.Enforce(params...)
// }

type PolicyAdapter struct {
	// menu Menu
}

func NewPolicyAdapter() *PolicyAdapter {
	return &PolicyAdapter{
		// menu: menu,
	}
}

type Adapter interface {
	LoadPolicy(model model.Model) error
	SavePolicy(model model.Model) error
	AddPolicy(sec string, ptype string, rule []string) error
	RemovePolicy(sec string, ptype string, rule []string) error
	RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error
}

func (s *PolicyAdapter) LoadPolicy(model model.Model) error {
	// menu, err := s.menu.GetAll(context.Background())
	// if err != nil {
	// 	return err
	// }

	// // lines := []string{}

	// // load api paths
	// for _, mf := range menu {
	// 	for _, mi := range mf.MenuItems {
	// 		for _, a := range mi.Api {
	// 			line := fmt.Sprintf("p, %s, %s, %s", mf.Role.Name, a.Path, a.Method)
	// 			// logger.Debug("api line ", line)
	// 			if err := persist.LoadPolicyLine(line, model); err != nil {
	// 				return fmt.Errorf("failed to load policy. error: %w", err)
	// 			}
	// 			// lines = append(lines, fmt.Sprintf("p, %s, %s, %s", mf.Role.Name, a.Path, a.Method))
	// 		}
	// 	}
	// }

	// // load users group (role)
	// for _, mf := range menu {
	// 	// lines = append(lines, fmt.Sprintf("g, %s, %s", mf.Role.Name, mf.Role.Extends))
	// 	if len(mf.Role.Extends) == 0 {
	// 		line := fmt.Sprintf("g, %s, ", mf.Role.Name)
	// 		// logger.Debug("role line ", line)
	// 		if err := persist.LoadPolicyLine(line, model); err != nil {
	// 			return fmt.Errorf("failed to load group policy. error: %w", err)
	// 		}
	// 	}

	// 	for _, v := range mf.Role.Extends {
	// 		line := fmt.Sprintf("g, %s, %s", mf.Role.Name, v)
	// 		// logger.Debug("role line ", line)
	// 		if err := persist.LoadPolicyLine(line, model); err != nil {
	// 			return fmt.Errorf("failed to load group policy. error: %w", err)
	// 		}
	// 	}
	// }

	return nil
}

// SavePolicy saves all policy rules to the storage.
func (s *PolicyAdapter) SavePolicy(model model.Model) error {
	return nil
}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (s *PolicyAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (s *PolicyAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (a *PolicyAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}
