package models

import (
	"github.com/hridss/code-explore/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const (
	Item_Delete_False = 0
	Item_Delete_True  = 1
)

const Table_Item_Name = "item"

type Item struct {
}

var ItemModel = Item{}

// get item by item_id
func (r *Item) GetItemByItemId(itemId string) (item map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Item_Name).Where(map[string]interface{}{
		"item_id":   itemId,
		"is_delete": Item_Delete_False,
	}))
	if err != nil {
		return
	}
	item = rs.Row()
	return
}

// item_id and name is exists
func (r *Item) HasSameName(itemId, name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Item_Name).Where(map[string]interface{}{
		"item_id <>": itemId,
		"name":       name,
		"is_delete":  Item_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() > 0 {
		has = true
	}
	return
}

// name is exists
func (r *Item) HasItemName(name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Item_Name).Where(map[string]interface{}{
		"name":      name,
		"is_delete": Item_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() > 0 {
		has = true
	}
	return
}

// get item by name
func (r *Item) GetItemByName(name string) (item map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Item_Name).Where(map[string]interface{}{
		"name":      name,
		"is_delete": Item_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	item = rs.Row()
	return
}

// delete item by item_id
func (r *Item) Delete(itemId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Update(Table_Item_Name, map[string]interface{}{
		"is_delete":   Item_Delete_True,
		"update_time": time.Now().Unix(),
	}, map[string]interface{}{
		"item_id": itemId,
	}))
	if err != nil {
		return
	}
	return
}

// insert item
func (r *Item) Insert(itemValue map[string]interface{}) (id int64, err error) {

	itemValue["create_time"] = time.Now().Unix()
	itemValue["update_time"] = time.Now().Unix()
	db := G.DB()
	tx, err := db.Begin(db.Config)
	if err != nil {
		return
	}
	var rs *mysql.ResultSet
	rs, err = db.ExecTx(db.AR().Insert(Table_Item_Name, itemValue), tx)
	if err != nil {
		tx.Rollback()
		return
	}
	id = rs.LastInsertId
	return
}

// update item by item_id
func (r *Item) Update(itemId string, itemValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	itemValue["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Item_Name, itemValue, map[string]interface{}{
		"item_id":   itemId,
		"is_delete": Item_Delete_False,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get limit items by search keyword
func (r *Item) GetItemsByKeywordAndLimit(keyword string, limit int, number int) (items []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Item_Name).Where(map[string]interface{}{
		"name LIKE": "%" + keyword + "%",
		"is_delete": Item_Delete_False,
	}).Limit(limit, number).OrderBy("item_id", "DESC"))
	if err != nil {
		return
	}
	items = rs.Rows()

	return
}

// get limit items
func (r *Item) GetItemsByLimit(limit int, number int) (items []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			From(Table_Item_Name).
			Where(map[string]interface{}{
				"is_delete": Item_Delete_False,
			}).
			Limit(limit, number).
			OrderBy("item_id", "DESC"))
	if err != nil {
		return
	}
	items = rs.Rows()

	return
}

// get all items
func (r *Item) GetItems() (items []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Item_Name).Where(map[string]interface{}{
			"is_delete": Item_Delete_False,
		}))
	if err != nil {
		return
	}
	items = rs.Rows()
	return
}

// get item count
func (r *Item) CountItems() (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Item_Name).
			Where(map[string]interface{}{
				"is_delete": Item_Delete_False,
			}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get item count by keyword
func (r *Item) CountItemsByKeyword(keyword string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().
		Select("count(*) as total").
		From(Table_Item_Name).
		Where(map[string]interface{}{
			"name LIKE": "%" + keyword + "%",
			"is_delete": Item_Delete_False,
		}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get item by name
func (r *Item) GetItemByLikeName(name string) (items []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Item_Name).Where(map[string]interface{}{
		"name Like": "%" + name + "%",
		"is_delete": Item_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	items = rs.Rows()
	return
}

// get item by many item_id
func (r *Item) GetItemByitemIds(itemIds []string) (items []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Item_Name).Where(map[string]interface{}{
		"item_id":   itemIds,
		"is_delete": Item_Delete_False,
	}))
	if err != nil {
		return
	}
	items = rs.Rows()
	return
}

// update item by name
func (r *Item) UpdateItemByName(item map[string]interface{}) (affect int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	item["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Item_Name, item, map[string]interface{}{
		"name":      item["name"],
		"is_delete": Item_Delete_False,
	}))
	if err != nil {
		return
	}
	affect = rs.RowsAffected
	return
}
