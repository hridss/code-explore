package controllers

import (
	"github.com/hridss/code-explore/app/models"
	"github.com/hridss/code-explore/app/utils"
	"regexp"
	"strings"
)

type ItemController struct {
	BaseController
}

func (this *ItemController) Add() {
	this.viewLayout("item/form", "item")
}

func (this *ItemController) Save() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/item/list")
	}
	name := strings.TrimSpace(this.GetString("name", ""))
	description := strings.TrimSpace(this.GetString("description", ""))
	sequence := strings.TrimSpace(this.GetString("sequence", "0"))

	if name == "" {
		this.jsonError("分类名称不能为空！")
	}
	match, err := regexp.MatchString(`[\\\\/:*?\"<>、|]`, name)
	if err != nil {
		this.jsonError("分类名称格式不正确！")
	}
	if match {
		this.jsonError("分类名称格式不正确！")
	}
	ok, err := models.ItemModel.HasItemName(name)
	if err != nil {
		this.ErrorLog("添加分类失败：" + err.Error())
		this.jsonError("添加分类失败！")
	}
	if ok {
		this.jsonError("分类名已经存在！")
	}

	// create Item database
	ItemId, err := models.ItemModel.Insert(map[string]interface{}{
		"name":        name,
		"description": description,
		"sequence":    sequence,
	})
	if err != nil {
		this.ErrorLog("添加分类失败：" + err.Error())
		this.jsonError("添加分类失败")
	}

	this.InfoLog("添加分类 " + utils.Convert.IntToString(ItemId, 10) + " 成功")
	this.jsonSuccess("添加分类成功", nil, "/system/item/list")
}

func (this *ItemController) List() {

	page, _ := this.GetInt("page", 1)
	keyword := strings.TrimSpace(this.GetString("keyword", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)
	limit := (page - 1) * number

	var err error
	var count int64
	var Items []map[string]string
	if keyword != "" {
		count, err = models.ItemModel.CountItemsByKeyword(keyword)
		Items, err = models.ItemModel.GetItemsByKeywordAndLimit(keyword, limit, number)
	} else {
		count, err = models.ItemModel.CountItems()
		Items, err = models.ItemModel.GetItemsByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("获取分类列表失败: " + err.Error())
		this.ViewError("获取分类列表失败", "/system/main/index")
	}

	this.Data["Items"] = Items
	this.Data["keyword"] = keyword
	this.SetPaginator(number, count)
	this.viewLayout("item/list", "item")
}

func (this *ItemController) Edit() {

	ItemId := this.GetString("item_id", "")
	if ItemId == "" {
		this.ViewError("分类不存在", "/system/item/list")
	}

	Item, err := models.ItemModel.GetItemByItemId(ItemId)
	if err != nil {
		this.ErrorLog("查找分类失败: " + err.Error())
		this.ViewError("查找分类失败", "/system/item/list")
	}
	if len(Item) == 0 {
		this.ViewError("分类不存在", "/system/item/list")
	}

	this.Data["Item"] = Item
	this.viewLayout("item/form", "item")
}

func (this *ItemController) Modify() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/item/list")
	}
	ItemId := this.GetString("item_id", "")
	name := strings.TrimSpace(this.GetString("name", ""))
	description := strings.TrimSpace(this.GetString("description", ""))
	sequence := strings.TrimSpace(this.GetString("sequence", "0"))

	if ItemId == "" {
		this.jsonError("分类不存在！")
	}
	if name == "" {
		this.jsonError("分类名称不能为空！")
	}
	match, err := regexp.MatchString(`[\\\\/:*?\"<>、|]`, name)
	if err != nil {
		this.jsonError("分类名称格式不正确！")
	}
	if match {
		this.jsonError("分类名称格式不正确！")
	}

	Item, err := models.ItemModel.GetItemByItemId(ItemId)
	if err != nil {
		this.ErrorLog("修改分类 " + ItemId + " 失败: " + err.Error())
		this.jsonError("修改分类失败！")
	}
	if len(Item) == 0 {
		this.jsonError("分类不存在！")
	}

	ok, _ := models.ItemModel.HasSameName(ItemId, name)
	if ok {
		this.jsonError("分类名已经存在！")
	}

	ItemValue := map[string]interface{}{
		"name":        name,
		"description": description,
		"sequence":    sequence,
	}
	// update Item document dir name if name update
	_, err = models.ItemModel.Update(ItemId, ItemValue)
	if err != nil {
		this.ErrorLog("修改分类 " + ItemId + " 失败：" + err.Error())
		this.jsonError("修改分类失败")
	}
	this.InfoLog("修改分类 " + ItemId + " 成功")
	this.jsonSuccess("修改分类成功", nil, "/system/item/list")
}

func (this *ItemController) Delete() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/item/list")
	}
	ItemId := this.GetString("item_id", "")
	if ItemId == "" {
		this.jsonError("没有选择分类！")
	}

	Item, err := models.ItemModel.GetItemByItemId(ItemId)
	if err != nil {
		this.ErrorLog("删除分类 " + ItemId + " 失败: " + err.Error())
		this.jsonError("删除分类失败")
	}
	if len(Item) == 0 {
		this.jsonError("分类不存在")
	}

	this.InfoLog("删除分类 " + ItemId + " 成功")
	this.jsonSuccess("删除分类成功", nil, "/system/item/list")
}