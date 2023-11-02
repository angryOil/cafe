package handler

import (
	"cafe/internal/controller"
	"cafe/internal/controller/member"
	"cafe/internal/controller/memberRole"
)

type MemberRoleHandler struct {
	cafeCon  controller.CafeController //
	memCon   member.Controller
	mRoleCon memberRole.Controller
}

func NewMemberRoleHandler(cafeCon controller.CafeController, memCon member.Controller, mRoleCon memberRole.Controller) MemberRoleHandler {
	return MemberRoleHandler{
		cafeCon:  cafeCon,
		memCon:   memCon,
		mRoleCon: mRoleCon,
	}
}
