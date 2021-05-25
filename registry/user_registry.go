package registry

import (
	ip "golang-clean-architecture/adapters/presenters"
	ir "golang-clean-architecture/adapters/repositories"
	"golang-clean-architecture/core/controllers"
	"golang-clean-architecture/core/interactors"
	up "golang-clean-architecture/core/ports/presenters"
	ur "golang-clean-architecture/core/ports/repositories"
)

func (r *registry) NewUserController() controllers.UserController {
	return controllers.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactors.UserInteractor {
	return interactors.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
