package models

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"` // gorm:"index" 用于声明这个字段为索引 外键
	Tag   Tag `json:"tag"`                 // 嵌套的struct, 它利用TagID与Tag模型相互关联，在执行查询的时候，能够达到Article、Tag关联查询的功能

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
	CoverImageUrl string `json:"cover_image_url"`
}

// === 注释的方法已经被全局覆盖，不再需要
// func (article *Article) BeforeCreate(scope *gorm.Scope) error {
// 	scope.SetColumn("CreatedOn", time.Now().Unix())

// 	return nil
// }

// func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
// 	scope.SetColumn("ModifiedOn", time.Now().Unix())

// 	return nil
// }

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	return article.ID > 0
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)

	// Article有一个结构体成员是TagID，就是外键。gorm会通过类名+ID 的方式去找到这两个类之间的关联关系
	// Article有一个结构体成员是Tag，就是我们嵌套在Article里的Tag结构体，我们可以通过Related进行关联查询

	return
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	// v.(I) 是什么？
	// v表示一个接口值，I表示接口类型。这个实际就是 Golang 中的类型断言，用于判断一个接口值的实际类型是否为某个类型，或一个非接口值的类型是否实现了某个接口类型

	db.Create(&Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	})

	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}

// 硬删除
func CleanAllArticle() bool {
	// 硬删除要使用 Unscoped()，这是 GORM 的约定
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{})

	return true
}
