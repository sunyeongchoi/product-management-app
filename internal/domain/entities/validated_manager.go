package entities

type ValidatedManager struct {
	Manager
	isValidated bool
}

func (vm *ValidatedManager) IsValid() bool {
	return vm.isValidated
}

func NewValidatedManager(manager *Manager) (*ValidatedManager, error) {
	if err := manager.validate(); err != nil {
		return nil, err
	}
	return &ValidatedManager{
		Manager: *manager,
		isValidated: true,
	}, nil
}
