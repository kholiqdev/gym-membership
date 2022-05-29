package member

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gym/cmd/domain/member/dto"
	"gym/cmd/domain/member/entity"
	"gym/cmd/domain/member/service"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

var memberService = &service.MemberServiceMock{Mock: mock.Mock{}}
var memberHandler = MemberHandlerImpl{SvcMember: memberService}

func TestMemberHandler_GetSuccess(t *testing.T) {
	member := entity.Member{
		ID:       1,
		Name:     "Abdul Kholiq",
		Phone:    "085959119",
		Email:    "kholiqdev@icloud.com",
		Password: "rahasia",
	}
	memberList := entity.MemberList{&member}
	//program mock
	memberService.Mock.On("GetMembers").Return(dto.CreateMemberListResponse(&memberList), nil).Once()

	dataResponse := struct {
		Message string `json:"message"`
		Result  struct {
			Members dto.MemberListResponse `json:"members"`
		} `json:"result"`
	}{}

	e := echo.New()
	r := httptest.NewRequest("GET", "/member", nil)
	r.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)
	memberHandler.Get(ctx)
	bodyRes, _ := ioutil.ReadAll(w.Result().Body)
	var err = json.Unmarshal(bodyRes, &dataResponse)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Result().StatusCode)
	assert.EqualValues(t, dataResponse.Message, "Success")
	assert.EqualValues(t, dataResponse.Result.Members[0].Email, member.Email)

}

func TestMemberHandler_CreateFail(t *testing.T) {
	//program mock
	memberService.Mock.On("Store", mock.Anything).Return(nil, errors.New("fail on get service")).Once()

	e := echo.New()
	bodyReq := make(map[string]interface{})
	bodyReq["name"] = "Abdul Koliq"
	bodyReq["email"] = "kholiqdev@icloud.com"
	bodyReq["password"] = "password"
	bodyJson, _ := json.Marshal(bodyReq)
	r := httptest.NewRequest("POST", "/member", bytes.NewReader(bodyJson))
	r.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)
	memberHandler.Create(ctx)
	bodyRes, _ := ioutil.ReadAll(w.Result().Body)
	defer r.Body.Close()

	assert.Equal(t, 400, w.Result().StatusCode)
	assert.NotContains(t, string(bodyRes), "Success", "Not Contains Success Message")
}

func TestMemberHandler_CreateSuccess(t *testing.T) {
	res := dto.MemberResponse{
		ID:    1,
		Name:  "Abdul Koliq",
		Email: "kholiqdev@icloud.com",
	}
	//program mock
	memberService.Mock.On("Store", mock.Anything).Return(res, nil).Once()

	e := echo.New()
	bodyReq := make(map[string]interface{})
	bodyReq["name"] = "Abdul Koliq"
	bodyReq["email"] = "kholiqdev@icloud.com"
	bodyReq["password"] = "password"
	bodyJson, _ := json.Marshal(bodyReq)
	r := httptest.NewRequest("POST", "/member", bytes.NewReader(bodyJson))
	r.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)
	memberHandler.Create(ctx)
	bodyRes, _ := ioutil.ReadAll(w.Result().Body)
	defer r.Body.Close()

	assert.Equal(t, 201, w.Result().StatusCode)
	assert.Contains(t, string(bodyRes), "Success", "Message Contains")
	assert.Contains(t, string(bodyRes), res.Email, "Email Contains")
}
