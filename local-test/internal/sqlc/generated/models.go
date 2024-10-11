// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlcgen

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type ReportsReason string

const (
	ReportsReasonSpam                 ReportsReason = "spam"
	ReportsReasonHarassment           ReportsReason = "harassment"
	ReportsReasonInappropriateContent ReportsReason = "inappropriate_content"
	ReportsReasonOther                ReportsReason = "other"
)

func (e *ReportsReason) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ReportsReason(s)
	case string:
		*e = ReportsReason(s)
	default:
		return fmt.Errorf("unsupported scan type for ReportsReason: %T", src)
	}
	return nil
}

type NullReportsReason struct {
	ReportsReason ReportsReason
	Valid         bool // Valid is true if ReportsReason is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullReportsReason) Scan(value interface{}) error {
	if value == nil {
		ns.ReportsReason, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ReportsReason.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullReportsReason) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ReportsReason), nil
}

type Account struct {
	ID          string
	UserID      string
	UserName    string
	IsSuspended bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Block struct {
	BlockerAccountID string
	BlockedAccountID string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Follow struct {
	FollowerAccountID  string
	FollowingAccountID string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type FollowRequest struct {
	RequesterAccountID string
	RequesteeAccountID string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Hashtag struct {
	ID        uint64
	Tag       string
	CreatedAt time.Time
}

type Interest struct {
	AccountID                          string
	NewsScore                          sql.NullInt16
	PoliticsScore                      sql.NullInt16
	EconomicsScore                     sql.NullInt16
	HealthScore                        sql.NullInt16
	SportsScore                        sql.NullInt16
	EntertainmentScore                 sql.NullInt16
	ArtScore                           sql.NullInt16
	CookingScore                       sql.NullInt16
	TravelScore                        sql.NullInt16
	FashionScore                       sql.NullInt16
	BeautyScore                        sql.NullInt16
	PetsScore                          sql.NullInt16
	ParentingScore                     sql.NullInt16
	EducationScore                     sql.NullInt16
	EnvironmentScore                   sql.NullInt16
	ClimateScore                       sql.NullInt16
	SpaceScore                         sql.NullInt16
	MentalHealthScore                  sql.NullInt16
	FitnessScore                       sql.NullInt16
	ReadingScore                       sql.NullInt16
	HistoryScore                       sql.NullInt16
	PhilosophyScore                    sql.NullInt16
	ReligionScore                      sql.NullInt16
	CultureScore                       sql.NullInt16
	VolunteeringScore                  sql.NullInt16
	SocialIssuesScore                  sql.NullInt16
	LawScore                           sql.NullInt16
	TaxesScore                         sql.NullInt16
	InvestmentScore                    sql.NullInt16
	RealEstateScore                    sql.NullInt16
	DiyScore                           sql.NullInt16
	GardeningScore                     sql.NullInt16
	InteriorDesignScore                sql.NullInt16
	AutomotiveScore                    sql.NullInt16
	GamingScore                        sql.NullInt16
	AnimeMangaScore                    sql.NullInt16
	CreativeWorksScore                 sql.NullInt16
	PhotographyVideoScore              sql.NullInt16
	MediaScore                         sql.NullInt16
	MarketingScore                     sql.NullInt16
	BrandingScore                      sql.NullInt16
	EntrepreneurshipScore              sql.NullInt16
	RemoteWorkScore                    sql.NullInt16
	DataScienceScore                   sql.NullInt16
	IotScore                           sql.NullInt16
	RoboticsEngineeringScore           sql.NullInt16
	BiotechnologyScore                 sql.NullInt16
	NanotechnologyScore                sql.NullInt16
	EnergyTechnologyScore              sql.NullInt16
	ArchaeologyScore                   sql.NullInt16
	PsychologyScore                    sql.NullInt16
	SociologyScore                     sql.NullInt16
	AnthropologyScore                  sql.NullInt16
	GeographyScore                     sql.NullInt16
	GeologyScore                       sql.NullInt16
	MeteorologyScore                   sql.NullInt16
	DisasterEmergencyManagementScore   sql.NullInt16
	UrbanPlanningScore                 sql.NullInt16
	ArchitectureScore                  sql.NullInt16
	AgricultureScore                   sql.NullInt16
	NutritionScienceScore              sql.NullInt16
	SleepScienceScore                  sql.NullInt16
	ProductivityScore                  sql.NullInt16
	LeadershipScore                    sql.NullInt16
	InternationalRelationsScore        sql.NullInt16
	FuturePredictionsScore             sql.NullInt16
	EventsScore                        sql.NullInt16
	CommunityScore                     sql.NullInt16
	TrendsScore                        sql.NullInt16
	LifestyleScore                     sql.NullInt16
	SoftwareDevelopmentScore           sql.NullInt16
	ProgrammingLanguagesScore          sql.NullInt16
	WebDevelopmentScore                sql.NullInt16
	MobileAppDevelopmentScore          sql.NullInt16
	DebuggingTechniquesScore           sql.NullInt16
	AlgorithmsMathematicsScore         sql.NullInt16
	DatabaseDesignScore                sql.NullInt16
	CloudComputingScore                sql.NullInt16
	ServerManagementScore              sql.NullInt16
	NetworkSecurityScore               sql.NullInt16
	CryptographyScore                  sql.NullInt16
	ArtificialIntelligenceScore        sql.NullInt16
	MachineLearningScore               sql.NullInt16
	DeepLearningScore                  sql.NullInt16
	ComputerVisionScore                sql.NullInt16
	NaturalLanguageProcessingScore     sql.NullInt16
	BlockchainTechnologyScore          sql.NullInt16
	QuantumComputingScore              sql.NullInt16
	EdgeComputingScore                 sql.NullInt16
	MicroservicesArchitectureScore     sql.NullInt16
	DevopsScore                        sql.NullInt16
	ContainerTechnologyScore           sql.NullInt16
	CiCdScore                          sql.NullInt16
	TestAutomationScore                sql.NullInt16
	UxUiDesignScore                    sql.NullInt16
	AgileDevelopmentMethodologiesScore sql.NullInt16
	OpenSourceScore                    sql.NullInt16
	VersionControlScore                sql.NullInt16
	ApiDesignScore                     sql.NullInt16
	PerformanceOptimizationScore       sql.NullInt16
}

type Label struct {
	TweetID   uint64
	Label1    string
	Label2    sql.NullString
	Label3    sql.NullString
	CreatedAt time.Time
}

type Like struct {
	LikingAccountID string
	OriginalTweetID uint64
	CreatedAt       time.Time
}

type Message struct {
	ID                 uint32
	SenderAccountID    string
	RecipientAccountID string
	Content            sql.NullString
	IsRead             bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Notification struct {
	ID                 uint32
	SenderAccountID    sql.NullString
	RecipientAccountID string
	Type               string
	Content            sql.NullString
	IsRead             bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Profile struct {
	AccountID       string
	Bio             sql.NullString
	ProfileImageUrl sql.NullString
	BannerImageUrl  sql.NullString
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Reply struct {
	ID                uint64
	OriginalTweetID   uint64
	ParentReplyID     sql.NullInt64
	ReplyingAccountID string
	CreatedAt         time.Time
}

type Report struct {
	ID                uint64
	ReporterAccountID string
	ReportedAccountID string
	Reason            ReportsReason
	Content           sql.NullString
	CreatedAt         time.Time
}

type RetweetsAndQuote struct {
	ID                  uint64
	RetweetingAccountID string
	OriginalTweetID     uint64
	CreatedAt           time.Time
}

type Setting struct {
	AccountID string
	IsPrivate bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tweet struct {
	ID              uint64
	AccountID       string
	IsPinned        bool
	Content         sql.NullString
	Code            sql.NullString
	LikesCount      uint32
	RepliesCount    uint32
	RetweetsCount   uint32
	IsRetweet       bool
	IsReply         bool
	IsQuote         bool
	EngagementScore uint32
	Media           json.RawMessage
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type TweetHashtag struct {
	TweetID   uint64
	HashtagID uint64
}