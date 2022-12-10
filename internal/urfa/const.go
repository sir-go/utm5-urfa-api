package urfa

// All URFA protocol constants

// packet attribute logintype
//noinspection GoUnusedConst
const (
	LoginTypeRegular uint32 = iota + 1 // 1 it's a regular user
	LoginTypeSystem                    // 2 it's a system user
	LoginTypeCard                      // 3 it's a card user
)

// packet attribute sslrequest (version of SSL)
//noinspection GoUnusedConst
const (
	SslReqTypeTLS1        = iota + 1 // 1
	SslReqTypeSSL3                   // 2
	SslReqTypeSSL3Cert               // 3
	SslReqTypeSSL3RsaCert            // 4
	SslReqTypeTLS12                  // 5
)

// termination attribute codes (reason of the session termination)
//noinspection GoUnusedConst
const (
	TermValueAccessRejected uint32 = 1 // access for this user is denied
	TermValueIllegalCode    uint32 = 3 // unknown call code was received
	TermValueDone           uint32 = 4 // session is just done... whatever
	TermValueForbidden      uint32 = 7 // user hasn't permission for using an API
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

// BinChunkSize - how many bytes to send at a time
const BinChunkSize = 1024
