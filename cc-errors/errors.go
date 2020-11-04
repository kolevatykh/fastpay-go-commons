package cc_errors

type BaseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

const (

	/** Код ошибки при успешном выполнении операции */
	ErrorSuccess = 0

	/** Аккаунт существует */
	ErrorAccountExist = 60100

	/** Идентификатор аккаунта существует */
	ErrorIdentifierExist = 60101

	/** Валюта существует */
	ErrorCurrencyExist = 60104

	/** Валюта не может быть созданна так как не существет чейнкодов для работы с ней */
	ErrorCurrencyCreateConflict = 60105

	/** Аккаунт не существует */
	ErrorAccountNotExist = 60110

	/** Идентификатор аккаунта не существует */
	ErrorIdentifierNotExist = 60111

	/** Банк не существует */
	ErrorBankNotExist = 60112

	/** Валюта не существует */
	ErrorCurrencyNotExist = 60114

	/** Клиентский банк заблокирован */
	ErrorClientBankIsBlocked = 60117

	/** Аккаунт не является эскроу-счётом */
	ErrorAccountNotEscrow = 60118

	/** Аккаунт не идентифицирован */
	ErrorAccountNotIdentified = 60119 //+

	/** Аккаунт не доступен */
	ErrorAccountNotAvailable = 60120

	/** Аккаунт заблокирован */
	ErrorAccountIsBlocked = 60121

	/** Банк не доступен */
	ErrorBankNotAvailable = 60122

	/** Банк заблокирован */
	ErrorBankIsBlocked = 60123

	/** Банк уже существует */
	ErrorBankExist = 60124

	/** Контракт не найден */
	ErrorContractNotExists = 60125

	/** Контракт уже существует */
	ErrorContractExists = 60126

	/** Информация о клиенте банка не найдена */
	ErrorCustomerNotFound = 60128

	/** Клиентский банк не доступен */
	ErrorClientBankNotAvailable = 60129

	/** Ошибка проверки сигнатуры */
	ErrorSignVerify = 60200 //++

	/** Недостаточно средств */
	ErrorFundsNotEnough = 60240 //++

	/** Ошибка валидации */
	ErrorValidateDefault = 60300

	/** Неверный формат адреса аккаунта */
	ErrorAddressNotFollowingRegex = 60301

	/** Требуется передать адрес аккаунта */
	ErrorAddressNotPassed = 60302

	/** Неверный формат идентификатора аккаунта */
	ErrorIdentifierNotFolowingRegex = 60303

	/** Требуется передать идентификатор */
	ErrorIdentifierNotPassed = 60304

	/** Неверный формат типа идентификации аккаунта */
	ErrorIdentityTypeIncorrect = 60305

	/** Неверный формат статуса аккаунта */
	ErrorStateIncorrect = 60306

	/** Неверный формат юридического типа аккаунта */
	ErrorJuridicalTypeIncorrect = 60307

	/** Неверный формат роли аккаунта */
	ErrorRoleIncorrect = 60308

	/** Требуется передать имя канала чейнкода */
	ErrorChannelNotPassed = 60309

	/** Требуется передать имя чейнкода */
	ErrorNameNotPassed = 60310

	/** Требуется передать версию чейнкода */
	ErrorVersionNotPassed = 60311

	/** Значение суммы не может быть отрицательным */
	ErrorAmountNegative = 60312

	/** Требуется передать сумму */
	ErrorAmountNotPassed = 60313

	/** Требуется передать тип идентификации аккаунта */
	ErrorIdentityTypeNotPassed = 60314

	/** Значение не может быть отрицательным */
	ErrorValueNegative = 60315

	/** Требуется передать значение */
	ErrorValueNotPassed = 60316

	/** Значение кода валюты должно входить в диапазон от 0 до 999 */
	ErrorCurrencyCodeRange = 60317

	/** Требуется передать код валюты */
	ErrorCurrencyCodeNotPassed = 60318

	/** Значение баланса не может быть отрицательным */
	ErrorBalanceNegative = 60319

	/** Требуется передать баланс */
	ErrorBalanceNotPassed = 60320

	/** Неверный формат сообщения */
	ErrorMsgHashNotFolowingRegex = 60321

	/** Требуется передать сообщение */
	ErrorMsgHashNotPassed = 60322

	/** Неверный формат R сигнатуры */
	ErrorSigRNotFolowingRegex = 60323

	/** Требуется передать R сигнатуры */
	ErrorSigRHashNotPassed = 60324

	/** Неверный формат S сигнатуры */
	ErrorSigSNotFolowingRegex = 60325

	/** Требуется передать S сигнатуры */
	ErrorSigSNotPassed = 60326

	/** Неверное значение V сигнатуры */
	ErrorSigVIncorrect = 60327

	/** Неверный формат V сигнатуры */
	ErrorSigVNotFollowingRegex = 60328

	/** Требуется передать идентификатор транзакции */
	ErrorTxIdSNotPassed = 60331

	/** Требуется передать время */
	ErrorTimestampNotPassed = 60332

	/** Значение времени не может быть отрицательным */
	ErrorTimestampNegative = 60333

	/** Требуется передать тип транзакции */
	ErrorTxTypeNotPassed = 60334

	/** Неверный формат типа транзакции */
	ErrorTxTypeIncorrect = 60335

	/** Значение размера страницы меньше минимально допустимого */
	ErrorPageSizeMinValue = 60336

	/** Требуется передать значение символа валюты */
	ErrorSymbolNotPassed = 60337

	/** Требуется передать значение количества знаков после запятой */
	ErrorDecimalsNotPassed = 60338

	/** Значение количества знаков должно входить в диапазон от 0 до 8 */
	ErrorDecimalsRange = 60339

	/** Значение количества знаков после запятой больше максимально допустимого */
	ErrorDecimalsMaxValue = 60340

	/** Требуется передать значение признака активности */
	ErrorEnabledNotPassed = 60341

	/** Требуется передать идентификатор аккаунта */
	ErrorIdentifiersNotPassed = 60342

	/** Количество идентификаторв больше максимально допустимого */
	ErrorIdentifiersMaxValue = 60343

	/** Требуется передать статус аккаунта */
	ErrorStateNotPassed = 60344

	/** Требуется передать юридический тип аккаунта */
	ErrorJuridicalTypeNotPassed = 60345

	/** Требуется передать юридический тип банка-владельца аккаунта */
	ErrorJuridicalTypeBankSetterNotPassed = 60346

	/** Требуется передать тип аккаунта */
	ErrorAccountTypeNotPassed = 60347

	/** Неверное значение типа аккаунта */
	ErrorAccountTypeIncorrect = 60348

	/** Требуется передать роль */
	ErrorRoleNotPassed = 60349

	/** Требуется передать ID */
	ErrorIdNotPassed = 60350

	/** Неверный формат ID */
	ErrorIdNotFolowingRegex = 60351

	/** Требуется передать номер счета */
	ErrorNumberNotPassed = 60352

	/** Неверный формат номера счета */
	ErrorNumberInvoiceNotFolowingRegex = 60353

	/** Неверное значение состояния счета */
	ErrorInvoiceStateIncorrect = 60354

	/** Требуется передать состояние счета */
	ErrorInvoiceStateNotPassed = 60355

	/** Значение лимита эмисси не может быть отрицательным */
	ErrorIssueLimitNegative = 60356

	/** Неверное значение статуса транзакции */
	ErrorStatusTransactionIncorrect = 60357

	/** Требуется передать статус транзакции */
	ErrorStatusTransactionPassed = 60358

	/** MSP ID не совпадают */
	ErrorMspIdNotMatch = 60359

	/** Адреса не совпадают */
	ErrorAddressNotMatch = 60360

	/** Неверный формат публичного ключа аккаунта */
	ErrorPublicKeyNotFolowingRegex = 60361

	/** Требуется передать публичный ключ аккаунта */
	ErrorPublicKeyNotPassed = 60362

	/** Требуется передать дополнительную конфигурацию банка */
	ErrorConfNotPassed = 60363

	/** Требуется передать адрес банка */
	ErrorBankAddressNotPassed = 60364

	/** Требуется передать код страны */
	ErrorCountryCodeNotPassed = 60365

	/** Требуется передать идентификатор транзакции */
	ErrorTransactionIdNotPassed = 60366

	/** Неверный формат идентификатора или адреса аккаунта */
	ErrorAddressOrIdentifierNotFolowingRegex = 60367

	/** Требуется передать идентификатор или адрес аккаунта */
	ErrorAddressOrIdentifierNotPassed = 60368

	/** Требуется передать отображаемое имя банка */
	ErrorBankDisplayNameNotPassed = 60369

	/** Ошибка при получении MSP_ID */
	ErrorGetMspId = 60370

	/** Время подписи истекло */
	ErrorExpirationSign = 60371

	/** Подпись уже была использована */
	ErrorSignUsed = 60372

	/** Действие запрещено на текущем этапе сделки */
	ErrorChangeStateSafeDealForbidden = 60373

	/** Невалидная конфигурация машины состояний для безапасной сделки */
	ErrorSettingsFSMSafeDeal = 60374

	/** Сделка с переданным идентификатором уже существет. */
	ErrorSafeDealExists = 60375

	/** Неверный адрес акцептанта безопасной сделки */
	ErrorNotValidAddressAcceptor = 60376

	/** Адрес аккаута не является контрагентом сделки */
	ErrorAddressNotCounterpartDeal = 60377

	/** Текущий контрагент уже проголосовал за расторжение(заключение) сделки */
	ErrorCounterpartMadeChoice = 60378

	/** Операция запрещена аккаунту, который не явлется 2ым участником или арбитром сделки */
	ErrorNotValidAddressAcceptorOrArbitrator = 60379

	/** Операция запрещена аккаунту, который не явлется инициатором */
	ErrorNotValidAddressInitiator = 60380

	/** Участник уже был приглашен */
	ErrorCounterpartWasInvited = 60381

	/** Участник не был приглашен */
	ErrorParticipantNotInvited = 60382

	/** Операция запрещена аккаунту, который не явлется владельцем сделки */
	ErrorNotValidAddressOwner = 60383

	/** Приглашено достаточно участников с такой ролью */
	ErrorEnoughParticipants = 60384

	/** Некорректная сумма сделки */
	ErrorIncorrectSafeDealAmount = 60385

	/** Эскроу-счет безопасной сделики не найден или уже был пополнен */
	ErrorSafeDealEscrowNotFound = 60386

	/** Некорректный код валюты аккаунта */
	ErrorIncorrectCurrencyCodeAccount = 60387

	/* Арбитр уже добавлен */
	ErrorArbitratorExist = 60388

	/* Арбитра не существует */
	ErrorArbitratorNotExist = 60389

	/* Арбитр не может частично исполнить сделку */
	ErrorArbitratorCannotPartialPerform = 60390

	/* В сделке не задействован физический актив */
	ErrorPhysicalAssetNotPresent = 60391

	/* Внести данные об частичном исполнении сделки может только первый заявитель об исполнении сделки */
	ErrorPartialPerformOrder = 60394

	/* В предложении можно указать только адрес аккаунта продажи */
	ErrorOfferAccountSellOnly = 60395

	/* Переданный контракт не является предложением */
	ErrorContractNotOffer = 60396

	/* В сделке не указан адрес владельца предложения */
	ErrorOfferOwnerMissing = 60397

	/* Коды валют не соответствует заявленным в предложении */
	ErrorCurrencyNotMatchOffer = 60399

	/* Курс не соответствует заявленному в предложении */
	ErrorPriceNotMatchOffer = 60400

	/* Количество актива не соответствует интервалу минимального и максимального количества. */
	ErrorAssetAmountNotCorrespondMaxMinInterval = 60401

	/* Физический актив не может начисляться на эскроу-счет */
	ErrorPhysicalAssetEscrow = 60402

	/* Количество переданного актива не соответствует переданному курсу */
	ErrorAmountNotCorrespondPrice = 60403

	/* Безопасная сделка не найдена */
	ErrorSafeDealNotFound = 60404

	/* Участник безопасной сделки уже проголосовал */
	ErrorCounterpartVoted = 60405

	/* Некорректная сумма эскроу-счета */
	ErrorIncorrectEscrowAmount = 60406

	/* Аккаунт с таким адресом уже принимает участие в сделке */
	ErrorAccountAlreadyInDeal = 60407

	/** Информация о доступных платформах не найдена */
	ErrorAvailablePlatformsNotFound = 60411

	/** Превышен операционный лимит перевода средств */
	ErrorLimitOperation = 60420

	/** Превышен дневной лимит перевода средств */
	ErrorLimitDaily = 60430

	/** Превышен месячный лимит перевода средств */
	ErrorLimitMonthly = 60440

	/** Превышен лимит баланса */
	ErrorLimitBalance = 60450

	/** Адреса отправителя и получателя совпадают */
	ErrorAddressFromToEqual = 60461

	/** Банк не может быть заблокирован */
	ErrorBankCannotBeBlocked = 60462

	/** История по трансграничному переводу не найдена */
	ErrorCrossTransactionHistoryNotFound = 60463

	/** Невозможно выполнить операцию с текущим статусом транзакции */
	ErrorStatusConflict = 60464

	/** Неверный адрес аккаунта для выполнения вывода средств в рамках трансграничного перевода */
	ErrorCrossTransactionWithdrawIncorrectAddress = 60465

	/** Неверная сумма для выполнения вывода средств в рамках трансграничномого перевода */
	ErrorCrossTransactionWithdrawIncorrectAmount = 60466

	/** Другая ошибка */
	ErrorDefault = 60500

	/** Счет не найден */
	ErrorInvoiceNotExist = 60501

	/** Cуммы cчета не совпадают */
	ErrorInvoiceTransactionAmount = 60504

	/** Попытка оплатить ранее оплаченный счет */
	ErrorInvoiceTransactionPaid = 60505

	/** Попытка оплатить счета с истекшим сроком жизни */
	ErrorInvoiceTransactionExpired = 60506

	/** Счет с таким номером уже существует */
	ErrorInvoiceExist = 60507

	/** Ошибка обновления статуса счета */
	ErrorInvoiceUpdate = 60508

	/** Указанный адрес платильщика не совпадает с тем, который указан в счете */
	ErrorPayerNotEqual = 60509

	/** Ошибка сохрания данных в БД */
	ErrorPutState = 60510

	/** Ошибка при получении данных в БД */
	ErrorGetState = 60511

	/** Ошибка при удалении данных в БД */
	ErrorDeleteState = 60512

	/** Ошибка создании композитного ключа */
	ErrorCreateCompositeKey = 60513

	/** Ошибка при серилизации JSON */
	ErrorJsonMarshal = 60514

	/** Ошибка при десерилизации JSON */
	ErrorJsonUnmarshal = 60515

	/** Ошибка при получении времени создания транзакции */
	ErrorGetTxTime = 60516

	/** Не доступен для этого владельца */
	ErrorForbidden = 60601

	/** Депозит не найден */
	ErrorDepositNotFound = 60602

	/** Ошибка валидации контракта */
	ErrorContractValidate = 60606

	/** Адрес банка в сертификате не совпадает */
	ErrorCertificateBankAddressNotMatch = 60609

	/** Клиринг. Расхождение требований и обязательств */
	ErrorClearingBadClaimsLiabilities = 60701

	/** Клиринг. Информация о клиринге не найдена */
	ErrorClearingInfoNotFound = 60702

	/** Сертификат. Отсутвует атрибут address в сертификате */
	ErrorCertificateBankAddressNotFound = 60801

	/** Клиентский банк уже существует */
	ErrorClientBankExists = 60802

	/** Информация о клиентском банке не найдена */
	ErrorClientBankNotFound = 60803

	/** Опорный банк(отправитель) не является владельцем клиентского банка */
	ErrorClientBankOwnerNotEqualSender = 60804

	/** Некорректное значение типа лимита */
	ErrorLimitTypeIncorrect = 60805

	/** Адрес банка отправителя не совпадает с адресом технического аккаунта */
	ErrorAccountTechnicalNotEqlSender = 60806
)

var ErrorCodeMessagesMap = map[int]string{
	ErrorAccountExist:                             "Аккаунт существует",
	ErrorIdentifierExist:                          "Идентификатор аккаунта существует",
	ErrorCurrencyExist:                            "Валюта существует",
	ErrorCurrencyCreateConflict:                   "Валюта не может быть созданна так как не существет чейнкодов для работы с ней",
	ErrorAccountNotExist:                          "Аккаунт не существует",
	ErrorIdentifierNotExist:                       "Идентификатор аккаунта не существует",
	ErrorBankNotExist:                             "Банк не существует",
	ErrorCurrencyNotExist:                         "Валюта не существует",
	ErrorClientBankIsBlocked:                      "Клиентский банк заблокирован",
	ErrorAccountNotIdentified:                     "Аккаунт не идентифицирован",
	ErrorAccountNotAvailable:                      "Аккаунт не доступен",
	ErrorAccountIsBlocked:                         "Аккаунт заблокирован",
	ErrorBankNotAvailable:                         "Банк не доступен",
	ErrorBankIsBlocked:                            "Банк заблокирован",
	ErrorBankExist:                                "Банк уже существует",
	ErrorContractNotExists:                        "Контракт не найден",
	ErrorContractExists:                           "Контракт уже существует",
	ErrorCustomerNotFound:                         "Информация о клиенте банка не найдена",
	ErrorClientBankNotAvailable:                   "Клиентский банк недоступен",
	ErrorSignVerify:                               "Ошибка проверки сигнатуры",
	ErrorFundsNotEnough:                           "Недостаточно средств",
	ErrorValidateDefault:                          "Ошибка валидации",
	ErrorAddressNotFollowingRegex:                 "Неверный формат адреса аккаунта",
	ErrorAddressNotPassed:                         "Требуется передать адрес аккаунта",
	ErrorIdentifierNotFolowingRegex:               "Неверный формат идентификатора аккаунта",
	ErrorIdentifierNotPassed:                      "Требуется передать идентификатор",
	ErrorIdentityTypeIncorrect:                    "Неверный формат типа идентификации аккаунта",
	ErrorStateIncorrect:                           "Неверный формат статуса аккаунта",
	ErrorJuridicalTypeIncorrect:                   "Неверный формат юридического типа аккаунта",
	ErrorRoleIncorrect:                            "Неверный формат роли аккаунта",
	ErrorChannelNotPassed:                         "Требуется передать имя канала чейнкода",
	ErrorNameNotPassed:                            "Требуется передать имя чейнкода",
	ErrorVersionNotPassed:                         "Требуется передать версию чейнкода",
	ErrorAmountNegative:                           "Значение суммы не может быть отрицательным",
	ErrorAmountNotPassed:                          "Требуется передать сумму",
	ErrorIdentityTypeNotPassed:                    "Требуется передать тип идентификации аккаунта",
	ErrorValueNegative:                            "Значение не может быть отрицательным",
	ErrorValueNotPassed:                           "Требуется передать значение",
	ErrorCurrencyCodeRange:                        "Значение кода валюты должно входить в диапазон от 0 до 999",
	ErrorCurrencyCodeNotPassed:                    "Требуется передать код валюты",
	ErrorBalanceNegative:                          "Значение баланса не может быть отрицательным",
	ErrorBalanceNotPassed:                         "Требуется передать баланс",
	ErrorMsgHashNotFolowingRegex:                  "Неверный формат сообщения",
	ErrorMsgHashNotPassed:                         "Требуется передать сообщение",
	ErrorSigRNotFolowingRegex:                     "Неверный формат R сигнатуры",
	ErrorSigRHashNotPassed:                        "Требуется передать R сигнатуры",
	ErrorSigSNotFolowingRegex:                     "Неверный формат S сигнатуры",
	ErrorSigSNotPassed:                            "Требуется передать S сигнатуры",
	ErrorSigVIncorrect:                            "Неверное значение V сигнатуры",
	ErrorSigVNotFollowingRegex:                    "Неверный формат V сигнатуры",
	ErrorTxIdSNotPassed:                           "Требуется передать идентификатор транзакции",
	ErrorTimestampNotPassed:                       "Требуется передать время",
	ErrorTimestampNegative:                        "Значение времени не может быть отрицательным",
	ErrorTxTypeNotPassed:                          "Требуется передать тип транзакции",
	ErrorTxTypeIncorrect:                          "Неверный формат типа транзакции",
	ErrorPageSizeMinValue:                         "Значение размера страницы меньше минимально допустимого",
	ErrorSymbolNotPassed:                          "Требуется передать значение символа валюты",
	ErrorDecimalsNotPassed:                        "Требуется передать значение количества знаков после запятой",
	ErrorDecimalsRange:                            "Значение количества знаков должно входить в диапазон от 0 до 8",
	ErrorDecimalsMaxValue:                         "Значение количества знаков после запятой больше максимально допустимого",
	ErrorEnabledNotPassed:                         "Требуется передать значение признака активности",
	ErrorIdentifiersNotPassed:                     "Требуется передать идентификатор аккаунта",
	ErrorIdentifiersMaxValue:                      "Количество идентификаторв больше максимально допустимого",
	ErrorStateNotPassed:                           "Требуется передать статус аккаунта",
	ErrorJuridicalTypeNotPassed:                   "Требуется передать юридический тип аккаунта",
	ErrorJuridicalTypeBankSetterNotPassed:         "Требуется передать юридический тип банка-владельца аккаунта",
	ErrorAccountTypeNotPassed:                     "Требуется передать тип аккаунта",
	ErrorAccountTypeIncorrect:                     "Неверное значение типа аккаунта",
	ErrorRoleNotPassed:                            "Требуется передать роль",
	ErrorIdNotPassed:                              "Требуется передать ID",
	ErrorIdNotFolowingRegex:                       "Неверный формат ID",
	ErrorNumberNotPassed:                          "Требуется передать номер счета",
	ErrorNumberInvoiceNotFolowingRegex:            "Неверный формат номера счета",
	ErrorInvoiceStateIncorrect:                    "Неверное значение состояния счета",
	ErrorInvoiceStateNotPassed:                    "Требуется передать состояние счета",
	ErrorIssueLimitNegative:                       "Значение лимита эмисси не может быть отрицательным",
	ErrorStatusTransactionIncorrect:               "Неверное значение статуса транзакции",
	ErrorStatusTransactionPassed:                  "Требуется передать статус транзакции",
	ErrorMspIdNotMatch:                            "MSP ID не совпадают",
	ErrorAddressNotMatch:                          "Адреса не совпадают",
	ErrorPublicKeyNotFolowingRegex:                "Неверный формат публичного ключа аккаунта",
	ErrorPublicKeyNotPassed:                       "Требуется передать публичный ключ аккаунта",
	ErrorConfNotPassed:                            "Требуется передать дополнительную конфигурацию банка",
	ErrorBankAddressNotPassed:                     "Требуется передать адрес банка",
	ErrorCountryCodeNotPassed:                     "Требуется передать код страны",
	ErrorTransactionIdNotPassed:                   "Требуется передать идентификатор транзакции",
	ErrorAddressOrIdentifierNotFolowingRegex:      "Неверный формат идентификатора или адреса аккаунта",
	ErrorAddressOrIdentifierNotPassed:             "Требуется передать идентификатор или адрес аккаунта",
	ErrorBankDisplayNameNotPassed:                 "Требуется передать отображаемое имя банка",
	ErrorGetMspId:                                 "Ошибка при получении MSP_ID",
	ErrorExpirationSign:                           "Время подписи истекло",
	ErrorSignUsed:                                 "Подпись уже была использована",
	ErrorChangeStateSafeDealForbidden:             "Действие запрещено на текущем этапе сделки",
	ErrorSettingsFSMSafeDeal:                      "Невалидная конфигурация машины состояний для безопасной сделки",
	ErrorSafeDealExists:                           "Сделка с переданным идентификатором уже существет.",
	ErrorNotValidAddressAcceptor:                  "Запрещено принимать заявку аккаунтом, который не явлется 2ым участником сделки",
	ErrorAddressNotCounterpartDeal:                "Адрес аккаута не является контрагентом сделки",
	ErrorCounterpartMadeChoice:                    "Текущий контрагент уже проголосовал за расторжение(заключение) сделки",
	ErrorNotValidAddressAcceptorOrArbitrator:      "Операция запрещена аккаунту, который не явлется 2ым участником или арбитром",
	ErrorNotValidAddressInitiator:                 "Операция запрещена аккаунту, который не явлется инициатором",
	ErrorCounterpartWasInvited:                    "Участник уже был приглашен",
	ErrorParticipantNotInvited:                    "Участник не был приглашен",
	ErrorNotValidAddressOwner:                     "Операция запрещена аккаунту, который не явлется владельцем сделки",
	ErrorEnoughParticipants:                       "Приглашено достаточно участников с такой ролью",
	ErrorIncorrectSafeDealAmount:                  "Некорректная сумма сделки",
	ErrorSafeDealEscrowNotFound:                   "Эскроу-счет безопасной сделики не найден или уже был пополнен",
	ErrorIncorrectCurrencyCodeAccount:             "Некорректный код валюты аккаунта",
	ErrorArbitratorExist:                          "Арбитр уже добавлен",
	ErrorArbitratorNotExist:                       "Арбитра не существует",
	ErrorArbitratorCannotPartialPerform:           "Арбитр не может частично исполнить сделку",
	ErrorPhysicalAssetNotPresent:                  "В сделке не задействован физический актив",
	ErrorPartialPerformOrder:                      "Внести данные об частичном исполнении сделки может только первый заявитель об исполнении сделки",
	ErrorOfferAccountSellOnly:                     "В предложении можно указать только адрес аккаунта продажи",
	ErrorContractNotOffer:                         "Переданный контракт не является предложением",
	ErrorOfferOwnerMissing:                        "В сделке не указан адрес владельца предложения",
	ErrorCurrencyNotMatchOffer:                    "Коды валют не соответствует заявленным в предложении",
	ErrorPriceNotMatchOffer:                       "Курс не соответствует заявленному в предложении",
	ErrorAssetAmountNotCorrespondMaxMinInterval:   "Количество актива не соответствует интервалу минимального и максимального количества",
	ErrorPhysicalAssetEscrow:                      "Физический актив не может начисляться на эскроу-счет",
	ErrorAmountNotCorrespondPrice:                 "Количество переданного актива не соответствует переданному курсу",
	ErrorSafeDealNotFound:                         "Безопасная сделка не найдена",
	ErrorCounterpartVoted:                         "Участник безопасной сделки уже проголосовал",
	ErrorIncorrectEscrowAmount:                    "Некорректная сумма эскроу-счета",
	ErrorAccountAlreadyInDeal:                     "Аккаунт с таким адресом уже принимает участие в сделке",
	ErrorAvailablePlatformsNotFound:               "Информация о доступных платформах не найдена",
	ErrorLimitOperation:                           "Превышен операционный лимит перевода средств",
	ErrorLimitDaily:                               "Превышен дневной лимит перевода средств",
	ErrorLimitMonthly:                             "Превышен месячный лимит перевода средств",
	ErrorLimitBalance:                             "Превышен лимит баланса",
	ErrorAddressFromToEqual:                       "Адреса отправителя и получателя совпадают",
	ErrorBankCannotBeBlocked:                      "Банк не может быть заблокирован",
	ErrorCrossTransactionHistoryNotFound:          "История по трансграничному переводу не найдена",
	ErrorStatusConflict:                           "Невозможно выполнить операцию с текущим статусом транзакции",
	ErrorCrossTransactionWithdrawIncorrectAddress: "Неверный адрес аккаунта для выполнения вывода средств в рамках трансграничного перевода",
	ErrorCrossTransactionWithdrawIncorrectAmount:  "Неверная сумма для выполнения вывода средств в рамках трансграничномого перевода",
	ErrorDefault:                                  "Другая ошибка",
	ErrorInvoiceNotExist:                          "Счет не найден",
	ErrorInvoiceTransactionAmount:                 "Cуммы cчета не совпадают",
	ErrorInvoiceTransactionPaid:                   "Попытка оплатить ранее оплаченный счет",
	ErrorInvoiceTransactionExpired:                "Попытка оплатить счета с истекшим сроком жизни",
	ErrorInvoiceExist:                             "Счет с таким номером уже существует",
	ErrorInvoiceUpdate:                            "Ошибка обновления статуса счета",
	ErrorPayerNotEqual:                            "Указанный адрес платильщика не совпадает с тем, который указан в счете",
	ErrorPutState:                                 "Ошибка сохрания данных в БД",
	ErrorGetState:                                 "Ошибка при получении данных в БД",
	ErrorDeleteState:                              "Ошибка при удалении данных в БД",
	ErrorCreateCompositeKey:                       "Ошибка создании композитного ключа",
	ErrorJsonMarshal:                              "Ошибка при серилизации JSON",
	ErrorJsonUnmarshal:                            "Ошибка при десерилизации JSON",
	ErrorGetTxTime:                                "Ошибка при получении времени создания транзакции",
	ErrorForbidden:                                "Не доступен для этого владельца",
	ErrorDepositNotFound:                          "Депозит не найден",
	ErrorContractValidate:                         "Ошибка валидации контракта",
	ErrorCertificateBankAddressNotMatch:           "Адрес банка в сертификате не совпадает",
	ErrorClearingBadClaimsLiabilities:             "Клиринг. Расхождение требований и обязательств",
	ErrorClearingInfoNotFound:                     "Клиринг. Информация о клиринге не найдена",
	ErrorCertificateBankAddressNotFound:           "Сертификат. Отсутвует атрибут address в сертификате",
	ErrorClientBankExists:                         "Клиентский банк уже существует",
	ErrorClientBankNotFound:                       "Информация о клиентском банке не найдена",
	ErrorClientBankOwnerNotEqualSender:            "Опорный банк(отправитель) не является владельцем клиентского банка",
	ErrorLimitTypeIncorrect:                       "Некорректное значение типа лимита",
	ErrorAccountTechnicalNotEqlSender:             "Адрес банка отправителя не совпадает с адресом технического аккаунта",
	ErrorAccountNotEscrow:                         "Аккаунт не является эскроу-счетом",
}

var ErrorStringCodeMap = map[string]int{
	"ErrorAccountExist":                             ErrorAccountExist,
	"ErrorIdentifierExist":                          ErrorIdentifierExist,
	"ErrorCurrencyExist":                            ErrorCurrencyExist,
	"ErrorCurrencyCreateConflict":                   ErrorCurrencyCreateConflict,
	"ErrorAccountNotExist":                          ErrorAccountNotExist,
	"ErrorIdentifierNotExist":                       ErrorIdentifierNotExist,
	"ErrorBankNotExist":                             ErrorBankNotExist,
	"ErrorCurrencyNotExist":                         ErrorCurrencyNotExist,
	"ErrorClientBankIsBlocked":                      ErrorClientBankIsBlocked,
	"ErrorAccountNotIdentified":                     ErrorAccountNotIdentified,
	"ErrorAccountNotAvailable":                      ErrorAccountNotAvailable,
	"ErrorAccountIsBlocked":                         ErrorAccountIsBlocked,
	"ErrorBankNotAvailable":                         ErrorBankNotAvailable,
	"ErrorBankIsBlocked":                            ErrorBankIsBlocked,
	"ErrorBankExist":                                ErrorBankExist,
	"ErrorContractNotExists":                        ErrorContractNotExists,
	"ErrorContractExists":                           ErrorContractExists,
	"ErrorCustomerNotFound":                         ErrorCustomerNotFound,
	"ErrorClientBankNotAvailable":                   ErrorClientBankNotAvailable,
	"ErrorSignVerify":                               ErrorSignVerify,
	"ErrorFundsNotEnough":                           ErrorFundsNotEnough,
	"ErrorValidateDefault":                          ErrorValidateDefault,
	"ErrorAddressNotFollowingRegex":                 ErrorAddressNotFollowingRegex,
	"ErrorAddressNotPassed":                         ErrorAddressNotPassed,
	"ErrorIdentifierNotFolowingRegex":               ErrorIdentifierNotFolowingRegex,
	"ErrorIdentifierNotPassed":                      ErrorIdentifierNotPassed,
	"ErrorIdentityTypeIncorrect":                    ErrorIdentityTypeIncorrect,
	"ErrorStateIncorrect":                           ErrorStateIncorrect,
	"ErrorJuridicalTypeIncorrect":                   ErrorJuridicalTypeIncorrect,
	"ErrorRoleIncorrect":                            ErrorRoleIncorrect,
	"ErrorChannelNotPassed":                         ErrorChannelNotPassed,
	"ErrorNameNotPassed":                            ErrorNameNotPassed,
	"ErrorVersionNotPassed":                         ErrorVersionNotPassed,
	"ErrorAmountNegative":                           ErrorAmountNegative,
	"ErrorAmountNotPassed":                          ErrorAmountNotPassed,
	"ErrorIdentityTypeNotPassed":                    ErrorIdentityTypeNotPassed,
	"ErrorValueNegative":                            ErrorValueNegative,
	"ErrorValueNotPassed":                           ErrorValueNotPassed,
	"ErrorCurrencyCodeRange":                        ErrorCurrencyCodeRange,
	"ErrorCurrencyCodeNotPassed":                    ErrorCurrencyCodeNotPassed,
	"ErrorBalanceNegative":                          ErrorBalanceNegative,
	"ErrorBalanceNotPassed":                         ErrorBalanceNotPassed,
	"ErrorMsgHashNotFolowingRegex":                  ErrorMsgHashNotFolowingRegex,
	"ErrorMsgHashNotPassed":                         ErrorMsgHashNotPassed,
	"ErrorSigRNotFolowingRegex":                     ErrorSigRNotFolowingRegex,
	"ErrorSigRHashNotPassed":                        ErrorSigRHashNotPassed,
	"ErrorSigSNotFolowingRegex":                     ErrorSigSNotFolowingRegex,
	"ErrorSigSNotPassed":                            ErrorSigSNotPassed,
	"ErrorSigVIncorrect":                            ErrorSigVIncorrect,
	"ErrorSigVNotFollowingRegex":                    ErrorSigVNotFollowingRegex,
	"ErrorTxIdSNotPassed":                           ErrorTxIdSNotPassed,
	"ErrorTimestampNotPassed":                       ErrorTimestampNotPassed,
	"ErrorTimestampNegative":                        ErrorTimestampNegative,
	"ErrorTxTypeNotPassed":                          ErrorTxTypeNotPassed,
	"ErrorTxTypeIncorrect":                          ErrorTxTypeIncorrect,
	"ErrorPageSizeMinValue":                         ErrorPageSizeMinValue,
	"ErrorSymbolNotPassed":                          ErrorSymbolNotPassed,
	"ErrorDecimalsNotPassed":                        ErrorDecimalsNotPassed,
	"ErrorDecimalsRange":                            ErrorDecimalsRange,
	"ErrorDecimalsMaxValue":                         ErrorDecimalsMaxValue,
	"ErrorEnabledNotPassed":                         ErrorEnabledNotPassed,
	"ErrorIdentifiersNotPassed":                     ErrorIdentifiersNotPassed,
	"ErrorIdentifiersMaxValue":                      ErrorIdentifiersMaxValue,
	"ErrorStateNotPassed":                           ErrorStateNotPassed,
	"ErrorJuridicalTypeNotPassed":                   ErrorJuridicalTypeNotPassed,
	"ErrorJuridicalTypeBankSetterNotPassed":         ErrorJuridicalTypeBankSetterNotPassed,
	"ErrorAccountTypeNotPassed":                     ErrorAccountTypeNotPassed,
	"ErrorAccountTypeIncorrect":                     ErrorAccountTypeIncorrect,
	"ErrorRoleNotPassed":                            ErrorRoleNotPassed,
	"ErrorIdNotPassed":                              ErrorIdNotPassed,
	"ErrorIdNotFolowingRegex":                       ErrorIdNotFolowingRegex,
	"ErrorNumberNotPassed":                          ErrorNumberNotPassed,
	"ErrorNumberInvoiceNotFolowingRegex":            ErrorNumberInvoiceNotFolowingRegex,
	"ErrorInvoiceStateIncorrect":                    ErrorInvoiceStateIncorrect,
	"ErrorInvoiceStateNotPassed":                    ErrorInvoiceStateNotPassed,
	"ErrorIssueLimitNegative":                       ErrorIssueLimitNegative,
	"ErrorStatusTransactionIncorrect":               ErrorStatusTransactionIncorrect,
	"ErrorStatusTransactionPassed":                  ErrorStatusTransactionPassed,
	"ErrorMspIdNotMatch":                            ErrorMspIdNotMatch,
	"ErrorAddressNotMatch":                          ErrorAddressNotMatch,
	"ErrorPublicKeyNotFolowingRegex":                ErrorPublicKeyNotFolowingRegex,
	"ErrorPublicKeyNotPassed":                       ErrorPublicKeyNotPassed,
	"ErrorConfNotPassed":                            ErrorConfNotPassed,
	"ErrorBankAddressNotPassed":                     ErrorBankAddressNotPassed,
	"ErrorCountryCodeNotPassed":                     ErrorCountryCodeNotPassed,
	"ErrorTransactionIdNotPassed":                   ErrorTransactionIdNotPassed,
	"ErrorAddressOrIdentifierNotFolowingRegex":      ErrorAddressOrIdentifierNotFolowingRegex,
	"ErrorAddressOrIdentifierNotPassed":             ErrorAddressOrIdentifierNotPassed,
	"ErrorBankDisplayNameNotPassed":                 ErrorBankDisplayNameNotPassed,
	"ErrorGetMspId":                                 ErrorGetMspId,
	"ErrorExpirationSign":                           ErrorExpirationSign,
	"ErrorSignUsed":                                 ErrorSignUsed,
	"ErrorChangeStateSafeDealForbidden":             ErrorChangeStateSafeDealForbidden,
	"ErrorSettingsFSMSafeDeal":                      ErrorSettingsFSMSafeDeal,
	"ErrorSafeDealExists":                           ErrorSafeDealExists,
	"ErrorNotValidAddressAcceptor":                  ErrorNotValidAddressAcceptor,
	"ErrorAddressNotCounterpartDeal":                ErrorAddressNotCounterpartDeal,
	"ErrorCounterpartMadeChoice":                    ErrorCounterpartMadeChoice,
	"ErrorNotValidAddressAcceptorOrArbitrator":      ErrorNotValidAddressAcceptorOrArbitrator,
	"ErrorNotValidAddressInitiator":                 ErrorNotValidAddressInitiator,
	"ErrorCounterpartWasInvited":                    ErrorCounterpartWasInvited,
	"ErrorParticipantNotInvited":                    ErrorParticipantNotInvited,
	"ErrorNotValidAddressOwner":                     ErrorNotValidAddressOwner,
	"ErrorEnoughParticipants":                       ErrorEnoughParticipants,
	"ErrorIncorrectSafeDealAmount":                  ErrorIncorrectSafeDealAmount,
	"ErrorSafeDealEscrowNotFound":                   ErrorSafeDealEscrowNotFound,
	"ErrorIncorrectCurrencyCodeAccount":             ErrorIncorrectCurrencyCodeAccount,
	"ErrorArbitratorExist":                          ErrorArbitratorExist,
	"ErrorArbitratorNotExist":                       ErrorArbitratorNotExist,
	"ErrorArbitratorCannotPartialPerform":           ErrorArbitratorCannotPartialPerform,
	"ErrorPhysicalAssetNotPresent":                  ErrorPhysicalAssetNotPresent,
	"ErrorPartialPerformOrder":                      ErrorPartialPerformOrder,
	"ErrorOfferAccountSellOnly":                     ErrorOfferAccountSellOnly,
	"ErrorContractNotOffer":                         ErrorContractNotOffer,
	"ErrorOfferOwnerMissing":                        ErrorOfferOwnerMissing,
	"ErrorCurrencyNotMatchOffer":                    ErrorCurrencyNotMatchOffer,
	"ErrorPriceNotMatchOffer":                       ErrorPriceNotMatchOffer,
	"ErrorAssetAmountNotCorrespondMaxMinInterval":   ErrorAssetAmountNotCorrespondMaxMinInterval,
	"ErrorPhysicalAssetEscrow":                      ErrorPhysicalAssetEscrow,
	"ErrorAmountNotCorrespondPrice":                 ErrorAmountNotCorrespondPrice,
	"ErrorSafeDealNotFound":                         ErrorSafeDealNotFound,
	"ErrorCounterpartVoted":                         ErrorCounterpartVoted,
	"ErrorIncorrectEscrowAmount":                    ErrorIncorrectEscrowAmount,
	"ErrorAccountAlreadyInDeal":                     ErrorAccountAlreadyInDeal,
	"ErrorAvailablePlatformsNotFound":               ErrorAvailablePlatformsNotFound,
	"ErrorLimitOperation":                           ErrorLimitOperation,
	"ErrorLimitDaily":                               ErrorLimitDaily,
	"ErrorLimitMonthly":                             ErrorLimitMonthly,
	"ErrorLimitBalance":                             ErrorLimitBalance,
	"ErrorAddressFromToEqual":                       ErrorAddressFromToEqual,
	"ErrorBankCannotBeBlocked":                      ErrorBankCannotBeBlocked,
	"ErrorCrossTransactionHistoryNotFound":          ErrorCrossTransactionHistoryNotFound,
	"ErrorStatusConflict":                           ErrorStatusConflict,
	"ErrorCrossTransactionWithdrawIncorrectAddress": ErrorCrossTransactionWithdrawIncorrectAddress,
	"ErrorCrossTransactionWithdrawIncorrectAmount":  ErrorCrossTransactionWithdrawIncorrectAmount,
	"ErrorDefault":                                  ErrorDefault,
	"ErrorInvoiceNotExist":                          ErrorInvoiceNotExist,
	"ErrorInvoiceTransactionAmount":                 ErrorInvoiceTransactionAmount,
	"ErrorInvoiceTransactionPaid":                   ErrorInvoiceTransactionPaid,
	"ErrorInvoiceTransactionExpired":                ErrorInvoiceTransactionExpired,
	"ErrorInvoiceExist":                             ErrorInvoiceExist,
	"ErrorInvoiceUpdate":                            ErrorInvoiceUpdate,
	"ErrorPayerNotEqual":                            ErrorPayerNotEqual,
	"ErrorPutState":                                 ErrorPutState,
	"ErrorGetState":                                 ErrorGetState,
	"ErrorDeleteState":                              ErrorDeleteState,
	"ErrorCreateCompositeKey":                       ErrorCreateCompositeKey,
	"ErrorJsonMarshal":                              ErrorJsonMarshal,
	"ErrorJsonUnmarshal":                            ErrorJsonUnmarshal,
	"ErrorGetTxTime":                                ErrorGetTxTime,
	"ErrorForbidden":                                ErrorForbidden,
	"ErrorDepositNotFound":                          ErrorDepositNotFound,
	"ErrorContractValidate":                         ErrorContractValidate,
	"ErrorCertificateBankAddressNotMatch":           ErrorCertificateBankAddressNotMatch,
	"ErrorClearingBadClaimsLiabilities":             ErrorClearingBadClaimsLiabilities,
	"ErrorClearingInfoNotFound":                     ErrorClearingInfoNotFound,
	"ErrorCertificateBankAddressNotFound":           ErrorCertificateBankAddressNotFound,
	"ErrorClientBankExists":                         ErrorClientBankExists,
	"ErrorClientBankNotFound":                       ErrorClientBankNotFound,
	"ErrorClientBankOwnerNotEqualSender":            ErrorClientBankOwnerNotEqualSender,
	"ErrorLimitTypeIncorrect":                       ErrorLimitTypeIncorrect,
	"ErrorAccountTechnicalNotEqlSender":             ErrorAccountTechnicalNotEqlSender,
	"ErrorAccountNotEscrow":                         ErrorAccountNotEscrow,
}
