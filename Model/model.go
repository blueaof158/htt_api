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
	AwardID          int    `json:"awardid"`
	AwardName        string `json:"awardname"`
	AwardDescription string `json:"awarddescription"`
	AwardInactive    int    `json:"awardInactive"`
	CreateBy         string `json:"createby"`
	CreateDate       string `json:"createdate"`
	UpdateBy         string `json:"updateby"`
	UpdateDate       string `json:"updatedate"`
	Image64          string `json:"image64"`
}

type Car struct {
	CarID          int    `json:"carid"`
	CarName        string `json:"carname"`
	CarDesctiption string `json:"cardesctiption"`
	CarTypeID      int    `json:"cartypeID"`
	CarInactive    int    `json:"carinactive"`
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

type CompanyInfo struct {
	Advise          string `json:"Advise"`
	CompanyAddress  string `json:"CompanyAddress"`
	CompanyName     string `json:"CompanyName"`
	ContactUs       string `json:"ContactUs"`
	Image641        string `json:"Image641"`
	Image642        string `json:"Image642"`
	Image643        string `json:"Image643"`
	JuristicID      string `json:"JuristicID"`
	OurService      string `json:"OurService"`
	AboutUs         string `json:"AboutUs"`
	ServiceLocation string `json:"ServiceLocation"`
	WhyUseUs1       string `json:"WhyUseUs1"`
	WhyUseUs2       string `json:"WhyUseUs2"`
	WhyUseUs3       string `json:"WhyUseUs3"`
	WhyUseUsTitl1   string `json:"WhyUseUsTitl1"`
	WhyUseUsTitl2   string `json:"WhyUseUsTitl2"`
	WhyUseUsTitl3   string `json:"WhyUseUsTitl3"`
	ContactNumber   string `json:"ContactNumber"`
	FacebookLink    string `json:"FacebookLink"`
	LineLink        string `json:"LineLink"`
	GoogleMap       string `json:"GoogleMap"`
}

type Content struct {
	ContentID        int    `json:"contentid"`
	ContentTitle     string `json:"contenttitle"`
	HyphenationTitle string `json:"Hyphenationtitle"`
	ContentText      string `json:"contenttext"`
	Content          string `json:"content"`
	ContentInactive  int    `json:"contentinactive"`
	CreateBy         string `json:"createby"`
	CreateDate       string `json:"createdate"`
	UpdateBy         string `json:"updateby"`
	UpdateDate       string `json:"updatedate"`
	Image64          string `json:"image64"`
}

type Executives struct {
	ExecutivesID        int    `json:"executivesid"`
	ExecutivesFirstName string `json:"executivesfirstname"`
	ExecutivesLastName  string `json:"executiveslastname"`
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
	Password   string `json:"password"`
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

type FrontendBannerTop struct {
	BannerTopImageLink string `json:"bannertopimagelink"`
	ImagePath          string `json:"imagepath"`
}

type FrontendAward struct {
	AwardName        string `json:"awardname"`
	AwardDescription string `json:"awarddescription"`
	Image64          string `json:"image64"`
}

type FrontendCarType struct {
	CarTypeID          int    `json:"cartypeid"`
	CarTypeName        string `json:"cartypename"`
	CarTypeDesctiption string `json:"cartypedesctiption"`
	CarTypeInactive    int    `json:"cartypeinactive"`
	Image64            string `json:"image64"`
}

type FrontendContent struct {
	ContentID    int    `json:"contentid"`
	ContentTitle string `json:"contenttitle"`
	ContentText  string `json:"contenttext"`
	Content      string `json:"content"`
	Image64      string `json:"image64"`
	CreateDate   string `json:"createdate"`
}

type FrontendCar struct {
	CarID          int      `json:"carid"`
	CarName        string   `json:"carname"`
	CarDesctiption string   `json:"cardesctiption"`
	CarTypeID      int      `json:"cartypeID"`
	Image64        []string `json:"image64"`
}

type Login struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type FrontendJobApplications struct {
	JobApplicationsID          int    `json:"jobapplicationsid"`
	JobApplicationsName        string `json:"jobapplicationsname"`
	JobApplicationsPosition    string `json:"jobapplicationsposition"`
	JobApplicationsDescription string `json:"jobapplicationsdescription"`
}
