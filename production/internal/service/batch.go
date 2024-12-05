package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
	"log"
)

func (s *Service) UpdateUserInterests(ctx context.Context) error {
	// Set flag
	terminated := false
	page := 0

	for {
		// Get account IDs
		page++
		log.Println("page:", page)
		accountIDs, err := s.repo.GetAccountIDs(ctx, &model.GetAccountIDsParams{
			Limit: 10,
			Offset: int32((page - 1) * 10),
		})
		if err != nil {
			return apperrors.NewInternalAppError("failed to get account IDs", err)
		}

		// Check if there are no more account IDs
		if len(accountIDs) == 0 {
			terminated = true
		}

		// Update user interests
		for _, accountID := range accountIDs {
			// Get liked tweet labels count
			labelCounts, err := s.repo.GetLikedTweetLabelsCount(ctx, accountID)
			if err != nil {
				return apperrors.NewInternalAppError("failed to get liked tweet labels count", err)
			}

			// Update user interests
			interests := convertToInterestScores(labelCounts)
			if err := s.repo.UpdateInterestScores(ctx, accountID, interests); err != nil {
				return apperrors.NewInternalAppError("failed to update user interests", err)
			}
		}

		// Check if terminated
		if terminated {
			break
		}
	}

	return nil
}

func convertToInterestScores(labelCounts []*model.LabelCount) *model.InterestScores {
	interests := &model.InterestScores{}
	for _, labelCount := range labelCounts {
		switch labelCount.Label {
		case model.LabelNews:
			interests.NewsScore += int16(labelCount.Count)
		case model.LabelPolitics:
			interests.PoliticsScore += int16(labelCount.Count)
		case model.LabelEconomics:
			interests.EconomicsScore += int16(labelCount.Count)
		case model.LabelHealth:
			interests.HealthScore += int16(labelCount.Count)
		case model.LabelSports:
			interests.SportsScore += int16(labelCount.Count)
		case model.LabelEntertainment:
			interests.EntertainmentScore += int16(labelCount.Count)
		case model.LabelArt:
			interests.ArtScore += int16(labelCount.Count)
		case model.LabelCooking:
			interests.CookingScore += int16(labelCount.Count)
		case model.LabelTravel:
			interests.TravelScore += int16(labelCount.Count)
		case model.LabelFashion:
			interests.FashionScore += int16(labelCount.Count)
		case model.LabelBeauty:
			interests.BeautyScore += int16(labelCount.Count)
		case model.LabelPets:
			interests.PetsScore += int16(labelCount.Count)
		case model.LabelParenting:
			interests.ParentingScore += int16(labelCount.Count)
		case model.LabelEducation:
			interests.EducationScore += int16(labelCount.Count)
		case model.LabelEnvironment:
			interests.EnvironmentScore += int16(labelCount.Count)
		case model.LabelClimate:
			interests.ClimateScore += int16(labelCount.Count)
		case model.LabelSpace:
			interests.SpaceScore += int16(labelCount.Count)
		case model.LabelMentalHealth:
			interests.MentalHealthScore += int16(labelCount.Count)
		case model.LabelFitness:
			interests.FitnessScore += int16(labelCount.Count)
		case model.LabelReading:
			interests.ReadingScore += int16(labelCount.Count)
		case model.LabelHistory:
			interests.HistoryScore += int16(labelCount.Count)
		case model.LabelPhilosophy:
			interests.PhilosophyScore += int16(labelCount.Count)
		case model.LabelReligion:
			interests.ReligionScore += int16(labelCount.Count)
		case model.LabelCulture:
			interests.CultureScore += int16(labelCount.Count)
		case model.LabelVolunteering:
			interests.VolunteeringScore += int16(labelCount.Count)
		case model.LabelSocialIssues:
			interests.SocialIssuesScore += int16(labelCount.Count)
		case model.LabelLaw:
			interests.LawScore += int16(labelCount.Count)
		case model.LabelTaxes:
			interests.TaxesScore += int16(labelCount.Count)
		case model.LabelInvestment:
			interests.InvestmentScore += int16(labelCount.Count)
		case model.LabelRealEstate:
			interests.RealEstateScore += int16(labelCount.Count)
		case model.LabelDIY:
			interests.DIYScore += int16(labelCount.Count)
		case model.LabelGardening:
			interests.GardeningScore += int16(labelCount.Count)
		case model.LabelInteriorDesign:
			interests.InteriorDesignScore += int16(labelCount.Count)
		case model.LabelAutomotive:
			interests.AutomotiveScore += int16(labelCount.Count)
		case model.LabelGaming:
			interests.GamingScore += int16(labelCount.Count)
		case model.LabelAnimeManga:
			interests.AnimeMangaScore += int16(labelCount.Count)
		case model.LabelCreativeWorks:
			interests.CreativeWorksScore += int16(labelCount.Count)
		case model.LabelPhotographyVideo:
			interests.PhotographyVideoScore += int16(labelCount.Count)
		case model.LabelMedia:
			interests.MediaScore += int16(labelCount.Count)
		case model.LabelMarketing:
			interests.MarketingScore += int16(labelCount.Count)
		case model.LabelBranding:
			interests.BrandingScore += int16(labelCount.Count)
		case model.LabelEntrepreneurship:
			interests.EntrepreneurshipScore += int16(labelCount.Count)
		case model.LabelRemoteWork:
			interests.RemoteWorkScore += int16(labelCount.Count)
		case model.LabelDataScience:
			interests.DataScienceScore += int16(labelCount.Count)
		case model.LabelIoT:
			interests.IoTScore += int16(labelCount.Count)
		case model.LabelRoboticsEngineering:
			interests.RoboticsEngineeringScore += int16(labelCount.Count)
		case model.LabelBiotechnology:
			interests.BiotechnologyScore += int16(labelCount.Count)
		case model.LabelNanotechnology:
			interests.NanotechnologyScore += int16(labelCount.Count)
		case model.LabelEnergyTechnology:
			interests.EnergyTechnologyScore += int16(labelCount.Count)
		case model.LabelArchaeology:
			interests.ArchaeologyScore += int16(labelCount.Count)
		case model.LabelPsychology:
			interests.PsychologyScore += int16(labelCount.Count)
		case model.LabelSociology:
			interests.SociologyScore += int16(labelCount.Count)
		case model.LabelAnthropology:
			interests.AnthropologyScore += int16(labelCount.Count)
		case model.LabelGeography:
			interests.GeographyScore += int16(labelCount.Count)
		case model.LabelGeology:
			interests.GeologyScore += int16(labelCount.Count)
		case model.LabelMeteorology:
			interests.MeteorologyScore += int16(labelCount.Count)
		case model.LabelDisasterEmergencyManagement:
			interests.DisasterEmergencyManagementScore += int16(labelCount.Count)
		case model.LabelUrbanPlanning:
			interests.UrbanPlanningScore += int16(labelCount.Count)
		case model.LabelArchitecture:
			interests.ArchitectureScore += int16(labelCount.Count)
		case model.LabelAgriculture:
			interests.AgricultureScore += int16(labelCount.Count)
		case model.LabelNutritionScience:
			interests.NutritionScienceScore += int16(labelCount.Count)
		case model.LabelSleepScience:
			interests.SleepScienceScore += int16(labelCount.Count)
		case model.LabelProductivity:
			interests.ProductivityScore += int16(labelCount.Count)
		case model.LabelLeadership:
			interests.LeadershipScore += int16(labelCount.Count)
		case model.LabelInternationalRelations:
			interests.InternationalRelationsScore += int16(labelCount.Count)
		case model.LabelFuturePredictions:
			interests.FuturePredictionsScore += int16(labelCount.Count)
		case model.LabelEvents:
			interests.EventsScore += int16(labelCount.Count)
		case model.LabelCommunity:
			interests.CommunityScore += int16(labelCount.Count)
		case model.LabelTrends:
			interests.TrendsScore += int16(labelCount.Count)
		case model.LabelLifestyle:
			interests.LifestyleScore += int16(labelCount.Count)
		case model.LabelSoftwareDevelopment:
			interests.SoftwareDevelopmentScore += int16(labelCount.Count)
		case model.LabelProgrammingLanguages:
			interests.ProgrammingLanguagesScore += int16(labelCount.Count)
		case model.LabelWebDevelopment:
			interests.WebDevelopmentScore += int16(labelCount.Count)
		case model.LabelMobileAppDevelopment:
			interests.MobileAppDevelopmentScore += int16(labelCount.Count)
		case model.LabelDebuggingTechniques:
			interests.DebuggingTechniquesScore += int16(labelCount.Count)
		case model.LabelAlgorithmsMathematics:
			interests.AlgorithmsMathematicsScore += int16(labelCount.Count)
		case model.LabelDatabaseDesign:
			interests.DatabaseDesignScore += int16(labelCount.Count)
		case model.LabelCloudComputing:
			interests.CloudComputingScore += int16(labelCount.Count)
		case model.LabelServerManagement:
			interests.ServerManagementScore += int16(labelCount.Count)
		case model.LabelNetworkSecurity:
			interests.NetworkSecurityScore += int16(labelCount.Count)
		case model.LabelCryptography:
			interests.CryptographyScore += int16(labelCount.Count)
		case model.LabelArtificialIntelligence:
			interests.ArtificialIntelligenceScore += int16(labelCount.Count)
		case model.LabelMachineLearning:
			interests.MachineLearningScore += int16(labelCount.Count)
		case model.LabelDeepLearning:
			interests.DeepLearningScore += int16(labelCount.Count)
		case model.LabelComputerVision:
			interests.ComputerVisionScore += int16(labelCount.Count)
		case model.LabelNaturalLanguageProcessing:
			interests.NaturalLanguageProcessingScore += int16(labelCount.Count)
		case model.LabelBlockchainTechnology:
			interests.BlockchainTechnologyScore += int16(labelCount.Count)
		case model.LabelQuantumComputing:
			interests.QuantumComputingScore += int16(labelCount.Count)
		case model.LabelEdgeComputing:
			interests.EdgeComputingScore += int16(labelCount.Count)
		case model.LabelMicroservicesArchitecture:
			interests.MicroservicesArchitectureScore += int16(labelCount.Count)
		case model.LabelDevOps:
			interests.DevOpsScore += int16(labelCount.Count)
		case model.LabelContainerTechnology:
			interests.ContainerTechnologyScore += int16(labelCount.Count)
		case model.LabelCICD:
			interests.CICDScore += int16(labelCount.Count)
		case model.LabelTestAutomation:
			interests.TestAutomationScore += int16(labelCount.Count)
		case model.LabelUXUIDesign:
			interests.UXUIDesignScore += int16(labelCount.Count)
		case model.LabelAgileDevelopmentMethodologies:
			interests.AgileDevelopmentMethodologiesScore += int16(labelCount.Count)
		case model.LabelOpenSource:
			interests.OpenSourceScore += int16(labelCount.Count)
		case model.LabelVersionControl:
			interests.VersionControlScore += int16(labelCount.Count)
		case model.LabelAPIDesign:
			interests.APIDesignScore += int16(labelCount.Count)
		case model.LabelPerformanceOptimization:
			interests.PerformanceOptimizationScore += int16(labelCount.Count)
		}
	}
	return interests
}

