package model

type CarType struct {
	CarTypeID          int    `json:"cartypeid"`
	CarTypeName        string `json:"cartypename"`
	CarTypeDesctiption string `json:"cartypedesctiption"`
	CarTypeInactive    int    `json:"cartypeinactive"`
	Image64            string `json:"image64"`
	CreateBy           string `json:"createby"`
	CreateDate         string `json:"createdate"`
	UpdateBy           string `json:"updateby"`
	UpdateDate         string `json:"updatedate"`
}

type CarTypeLst struct {
	CarTypeID   int    `json:"cartypeid"`
	CarTypeName string `json:"cartypename"`
}

type Award struct {
	AwardID          int    `json:"awardID"`
	AwardName        string `json:"awardname"`
	AwardDescription string `json:"awarddescription"`
	AwardInactive    int    `json:"awardInactive"`
	CreateBy         string `json:"createby"`
	CreateDate       string `json:"createdate"`
	UpdateBy         string `json:"updateby"`
	UpdateDate       string `json:"updatedate"`
}

type Car struct {
	CarID          int    `json:"carid"`
	CarName        string `json:"carname"`
	CarDesctiption string `json:"cardesctiption"`
	CarTypeID      int    `json:"cartypeID"`
	CarInactive    int    `json:"carInactive"`
	CreateBy       string `json:"createby"`
	CreateDate     string `json:"createdate"`
	UpdateBy       string `json:"updateby"`
	UpdateDate     string `json:"updatedate"`
	Image64        string `json:"image64"`
	CarTypeName    string `json:"cartypename"`
}

type ImageJson struct {
	Id        int      `json:"id"`
	ImageType string   `json:"imagetype"`
	Image64   []string `json:"image64"`
}

type Content struct {
	ContentID        int    `json:"contentid"`
	ContentTitle     string `json:"contenttitle"`
	HyphenationTitle string `json:"cyphenationtitle"`
	ContentText      string `json:"contenttext"`
	Content          string `json:"content"`
	ContentInactive  int    `json:"contentInactive"`
	CreateBy         string `json:"createby"`
	CreateDate       string `json:"createdate"`
	UpdateBy         string `json:"updateby"`
	UpdateDate       string `json:"updatedate"`
}

type Executives struct {
	ExecutivesID        int    `json:"executivesid"`
	ExecutivesFirstName string `json:"executivesfirstname"`
	ExecutivesLastName  string `json:"executiveslastName"`
	ExecutivesPosition  string `json:"executivesposition"`
	ExecutivesBio       string `json:"executivesbio"`
	ExecutivesInactive  int    `json:"executivesinactive"`
	CreateBy            string `json:"createby"`
	CreateDate          string `json:"createdate"`
	UpdateBy            string `json:"updateby"`
	UpdateDate          string `json:"updatedate"`
}

type Images struct {
	ImagesID     int    `json:"imagesid"`
	CarTypeID    int    `json:"cartypeid"`
	CarID        int    `json:"carid"`
	AwardID      int    `json:"awardID"`
	ContentID    int    `json:"contentid"`
	ExecutivesID int    `json:"executivesid"`
	BannerTopID  int    `json:"bannertopid"`
	ImagePath    string `json:"imagepath"`
	CreateBy     string `json:"createby"`
	CreateDate   string `json:"createdate"`
	UpdateBy     string `json:"updateby"`
	UpdateDate   string `json:"updatedate"`
}

type JobApplications struct {
	JobApplicationsID          int    `json:"jobapplicationsid"`
	JobApplicationsName        string `json:"jobapplicationsname"`
	JobApplicationsPosition    string `json:"jobapplicationsposition"`
	JobApplicationsDescription string `json:"jobapplicationsdescription"`
	JobApplicationsInactive    int    `json:"jobapplicationsinactive"`
	JobApplicationFilePath     string `json:"jobapplicationfilepath"`
	CreateBy                   string `json:"createby"`
	CreateDate                 string `json:"createdate"`
	UpdateBy                   string `json:"updateby"`
	UpdateDate                 string `json:"updatedate"`
}

type OtherConfigs struct {
	OtherConfigsID int    `json:"otherconfigsid"`
	ConfigName     string `json:"configname"`
	Value          string `json:"value"`
	CreateBy       string `json:"createby"`
	CreateDate     string `json:"createdate"`
	UpdateBy       string `json:"updateby"`
	UpdateDate     string `json:"updatedate"`
}

type Seo struct {
	SeoID      int    `json:"seoid"`
	ConfigName string `json:"configname"`
	Value      string `json:"value"`
	CreateBy   string `json:"createby"`
	CreateDate string `json:"createdate"`
	UpdateBy   string `json:"updateby"`
	UpdateDate string `json:"updatedate"`
}

type User struct {
	UserID     int    `json:"userid"`
	User       string `json:"user"`
	Password   string `json:"uassword"`
	Inactive   int    `json:"inactive"`
	CreateBy   string `json:"createby"`
	CreateDate string `json:"createdate"`
	UpdateBy   string `json:"updateby"`
	UpdateDate string `json:"updatedate"`
}

type BannerTop struct {
	BannerTopID        int    `json:"bannertopid"`
	BannerTopImageLink string `json:"bannertopimagelink"`
	BannerTopInactive  int    `json:"bannertopinactive"`
	ImagePath          string `json:"imagepath"`
	CreateBy           string `json:"createby"`
	CreateDate         string `json:"createdate"`
	UpdateBy           string `json:"updateby"`
	UpdateDate         string `json:"updatedate"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
