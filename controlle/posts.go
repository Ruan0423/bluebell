package controlle

import (
	"fmt"
	"strconv"
	"web_framework/logic"
	"web_framework/models"

	"github.com/gin-gonic/gin"
)

func CreatePostHandler(c *gin.Context) {
	//获取参数
	postparam := new(models.Post)
	if err := c.ShouldBindJSON(postparam); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	//获取创建用户的id
	userid, err := GetUserId(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	postparam.AuthorID = userid

	// 创建帖子
	if err := logic.CreatePost(postparam); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)

}

func GetPostHandle(c *gin.Context) {
	//获取参数
	postidstr := c.Param("id")
	post_id, err := strconv.ParseInt(postidstr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//从数据库获取帖子信息
	postdata, err := logic.GetPostApi(post_id)
	if err != nil {
		ResponseErrorwithMsg(c, CodeServerBusy, err)
		return
	}

	//获取帖子作者名字
	Athor_name, err := logic.GetAthorByAthorid(postdata.AuthorID)
	if err != nil {
		ResponseErrorwithMsg(c, CodeServerBusy, err)
		return
	}

	//获取帖子社区名字
	Community_Detail, err := logic.GetCommunityDetatl(postdata.CommunityID)
	if err != nil {
		ResponseErrorwithMsg(c, CodeServerBusy, err)
		return
	}
	fmt.Println("社区查询成功：", Community_Detail)
	PostApi := new(models.PostApi)
	PostApi.Post = postdata
	PostApi.Athor = Athor_name
	PostApi.CommunityDetail = Community_Detail

	ResponseSuccess(c, PostApi)

}

// GetPostListHandler 获取帖子列表
func GetPostListHandler(c *gin.Context) {

	// 获取参数
	pagenum , pageSize := GetPOstParam(c)

	//获取post列表
	postList, err := logic.GetPostList(pagenum, pageSize)

	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	PostDetaiList := make([]*models.PostApi, 0, len(postList))

	for i := 0; i < len(postList); i++ {

		//获取帖子作者名字
		Athor_name, err := logic.GetAthorByAthorid(postList[i].AuthorID)
		if err != nil {
			ResponseErrorwithMsg(c, CodeServerBusy, err)
			return
		}

		//获取帖子社区名字
		Community_Detail, err := logic.GetCommunityDetatl(postList[i].CommunityID)
		if err != nil {
			ResponseErrorwithMsg(c, CodeServerBusy, err)
			return
		}
		fmt.Println("社区查询成功：", Community_Detail)
		PostApi := new(models.PostApi)
		PostApi.Post = postList[i]
		PostApi.Athor = Athor_name
		PostApi.CommunityDetail = Community_Detail

		PostDetaiList = append(PostDetaiList, PostApi)

	}
	ResponseSuccess(c, PostDetaiList)

}
