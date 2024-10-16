//go:generate oapi-codegen -generate types,gin,spec,skip-prune -import-mapping ./common.yaml:github.com/supsysjp/engage/controller/common -package email -o schema.gen.go ./../../reference/email.yaml
package email

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/supsysjp/engage/controller/common"
	"github.com/supsysjp/engage/logger"
	"github.com/supsysjp/engage/usecase"
	"github.com/supsysjp/engage/util"
)

type emailController struct {
	emailu usecase.EmailUsecase
	unsubu usecase.EmailUnsubscribeUsecase
}

type EmailControllerInput struct {
	Emailu usecase.EmailUsecase
	Unsubu usecase.EmailUnsubscribeUsecase
}

func NewEmail(input EmailControllerInput) ServerInterface {
	return &emailController{
		emailu: input.Emailu,
		unsubu: input.Unsubu,
	}
}

func (c *emailController) PostEngageV1EmailPublishRegistrations(ctx *gin.Context) {
	//cognito認証を用いてheaderからユーザ情報を取得
	au, err := common.GetAuthenticatedUser(ctx)
	if err != nil {
		if errors.Is(err, common.ErrAuthenticatedUserIsNil) {
			ctx.JSON(common.NewUnauthorized("", "failed to authenticate user", err))
			return
		}
		ctx.JSON(common.NewInternalServerError("", "Get Authenticated User failed", err))
		return
	}

	var reqBody PostEngageV1EmailPublishRegistrationsJSONRequestBody
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(common.NewBadRequest2("", fmt.Sprintf("invalid request: %s", err.Error())))
		return
	}
	m, err := ToEmailPublishRegistration(ctx, au.TenantID, reqBody)
	if err != nil {
		ctx.JSON(common.NewBadRequest2("", fmt.Sprintf("invalid request: %s", err.Error())))
		return
	}
	go func() {
		if err := c.emailu.Create(context.Background(), *m, reqBody.PathToPublishList); err != nil {
			logger.App.Error(
				au.TenantID,
				"failed to create email publish registration",
				map[string]any{
					"target_group_id": m.TargetGroupID,
				},
				err,
			)
		}
	}()

	var resParam EmailPublishRegistrationsPostResponse
	resParam.TargetGroupId = m.TargetGroupID
	ctx.JSON(http.StatusOK, resParam)
}

func (c *emailController) GetEngageV1EmailPublishRegistrationsTargetGroupId(ctx *gin.Context, targetGroupID string) {
	//cognito認証を用いてheaderからユーザ情報を取得
	au, err := common.GetAuthenticatedUser(ctx)
	if err != nil {
		if errors.Is(err, common.ErrAuthenticatedUserIsNil) {
			ctx.JSON(common.NewUnauthorized("", "failed to authenticate user", err))
			return
		}
		ctx.JSON(common.NewInternalServerError("", "Get Authenticated User failed", err))
		return
	}

	m, err := c.emailu.FetchByID(ctx, au.TenantID, targetGroupID)
	if err != nil {
		uerr := common.ToUsecaseError(err)
		if uerr == nil {
			// UsecaseErrorではない場合はInternalServerError
			ctx.JSON(common.NewInternalServerError("", "internal server error", err))
			return
		} else if uerr.Code() == usecase.EmailErrCodeRecordNotFound {
			ctx.JSON(common.NewNotFound(uerr.Code().ToStr(), "failed to fetch email publish registration", uerr))
			return
		} else {
			ctx.JSON(common.NewBadRequest2(uerr.Code().ToStr(), fmt.Sprintf("invalid request: %s", uerr.Error())))
			return
		}
	}
	resParam := EmailPublishRegistrationsFetchByIDResponse{
		TargetGroupId:   m.TargetGroupID,
		ExecutionStatus: EmailPublishRegistrationsExecutionStatus(m.ExecutionStatus),
		ErrorMsg:        &m.ErrorMsg.String,
	}
	ctx.JSON(http.StatusOK, resParam)
}

func (c *emailController) PostEngageV1EmailTestPublishes(ctx *gin.Context) {
	//cognito認証を用いてheaderからユーザ情報を取得
	au, err := common.GetAuthenticatedUser(ctx)
	if err != nil {
		if errors.Is(err, common.ErrAuthenticatedUserIsNil) {
			ctx.JSON(common.NewUnauthorized("", "failed to authenticate user", err))
			return
		}
		ctx.JSON(common.NewInternalServerError("", "Get Authenticated User failed", err))
		return
	}

	fh, err := ctx.FormFile("publish_list")
	if err != nil && err != http.ErrMissingFile {
		ctx.JSON(common.NewBadRequest2("", fmt.Sprintf("invalid file: %s", err.Error())))
		return
	}
	var csv *os.File
	defer func() {
		if csv != nil {
			csv.Close()
			os.Remove(csv.Name())
		}
	}()
	if fh != nil {
		csv, err = util.CreateTemp()
		if err != nil {
			ctx.JSON(common.NewInternalServerError("", "failed to create temp file", err))
			return
		}
		if err = util.FileHeaderToFile(fh, csv); err != nil {
			ctx.JSON(common.NewBadRequest2("", fmt.Sprintf("invalid file: %s", err.Error())))
			return
		}
		//if _, err = csv.Seek(0, 0); err != nil {
		//	ctx.JSON(common.NewInternalServerError("", "failed to seek file", err))
		//}
	}
	var reqBody PostEngageV1EmailTestPublishesMultipartRequestBody
	if err = ctx.Bind(&reqBody); err != nil {
		ctx.JSON(common.NewBadRequest2("", fmt.Sprintf("invalid request: %s", err.Error())))
		return
	}
	m, err := ToEmailTestPublishInput(ctx.Request.Context(), reqBody)
	if err != nil {
		ctx.JSON(common.NewBadRequest2("", fmt.Sprintf("invalid request: %s", err.Error())))
		return
	}
	if err = c.emailu.ExecTestPublishes(ctx.Request.Context(), au.TenantID, *m, csv); err != nil {
		uerr := common.ToUsecaseError(err) // UsecaseErrorに変換してみる
		if uerr == nil {
			// nil の時はInternal Server Error
			ctx.JSON(common.NewInternalServerError("", "failed to exec test publishes", err))
			return
		} else {
			// レスポンスデータには、UsecaseError が持つエラーコードをセットする
			ctx.JSON(common.NewBadRequest2(uerr.Code().ToStr(), fmt.Sprintf("invalid request: %s", uerr.Error())))
			return
		}
	}
	ctx.Writer.WriteHeader(http.StatusOK)
}

func (c *emailController) PostEngageV1EmailPublishRegistrationsTargetGroupIdCancel(ctx *gin.Context, targetGroupID string) {

	//cognito認証を用いてheaderからユーザ情報を取得
	au, err := common.GetAuthenticatedUser(ctx)
	if err != nil {
		if errors.Is(err, common.ErrAuthenticatedUserIsNil) {
			ctx.JSON(common.NewUnauthorized("", "failed to authenticate user", err))
			return
		}
		ctx.JSON(common.NewInternalServerError("", "Get Authenticated User failed", err))
		return
	}

	if err := c.emailu.Cancel(ctx, au.TenantID, targetGroupID); err != nil {
		uerr := common.ToUsecaseError(err)
		if uerr == nil {
			// nil の時はInternal Server Error
			ctx.JSON(common.NewInternalServerError("", "failed to update execution_status to canceled", err))
			return
		} else {
			// レスポンスデータには、UsecaseError が持つエラーコードをセットする
			ctx.JSON(common.NewBadRequest2(uerr.Code().ToStr(), fmt.Sprintf("failed to update execution_status to canceled: %s", uerr.Error())))
			return
		}
	}
	ctx.Writer.WriteHeader(http.StatusOK)

}

func (c *emailController) GetEngageV1EmailPublishRegistrationsStatus(ctx *gin.Context, s1 GetEngageV1EmailPublishRegistrationsStatusParams) {
	//cognito認証を用いてheaderからユーザ情報を取得
	au, err := common.GetAuthenticatedUser(ctx)
	if err != nil {
		if errors.Is(err, common.ErrAuthenticatedUserIsNil) {
			ctx.JSON(common.NewUnauthorized("", "failed to authenticate user", err))
			return
		}
		ctx.JSON(common.NewInternalServerError("", "Get Authenticated User failed", err))
		return
	}

	var resParamList []EmailPublishRegistrationStatus
	targetGroupIdList := strings.Split(s1.TargetGroupIds, ",")

	//配信予約情報を取得する
	//エラー時はInternalServerErrorを返す
	m, err := c.emailu.FetchByIDs(ctx, au.TenantID, targetGroupIdList)
	if err != nil {
		ctx.JSON(common.NewInternalServerError("", "internal server error", err))
		return
	}
	//返すデータが存在しないときは空の配列を返す
	if m == nil {
		resParamList = []EmailPublishRegistrationStatus{}
	}

	//配信予約情報を配信予約ステータスに変換する
	//エラー時はInternalServerErrorを返す
	for _, element := range m {
		resParam := EmailPublishRegistrationStatus{
			TargetGroupId:       element.TargetGroupID,
			ExecutionStatus:     string(element.ExecutionStatus),
			TargetCustomerCount: int(element.TargetCustomerCount.Int32),
		}
		resParamList = append(resParamList, resParam)
	}

	ctx.JSON(http.StatusOK, resParamList)
}

// メール即時配信
// (POST /engage/v1/email/immediate_publishes)
func (c *emailController) PostEngageV1EmailImmediatePublishes(ctx *gin.Context) {
	var reqBody PostEngageV1EmailImmediatePublishesJSONRequestBody
	if err := ctx.BindJSON(&reqBody); err != nil {
		// 400 Bad Request
		ctx.JSON(common.NewBadRequest2("", fmt.Sprintf("invalid request: %s", err.Error())))
		return
	}
	if errInfos := c.emailu.RequestImmediatePublish(ctx, reqBody.PublishListNames); len(errInfos) > 0 {
		var res EmailImmediatePublishesPostInternalServerErrorResponse
		res.Errors = make([]EmailImmediatePublishErrInfo, len(errInfos))
		for i, errInfo := range errInfos {
			res.Errors[i] = EmailImmediatePublishErrInfo{
				PublishListName: errInfo.PublishListName,
				Code:            errInfo.Code,
				Msg:             errInfo.Msg,
				Detail:          errInfo.Detail,
			}
		}
		// 500 Internal Server Error
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	// 204 No Content
	ctx.Writer.WriteHeader(http.StatusNoContent)
}

// 購読解除リストインポート
// (POST /engage/v1/email/list_unsubscribe/import)
func (c *emailController) PostEngageV1EmailListUnsubscribeImport(ctx *gin.Context) {
	// 非同期で購読解除リストをインポートする
	go func() {
		if err := c.unsubu.ImportUnsubscribeList(ctx); err != nil {
			logger.App.Error("system", "failed to import unsubscribe list", nil, err)
		}
	}()
	// 204 No Content
	ctx.Writer.WriteHeader(http.StatusNoContent)
}