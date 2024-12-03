// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: interests.sql

package sqlcgen

import (
	"context"
)

const createInterestsWithDefaultValues = `-- name: CreateInterestsWithDefaultValues :exec
INSERT INTO interests (account_id)
VALUES ($1)
`

func (q *Queries) CreateInterestsWithDefaultValues(ctx context.Context, accountID string) error {
	_, err := q.db.ExecContext(ctx, createInterestsWithDefaultValues, accountID)
	return err
}

const deleteInterests = `-- name: DeleteInterests :exec
DELETE FROM interests
WHERE account_id = $1
`

func (q *Queries) DeleteInterests(ctx context.Context, accountID string) error {
	_, err := q.db.ExecContext(ctx, deleteInterests, accountID)
	return err
}

const getInterestScores = `-- name: GetInterestScores :one
SELECT account_id, news_score, politics_score, economics_score, health_score, sports_score, entertainment_score, art_score, cooking_score, travel_score, fashion_score, beauty_score, pets_score, parenting_score, education_score, environment_score, climate_score, space_score, mental_health_score, fitness_score, reading_score, history_score, philosophy_score, religion_score, culture_score, volunteering_score, social_issues_score, law_score, taxes_score, investment_score, real_estate_score, diy_score, gardening_score, interior_design_score, automotive_score, gaming_score, anime_manga_score, creative_works_score, photography_video_score, media_score, marketing_score, branding_score, entrepreneurship_score, remote_work_score, data_science_score, iot_score, robotics_engineering_score, biotechnology_score, nanotechnology_score, energy_technology_score, archaeology_score, psychology_score, sociology_score, anthropology_score, geography_score, geology_score, meteorology_score, disaster_emergency_management_score, urban_planning_score, architecture_score, agriculture_score, nutrition_science_score, sleep_science_score, productivity_score, leadership_score, international_relations_score, future_predictions_score, events_score, community_score, trends_score, lifestyle_score, software_development_score, programming_languages_score, web_development_score, mobile_app_development_score, debugging_techniques_score, algorithms_mathematics_score, database_design_score, cloud_computing_score, server_management_score, network_security_score, cryptography_score, artificial_intelligence_score, machine_learning_score, deep_learning_score, computer_vision_score, natural_language_processing_score, blockchain_technology_score, quantum_computing_score, edge_computing_score, microservices_architecture_score, devops_score, container_technology_score, ci_cd_score, test_automation_score, ux_ui_design_score, agile_development_methodologies_score, open_source_score, version_control_score, api_design_score, performance_optimization_score FROM interests
WHERE account_id = $1
`

func (q *Queries) GetInterestScores(ctx context.Context, accountID string) (Interest, error) {
	row := q.db.QueryRowContext(ctx, getInterestScores, accountID)
	var i Interest
	err := row.Scan(
		&i.AccountID,
		&i.NewsScore,
		&i.PoliticsScore,
		&i.EconomicsScore,
		&i.HealthScore,
		&i.SportsScore,
		&i.EntertainmentScore,
		&i.ArtScore,
		&i.CookingScore,
		&i.TravelScore,
		&i.FashionScore,
		&i.BeautyScore,
		&i.PetsScore,
		&i.ParentingScore,
		&i.EducationScore,
		&i.EnvironmentScore,
		&i.ClimateScore,
		&i.SpaceScore,
		&i.MentalHealthScore,
		&i.FitnessScore,
		&i.ReadingScore,
		&i.HistoryScore,
		&i.PhilosophyScore,
		&i.ReligionScore,
		&i.CultureScore,
		&i.VolunteeringScore,
		&i.SocialIssuesScore,
		&i.LawScore,
		&i.TaxesScore,
		&i.InvestmentScore,
		&i.RealEstateScore,
		&i.DiyScore,
		&i.GardeningScore,
		&i.InteriorDesignScore,
		&i.AutomotiveScore,
		&i.GamingScore,
		&i.AnimeMangaScore,
		&i.CreativeWorksScore,
		&i.PhotographyVideoScore,
		&i.MediaScore,
		&i.MarketingScore,
		&i.BrandingScore,
		&i.EntrepreneurshipScore,
		&i.RemoteWorkScore,
		&i.DataScienceScore,
		&i.IotScore,
		&i.RoboticsEngineeringScore,
		&i.BiotechnologyScore,
		&i.NanotechnologyScore,
		&i.EnergyTechnologyScore,
		&i.ArchaeologyScore,
		&i.PsychologyScore,
		&i.SociologyScore,
		&i.AnthropologyScore,
		&i.GeographyScore,
		&i.GeologyScore,
		&i.MeteorologyScore,
		&i.DisasterEmergencyManagementScore,
		&i.UrbanPlanningScore,
		&i.ArchitectureScore,
		&i.AgricultureScore,
		&i.NutritionScienceScore,
		&i.SleepScienceScore,
		&i.ProductivityScore,
		&i.LeadershipScore,
		&i.InternationalRelationsScore,
		&i.FuturePredictionsScore,
		&i.EventsScore,
		&i.CommunityScore,
		&i.TrendsScore,
		&i.LifestyleScore,
		&i.SoftwareDevelopmentScore,
		&i.ProgrammingLanguagesScore,
		&i.WebDevelopmentScore,
		&i.MobileAppDevelopmentScore,
		&i.DebuggingTechniquesScore,
		&i.AlgorithmsMathematicsScore,
		&i.DatabaseDesignScore,
		&i.CloudComputingScore,
		&i.ServerManagementScore,
		&i.NetworkSecurityScore,
		&i.CryptographyScore,
		&i.ArtificialIntelligenceScore,
		&i.MachineLearningScore,
		&i.DeepLearningScore,
		&i.ComputerVisionScore,
		&i.NaturalLanguageProcessingScore,
		&i.BlockchainTechnologyScore,
		&i.QuantumComputingScore,
		&i.EdgeComputingScore,
		&i.MicroservicesArchitectureScore,
		&i.DevopsScore,
		&i.ContainerTechnologyScore,
		&i.CiCdScore,
		&i.TestAutomationScore,
		&i.UxUiDesignScore,
		&i.AgileDevelopmentMethodologiesScore,
		&i.OpenSourceScore,
		&i.VersionControlScore,
		&i.ApiDesignScore,
		&i.PerformanceOptimizationScore,
	)
	return i, err
}

const updateInterestsScore = `-- name: UpdateInterestsScore :exec
UPDATE interests
SET
    news_score = COALESCE($1, news_score),
    politics_score = COALESCE($2, politics_score),
    economics_score = COALESCE($3, economics_score),
    health_score = COALESCE($4, health_score),
    sports_score = COALESCE($5, sports_score),
    entertainment_score = COALESCE($6, entertainment_score),
    art_score = COALESCE($7, art_score),
    cooking_score = COALESCE($8, cooking_score),
    travel_score = COALESCALE($9, travel_score),
    fashion_score = COALESCE($10, fashion_score),
    beauty_score = COALESCE($11, beauty_score),
    pets_score = COALESCE($12, pets_score),
    parenting_score = COALESCE($13, parenting_score),
    education_score = COALESCE($14, education_score),
    environment_score = COALESCE($15, environment_score),
    climate_score = COALESCE($16, climate_score),
    space_score = COALESCE($17, space_score),
    mental_health_score = COALESCE($18, mental_health_score),
    fitness_score = COALESCE($19, fitness_score),
    reading_score = COALESCE($20, reading_score),
    history_score = COALESCE($21, history_score),
    philosophy_score = COALESCE($22, philosophy_score),
    religion_score = COALESCE($23, religion_score),
    culture_score = COALESCE($24, culture_score),
    volunteering_score = COALESCE($25, volunteer_score),
    social_issues_score = COALESCE($26, social_issues_score),
    law_score = COALESCE($27, law_score),
    taxes_score = COALESCE($28, taxes_score),
    investment_score = COALESCE($29, investment_score),
    real_estate_score = COALESCE($30, real_estate_score),
    diy_score = COALESCE($31, diy_score),
    gardening_score = COALESCE($32, gardening_score),
    interior_design_score = COALESCE($33, interior_design_score),
    automotive_score = COALESCE($34, automotive_score),
    gaming_score = COALESCE($35, gaming_score),
    anime_manga_score = COALESCE($36, anime_manga_score),
    creative_works_score = COALESCE($37, creative_works_score),
    photography_video_score = COALESCE($38, photography_video_score),
    media_score = COALESCE($39, media_score),
    marketing_score = COALESCE($40, marketing_score),
    branding_score = COALESCE($41, branding_score),
    entrepreneurship_score = COALESCE($42, entrepreneurship_score),
    remote_work_score = COALESCE($43, remote_work_score),
    data_science_score = COALESCE($44, data_science_score),
    iot_score = COALESCE($45, iot_score),
    robotics_engineering_score = COALESCE($46, robotics_engineering_score),
    biotechnology_score = COALESCE($47, biotechnology_score),
    nanotechnology_score = COALESCE($48, nanotechnology_score),
    energy_technology_score = COALESCE($49, energy_technology_score),
    archaeology_score = COALESCE($50, archaeology_score),
    psychology_score = COALESCE($51, psychology_score),
    sociology_score = COALESCE($52, sociology_score),
    anthropology_score = COALESCE($53, anthropology_score),
    geography_score = COALESCE($54, geography_score),
    geology_score = COALESCE($55, geology_score),
    meteorology_score = COALESCE($56, meteorology_score),
    disaster_emergency_management_score = COALESCE($57, disaster_emergency_management_score),
    urban_planning_score = COALESCE($58, urban_planning_score),
    architecture_score = COALESCE($59, architecture_score),
    agriculture_score = COALESCE($60, agriculture_score),
    nutrition_science_score = COALESCE($61, nutrition_science_score),
    sleep_science_score = COALESCE($62, sleep_science_score),
    productivity_score = COALESCE($63, productivity_score),
    leadership_score = COALESCE($64, leadership_score),
    international_relations_score = COALESCE($65, international_relations_score),
    future_predictions_score = COALESCE($66, future_predictions_score),
    events_score = COALESCE($67, events_score),
    community_score = COALESCE($68, community_score),
    trends_score = COALESCE($69, trends_score),
    lifestyle_score = COALESCE($70, lifestyle_score),
    software_development_score = COALESCE($71, software_development_score),
    programming_languages_score = COALESCE($72, programming_languages_score),
    web_development_score = COALESCE($73, web_development_score),
    mobile_app_development_score = COALESCE($74, mobile_app_development_score),
    debugging_techniques_score = COALESCE($75, debugging_techniques_score),
    algorithms_mathematics_score = COALESCE($76, algorithms_mathematics_score),
    database_design_score = COALESCE($77, database_design_score),
    cloud_computing_score = COALESCE($78, cloud_computing_score),
    server_management_score = COALESCE($79, server_management_score),
    network_security_score = COALESCE($80, network_security_score),
    cryptography_score = COALESCE($81, cryptography_score),
    artificial_intelligence_score = COALESCE($82, artificial_intelligence_score),
    machine_learning_score = COALESCE($83, machine_learning_score),
    deep_learning_score = COALESCE($84, deep_learning_score),
    computer_vision_score = COALESCE($85, computer_vision_score),
    natural_language_processing_score = COALESCE($86, natural_language_processing_score),
    blockchain_technology_score = COALESCE($87, blockchain_technology_score),
    quantum_computing_score = COALESCE($88, quantum_computing_score),
    edge_computing_score = COALESCE($89, edge_computing_score),
    microservices_architecture_score = COALESCE($90, microservices_architecture_score),
    devops_score = COALESCE($91, devops_score),
    container_technology_score = COALESCE($92, container_technology_score),
    ci_cd_score = COALESCE($93, ci_cd_score),
    test_automation_score = COALESCE($94, test_automation_score),
    ux_ui_design_score = COALESCE($95, ux_ui_design_score),
    agile_development_methodologies_score = COALESCE($96, agile_development_methodologies_score),
    open_source_score = COALESCE($97, open_source_score),
    version_control_score = COALESCE($98, version_control_score),
    api_design_score = COALESCE($99, api_design_score),
    performance_optimization_score = COALESCE($100, performance_optimization_score)
WHERE account_id = $101
`

type UpdateInterestsScoreParams struct {
	NewsScore                          int16
	PoliticsScore                      int16
	EconomicsScore                     int16
	HealthScore                        int16
	SportsScore                        int16
	EntertainmentScore                 int16
	ArtScore                           int16
	CookingScore                       int16
	Coalescale                         interface{}
	FashionScore                       int16
	BeautyScore                        int16
	PetsScore                          int16
	ParentingScore                     int16
	EducationScore                     int16
	EnvironmentScore                   int16
	ClimateScore                       int16
	SpaceScore                         int16
	MentalHealthScore                  int16
	FitnessScore                       int16
	ReadingScore                       int16
	HistoryScore                       int16
	PhilosophyScore                    int16
	ReligionScore                      int16
	CultureScore                       int16
	VolunteeringScore                  int16
	SocialIssuesScore                  int16
	LawScore                           int16
	TaxesScore                         int16
	InvestmentScore                    int16
	RealEstateScore                    int16
	DiyScore                           int16
	GardeningScore                     int16
	InteriorDesignScore                int16
	AutomotiveScore                    int16
	GamingScore                        int16
	AnimeMangaScore                    int16
	CreativeWorksScore                 int16
	PhotographyVideoScore              int16
	MediaScore                         int16
	MarketingScore                     int16
	BrandingScore                      int16
	EntrepreneurshipScore              int16
	RemoteWorkScore                    int16
	DataScienceScore                   int16
	IotScore                           int16
	RoboticsEngineeringScore           int16
	BiotechnologyScore                 int16
	NanotechnologyScore                int16
	EnergyTechnologyScore              int16
	ArchaeologyScore                   int16
	PsychologyScore                    int16
	SociologyScore                     int16
	AnthropologyScore                  int16
	GeographyScore                     int16
	GeologyScore                       int16
	MeteorologyScore                   int16
	DisasterEmergencyManagementScore   int16
	UrbanPlanningScore                 int16
	ArchitectureScore                  int16
	AgricultureScore                   int16
	NutritionScienceScore              int16
	SleepScienceScore                  int16
	ProductivityScore                  int16
	LeadershipScore                    int16
	InternationalRelationsScore        int16
	FuturePredictionsScore             int16
	EventsScore                        int16
	CommunityScore                     int16
	TrendsScore                        int16
	LifestyleScore                     int16
	SoftwareDevelopmentScore           int16
	ProgrammingLanguagesScore          int16
	WebDevelopmentScore                int16
	MobileAppDevelopmentScore          int16
	DebuggingTechniquesScore           int16
	AlgorithmsMathematicsScore         int16
	DatabaseDesignScore                int16
	CloudComputingScore                int16
	ServerManagementScore              int16
	NetworkSecurityScore               int16
	CryptographyScore                  int16
	ArtificialIntelligenceScore        int16
	MachineLearningScore               int16
	DeepLearningScore                  int16
	ComputerVisionScore                int16
	NaturalLanguageProcessingScore     int16
	BlockchainTechnologyScore          int16
	QuantumComputingScore              int16
	EdgeComputingScore                 int16
	MicroservicesArchitectureScore     int16
	DevopsScore                        int16
	ContainerTechnologyScore           int16
	CiCdScore                          int16
	TestAutomationScore                int16
	UxUiDesignScore                    int16
	AgileDevelopmentMethodologiesScore int16
	OpenSourceScore                    int16
	VersionControlScore                int16
	ApiDesignScore                     int16
	PerformanceOptimizationScore       int16
	AccountID                          string
}

func (q *Queries) UpdateInterestsScore(ctx context.Context, arg UpdateInterestsScoreParams) error {
	_, err := q.db.ExecContext(ctx, updateInterestsScore,
		arg.NewsScore,
		arg.PoliticsScore,
		arg.EconomicsScore,
		arg.HealthScore,
		arg.SportsScore,
		arg.EntertainmentScore,
		arg.ArtScore,
		arg.CookingScore,
		arg.Coalescale,
		arg.FashionScore,
		arg.BeautyScore,
		arg.PetsScore,
		arg.ParentingScore,
		arg.EducationScore,
		arg.EnvironmentScore,
		arg.ClimateScore,
		arg.SpaceScore,
		arg.MentalHealthScore,
		arg.FitnessScore,
		arg.ReadingScore,
		arg.HistoryScore,
		arg.PhilosophyScore,
		arg.ReligionScore,
		arg.CultureScore,
		arg.VolunteeringScore,
		arg.SocialIssuesScore,
		arg.LawScore,
		arg.TaxesScore,
		arg.InvestmentScore,
		arg.RealEstateScore,
		arg.DiyScore,
		arg.GardeningScore,
		arg.InteriorDesignScore,
		arg.AutomotiveScore,
		arg.GamingScore,
		arg.AnimeMangaScore,
		arg.CreativeWorksScore,
		arg.PhotographyVideoScore,
		arg.MediaScore,
		arg.MarketingScore,
		arg.BrandingScore,
		arg.EntrepreneurshipScore,
		arg.RemoteWorkScore,
		arg.DataScienceScore,
		arg.IotScore,
		arg.RoboticsEngineeringScore,
		arg.BiotechnologyScore,
		arg.NanotechnologyScore,
		arg.EnergyTechnologyScore,
		arg.ArchaeologyScore,
		arg.PsychologyScore,
		arg.SociologyScore,
		arg.AnthropologyScore,
		arg.GeographyScore,
		arg.GeologyScore,
		arg.MeteorologyScore,
		arg.DisasterEmergencyManagementScore,
		arg.UrbanPlanningScore,
		arg.ArchitectureScore,
		arg.AgricultureScore,
		arg.NutritionScienceScore,
		arg.SleepScienceScore,
		arg.ProductivityScore,
		arg.LeadershipScore,
		arg.InternationalRelationsScore,
		arg.FuturePredictionsScore,
		arg.EventsScore,
		arg.CommunityScore,
		arg.TrendsScore,
		arg.LifestyleScore,
		arg.SoftwareDevelopmentScore,
		arg.ProgrammingLanguagesScore,
		arg.WebDevelopmentScore,
		arg.MobileAppDevelopmentScore,
		arg.DebuggingTechniquesScore,
		arg.AlgorithmsMathematicsScore,
		arg.DatabaseDesignScore,
		arg.CloudComputingScore,
		arg.ServerManagementScore,
		arg.NetworkSecurityScore,
		arg.CryptographyScore,
		arg.ArtificialIntelligenceScore,
		arg.MachineLearningScore,
		arg.DeepLearningScore,
		arg.ComputerVisionScore,
		arg.NaturalLanguageProcessingScore,
		arg.BlockchainTechnologyScore,
		arg.QuantumComputingScore,
		arg.EdgeComputingScore,
		arg.MicroservicesArchitectureScore,
		arg.DevopsScore,
		arg.ContainerTechnologyScore,
		arg.CiCdScore,
		arg.TestAutomationScore,
		arg.UxUiDesignScore,
		arg.AgileDevelopmentMethodologiesScore,
		arg.OpenSourceScore,
		arg.VersionControlScore,
		arg.ApiDesignScore,
		arg.PerformanceOptimizationScore,
		arg.AccountID,
	)
	return err
}
