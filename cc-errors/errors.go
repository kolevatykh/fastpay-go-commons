package cc_errors

type BaseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

const (
	/**
	 * Код ошибки при успешном выполнении операции
	 */
	ErrorSuccess = 0

	/**
	 * Минимальная сумма перевода средств
	 */
	ErrorAmountTxMin = 60001

	/**
	 * Аккаунт существует
	 */
	ErrorAccountExist = 60100

	/**
	 * Идентификатор аккаунта существует
	 */
	ErrorIdentifierExist = 60101

	/**
	 * Валюта существует
	 */
	ErrorCurrencyExist = 60104

	/**
	 * Валюта не может быть созданна так как не существет чейнкодов для работы с ней
	 */
	ErrorCurrencyCreateConflict = 60105

	/**
	 * Аккаунт не существует
	 */
	ErrorAccountNotExist = 60110

	/**
	 * Идентификатор аккаунта не существует
	 */
	ErrorIdentifierNotExist = 60111

	/**
	 * Банк не существует
	 */
	ErrorBankNotExist = 60112

	/**
	 * Валюта не существует
	 */
	ErrorCurrencyNotExist = 60114

	/**
	 * Аккаунт не доступен
	 */
	ErrorAccountNotAvailable = 60120

	/**
	 * Аккаунт заблокирован
	 */
	ErrorAccountIsBlocked = 60121

	/**
	 * Банк не доступен
	 */
	ErrorBankNotAvailable = 60122

	/**
	 * Банк заблокирован
	 */
	ErrorBankIsBlocked = 60123

	/**
	 * Банк уже существует
	 */
	ErrorBankExist = 60124

	/**
	 * Контракт не найден
	 */
	ErrorContractNotExists = 60125

	/**
	 * Контракт уже существует
	 */
	ErrorContractExists = 60126

	/**
	 * Клиент банка уже существует
	 */
	ErrorCustomerExists = 60127

	/**
	 * Информация о клиенте банка не найдена
	 */
	ErrorCustomerNotFound = 60128

	/**
	 * Клиентский банк не доступен
	 */
	ErrorClientBankNotAvailable = 60129

	/**
	 * Аккаунт создан без идентификатора
	 */
	ErrorAccountCreatedWithoutIdentifier = 60130

	/**
	 * Ошибка проверки сигнатуры
	 */
	ErrorSignVerify = 60200

	/**
	 * Не достаточно средств
	 */
	ErrorFundsNotEnough = 60240

	/**
	 * Превышен допустимый баланс средств
	 */
	ErrorBalanceOperation = 60250

	/**
	 * Ошибка валидации
	 */
	ErrorValidateDefault = 60300

	/**
	 * Неверный формат адреса аккаунта
	 */
	ErrorAddressNotFollowingRegex = 60301

	/**
	 * Требуется передать адрес аккаунта
	 */
	ErrorAddressNotPassed = 60302

	/**
	 * Неверный формат идентификатора аккаунта
	 */
	ErrorIdentifierNotFolowingRegex = 60303

	/**
	 * Требуется передать идентификатор
	 */
	ErrorIdentifierNotPassed = 60304

	/**
	 * Неверный формат типа идентификации аккаунта
	 */
	ErrorIdentityTypeIncorrect = 60305

	/**
	 * Неверный формат статуса аккаунта
	 */
	ErrorStateIncorrect = 60306

	/**
	 * Неверный формат юридического типа аккаунта
	 */
	ErrorJuridicalTypeIncorrect = 60307

	/**
	 * Неверный формат роли аккаунта
	 */
	ErrorRoleIncorrect = 60308

	/**
	 * Требуется передать имя канала чейнкода
	 */
	ErrorChannelNotPassed = 60309

	/**
	 * Требуется передать имя чейнкода
	 */
	ErrorNameNotPassed = 60310

	/**
	 * Требуется передать версию чейнкода
	 */
	ErrorVersionNotPassed = 60311

	/**
	 * Значение суммы не может быть отрицательным
	 */
	ErrorAmountNegative = 60312

	/**
	 * Требуется передать сумму
	 */
	ErrorAmountNotPassed = 60313

	/**
	 * Требуется передать тип идентификации аккаунта
	 */
	ErrorIdentityTypeNotPassed = 60314

	/**
	 * Значение не может быть отрицательным
	 */
	ErrorValueNegative = 60315

	/**
	 * Требуется передать значение
	 */
	ErrorValueNotPassed = 60316

	/**
	 * Значение кода валюты должно входить в диапазон от 0 до 999
	 */
	ErrorCurrencyCodeRange = 60317

	/**
	 * Требуется передать код валюты
	 */
	ErrorCurrencyCodeNotPassed = 60318

	/**
	 * Значение баланса не может быть отрицательным
	 */
	ErrorBalanceNegative = 60319

	/**
	 * Требуется передать баланс
	 */
	ErrorBalanceNotPassed = 60320

	/**
	 * Неверный формат сообщения
	 */
	ErrorMsgHashNotFolowingRegex = 60321

	/**
	 * Требуется передать сообщение
	 */
	ErrorMsgHashNotPassed = 60322

	/**
	 * Неверный формат R сигнатуры
	 */
	ErrorSigRNotFolowingRegex = 60323

	/**
	 * Требуется передать R сигнатуры
	 */
	ErrorSigRHashNotPassed = 60324

	/**
	 * Неверный формат S сигнатуры
	 */
	ErrorSigSNotFolowingRegex = 60325

	/**
	 * Требуется передать S сигнатуры
	 */
	ErrorSigSNotPassed = 60326

	/**
	 * Неверное значение V сигнатуры
	 */
	ErrorSigVIncorrect = 60327

	/**
	 * Неверный формат V сигнатуры
	 */
	ErrorSigVNotFollowingRegex = 60328

	/**
	 * Значение atc не может быть отрицательным
	 */
	ErrorAtcNegative = 60329

	/**
	 * Требуется передать atc
	 */
	ErrorAtcNotPassed = 60330

	/**
	 * Требуется передать идентификатор транзакции
	 */
	ErrorTxIdSNotPassed = 60331

	/**
	 * Требуется передать время
	 */
	ErrorTimestampNotPassed = 60332

	/**
	 * Значение времени не может быть отрицательным
	 */
	ErrorTimestampNegative = 60333

	/**
	 * Требуется передать тип транзакции
	 */
	ErrorTxTypeNotPassed = 60334

	/**
	 * Неверный формат типа транзакции
	 */
	ErrorTxTypeIncorrect = 60335

	/**
	 * Значение размера страницы меньше минимально допустимого
	 */
	ErrorPageSizeMinValue = 60336

	/**
	 * Требуется передать значение символа валюты
	 */
	ErrorSymbolNotPassed = 60337

	/**
	 * Требуется передать значение количества знаков после запятой
	 */
	ErrorDecimalsNotPassed = 60338

	/**
	 * Значение количества знаков должно входить в диапазон от 0 до 8
	 */
	ErrorDecimalsRange = 60339

	/**
	 * Значение количества знаков после запятой больше максимально допустимого
	 */
	ErrorDecimalsMaxValue = 60340

	/**
	 * Требуется передать значение признака активности
	 */
	ErrorEnabledNotPassed = 60341

	/**
	 * Требуется передать идентификатор аккаунта
	 */
	ErrorIdentifiersNotPassed = 60342

	/**
	 * Количество идентификаторв больше максимально допустимого
	 */
	ErrorIdentifiersMaxValue = 60343

	/**
	 * Требуется передать статус аккаунта
	 */
	ErrorStateNotPassed = 60344

	/**
	 * Требуется передать юридический тип аккаунта
	 */
	ErrorJuridicalTypeNotPassed = 60345

	/**
	 * Требуется передать юридический тип банка-владельца аккаунта
	 */
	ErrorJuridicalTypeBankSetterNotPassed = 60346

	/**
	 * Требуется передать тип аккаунта
	 */
	ErrorAccountTypeNotPassed = 60347

	/**
	 * Неверное значение типа аккаунта
	 */
	ErrorAccountTypeIncorrect = 60348

	/**
	 * Требуется передать роль
	 */
	ErrorRoleNotPassed = 60349

	/**
	 * Требуется передать ID
	 */
	ErrorIdNotPassed = 60350
	/**
	 * Неверный формат ID
	 */
	ErrorIdNotFolowingRegex = 60351
	/**
	 * Требуется передать номер счета
	 */
	ErrorNumberNotPassed = 60352
	/**
	 * Неверный формат номера счета
	 */
	ErrorNumberInvoiceNotFolowingRegex = 60353
	/**
	 * Неверное значение состояния счета
	 */
	ErrorInvoiceStateIncorrect = 60354

	/**
	 * Требуется передать состояние счета
	 */
	ErrorInvoiceStateNotPassed = 60355

	/**
	 * Значение лимита эмисси не может быть отрицательным
	 */
	ErrorIssueLimitNegative = 60356

	/**
	 * Неверное значение статуса транзакции
	 */
	ErrorStatusTransactionIncorrect = 60357

	/**
	 * Требуется передать статус транзакции
	 */
	ErrorStatusTransactionPassed = 60358

	/**
	 * MSP ID не совпадают
	 */
	ErrorMspIdNotMatch = 60359

	/**
	 * Адреса не совпадают
	 */
	ErrorAddressNotMatch = 60360

	/**
	 * Неверный формат публичного ключа аккаунта
	 */
	ErrorPublicKeyNotFolowingRegex = 60361

	/**
	 * Требуется передать публичный ключ аккаунта
	 */
	ErrorPublicKeyNotPassed = 60362

	/**
	 * Требуется передать дополнительную конфигурацию банка
	 */
	ErrorConfNotPassed = 60363

	/**
	 * Требуется передать адрес банка
	 */
	ErrorBankAddressNotPassed = 60364

	/**
	 * Требуется передать код страны
	 */
	ErrorCountryCodeNotPassed = 60365

	/**
	 * Требуется передать идентификатор транзакции
	 */
	ErrorTransactionIdNotPassed = 60366

	/**
	 * Неверный формат идентификатора или адреса аккаунта
	 */
	ErrorAddressOrIdentifierNotFolowingRegex = 60367

	/**
	 * Требуется передать идентификатор или адрес аккаунта
	 */
	ErrorAddressOrIdentifierNotPassed = 60368

	/**
	 * Требуется передать отображаемое имя банка
	 */
	ErrorBankDisplayNameNotPassed = 60369

	/**
	 * Ошибка при получении MSP_ID
	 */
	ErrorGetMspId = 60370

	/**
	 * Время подписи истекло
	 */
	ErrorExpirationSign = 60371

	/**
	 * Подпись уже была использована
	 */
	ErrorSignUsed = 60372

	/**
	 * Действие запрещено на текущем этапе сделки
	 */
	ErrorChangeStateSafeDealForbidden = 60373

	/**
	 * Невалидная конфигурация машины состояний для безапасной сделки
	 */
	ErrorSettingsFSMSafeDeal = 60374

	/**
	 * Сделка с переданным идентификатором уже существет.
	 */
	ErrorSafeDealExists = 60375

	/**
	 * Неверный адрес акцептанта безопасной сделки
	 */
	ErrorNotValidAddressAcceptor = 60376

	/**
	 * Адрес аккаута не является контрагентом сделки
	 */
	ErrorAddressNotCounterpartDeal = 60377

	/**
	 * Текущий контрагент уже проголосовал за расторжение(заключение) сделки
	 */
	ErrorCounterpartMadeChoice = 60378

	/**
	 * Операция запрещена аккаунту, который не явлется 2ым участником или арбитром сделки
	 */
	ErrorNotValidAddressAcceptorOrArbitrator = 60379

	/**
	 * Операция запрещена аккаунту, который не явлется инициатором
	 */
	ErrorNotValidAddressInitiator = 60380

	/**
	 * Операция запрещена аккаунту, который не явлется инициатором
	 */
	ErrorCounterpartWasInvited = 60381

	/**
	 * Участник не был приглашен
	 */
	ErrorParticipantNotInvited = 60382

	/**
	 * Операция запрещена аккаунту, который не явлется владельцем сделки
	 */
	ErrorNotValidAddressOwner = 60383

	/**
	 * Приглашено достаточно участников с такой ролью
	 */
	ErrorEnoughParticipants = 60384

	/**
	 * Некорректная сумма сделки
	 */
	ErrorIncorrectSafeDealAmount = 60385

	/**
	 * Депозит безопасной сделики не найден или уже был пополнен
	 */
	ErrorSafeDealDepositNotFound = 60386

	/**
	 * Некорректный код валюты аккаунта
	 */
	ErrorIncorrectCurrencyCodeAccount = 60387

	/*
	 * Арбитр уже добавлен
	 */
	ErrorArbitratorExist = 60388

	/*
	 * Арбитра не существует
	 */
	ErrorArbitratorNotExist = 60389

	/*
	 * Арбитр не может частично исполнить сделку
	 */
	ErrorArbitratorCannotPartialPerform = 60390

	/*
	 * В сделке не задействован физический актив
	 */
	ErrorPhysicalAssetNotPresent = 60391

	/*
	 * Инициатор не является владельцем физического актива
	 */
	ErrorInitiatorNotOwnPhysicalAsset = 60392

	/*
	 * Акцептант не является владельцем физического актива
	 */
	ErrorAcceptorNotOwnPhysicalAsset = 60393

	/*
	 * Внести данные об частичном исполнении сделки может только первый заявитель об исполнении сделки
	 */
	ErrorPartialPerformOrder = 60394

	/*
	 * В предложении можно указать только адрес аккаунта покупки
	 */
	ErrorOfferAccountBuyOnly = 60395

	/*
	 * Переданный контракт не является предложением
	 */
	ErrorContractNotOffer = 60396

	/*
	 * В сделке не указан адрес владельца предложения
	 */
	ErrorOfferOwnerMissing = 60397

	/*
	 * Минимальное и максимальное количество актива не соответствует предложению
	 */
	ErrorAssetAmountNotMatchOffer = 60398

	/*
	 * Коды валют не соответствует заявленным в предложении
	 */
	ErrorCurrencyNotMatchOffer = 60399

	/*
	 * Курс не соответствует заявленному в предложении
	 */
	ErrorPriceNotMatchOffer = 60400

	/*
	 * Количество актива не соответствует интервалу минимального и максимального количества.
	 */
	ErrorAssetAmountNotCorrespondMaxMinInterval = 60401

	/*
	 * Физический актив не может начисляться на депозит
	 */
	ErrorPhysicalAssetDeposit = 60402

	/*
	 * Количество переданного актива не соответствует переданному курсу
	 */
	ErrorAmountNotCorrespondPrice = 60403

	/**
	 * Попытка перевода средств между юр. лицами
	 */
	ErrorLimitLegalEntry = 60410

	/**
	 * Информация о доступных платформах не найдена
	 */
	ErrorAvailablePlatformsNotFound = 60411

	/**
	 * Превышен операционный лимит перевода средств
	 */
	ErrorLimitOperation = 60420

	/**
	 * Превышен дневной лимит перевода средств
	 */
	ErrorLimitDaily = 60430

	/**
	 * Превышен месячный лимит перевода средств
	 */
	ErrorLimitMonthly = 60440

	/**
	 * Превышен лимит баланса
	 */
	ErrorLimitBalance = 60450

	/**
	 * Лимит не наден
	 */
	ErrorLimitNotFound = 60460

	/**
	 * Адреса отправителя и получателя совпадают
	 */
	ErrorAddressFromToEqual = 60461

	/**
	 * Банк не может быть заблокирован
	 */
	ErrorBankCannotBeBlocked = 60462

	/**
	 * История по трансграничному переводу не найдена
	 */
	ErrorCrossTransactionHistoryNotFound = 60463

	/**
	 * Невозможно выполнить операцию с текущим статусом транзакции
	 */
	ErrorStatusConflict = 60464

	/**
	 * Неверный адрес аккаунта для выполнения вывода средств в рамках трансграничного перевода
	 */
	ErrorCrossTransactionWithdrawIncorrectAddress = 60465

	/**
	 * Неверная сумма для выполнения вывода средств в рамках трансграничномого перевода
	 */
	ErrorCrossTransactionWithdrawIncorrectAmount = 60466

	/**
	 * Другая ошибка
	 */
	ErrorDefault = 60500

	/**
	 * Счет не найден
	 */
	ErrorInvoiceNotExist = 60501

	/**
	 * Счет плательщики не совпадают
	 */
	ErrorInvoiceTransactionPayer = 60502

	/**
	 * Счет получатели не совпадают
	 */
	ErrorInvoiceTransactionRecipient = 60503

	/**
	 * Счет суммы не совпадают
	 */
	ErrorInvoiceTransactionAmount = 60504

	/**
	 * Счет попытка оплатить ранее оплаченный счет
	 */
	ErrorInvoiceTransactionPaid = 60505

	/**
	 * Счет попытка оплатить счета с истекшим сроком жизни
	 */
	ErrorInvoiceTransactionExpired = 60506
	/**
	 * Счет с таким номером уже существует
	 */
	ErrorInvoiceExist = 60507
	/**
	 * Ошибка обновления статуса счета
	 */
	ErrorInvoiceUpdate = 60508

	/**
	 * Указанный адрес платильщика не совпадает с тем, который указан в счете
	 */
	ErrorPayerNotEqual = 60509

	/**
	 * Ошибка сохрания данных в БД
	 */
	ErrorPutState = 60510

	/**
	 * Ошибка при получении данных в БД
	 */
	ErrorGetState = 60511

	/**
	 * Ошибка при удалении данных в БД
	 */
	ErrorDeleteState = 60512

	/**
	 * Ошибка создании композитного ключа
	 */
	ErrorCreateCompositeKey = 60513

	/**
	 * Ошибка при серилизации JSON
	 */
	ErrorJsonMarshal = 60514

	/**
	 * Ошибка при десерилизации JSON
	 */
	ErrorJsonUnmarshal = 60515

	/**
	 * Ошибка при получении времени создания транзакции
	 */
	ErrorGetTxTime = 60516

	/**
	 * Не доступен для этого владельца
	 */
	ErrorForbidden = 60601

	/**
	 * Депозит не найден
	 */
	ErrorDepositNotFound = 60602

	/**
	 * История по транзакции не найдена
	 */
	ErrorTransactionHistoryNotFound = 60603

	/**
	 * Маршрут контрактов не найден
	 */
	ErrorContractRouteNotFound = 60604

	/**
	 * Контракт не найден
	 */
	ErrorContractNotFound = 60605

	/**
	 * Ошибка валидации контракта
	 */
	ErrorContractValidate = 60606

	/**
	 * Комиссия превышает сумму перевода
	 */
	ErrorContractRecoverCommission = 60607
	/**
	 * Ошибка выполнения транзакии
	 */
	ErrorTransactionExecute = 60608
	/**
	 * Адрес банка в сертификате не найден
	 */
	ErrorCertificateBankAddressNoFound = 60609

	/**
	 * Клиринг. Расхождение требований и обязательств
	 */
	ErrorClearingBadClaimsLiabilities = 60701

	/**
	 * Клиринг. Информация о клиринге не найдена
	 */
	ErrorClearingInfoNotFound = 60702

	/**
	 * Сертификат. Ошибка валидация сертификаты
	 */
	ErrorCertificateNotValid = 60801

	/**
	 * Клиентский банк уже существует
	 */
	ErrorClientBankExists = 60802

	/**
	 * Информация о клиентском банке не найдена
	 */
	ErrorClientBankNotFound = 60803

	/**
	 * Опорный банк(отправитель) не является владельцем клиентского банка
	 */
	ErrorClientBankOwnerNotEqualSender = 60804

	/**
	 * Некорректное значение типа лимита
	 */
	ErrorLimitTypeIncorrect = 60805

	/**
	 * Адрес банка отправителя не совпадает с адресом технического аккаунта
	 */
	ErrorAccountTechnicalNotEqlSender = 60806
)

type Error struct {
	Code    int
	Message string
}

var ErrorMessages = map[string]Error{
	"ErrorValidateDefault":                     {60300, "Ошибка валидации"},
	"ErrorAddressNotFollowingRegex":            {60301, "Неверный формат адреса аккаунта"},
	"ErrorAddressNotPassed":                    {60302, "Требуется передать адрес аккаунта"},
	"ErrorIdentifierNotFolowingRegex":          {60303, "Неверный формат идентификатора аккаунта"},
	"ErrorIdentifierNotPassed":                 {60304, "Требуется передать идентификатор"},
	"ErrorIdentityTypeIncorrect":               {60305, "Неверный формат типа идентификации аккаунта"},
	"ErrorStateIncorrect":                      {60306, "Неверный формат статуса аккаунта"},
	"ErrorJuridicalTypeIncorrect":              {60307, "Неверный формат юридического типа аккаунта"},
	"ErrorRoleIncorrect":                       {60308, "Неверный формат роли аккаунта"},
	"ErrorChannelNotPassed":                    {60309, "Требуется передать имя канала чейнкода"},
	"ErrorNameNotPassed":                       {60310, "Требуется передать имя чейнкода"},
	"ErrorVersionNotPassed":                    {60311, "Требуется передать версию чейнкода"},
	"ErrorAmountNegative":                      {60312, "Значение суммы не может быть отрицательным"},
	"ErrorAmountNotPassed":                     {60313, "Требуется передать сумму"},
	"ErrorIdentityTypeNotPassed":               {60314, "Требуется передать тип идентификации аккаунта"},
	"ErrorValueNegative":                       {60315, "Значение не может быть отрицательным"},
	"ErrorValueNotPassed":                      {60316, "Требуется передать значение"},
	"ErrorCurrencyCodeRange":                   {60317, "Значение кода валюты должно входить в диапазон от 0 до 999"},
	"ErrorCurrencyCodeNotPassed":               {60318, "Требуется передать код валюты"},
	"ErrorBalanceNegative":                     {60319, "Значение баланса не может быть отрицательным"},
	"ErrorBalanceNotPassed":                    {60320, "Требуется передать баланс"},
	"ErrorMsgHashNotFolowingRegex":             {60321, "Неверный формат сообщения"},
	"ErrorMsgHashNotPassed":                    {60322, "Требуется передать сообщение"},
	"ErrorSigRNotFolowingRegex":                {60323, "Неверный формат R сигнатуры"},
	"ErrorSigRHashNotPassed":                   {60324, "Требуется передать R сигнатуры"},
	"ErrorSigSNotFolowingRegex":                {60325, "Неверный формат S сигнатуры"},
	"ErrorSigSNotPassed":                       {60326, "Требуется передать S сигнатуры"},
	"ErrorSigVIncorrect":                       {60327, "Неверное значение V сигнатуры"},
	"ErrorSigVNotFollowingRegex":               {60328, "Неверный формат V сигнатуры"},
	"ErrorAtcNegative":                         {60329, "Значение atc не может быть отрицательным"},
	"ErrorAtcNotPassed":                        {60330, "Требуется передать atc"},
	"ErrorTxIdSNotPassed":                      {60331, "Требуется передать идентификатор транзакции"},
	"ErrorTimestampNotPassed":                  {60332, "Требуется передать время"},
	"ErrorTimestampNegative":                   {60333, "Значение времени не может быть отрицательным"},
	"ErrorTxTypeNotPassed":                     {60334, "Требуется передать тип транзакции"},
	"ErrorTxTypeIncorrect":                     {60335, "Неверный формат типа транзакции"},
	"ErrorPageSizeMinValue":                    {60336, "Значение размера страницы меньше минимально допустимого"},
	"ErrorSymbolNotPassed":                     {60337, "Требуется передать значение символа валюты"},
	"ErrorDecimalsNotPassed":                   {60338, "Требуется передать значение количества знаков после запятой"},
	"ErrorDecimalsRange":                       {60339, "Значение количества знаков должно входить в диапазон от 0 до 8"},
	"ErrorDecimalsMaxValue":                    {60340, "Значение количества знаков после запятой больше максимально допустимого"},
	"ErrorEnabledNotPassed":                    {60341, "Требуется передать значение признака активности"},
	"ErrorIdentifiersNotPassed":                {60342, "Требуется передать идентификатор аккаунта"},
	"ErrorIdentifiersMaxValue":                 {60343, "Количество идентификаторв больше максимально допустимого"},
	"ErrorStateNotPassed":                      {60344, "Требуется передать статус аккаунта"},
	"ErrorJuridicalTypeNotPassed":              {60345, "Требуется передать юридический тип аккаунта"},
	"ErrorJuridicalTypeBankSetterNotPassed":    {60346, "Требуется передать юридический тип банка-владельца аккаунта"},
	"ErrorAccountTypeNotPassed":                {60347, "Требуется передать тип аккаунта"},
	"ErrorAccountTypeIncorrect":                {60348, "Неверное значение типа аккаунта"},
	"ErrorRoleNotPassed":                       {60349, "Требуется передать роль"},
	"ErrorIdNotPassed":                         {60350, "Требуется передать ID"},
	"ErrorIdNotFolowingRegex":                  {60351, "Неверный формат ID"},
	"ErrorNumberNotPassed":                     {60352, "Требуется передать номер счета"},
	"ErrorNumberInvoiceNotFolowingRegex":       {60353, "Неверный формат номера счета"},
	"ErrorInvoiceStateIncorrect":               {60354, "Неверное значение состояния счета"},
	"ErrorInvoiceStateNotPassed":               {60355, "Требуется передать состояние счета"},
	"ErrorIssueLimitNegative":                  {60356, "Значение лимита эмисси не может быть отрицательным"},
	"ErrorStatusTransactionIncorrect":          {60357, "Неверное значение статуса транзакции"},
	"ErrorStatusTransactionPassed":             {60358, "Требуется передать статус транзакции"},
	"ErrorMspIdNotMatch":                       {60359, "MSP ID не совпадают"},
	"ErrorAddressNotMatch":                     {60360, "Адреса не совпадают"},
	"ErrorPublicKeyNotFolowingRegex":           {60361, "Неверный формат публичного ключа аккаунта"},
	"ErrorPublicKeyNotPassed":                  {60362, "Требуется передать публичный ключ аккаунта"},
	"ErrorConfNotPassed":                       {60363, "Требуется передать дополнительную конфигурацию банка"},
	"ErrorBankAddressNotPassed":                {60364, "Требуется передать адрес банка"},
	"ErrorCountryCodeNotPassed":                {60365, "Требуется передать код страны"},
	"ErrorTransactionIdNotPassed":              {60366, "Требуется передать идентификатор транзакции"},
	"ErrorAddressOrIdentifierNotFolowingRegex": {60367, "Неверный формат идентификатора или адреса аккаунта"},
	"ErrorAddressOrIdentifierNotPassed":        {60368, "Требуется передать идентификатор или адрес аккаунта"},
	"ErrorBankDisplayNameNotPassed":            {60369, "Требуется передать отображаемое имя банка"},
	"ErrorGetMspId":                            {60370, "Ошибка при получении MSP_ID"},
	"ErrorExpirationSign":                      {60371, "Время подписи истекло"},
	"ErrorSignUsed":                            {60372, "Подпись уже была использована"},
	"ErrorChangeStateSafeDealForbidden":        {60373, "Действие запрещено на текущем этапе сделки"},
	"ErrorSettingsFSMSafeDeal":                 {60374, "Невалидная конфигурация машины состояний для безапасной сделки"},
	"ErrorSafeDealExists":                      {60375, "Сделка с переданным идентификатором уже существет."},
	"ErrorNotValidAddressAcceptor":             {60376, "Запрещено принимать заявку аккаунтом, который не явлется 2ым участником сделки"},
	"ErrorAddressNotCounterpartDeal":           {60377, "Адрес аккаута не является контрагентом сделки"},
	"ErrorCounterpartMadeChoice":               {60378, "Текущий контрагент уже проголосовал за расторжение(заключение) сделки"},
	"ErrorNotValidAddressAcceptorOrArbitrator": {60379, "Операция запрещена аккаунту, который не явлется 2ым участником или арбитром сделки"},
	"ErrorNotValidAddressInitiator":            {60380, "Операция запрещена аккаунту, который не явлется инициатором"},
	"ErrorLimitTypeIncorrect":                  {60805, "Некорректное значение типа лимита"},
}
