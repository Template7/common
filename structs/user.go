package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type userStatus int
type LoginClientOs int
type LoginChannel int
type Gender int

const (
	UserStatusBlock       userStatus = -1
	UserStatusInitialized userStatus = 0 // not finish basic info yet
	UserStatusActivate    userStatus = 1 // sign up finished, able to use the app

	LoginClientOsUnknown LoginClientOs = 0
	LoginClientOsIos     LoginClientOs = 1
	LoginClientOsAndroid LoginClientOs = 2

	LoginChannelUnknown  LoginChannel = 0
	LoginChannelMobile   LoginChannel = 1
	LoginChannelFacebook LoginChannel = 2
	LoginChannelGoogle   LoginChannel = 3

	GenderUnknown Gender = 0
	GenderMale    Gender = 1
	GenderFemale  Gender = 2
	GenderGay     Gender = 3 // distinguish 0 and 1 ?
	GenderLesbian Gender = 4 // distinguish 0 and 1 ?
)

type User struct {
	Id          *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" swaggerignore:"true"` // mongo default document id
	UserId      string              `json:"user_id" bson:"user_id"`
	BasicInfo   UserInfo            `json:"basic_info" bson:"basic_info"`
	Mobile      string              `json:"mobile" bson:"mobile" example:"+886987654321"` // +886987654321
	Email       string              `json:"email" bson:"email" example:"username@mail.com"`
	Status      userStatus          `json:"status" bson:"status"`
	LoginClient LoginInfo           `json:"login_info" bson:"login_info"`
	LastUpdate  int64               `json:"last_update" bson:"last_update"` // unix time in second
}

// fields which allows user to change by themselves
type UserInfo struct {
	NickName        string   `json:"nick_name" bson:"nick_name" binding:"required"`
	Avatar          string   `json:"Avatar" bson:"Avatar"` // s3 object url
	ProfilePictures []string `json:"profile_pictures" bson:"profile_pictures"`
	Birthday        int64    `json:"birthday" bson:"birthday"`
	Gender          Gender   `json:"gender" bson:"gender,omitempty"` // configable?
	Hobbies         []string `json:"hobbies" bson:"hobbies"`
	Bio             string   `json:"bio" bson:"bio"`
}

type LoginInfo struct {
	Os            LoginClientOs `json:"os" bson:"os"`
	Device        string        `json:"device" bson:"device" binding:"required"` // iPhoneN, PixelN, NoteN, ...
	Channel       LoginChannel  `json:"channel" bson:"channel"`
	ChannelUserId string        `json:"channel_user_id" bson:"channel_user_id"` // user id of the channel
}
