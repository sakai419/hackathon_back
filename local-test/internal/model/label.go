package model

import (
	"local-test/pkg/apperrors"
)

type Label string

const (
	LabelNews	Label = "news"
	LabelPolitics Label = "politics"
	LabelEconomics Label = "economics"
	LabelHealth Label = "health"
	LabelSports Label = "sports"
	LabelEntertainment Label = "entertainment"
	LabelArt Label = "art"
	LabelCooking Label = "cooking"
	LabelTravel Label = "travel"
	LabelFashion Label = "fashion"
	LabelBeauty Label = "beauty"
	LabelPets Label = "pets"
	LabelParenting Label = "parenting"
	LabelEducation Label = "education"
	LabelEnvironment Label = "environment"
	LabelClimate Label = "climate"
	LabelSpace Label = "space"
	LabelMentalHealth Label = "mental_health"
	LabelFitness Label = "fitness"
	LabelReading Label = "reading"
	LabelHistory Label = "history"
	LabelPhilosophy Label = "philosophy"
	LabelReligion Label = "religion"
	LabelCulture Label = "culture"
	LabelVolunteering Label = "volunteering"
	LabelSocialIssues Label = "social_issues"
	LabelLaw Label = "law"
	LabelTaxes Label = "taxes"
	LabelInvestment Label = "investment"
	LabelRealEstate Label = "real_estate"
	LabelDIY Label = "diy"
	LabelGardening Label = "gardening"
	LabelInteriorDesign Label = "interior_design"
	LabelAutomotive Label = "automotive"
	LabelGaming Label = "gaming"
	LabelAnimeManga Label = "anime_manga"
	LabelCreativeWorks Label = "creative_works"
	LabelPhotographyVideo Label = "photography_video"
	LabelMedia Label = "media"
	LabelMarketing Label = "marketing"
	LabelBranding Label = "branding"
	LabelEntrepreneurship Label = "entrepreneurship"
	LabelRemoteWork Label = "remote_work"
	LabelDataScience Label = "data_science"
	LabelIoT Label = "iot"
	LabelRoboticsEngineering Label = "robotics_engineering"
	LabelBiotechnology Label = "biotechnology"
	LabelNanotechnology Label = "nanotechnology"
	LabelEnergyTechnology Label = "energy_technology"
	LabelArchaeology Label = "archaeology"
	LabelPsychology Label = "psychology"
	LabelSociology Label = "sociology"
	LabelAnthropology Label = "anthropology"
	LabelGeography Label = "geography"
	LabelGeology Label = "geology"
	LabelMeteorology Label = "meteorology"
	LabelDisasterEmergencyManagement Label = "disaster_emergency_management"
	LabelUrbanPlanning Label = "urban_planning"
	LabelArchitecture Label = "architecture"
	LabelAgriculture Label = "agriculture"
	LabelNutritionScience Label = "nutrition_science"
	LabelSleepScience Label = "sleep_science"
	LabelProductivity Label = "productivity"
	LabelLeadership Label = "leadership"
	LabelInternationalRelations Label = "international_relations"
	LabelFuturePredictions Label = "future_predictions"
	LabelEvents Label = "events"
	LabelCommunity Label = "community"
	LabelTrends Label = "trends"
	LabelLifestyle Label = "lifestyle"
	LabelSoftwareDevelopment Label = "software_development"
	LabelProgrammingLanguages Label = "programming_languages"
	LabelWebDevelopment Label = "web_development"
	LabelMobileAppDevelopment Label = "mobile_app_development"
	LabelDebuggingTechniques Label = "debugging_techniques"
	LabelAlgorithmsMathematics Label = "algorithms_mathematics"
	LabelDatabaseDesign Label = "database_design"
	LabelCloudComputing Label = "cloud_computing"
	LabelServerManagement Label = "server_management"
	LabelNetworkSecurity Label = "network_security"
	LabelCryptography Label = "cryptography"
	LabelArtificialIntelligence Label = "artificial_intelligence"
	LabelMachineLearning Label = "machine_learning"
	LabelDeepLearning Label = "deep_learning"
	LabelComputerVision Label = "computer_vision"
	LabelNaturalLanguageProcessing Label = "natural_language_processing"
	LabelBlockchainTechnology Label = "blockchain_technology"
	LabelQuantumComputing Label = "quantum_computing"
	LabelEdgeComputing Label = "edge_computing"
	LabelMicroservicesArchitecture Label = "microservices_architecture"
	LabelDevOps Label = "devops"
	LabelContainerTechnology Label = "container_technology"
	LabelCICD Label = "ci_cd"
	LabelTestAutomation Label = "test_automation"
	LabelUXUIDesign Label = "ux_ui_design"
	LabelAgileDevelopmentMethodologies Label = "agile_development_methodologies"
	LabelOpenSource Label = "open_source"
	LabelVersionControl Label = "version_control"
	LabelAPIDesign Label = "api_design"
	LabelPerformanceOptimization Label = "performance_optimization"
)

func (l Label) Validate() error {
	switch l {
	case LabelNews,
		LabelPolitics,
		LabelEconomics,
		LabelHealth,
		LabelSports,
		LabelEntertainment,
		LabelArt,
		LabelCooking,
		LabelTravel,
		LabelFashion,
		LabelBeauty,
		LabelPets,
		LabelParenting,
		LabelEducation,
		LabelEnvironment,
		LabelClimate,
		LabelSpace,
		LabelMentalHealth,
		LabelFitness,
		LabelReading,
		LabelHistory,
		LabelPhilosophy,
		LabelReligion,
		LabelCulture,
		LabelVolunteering,
		LabelSocialIssues,
		LabelLaw,
		LabelTaxes,
		LabelInvestment,
		LabelRealEstate,
		LabelDIY,
		LabelGardening,
		LabelInteriorDesign,
		LabelAutomotive,
		LabelGaming,
		LabelAnimeManga,
		LabelCreativeWorks,
		LabelPhotographyVideo,
		LabelMedia,
		LabelMarketing,
		LabelBranding,
		LabelEntrepreneurship,
		LabelRemoteWork,
		LabelDataScience,
		LabelIoT,
		LabelRoboticsEngineering,
		LabelBiotechnology,
		LabelNanotechnology,
		LabelEnergyTechnology,
		LabelArchaeology,
		LabelPsychology,
		LabelSociology,
		LabelAnthropology,
		LabelGeography,
		LabelGeology,
		LabelMeteorology,
		LabelDisasterEmergencyManagement,
		LabelUrbanPlanning,
		LabelArchitecture,
		LabelAgriculture,
		LabelNutritionScience,
		LabelSleepScience,
		LabelProductivity,
		LabelLeadership,
		LabelInternationalRelations,
		LabelFuturePredictions,
		LabelEvents,
		LabelCommunity,
		LabelTrends,
		LabelLifestyle,
		LabelSoftwareDevelopment,
		LabelProgrammingLanguages,
		LabelWebDevelopment,
		LabelMobileAppDevelopment,
		LabelDebuggingTechniques,
		LabelAlgorithmsMathematics,
		LabelDatabaseDesign,
		LabelCloudComputing,
		LabelServerManagement,
		LabelNetworkSecurity,
		LabelCryptography,
		LabelArtificialIntelligence,
		LabelMachineLearning,
		LabelDeepLearning,
		LabelComputerVision,
		LabelNaturalLanguageProcessing,
		LabelBlockchainTechnology,
		LabelQuantumComputing,
		LabelEdgeComputing,
		LabelMicroservicesArchitecture,
		LabelDevOps,
		LabelContainerTechnology,
		LabelCICD,
		LabelTestAutomation,
		LabelUXUIDesign,
		LabelAgileDevelopmentMethodologies,
		LabelOpenSource,
		LabelVersionControl,
		LabelAPIDesign,
		LabelPerformanceOptimization:
		return nil
	default:
		return &apperrors.ErrInvalidInput{
			Message: "label is invalid",
		}
	}
}

func GetLabels() []string{
	// Get all labels
	labels := []Label{
		LabelNews,
		LabelPolitics,
		LabelEconomics,
		LabelHealth,
		LabelSports,
		LabelEntertainment,
		LabelArt,
		LabelCooking,
		LabelTravel,
		LabelFashion,
		LabelBeauty,
		LabelPets,
		LabelParenting,
		LabelEducation,
		LabelEnvironment,
		LabelClimate,
		LabelSpace,
		LabelMentalHealth,
		LabelFitness,
		LabelReading,
		LabelHistory,
		LabelPhilosophy,
		LabelReligion,
		LabelCulture,
		LabelVolunteering,
		LabelSocialIssues,
		LabelLaw,
		LabelTaxes,
		LabelInvestment,
		LabelRealEstate,
		LabelDIY,
		LabelGardening,
		LabelInteriorDesign,
		LabelAutomotive,
		LabelGaming,
		LabelAnimeManga,
		LabelCreativeWorks,
		LabelPhotographyVideo,
		LabelMedia,
		LabelMarketing,
		LabelBranding,
		LabelEntrepreneurship,
		LabelRemoteWork,
		LabelDataScience,
		LabelIoT,
		LabelRoboticsEngineering,
		LabelBiotechnology,
		LabelNanotechnology,
		LabelEnergyTechnology,
		LabelArchaeology,
		LabelPsychology,
		LabelSociology,
		LabelAnthropology,
		LabelGeography,
		LabelGeology,
		LabelMeteorology,
		LabelDisasterEmergencyManagement,
		LabelUrbanPlanning,
		LabelArchitecture,
		LabelAgriculture,
		LabelNutritionScience,
		LabelSleepScience,
		LabelProductivity,
		LabelLeadership,
		LabelInternationalRelations,
		LabelFuturePredictions,
		LabelEvents,
		LabelCommunity,
		LabelTrends,
		LabelLifestyle,
		LabelSoftwareDevelopment,
		LabelProgrammingLanguages,
		LabelWebDevelopment,
		LabelMobileAppDevelopment,
		LabelDebuggingTechniques,
		LabelAlgorithmsMathematics,
		LabelDatabaseDesign,
		LabelCloudComputing,
		LabelServerManagement,
		LabelNetworkSecurity,
		LabelCryptography,
		LabelArtificialIntelligence,
		LabelMachineLearning,
		LabelDeepLearning,
		LabelComputerVision,
		LabelNaturalLanguageProcessing,
		LabelBlockchainTechnology,
		LabelQuantumComputing,
		LabelEdgeComputing,
		LabelMicroservicesArchitecture,
		LabelDevOps,
		LabelContainerTechnology,
		LabelCICD,
		LabelTestAutomation,
		LabelUXUIDesign,
		LabelAgileDevelopmentMethodologies,
		LabelOpenSource,
		LabelVersionControl,
		LabelAPIDesign,
		LabelPerformanceOptimization,
	}

	ret := make([]string, 0)
	for _, label := range labels {
		ret = append(ret, string(label))
	}
	return ret
}

type LabelTweetParams struct {
	TweetID int64
	Label1  *Label
	Label2  *Label
	Label3  *Label
}