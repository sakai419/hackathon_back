package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)
func (s *Service) GetReportedUserInfosOrderByReportCount(ctx context.Context, params *model.GetReportedUserInfosOrderByReportCountParams) ([]*model.ReportedUserInfo, error) {
	// Validate input parameters
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Get reported accountIDs order by report count
	reportedAccountIDs, err := s.repo.GetReportedAccountIDsOrderByReportCount(ctx, &model.GetReportedAccountIDsOrderByReportCountParams{
		Limit:  params.Limit,
		Offset: params.Offset,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get reported user infos order by report count", err)
	}

	// Extract accountIDs
	accountIDs := make([]string, 0, len(reportedAccountIDs))
	for _, reportedAccountID := range reportedAccountIDs {
		accountIDs = append(accountIDs, reportedAccountID.AccountID)
	}

	// Get user infos
	userInfos, err := s.repo.GetUserInfos(ctx, &model.GetUserInfosParams{
		TargetAccountIDs: accountIDs,
		ClientAccountID: "dammy",
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get user infos by account IDs", err)
	}

	// Create user info map
	userInfoMap := make(map[string]*model.UserInfoInternal, len(userInfos))
	for _, userInfo := range userInfos {
		userInfoMap[userInfo.ID] = userInfo
	}

	// Create reported user infos
	reportedUserInfos := make([]*model.ReportedUserInfo, 0, len(reportedAccountIDs))
	for _, reportedAccountID := range reportedAccountIDs {
		userInfo, ok := userInfoMap[reportedAccountID.AccountID]
		if !ok {
			continue
		}

		reportedUserInfos = append(reportedUserInfos, &model.ReportedUserInfo{
			UserInfo:    *convertToUserInfo(userInfo),
			ReportCount: reportedAccountID.ReportCount,
		})
	}

	return reportedUserInfos, nil
}

func (s *Service) GetReportsByReportedAccountID(ctx context.Context, params *model.GetReportsByReportedAccountIDParams) ([]*model.Report, error) {
	// Validate input parameters
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Get reports
	reports, err := s.repo.GetReportsByReportedAccountID(ctx, &model.GetReportsByReportedAccountIDParams{
		ReportedAccountID: params.ReportedAccountID,
		Limit:             params.Limit,
		Offset:            params.Offset,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get reports by reported account ID", err)
	}

	// Extract reporter accountIDs
	reporterAccountIDsMap := make(map[string]struct{}, len(reports))
	for _, report := range reports {
		reporterAccountIDsMap[report.ReporterAccountID] = struct{}{}
	}
	reporterAccountIDs := make([]string, 0, len(reports))
	for reporterAccountID := range reporterAccountIDsMap {
		reporterAccountIDs = append(reporterAccountIDs, reporterAccountID)
	}

	// Get user infos
	userInfos, err := s.repo.GetUserInfos(ctx, &model.GetUserInfosParams{
		TargetAccountIDs: reporterAccountIDs,
		ClientAccountID: "dammy",
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get user infos by account IDs", err)
	}

	// Create user info map
	userInfoMap := make(map[string]*model.UserInfoInternal, len(userInfos))
	for _, userInfo := range userInfos {
		userInfoMap[userInfo.ID] = userInfo
	}

	// Create reports
	var resultReports []*model.Report
	for _, report := range reports {
		userInfo, ok := userInfoMap[report.ReporterAccountID]
		if !ok {
			continue
		}

		resultReports = append(resultReports, &model.Report{
			ReportID:          report.ReportID,
			ReporterInfo:      *convertToUserInfo(userInfo),
			Reason:            report.Reason,
			Content:           report.Content,
			CreatedAt:         report.CreatedAt,
		})
	}

	return resultReports, nil
}