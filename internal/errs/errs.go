package errs

import "errors"

var (
	ErrUserNotFound                = errors.New("пользователь не найден")
	ErrUserAlreadyExists           = errors.New("пользователь уже существует")
	ErrIncorrectUsernameOrPassword = errors.New("неверный email или пароль")
	ErrNoAccessRights              = errors.New("недостаточно прав для выполнения действия")
	ErrInvalidRole                 = errors.New("недопустимая роль пользователя")
	ErrBusNotFound                 = errors.New("автобус не найден")
	ErrBusAlreadyExists            = errors.New("автобус с таким номером уже существует")
	ErrMaintenanceNotFound         = errors.New("запись техобслуживания не найдена")
	ErrInvalidMaintenance          = errors.New("неверные данные техобслуживания")
	ErrReportNotFound              = errors.New("отчёт не найден")
	ErrReportGenerationFail        = errors.New("не удалось сформировать отчёт")
	ErrValidationFailed            = errors.New("ошибка валидации данных")
	ErrNotFound                    = errors.New("не найдено")
	ErrSomethingWentWrong          = errors.New("что-то пошло не так")
)
