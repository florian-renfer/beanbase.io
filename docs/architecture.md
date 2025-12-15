## Architecture guid for Beanbase

Beanbase uses a hexagonal architecture for the RESTful API implmented using Go.
Therfore, outer layers can rely on inner layers but not the other way around.
To enable the application to abstract dependencies in such way, interfaces
play a huge role.

### Domain Layer

The domain layer defines core entities, constants, errors and interfaces
required to persist the data without providing an implementation.
The actual implementation is treated as detail and is part of the adapters.

### Usecase Layer

Usecases themselves are just interfaces that define an Execute function with
proper parameters and return values for the given context.

Basically every usecase requires an input format, which is defined as
UsecaseInput, and an output format, which is defined as UsecaseOutput, and
provided by a UsecasePresenter. The UsecasePresenter is responsible for
mapping the given data to the required  output format and is implemented
as an adapter. This results in UsecasePresenter being an interface.
Yet, an implementation of the usecase is required. This is achieved using a
dedicated interactor, that is only visible to the package.
