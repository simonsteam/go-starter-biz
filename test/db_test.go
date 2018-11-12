package test_test

import (
	"fmt"
	"github.com/go-pg/pg/orm"
	"local/biz/test"
	"testing"
)

// demo from go-pg ,@see https://github.com/go-pg/pg/wiki/Writing-Queries#has-many-to-many
func TestMany2ManyQuery(t *testing.T) {
	env := test.CreateEnv(t, test.GetTestDatabaseNameForCaller(), true)
	defer env.Release(t, false)

	type Item struct {
		Id    int
		Items []Item `pg:"many2many:item_to_items,joinFK:sub_id"`
	}

	type ItemToItem struct {
		ItemId int
		SubId  int
	}

	db := env.TestDB

	// Register many to many model so ORM can better recognize m2m relation.
	orm.RegisterTable((*ItemToItem)(nil))

	models := []interface{}{
		(*Item)(nil),
		(*ItemToItem)(nil),
	}
	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			panic(err)
		}
	}

	values := []interface{}{
		&Item{Id: 1},
		&Item{Id: 2},
		&Item{Id: 3},
		&ItemToItem{ItemId: 1, SubId: 2},
		&ItemToItem{ItemId: 1, SubId: 3},
	}
	for _, v := range values {
		err := db.Insert(v)
		if err != nil {
			panic(err)
		}
	}

	// Select item and all subitems with following queries:
	//
	// SELECT "item".* FROM "items" AS "item" ORDER BY "item"."id" LIMIT 1
	//
	// SELECT * FROM "items" AS "item"
	// JOIN "item_to_items" ON ("item_to_items"."item_id") IN ((1))
	// WHERE ("item"."id" = "item_to_items"."sub_id")

	var item Item
	err := db.Model(&item).Column("item.*", "Items").First()
	if err != nil {
		panic(err)
	}
	fmt.Println("Item", item.Id)
	fmt.Println("Subitems", item.Items[0].Id, item.Items[1].Id)
}
