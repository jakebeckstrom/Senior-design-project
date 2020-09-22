package models

type FileArtifact struct {
	TicketId uint   `json:"ticketID" gorm:"primary_key"`
	Filename string `json:"filename" gorm:"primary_key"`
	Data     []byte `json:"-"`
}

func GetAssociatedArtifacts(TicketID uint) (*[]FileArtifact, error) {
	var fileArtifacts []FileArtifact

	err := db.Where("ticket_id = ?", TicketID).Find(&fileArtifacts).Error

	return &fileArtifacts, err
}

func GetArtifact(TicketID uint, FileName string) (*FileArtifact, error) {
	var fileArtifact FileArtifact

	err := db.Where("ticket_id = ? AND filename = ?", TicketID, FileName).First(&fileArtifact).Error

	return &fileArtifact, err
}
