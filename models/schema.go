package models

import "time"

// Library: (ID, Name)
type Library struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `form:"name" json:"name" validate:"required" binding:"required"`
}

// Users: (ID, Name, Email, ContactNumber, Role, LibID)
type Users struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `json:"name" validate:"required"`
	Email         string `json:"email" validate:"required, email"`
	ContactNumber string `json:"contactNumber" validate:"required"`
	Role          string `json:"role" validate:"required, oneof reader admin owner"`
	LibID         uint   `json:"libId" validate:"required"`
}

// BookInventory : (ISBN, LibID, Title, Authors, Publisher, Version, TotalCopies, AvailableCopies)
type BookInventory struct {
	ISBN            uint   `json:"isbn" gorm:"primaryKey;autoIncrement:true"`
	LibID           uint   `json:"libId" validate:"required"`
	Title           string `json:"title" validate:"required"`
	Authors         string `json:"authors" validate:"required"`
	Publisher       string `json:"publisher" validate:"required"`
	Version         uint   `json:"version" validate:"required"`
	TotalCopies     uint   `json:"totalCopies" validate:"required"`
	AvailableCopies uint   `json:"availableCopies"`
}

// RequestEvents: (ReqID, BookID, ReaderID, RequestDate, ApprovalDate, ApproverID, RequestType)
type RequestEvents struct {
	ReqID        uint      `gorm:"primaryKey;autoIncrement:true"`
	BookID       uint      `json:"bookId" validate:"required"`
	ReaderID     string    `json:"readerId" validate:"required"`
	RequestDate  time.Time `gorm:"default:current_timestamp"`
	ApprovalDate time.Time `gorm:"type:timestamp;default:null"`
	ApproverID   string    `gorm:"default:''"`
	RequestType  string    `json:"requestType" validate:"required"`
}

// IssueRegistery: (IssueID, ISBN, ReaderID, IssueApproverID, IssueStatus, IssueDate
type IssueRegistery struct {
	IssueID            uint `gorm:"primaryKey;autoIncrement:true"`
	ISBN               uint
	ReaderID           string
	IssueApproverID    uint
	IssueStatus        string
	IssueDate          time.Time
	ExpectedReturnDate time.Time
	ReturnDate         time.Time
	ReturnApproverID   uint
}
