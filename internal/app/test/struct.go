package test

type RobotTicket struct {
	//gorm.Model
	Id               int    `gorm:"INT(11) 'id'" json:"id"`
	Status           string `gorm:"INT(11) 'status'" json:"status"`
	MessageId        string `gorm:"VARCHAR(128) 'message_id'" json:"message_id"`
	ProcessingPerson string `gorm:"VARCHAR(32) 'processing_person'" json:"processing_person"`
	Desc             string `gorm:"VARCHAR(1024) 'desc'" json:"desc"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
