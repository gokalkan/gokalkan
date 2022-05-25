package gokalkan

// KeyUser - это пользователь ключа. По этому параметру можно определить:
// - держателя сертификата (ключа)
// - является юр. лицом или нет
type KeyUser string

const (
	KeyUserUnknown                       KeyUser = ""                   // Неизвестный пользователь ключа
	KeyUserIndividual                    KeyUser = "INDIVIDUAL"         // Физическое лицо
	KeyUserOrganization                  KeyUser = "ORGANIZATION"       // Юридическое лицо
	KeyUserCEO                           KeyUser = "CEO"                // Первый руководитель юридического лица, имеющий право подписи
	KeyUserCanSign                       KeyUser = "CAN_SIGN"           // Лицо, наделенное правом подписи
	KeyUserCanSignFinancial              KeyUser = "CAN_SIGN_FINANCIAL" // Лицо, наделенное правом подписи финансовых документов
	KeyUserHR                            KeyUser = "HR"                 // Сотрудник отдела кадров, наделенный правом подтверждать заявки на выпуск регистрационных свидетельств поданные от сотрудников юридического лица
	KeyUserEmployee                      KeyUser = "EMPLOYEE"           // Сотрудник организации
	KeyUserNCAPrivileges                 KeyUser = "NCA_PRIVILEGES"     // Полномочия в информационной системе НУЦ РК
	KeyUserNCAAdmin                      KeyUser = "NCA_ADMIN"          // Администратор НУЦ РК
	KeyUserNCAManager                    KeyUser = "NCA_MANAGER"        // Менеджер НУЦ РК
	KeyUserNCAOperator                   KeyUser = "NCA_OPERATOR"       // Оператор НУЦ РК
	KeyUserIdentification                KeyUser = "IDENTIFICATION"
	KeyUserIdentificationCON             KeyUser = "IDENTIFICATION_CON"
	KeyUserIdentificationRemote          KeyUser = "IDENTIFICATION_REMOTE"
	KeyUserIdentificationRemoteDigitalID KeyUser = "IDENTIFICATION_REMOTE_DIGITAL_ID"
)

//nolint:gochecknoglobals
var keyUserMap = map[KeyUser]string{
	KeyUserIndividual:                    "1.2.398.3.3.4.1.1",
	KeyUserOrganization:                  "1.2.398.3.3.4.1.2",
	KeyUserCEO:                           "1.2.398.3.3.4.1.2.1",
	KeyUserCanSign:                       "1.2.398.3.3.4.1.2.2",
	KeyUserCanSignFinancial:              "1.2.398.3.3.4.1.2.3",
	KeyUserHR:                            "1.2.398.3.3.4.1.2.4",
	KeyUserEmployee:                      "1.2.398.3.3.4.1.2.5",
	KeyUserNCAPrivileges:                 "1.2.398.3.3.4.2",
	KeyUserNCAAdmin:                      "1.2.398.3.3.4.2.1",
	KeyUserNCAManager:                    "1.2.398.3.3.4.2.2",
	KeyUserNCAOperator:                   "1.2.398.3.3.4.2.3",
	KeyUserIdentification:                "1.2.398.3.3.4.3",
	KeyUserIdentificationCON:             "1.2.398.3.3.4.3.1",
	KeyUserIdentificationRemote:          "1.2.398.3.3.4.3.2",
	KeyUserIdentificationRemoteDigitalID: "1.2.398.3.3.4.3.2.1",
}
