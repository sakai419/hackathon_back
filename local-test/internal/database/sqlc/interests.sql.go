// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: interests.sql

package database

import (
	"context"
	"database/sql"
)

const createInterestsWithDefaultScores = `-- name: CreateInterestsWithDefaultScores :exec
INSERT INTO interests (account_id)
VALUES (?)
`

func (q *Queries) CreateInterestsWithDefaultScores(ctx context.Context, accountID string) error {
	_, err := q.db.ExecContext(ctx, createInterestsWithDefaultScores, accountID)
	return err
}

const deleteInterests = `-- name: DeleteInterests :exec
DELETE FROM interests
WHERE account_id = ?
`

func (q *Queries) DeleteInterests(ctx context.Context, accountID string) error {
	_, err := q.db.ExecContext(ctx, deleteInterests, accountID)
	return err
}

const getInterestByAccountId = `-- name: GetInterestByAccountId :one
SELECT account_id, news_score, politics_score, economics_score, health_score, sports_score, entertainment_score, art_score, cooking_score, travel_score, fashion_score, beauty_score, pets_score, parenting_score, education_score, environment_score, climate_score, space_score, mental_health_score, fitness_score, reading_score, history_score, philosophy_score, religion_score, culture_score, volunteering_score, social_issues_score, law_score, taxes_score, investment_score, real_estate_score, diy_score, gardening_score, interior_design_score, automotive_score, gaming_score, anime_manga_score, creative_works_score, photography_video_score, media_score, marketing_score, branding_score, entrepreneurship_score, remote_work_score, data_science_score, iot_score, robotics_engineering_score, biotechnology_score, nanotechnology_score, energy_technology_score, archaeology_score, psychology_score, sociology_score, anthropology_score, geography_score, geology_score, meteorology_score, disaster_emergency_management_score, urban_planning_score, architecture_score, agriculture_score, nutrition_science_score, sleep_science_score, productivity_score, leadership_score, international_relations_score, future_predictions_score, events_score, community_score, trends_score, lifestyle_score, software_development_score, programming_languages_score, web_development_score, mobile_app_development_score, debugging_techniques_score, algorithms_mathematics_score, database_design_score, cloud_computing_score, server_management_score, network_security_score, cryptography_score, artificial_intelligence_score, machine_learning_score, deep_learning_score, computer_vision_score, natural_language_processing_score, blockchain_technology_score, quantum_computing_score, edge_computing_score, microservices_architecture_score, devops_score, container_technology_score, ci_cd_score, test_automation_score, ux_ui_design_score, agile_development_methodologies_score, open_source_score, version_control_score, api_design_score, performance_optimization_score FROM interests
WHERE account_id = ?
`

func (q *Queries) GetInterestByAccountId(ctx context.Context, accountID string) (Interest, error) {
	row := q.db.QueryRowContext(ctx, getInterestByAccountId, accountID)
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

const updateInterestScore = `-- name: UpdateInterestScore :exec
UPDATE interests
SET
    news_score = COALESCE(?, news_score),
    politics_score = COALESCE(?, politics_score),
    economics_score = COALESCE(?, economics_score),
    health_score = COALESCE(?, health_score),
    sports_score = COALESCE(?, sports_score),
    entertainment_score = COALESCE(?, entertainment_score),
    art_score = COALESCE(?, art_score),
    cooking_score = COALESCE(?, cooking_score),
    travel_score = COALESCALE(?, travel_score),
    fashion_score = COALESCE(?, fashion_score),
    beauty_score = COALESCE(?, beauty_score),
    pets_score = COALESCE(?, pets_score),
    parenting_score = COALESCE(?, parenting_score),
    education_score = COALESCE(?, education_score),
    environment_score = COALESCE(?, environment_score),
    climate_score = COALESCE(?, climate_score),
    space_score = COALESCE(?, space_score),
    mental_health_score = COALESCE(?, mental_health_score),
    fitness_score = COALESCE(?, fitness_score),
    reading_score = COALESCE(?, reading_score),
    history_score = COALESCE(?, history_score),
    philosophy_score = COALESCE(?, philosophy_score),
    religion_score = COALESCE(?, religion_score),
    culture_score = COALESCE(?, culture_score),
    volunteering_score = COALESCE(?, volunteer_score),
    social_issues_score = COALESCE(?, social_issues_score),
    law_score = COALESCE(?, law_score),
    taxes_score = COALESCE(?, taxes_score),
    investment_score = COALESCE(?, investment_score),
    real_estate_score = COALESCE(?, real_estate_score),
    diy_score = COALESCE(?, diy_score),
    gardening_score = COALESCE(?, gardening_score),
    interior_design_score = COALESCE(?, interior_design_score),
    automotive_score = COALESCE(?, automotive_score),
    gaming_score = COALESCE(?, gaming_score),
    anime_manga_score = COALESCE(?, anime_manga_score),
    creative_works_score = COALESCE(?, creative_works_score),
    photography_video_score = COALESCE(?, photography_video_score),
    media_score = COALESCE(?, media_score),
    marketing_score = COALESCE(?, marketing_score),
    branding_score = COALESCE(?, branding_score),
    entrepreneurship_score = COALESCE(?, entrepreneurship_score),
    remote_work_score = COALESCE(?, remote_work_score),
    data_science_score = COALESCE(?, data_science_score),
    iot_score = COALESCE(?, iot_score),
    robotics_engineering_score = COALESCE(?, robotics_engineering_score),
    biotechnology_score = COALESCE(?, biotechnology_score),
    nanotechnology_score = COALESCE(?, nanotechnology_score),
    energy_technology_score = COALESCE(?, energy_technology_score),
    archaeology_score = COALESCE(?, archaeology_score),
    psychology_score = COALESCE(?, psychology_score),
    sociology_score = COALESCE(?, sociology_score),
    anthropology_score = COALESCE(?, anthropology_score),
    geography_score = COALESCE(?, geography_score),
    geology_score = COALESCE(?, geology_score),
    meteorology_score = COALESCE(?, meteorology_score),
    disaster_emergency_management_score = COALESCE(?, disaster_emergency_management_score),
    urban_planning_score = COALESCE(?, urban_planning_score),
    architecture_score = COALESCE(?, architecture_score),
    agriculture_score = COALESCE(?, agriculture_score),
    nutrition_science_score = COALESCE(?, nutrition_science_score),
    sleep_science_score = COALESCE(?, sleep_science_score),
    productivity_score = COALESCE(?, productivity_score),
    leadership_score = COALESCE(?, leadership_score),
    international_relations_score = COALESCE(?, international_relations_score),
    future_predictions_score = COALESCE(?, future_predictions_score),
    events_score = COALESCE(?, events_score),
    community_score = COALESCE(?, community_score),
    trends_score = COALESCE(?, trends_score),
    lifestyle_score = COALESCE(?, lifestyle_score),
    software_development_score = COALESCE(?, software_development_score),
    programming_languages_score = COALESCE(?, programming_languages_score),
    web_development_score = COALESCE(?, web_development_score),
    mobile_app_development_score = COALESCE(?, mobile_app_development_score),
    debugging_techniques_score = COALESCE(?, debugging_techniques_score),
    algorithms_mathematics_score = COALESCE(?, algorithms_mathematics_score),
    database_design_score = COALESCE(?, database_design_score),
    cloud_computing_score = COALESCE(?, cloud_computing_score),
    server_management_score = COALESCE(?, server_management_score),
    network_security_score = COALESCE(?, network_security_score),
    cryptography_score = COALESCE(?, cryptography_score),
    artificial_intelligence_score = COALESCE(?, artificial_intelligence_score),
    machine_learning_score = COALESCE(?, machine_learning_score),
    deep_learning_score = COALESCE(?, deep_learning_score),
    computer_vision_score = COALESCE(?, computer_vision_score),
    natural_language_processing_score = COALESCE(?, natural_language_processing_score),
    blockchain_technology_score = COALESCE(?, blockchain_technology_score),
    quantum_computing_score = COALESCE(?, quantum_computing_score),
    edge_computing_score = COALESCE(?, edge_computing_score),
    microservices_architecture_score = COALESCE(?, microservices_architecture_score),
    devops_score = COALESCE(?, devops_score),
    container_technology_score = COALESCE(?, container_technology_score),
    ci_cd_score = COALESCE(?, ci_cd_score),
    test_automation_score = COALESCE(?, test_automation_score),
    ux_ui_design_score = COALESCE(?, ux_ui_design_score),
    agile_development_methodologies_score = COALESCE(?, agile_development_methodologies_score),
    open_source_score = COALESCE(?, open_source_score),
    version_control_score = COALESCE(?, version_control_score),
    api_design_score = COALESCE(?, api_design_score),
    performance_optimization_score = COALESCE(?, performance_optimization_score)
WHERE account_id = ?
`

type UpdateInterestScoreParams struct {
	NewsScore                          sql.NullInt16
	PoliticsScore                      sql.NullInt16
	EconomicsScore                     sql.NullInt16
	HealthScore                        sql.NullInt16
	SportsScore                        sql.NullInt16
	EntertainmentScore                 sql.NullInt16
	ArtScore                           sql.NullInt16
	CookingScore                       sql.NullInt16
	Coalescale                         interface{}
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
	AccountID                          string
}

func (q *Queries) UpdateInterestScore(ctx context.Context, arg UpdateInterestScoreParams) error {
	_, err := q.db.ExecContext(ctx, updateInterestScore,
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
