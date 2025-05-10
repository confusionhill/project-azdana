package auth

import (
	"encoding/xml"

	"github.com/google/go-querystring/query"
)

type LoginUserRequestDTO struct {
	Username string `form:"strUsername"`
	Password string `form:"strPassword"`
}

type RegisterUserRequestDTO struct {
	Email     string `form:"strEmail"`
	Age       int    `form:"intAge"`
	Gender    string `form:"strGender"`
	HairColor int    `form:"intColorHair"`
	SkinColor int    `form:"intColorSkin"`
	HairID    int    `form:"HairID"`
	EyeColor  int    `form:"intColorEye"`
	DOB       string `form:"strDOB"`
	Password  string `form:"strPassword"`
	Username  string `form:"strUsername"`
	ClassID   int    `form:"ClassID"`
}

type RegisterUserResponseDTO struct {
	Status          string `url:"status"`
	UserId          *int64 `url:"userid"`
	StrErr          string `url:"strErr"`
	StrReason       string `url:"strReason"`
	StrButtonName   string `url:"strButtonName"`
	StrButtonAction string `url:"strButtonAction"`
	StrMsg          string `url:"strMsg"`
}

func (r RegisterUserResponseDTO) ToString() string {
	values, _ := query.Values(r)
	return values.Encode()
}

type LoginResponseDTO struct {
	XMLName    xml.Name            `xml:"login"`
	BSuccess   int64               `xml:"bSuccess,attr"`
	SMsg       string              `xml:"sMsg,attr,omitempty"`
	IAccess    int64               `xml:"iAccess,attr,omitempty"`
	IUpg       int64               `xml:"iUpg,attr,omitempty"`
	IAge       int64               `xml:"iAge,attr,omitempty"`
	SToken     string              `xml:"sToken,attr,omitempty"`
	DUpgExp    int64               `xml:"dUpgExp,attr,omitempty"`
	IUpgDays   int64               `xml:"iUpgDays,attr,omitempty"`
	ISendEmail int64               `xml:"iSendEmail,attr,omitempty"`
	StrEmail   string              `xml:"strEmail,attr,omitempty"`
	BCCOnly    int64               `xml:"bCCOnly,attr"`
	Servers    []ServerResponseDTO `xml:"servers,omitempty"`
}

type ServerResponseDTO struct {
	SName   string `xml:"sName,attr"`
	SIP     string `xml:"sIP,attr"`
	ICount  int64  `xml:"iCount,attr"`
	IMax    int64  `xml:"iMax,attr"`
	BOnline int64  `xml:"bOnline,attr"`
	BChat   int64  `xml:"bChat,attr"`
	IChat   int64  `xml:"iChat,attr"`
	BUpg    int64  `xml:"bUpg,attr"`
}

type GameLoginRequestDTO struct {
	XMLName xml.Name         `xml:"msg"`
	T       string           `xml:"t,attr"`
	Body    GameLoginBodyDTO `xml:"body"`
}

type GameLoginBodyDTO struct {
	Action string           `xml:"action,attr"`
	R      int              `xml:"r,attr"`
	Login  GameLoginDataDTO `xml:"login"`
}

type GameLoginDataDTO struct {
	Z     string `xml:"z,attr"`
	Nick  string `xml:"nick"`
	Pword string `xml:"pword"`
}

// func main() {
// 	xmlData := `
// 	<msg t='sys'>
// 		<body action='login' r='0'>
// 			<login z='zone_master'>
// 				<nick><![CDATA[dukun]]></nick>
// 				<pword><![CDATA[dukun123]]></pword>
// 			</login>
// 		</body>
// 	</msg>`

// 	var msg Msg
// 	err := xml.Unmarshal([]byte(xmlData), &msg)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Printf("Parsed struct: %+v\n", msg)
// }
