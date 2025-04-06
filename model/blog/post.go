package blog

import (
	"time"

	"gorm.io/gorm"
)

// Post 文章模型
type Post struct {
	ID        int64          `gorm:"primaryKey;column:id" json:"id"`
	Title     string         `gorm:"column:title;not null;size:255" json:"title"`
	Content   string         `gorm:"column:content;not null;type:text" json:"content"`
	Slug      string         `gorm:"column:slug;not null;size:255;unique" json:"slug"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at"`
}

// TableName 指定表名
func (Post) TableName() string {
	return "post"
}

// PostModel 文章模型操作
type PostModel struct {
	db *gorm.DB
}

// NewPostModel 创建文章模型实例
func NewPostModel(db *gorm.DB) *PostModel {
	return &PostModel{
		db: db,
	}
}

// Create 创建文章
func (m *PostModel) Create(post *Post) error {
	return m.db.Create(post).Error
}

// GetByID 根据ID获取文章
func (m *PostModel) GetByID(id int64) (*Post, error) {
	var post Post
	err := m.db.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// GetBySlug 根据Slug获取文章
func (m *PostModel) GetBySlug(slug string) (*Post, error) {
	var post Post
	err := m.db.Where("slug = ?", slug).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// GetAll 获取所有文章
func (m *PostModel) GetAll() ([]Post, error) {
	var posts []Post
	err := m.db.Order("id desc").Find(&posts).Error
	return posts, err
}

// Update 更新文章
func (m *PostModel) Update(id int64, updates map[string]interface{}) error {
	return m.db.Model(&Post{}).Where("id = ?", id).Updates(updates).Error
}

// UpdateBySlug 根据slug更新文章
func (m *PostModel) UpdateBySlug(slug string, updates map[string]interface{}) error {
	return m.db.Model(&Post{}).Where("slug = ?", slug).Updates(updates).Error
}

// Delete 软删除文章
func (m *PostModel) Delete(id int64) error {
	return m.db.Delete(&Post{}, id).Error
}

// DeleteBySlug 根据slug软删除文章
func (m *PostModel) DeleteBySlug(slug string) error {
	// 软删除
	post := &Post{
		Slug: slug,
	}
	return m.db.Delete(post).Error
}

// HardDelete 硬删除文章
func (m *PostModel) HardDelete(id int64) error {
	return m.db.Unscoped().Delete(&Post{}, id).Error
}

// GetDeleted 获取已删除的文章
func (m *PostModel) GetDeleted() ([]Post, error) {
	var posts []Post
	err := m.db.Unscoped().Where("deleted_at IS NOT NULL").Find(&posts).Error
	return posts, err
}

// Restore 恢复软删除的文章
func (m *PostModel) Restore(id int64) error {
	return m.db.Model(&Post{}).Where("id = ?", id).Update("deleted_at", nil).Error
}

// Search 搜索文章
func (m *PostModel) Search(keyword string) ([]Post, error) {
	var posts []Post
	err := m.db.Where("title LIKE ? OR content LIKE ?",
		"%"+keyword+"%", "%"+keyword+"%").Find(&posts).Error
	return posts, err
}

// GetByPage 分页获取文章
func (m *PostModel) GetByPage(page, pageSize int) ([]Post, int64, error) {
	var posts []Post
	var total int64

	err := m.db.Model(&Post{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = m.db.Offset((page - 1) * pageSize).Limit(pageSize).
		Order("id desc").Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}
