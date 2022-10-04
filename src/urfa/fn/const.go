package fn

// search criteria
//noinspection GoUnusedConst
const (
	SearchCriteriaLike    = iota + 1 // 1
	_                                //
	SearchCriteriaEq                 // 3
	SearchCriteriaNotEq              // 4
	_                                //
	_                                //
	SearchCriteriaGt                 // 7
	SearchCriteriaLt                 // 8
	SearchCriteriaGorEt              // 9
	SearchCriteriaLorEt              // 10
	SearchCriteriaNotLike            // 11
)

// search type
//noinspection GoUnusedConst
const (
	SelectTypeAnd = iota // 0
	SelectTypeOr         // 1
)

// user fields
//noinspection GoUnusedConst
const (
	UserFieldUserId           = iota + 1 // 1
	UserFieldUserLogin                   // 2
	UserFieldUserAccount                 // 3
	UserFieldDiscountPeriodId            // 4
	UserFieldFullname                    // 5
	UserFieldCreateDate                  // 6
	UserFieldLastChangeDate              // 7
	UserFieldWhoCreate                   // 8
	UserFieldWhoChange                   // 9
	UserFieldIsJur                       // 10
	UserFieldJurAddress                  // 11
	UserFieldActualAddress               // 12
	UserFieldWorkPhone                   // 13
	UserFieldHomePhone                   // 14
	UserFieldMobilePhone                 // 15
	UserFieldWebPage                     // 16
	UserFieldIcqUin                      // 17
	UserFieldTaxNumber                   // 18
	UserFieldKppNumber                   // 19
	_                                    //
	UserFieldHouseId                     // 21
	UserFieldFlatNumber                  // 22
	UserFieldEntrance                    // 23
	UserFieldFloor                       // 24
	UserFieldEmail                       // 25
	UserFieldPassport                    // 26
	_                                    //
	UserFieldIp                          // 28
	_                                    //
	UserFieldGroupId                     // 30
	UserFieldBalance                     // 31
	UserFieldPersonalManager             // 32
	UserFieldConnectDate                 // 33
	UserFieldComments                    // 34
	UserFieldInternetStatus              // 35
	UserFieldTariffId                    // 36
	UserFieldServiceId                   // 37
	UserFieldSlinkId                     // 38
	UserFieldTplinkId                    // 39
	UserFieldDistrict                    // 40
	UserFieldBuilding                    // 41
	UserFieldMac                         // 42
	UserFieldSlinkLogin                  // 43
	UserFieldExternalId                  // 44
)

// contact boss position
//noinspection GoUnusedConst
const (
	ContactBossPositionCEO = iota + 1 // 1
	ContactBossPositionBKR            // 2
)

// user log actions
//noinspection GoUnusedConst
const (
	ActUnknown                          = iota // 0
	ActAccountReversing                        // 1
	ActBalanceCorrection                       // 2
	ActAddUser                                 // 3
	ActEditUser                                // 4
	ActDeleteUser                              // 5
	ActAddIpGroup                              // 6
	ActEditIpGroup                             // 7
	ActAddUserGroup                            // 8
	ActDeleteUserGroup                         // 9
	ActEditUserGroup                           // 10
	ActAddUserToGroup                          // 11
	ActRemoveUserFromGroup                     // 12
	ActAddAccount                              // 13
	ActEditAccount                             // 14
	ActAddService                              // 15
	ActEditService                             // 16
	ActUsersGroupOperations                    // 17
	ActAddTariff                               // 18
	ActEditTariff                              // 19
	ActDeleteTariff                            // 20
	ActAddServiceToTariff                      // 21
	ActDeleteServieceFromTariff                // 22
	ActUnlinkTariffFromAccount                 // 23
	ActLinkTariffToAccount                     // 24
	ActAddAdditionalSettingsParamter           // 25
	ActEditAdditionalSettingsParamter          // 26
	ActDeleteAdditionalSettingsParamter        // 27
	ActInstantTariffChange                     // 28
	ActPolicyChanged                           // 29
	ActAddTecParam                             // 30
	ActEditTecParam                            // 31
	ActDeleteTecParam                          // 32
	ActSetVoluntaryBlocking                    // 33
	ActDeleteVoluntaryBlocking                 // 34
	ActAddServiceLink                          // 35
	ActDeleteServiceLink                       // 36
	ActCostCoefChange                          // 37
)

// service type
//noinspection GoUnusedConst
const (
	ServiceTypeOnce      = iota + 1 // 1
	ServiceTypePeriodic             // 2
	ServiceTypeIpTraffic            // 3
	ServiceTypeHotspot              // 4
	ServiceTypeDialup               // 5
	ServiceTypeVoice                // 6
)

// traffic local policy
//noinspection GoUnusedConst
const (
	TrafficLocalPolicyToReciever = iota // 0
	TrafficLocalPolicyToSender          // 1
	TrafficLocalPolicyBoth              // 2
)

// period type
//noinspection GoUnusedConst
const (
	PeriodTypeDaily     = iota + 1 // 1
	PeriodTypeWeekly               // 2
	PeriodTypeMonthly              // 3
	PeriodTypeQuarterly            // 4
	PeriodTypeYearly               // 5
	PeriodTypeCustom    = 0x100000
)

// traffic report type
//noinspection GoUnusedConst
const (
	TrafficReportTypeByHours  = iota + 1 // 1
	TrafficReportTypeByDays              // 2
	TrafficReportTypeByMonths            // 3
	TrafficReportTypeByIp                // 4
)

// charge method
//noinspection GoUnusedConst
const (
	ChargeMethodAtPeriodStart = iota + 1 // 1
	ChargeMethodAtPeriodEnd              // 2
	ChargeMethodFlow                     // 3
)

// dynam shape
//noinspection GoUnusedConst
const (
	DynShapeApplyForVPN    = iota + 1 // 1
	DynShapeApplyForNonVPN            // 2
	DynShapeApplyForBoth              // 3
)

// message reciever type
//noinspection GoUnusedConst
const (
	MsgRecieverTypeUser         = iota // 0
	MsgRecieverTypeUserGroup           // 1
	MsgRecieverTypeSysUser             // 2
	MsgRecieverTypeSysUserGroup        // 3
	MsgRecieverTypeAll                 // 4
)

// message flag
//noinspection GoUnusedConst
const (
	MsgFlagUnreaded  = 1 << 0
	MsgFlagAnswered  = 1 << 1
	MsgFlagForwarded = 1 << 2
	MsgFlagImportant = 1 << 3
	MsgFlagDeleted   = 1 << 4
)

// event type
//noinspection GoUnusedConst
const (
	EventTypeUndefined         = 0x00000000
	EventTypeInternetOn        = 0x00000001
	EventTypeInternetOff       = 0x00000002
	EventTypeBlockTypeChanged  = 0x00000020
	EventTypeUserAdded         = 0x00000040
	EventTypeUserModifed       = 0x00000080
	EventTypeUserDeleted       = 0x00000100
	EventTypeTparamAdded       = 0x00000200
	EventTypeTparamModifed     = 0x00000400
	EventTypeTparamDeleted     = 0x00000800
	EventTypeTplinkAdded       = 0x00001000
	EventTypeTplinkModifed     = 0x00002000
	EventTypeTplinkDeleted     = 0x00004000
	EventTypeRawTrafFileClosed = 0x00008000
	EventTypeLogFileClosed     = 0x00010000
	EventTypeHotspotEnabled    = 0x00020000
	EventTypeHotspotDisabled   = 0x00040000
	EventTypeSessionOpened     = 0x00080000
	EventTypeSessionClosed     = 0x00100000
	EventTypeSetBwLimitIn      = 0x00200000
	EventTypeEditBwLimitIn     = 0x00400000
	EventTypeDelBwLimitIn      = 0x00800000
	EventTypeSetBwLimitOut     = 0x01000000
	EventTypeEditBwLimitOut    = 0x02000000
	EventTypeDelBwLimitOut     = 0x04000000
	EventTypeBalanceNotifSent  = 0x40000000
)

// block type
//noinspection GoUnusedConst
const (
	BlockTypeNone         = 0x000
	BlockTypeUser         = 0x001
	BlockTypeSystem       = 0x010
	BlockTypeSystemRecAb  = 0x020
	BlockTypeSystemRecPay = 0x040
	BlockTypeAdmin        = 0x100
	BlockTypeAdminRecAb   = 0x200
	BlockTypeAdminRecPay  = 0x400
)

// payment type
//noinspection GoUnusedConst
const (
	PaymentTypeCash         = 1
	PaymentTypeWireTransfer = 2
	PaymentTypeCard         = 3
	PaymentTypeCredit       = 7
	PaymentTypeRollback     = 8
)

// time
//noinspection GoUnusedConst
const (
	TimeMax = 2e9
	TimeMin = 0
)

// packet attribute logintype
//noinspection GoUnusedConst
const (
	LoginTypeRegular uint32 = iota + 1 // 1
	LoginTypeSystem                    // 2
	LoginTypeCard                      // 3
)
