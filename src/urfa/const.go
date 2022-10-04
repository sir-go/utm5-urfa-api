package urfa

// packet attribute logintype
//noinspection GoUnusedConst
const (
	LoginTypeRegular uint32 = iota + 1 // 1
	LoginTypeSystem                    // 2
	LoginTypeCard                      // 3
)

// packet attribute sslrequest
//noinspection GoUnusedConst
const (
	SslReqTypeTLS1        = iota + 1 // 1
	SslReqTypeSSL3                   // 2
	SslReqTypeSSL3Cert               // 3
	SslReqTypeSSL3RsaCert            // 4
	SslReqTypeTLS12                  // 5
)

// termination attribute codes
//noinspection GoUnusedConst
const (
	TermValueAccessRejected uint32 = 1
	TermValueIllegalCode    uint32 = 3
	TermValueDone           uint32 = 4
	TermValueForbidden      uint32 = 7
)

// urfa packet
//noinspection GoUnusedConst
const (
	PacketHeaderSize uint16 = 4
	PacketVersion    uint8  = 35

	PacketTypeSessionInit   uint8 = 192 // c0
	PacketTypeAccessRequest uint8 = 193 // c1
	PacketTypeAccessAccept  uint8 = 194 // c2
	PacketTypeAccessReject  uint8 = 195 // c3
	PacketTypeSessionData   uint8 = 200 // c8
	PacketTypeSessionCall   uint8 = 201 // c9
	PacketTypeSessionTerm   uint8 = 203 // cb
)

// urfa packet attribute
//noinspection GoUnusedConst
const (
	AttrTypeLoginType     uint8 = iota + 1 // 1
	AttrTypeLogin                          // 2
	AttrTypeCall                           // 3
	AttrTypeTermination                    // 4
	AttrTypeData                           // 5
	AttrTypeKey                            // 6
	_                                      // 7
	AttrTypeCHAPChallenge                  // 8
	AttrTypeCHAPResponse                   // 9
	AttrTypeSSLRequest                     // 10

	AttrHeaderSize = 4
)

const BinChunkSize = 1024
