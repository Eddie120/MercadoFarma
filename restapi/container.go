package restapi

import (
	"github.com/mercadofarma/services/controllers"
	"github.com/mercadofarma/services/db/mysql"
	"github.com/mercadofarma/services/repos/business"
	"github.com/mercadofarma/services/repos/users"
	businessService "github.com/mercadofarma/services/services/business"
	userService "github.com/mercadofarma/services/services/users"
	"go.uber.org/dig"
	"os"
)

func buildContainer() *dig.Container {
	container := dig.New()

	provider := func(constructor interface{}, opt ...dig.ProvideOption) {
		err := container.Provide(constructor, opt...)
		if err != nil {
			panic(err)
		}
	}

	provider(func() mysql.DataAccess {
		dataSource := os.Getenv("ENDPOINT_URL")
		if dataSource == "" {
			dataSource = "root:@tcp(127.0.0.1:3306)/mercadofarma" // localhost
		}

		const driverName = "mysql"

		db, err := mysql.CreateDBConnection(driverName, dataSource)
		if err != nil {
			panic(err)
		}

		return mysql.NewDataAccess(db, driverName, dataSource)
	})

	// repos
	provider(func(db mysql.DataAccess) users.UserRepo {
		return users.NewUserRepo(db)
	})
	provider(func(db mysql.DataAccess) business.BusinessRepo {
		return business.NewBusinessRepo(db)
	})

	// services
	provider(func(userRepo users.UserRepo) userService.UserService {
		return userService.NewUserService(userRepo)
	})
	provider(func(businessRepo business.BusinessRepo, userService userService.UserService) businessService.BusinessService {
		return businessService.NewBusinessService(businessRepo, userService)
	})

	// controllers
	provider(func(userService userService.UserService, businessService businessService.BusinessService) *controllers.BusinessController {
		return controllers.NewBusinessController(userService, businessService)
	})

	return container
}
