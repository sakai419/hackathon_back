package repository

import (
	"context"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"
)

func (r *Repository) GetInterestScores(ctx context.Context, accountID string) (*model.InterestScores, error) {
	// Get interest scores
	interests, err := r.q.GetInterestScores(ctx, accountID)
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get interest scores",
				Err: err,
			},
		)
	}

	// Convert to model
	ret := convertToInterestScores(interests)

	return ret, nil
}

func convertToInterestScores(interests sqlcgen.Interest) *model.InterestScores {
	ret := &model.InterestScores{
		NewsScore: interests.NewsScore,
		PoliticsScore: interests.PoliticsScore,
		EconomicsScore: interests.EconomicsScore,
		HealthScore: interests.HealthScore,
		SportsScore: interests.SportsScore,
		EntertainmentScore: interests.EntertainmentScore,
		ArtScore: interests.ArtScore,
		CookingScore: interests.CookingScore,
		TravelScore: interests.TravelScore,
		FashionScore: interests.FashionScore,
		BeautyScore: interests.BeautyScore,
		PetsScore: interests.PetsScore,
		ParentingScore: interests.ParentingScore,
		EducationScore: interests.EducationScore,
		EnvironmentScore: interests.EnvironmentScore,
		ClimateScore: interests.ClimateScore,
		SpaceScore: interests.SpaceScore,
		MentalHealthScore: interests.MentalHealthScore,
		FitnessScore: interests.FitnessScore,
		ReadingScore: interests.ReadingScore,
		HistoryScore: interests.HistoryScore,
		PhilosophyScore: interests.PhilosophyScore,
		ReligionScore: interests.ReligionScore,
		CultureScore: interests.CultureScore,
		VolunteeringScore: interests.VolunteeringScore,
		SocialIssuesScore: interests.SocialIssuesScore,
		LawScore: interests.LawScore,
		TaxesScore: interests.TaxesScore,
		InvestmentScore: interests.InvestmentScore,
		RealEstateScore: interests.RealEstateScore,
		DIYScore: interests.DiyScore,
		GardeningScore: interests.GardeningScore,
		InteriorDesignScore: interests.InteriorDesignScore,
		AutomotiveScore: interests.AutomotiveScore,
		GamingScore: interests.GamingScore,
		AnimeMangaScore: interests.AnimeMangaScore,
		CreativeWorksScore: interests.CreativeWorksScore,
		PhotographyVideoScore: interests.PhotographyVideoScore,
		MediaScore: interests.MediaScore,
		MarketingScore: interests.MarketingScore,
		BrandingScore: interests.BrandingScore,
		EntrepreneurshipScore: interests.EntrepreneurshipScore,
		RemoteWorkScore: interests.RemoteWorkScore,
		DataScienceScore: interests.DataScienceScore,
		IoTScore: interests.IotScore,
		RoboticsEngineeringScore: interests.RoboticsEngineeringScore,
		BiotechnologyScore: interests.BiotechnologyScore,
		NanotechnologyScore: interests.NanotechnologyScore,
		EnergyTechnologyScore: interests.EnergyTechnologyScore,
		ArchaeologyScore: interests.ArchaeologyScore,
		PsychologyScore: interests.PsychologyScore,
		SociologyScore: interests.SociologyScore,
		AnthropologyScore: interests.AnthropologyScore,
		GeographyScore: interests.GeographyScore,
		GeologyScore: interests.GeologyScore,
		MeteorologyScore: interests.MeteorologyScore,
		DisasterEmergencyManagementScore: interests.DisasterEmergencyManagementScore,
		UrbanPlanningScore: interests.UrbanPlanningScore,
		ArchitectureScore: interests.ArchitectureScore,
		AgricultureScore: interests.AgricultureScore,
		NutritionScienceScore: interests.NutritionScienceScore,
		SleepScienceScore: interests.SleepScienceScore,
		ProductivityScore: interests.ProductivityScore,
		LeadershipScore: interests.LeadershipScore,
		InternationalRelationsScore: interests.InternationalRelationsScore,
		FuturePredictionsScore: interests.FuturePredictionsScore,
		EventsScore: interests.EventsScore,
		CommunityScore: interests.CommunityScore,
		TrendsScore: interests.TrendsScore,
		LifestyleScore: interests.LifestyleScore,
		SoftwareDevelopmentScore: interests.SoftwareDevelopmentScore,
		ProgrammingLanguagesScore: interests.ProgrammingLanguagesScore,
		WebDevelopmentScore: interests.WebDevelopmentScore,
		MobileAppDevelopmentScore: interests.MobileAppDevelopmentScore,
		DebuggingTechniquesScore: interests.DebuggingTechniquesScore,
		AlgorithmsMathematicsScore: interests.AlgorithmsMathematicsScore,
		DatabaseDesignScore: interests.DatabaseDesignScore,
		CloudComputingScore: interests.CloudComputingScore,
		ServerManagementScore: interests.ServerManagementScore,
		NetworkSecurityScore: interests.NetworkSecurityScore,
		CryptographyScore: interests.CryptographyScore,
		ArtificialIntelligenceScore: interests.ArtificialIntelligenceScore,
		MachineLearningScore: interests.MachineLearningScore,
		DeepLearningScore: interests.DeepLearningScore,
		ComputerVisionScore: interests.ComputerVisionScore,
		NaturalLanguageProcessingScore: interests.NaturalLanguageProcessingScore,
		BlockchainTechnologyScore: interests.BlockchainTechnologyScore,
		QuantumComputingScore: interests.QuantumComputingScore,
		EdgeComputingScore: interests.EdgeComputingScore,
		MicroservicesArchitectureScore: interests.MicroservicesArchitectureScore,
		DevOpsScore: interests.DevopsScore,
		ContainerTechnologyScore: interests.ContainerTechnologyScore,
		CICDScore: interests.CiCdScore,
		TestAutomationScore: interests.TestAutomationScore,
		UXUIDesignScore: interests.UxUiDesignScore,
		AgileDevelopmentMethodologiesScore: interests.AgileDevelopmentMethodologiesScore,
		OpenSourceScore: interests.OpenSourceScore,
		VersionControlScore: interests.VersionControlScore,
		APIDesignScore: interests.ApiDesignScore,
		PerformanceOptimizationScore: interests.PerformanceOptimizationScore,
	}

	return ret
}