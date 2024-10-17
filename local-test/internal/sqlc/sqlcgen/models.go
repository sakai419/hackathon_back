// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlcgen

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/sqlc-dev/pqtype"
)

type FollowStatus string

const (
	FollowStatusPending  FollowStatus = "pending"
	FollowStatusAccepted FollowStatus = "accepted"
)

func (e *FollowStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = FollowStatus(s)
	case string:
		*e = FollowStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for FollowStatus: %T", src)
	}
	return nil
}

type NullFollowStatus struct {
	FollowStatus FollowStatus
	Valid        bool // Valid is true if FollowStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFollowStatus) Scan(value interface{}) error {
	if value == nil {
		ns.FollowStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.FollowStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFollowStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.FollowStatus), nil
}

type NotificationType string

const (
	NotificationTypeFollow          NotificationType = "follow"
	NotificationTypeLike            NotificationType = "like"
	NotificationTypeRetweet         NotificationType = "retweet"
	NotificationTypeReply           NotificationType = "reply"
	NotificationTypeMessage         NotificationType = "message"
	NotificationTypeQuote           NotificationType = "quote"
	NotificationTypeFollowRequest   NotificationType = "follow_request"
	NotificationTypeRequestAccepted NotificationType = "request_accepted"
	NotificationTypeReport          NotificationType = "report"
	NotificationTypeWarning         NotificationType = "warning"
	NotificationTypeOther           NotificationType = "other"
)

func (e *NotificationType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = NotificationType(s)
	case string:
		*e = NotificationType(s)
	default:
		return fmt.Errorf("unsupported scan type for NotificationType: %T", src)
	}
	return nil
}

type NullNotificationType struct {
	NotificationType NotificationType
	Valid            bool // Valid is true if NotificationType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullNotificationType) Scan(value interface{}) error {
	if value == nil {
		ns.NotificationType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.NotificationType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullNotificationType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.NotificationType), nil
}

type ReportReason string

const (
	ReportReasonSpam                 ReportReason = "spam"
	ReportReasonHarassment           ReportReason = "harassment"
	ReportReasonInappropriateContent ReportReason = "inappropriate_content"
	ReportReasonOther                ReportReason = "other"
)

func (e *ReportReason) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ReportReason(s)
	case string:
		*e = ReportReason(s)
	default:
		return fmt.Errorf("unsupported scan type for ReportReason: %T", src)
	}
	return nil
}

type NullReportReason struct {
	ReportReason ReportReason
	Valid        bool // Valid is true if ReportReason is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullReportReason) Scan(value interface{}) error {
	if value == nil {
		ns.ReportReason, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ReportReason.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullReportReason) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ReportReason), nil
}

type Account struct {
	ID          string
	UserID      string
	UserName    string
	IsSuspended bool
	IsAdmin     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Block struct {
	BlockerAccountID string
	BlockedAccountID string
	CreatedAt        time.Time
}

type Follow struct {
	FollowerAccountID  string
	FollowingAccountID string
	Status             FollowStatus
	CreatedAt          time.Time
}

type Hashtag struct {
	ID        int64
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
	TweetID   int64
	Label1    string
	Label2    sql.NullString
	Label3    sql.NullString
	CreatedAt time.Time
}

type Like struct {
	LikingAccountID string
	OriginalTweetID int64
	CreatedAt       time.Time
}

type Message struct {
	ID                 int64
	SenderAccountID    string
	RecipientAccountID string
	Content            sql.NullString
	IsRead             bool
	CreatedAt          time.Time
}

type Notification struct {
	ID                 int64
	SenderAccountID    sql.NullString
	RecipientAccountID string
	Type               NotificationType
	Content            sql.NullString
	IsRead             bool
	CreatedAt          time.Time
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
	TweetID           int64
	OriginalTweetID   int64
	ParentReplyID     sql.NullInt64
	ReplyingAccountID string
	CreatedAt         time.Time
}

type Report struct {
	ID                int64
	ReporterAccountID string
	ReportedAccountID string
	Reason            ReportReason
	Content           sql.NullString
	CreatedAt         time.Time
}

type RetweetsAndQuote struct {
	TweetID             int64
	RetweetingAccountID string
	OriginalTweetID     int64
	CreatedAt           time.Time
}

type Setting struct {
	AccountID string
	IsPrivate bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tweet struct {
	ID              int64
	AccountID       string
	IsPinned        bool
	Content         sql.NullString
	Code            sql.NullString
	LikesCount      int32
	RepliesCount    int32
	RetweetsCount   int32
	IsRetweet       bool
	IsReply         bool
	IsQuote         bool
	EngagementScore int32
	Media           pqtype.NullRawMessage
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type TweetHashtag struct {
	TweetID   int64
	HashtagID int64
}
