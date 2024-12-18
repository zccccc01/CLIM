// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"CLIM/app/user/dao/model"
)

func newUserProfile(db *gorm.DB, opts ...gen.DOOption) userProfile {
	_userProfile := userProfile{}

	_userProfile.userProfileDo.UseDB(db, opts...)
	_userProfile.userProfileDo.UseModel(&model.UserProfile{})

	tableName := _userProfile.userProfileDo.TableName()
	_userProfile.ALL = field.NewAsterisk(tableName)
	_userProfile.UserID = field.NewBytes(tableName, "user_id")
	_userProfile.AvatarURL = field.NewString(tableName, "avatar_url")
	_userProfile.OnlineStatus = field.NewInt32(tableName, "online_status")
	_userProfile.Bio = field.NewString(tableName, "bio")
	_userProfile.Birthday = field.NewTime(tableName, "birthday")
	_userProfile.Gender = field.NewString(tableName, "gender")
	_userProfile.Location = field.NewString(tableName, "location")
	_userProfile.LastSeenAt = field.NewTime(tableName, "last_seen_at")

	_userProfile.fillFieldMap()

	return _userProfile
}

type userProfile struct {
	userProfileDo userProfileDo

	ALL          field.Asterisk
	UserID       field.Bytes // User UUID (foreign key to users table)
	AvatarURL    field.String
	OnlineStatus field.Int32  // 0: offline, 1: online
	Bio          field.String // Short biography or status message
	Birthday     field.Time   // Users date of birth
	Gender       field.String // User gender
	Location     field.String // Users location or timezone
	LastSeenAt   field.Time   // Last seen timestamp for the user

	fieldMap map[string]field.Expr
}

func (u userProfile) Table(newTableName string) *userProfile {
	u.userProfileDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userProfile) As(alias string) *userProfile {
	u.userProfileDo.DO = *(u.userProfileDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userProfile) updateTableName(table string) *userProfile {
	u.ALL = field.NewAsterisk(table)
	u.UserID = field.NewBytes(table, "user_id")
	u.AvatarURL = field.NewString(table, "avatar_url")
	u.OnlineStatus = field.NewInt32(table, "online_status")
	u.Bio = field.NewString(table, "bio")
	u.Birthday = field.NewTime(table, "birthday")
	u.Gender = field.NewString(table, "gender")
	u.Location = field.NewString(table, "location")
	u.LastSeenAt = field.NewTime(table, "last_seen_at")

	u.fillFieldMap()

	return u
}

func (u *userProfile) WithContext(ctx context.Context) *userProfileDo {
	return u.userProfileDo.WithContext(ctx)
}

func (u userProfile) TableName() string { return u.userProfileDo.TableName() }

func (u userProfile) Alias() string { return u.userProfileDo.Alias() }

func (u userProfile) Columns(cols ...field.Expr) gen.Columns { return u.userProfileDo.Columns(cols...) }

func (u *userProfile) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userProfile) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 8)
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["avatar_url"] = u.AvatarURL
	u.fieldMap["online_status"] = u.OnlineStatus
	u.fieldMap["bio"] = u.Bio
	u.fieldMap["birthday"] = u.Birthday
	u.fieldMap["gender"] = u.Gender
	u.fieldMap["location"] = u.Location
	u.fieldMap["last_seen_at"] = u.LastSeenAt
}

func (u userProfile) clone(db *gorm.DB) userProfile {
	u.userProfileDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userProfile) replaceDB(db *gorm.DB) userProfile {
	u.userProfileDo.ReplaceDB(db)
	return u
}

type userProfileDo struct{ gen.DO }

func (u userProfileDo) Debug() *userProfileDo {
	return u.withDO(u.DO.Debug())
}

func (u userProfileDo) WithContext(ctx context.Context) *userProfileDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userProfileDo) ReadDB() *userProfileDo {
	return u.Clauses(dbresolver.Read)
}

func (u userProfileDo) WriteDB() *userProfileDo {
	return u.Clauses(dbresolver.Write)
}

func (u userProfileDo) Session(config *gorm.Session) *userProfileDo {
	return u.withDO(u.DO.Session(config))
}

func (u userProfileDo) Clauses(conds ...clause.Expression) *userProfileDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userProfileDo) Returning(value interface{}, columns ...string) *userProfileDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userProfileDo) Not(conds ...gen.Condition) *userProfileDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userProfileDo) Or(conds ...gen.Condition) *userProfileDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userProfileDo) Select(conds ...field.Expr) *userProfileDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userProfileDo) Where(conds ...gen.Condition) *userProfileDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userProfileDo) Order(conds ...field.Expr) *userProfileDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userProfileDo) Distinct(cols ...field.Expr) *userProfileDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userProfileDo) Omit(cols ...field.Expr) *userProfileDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userProfileDo) Join(table schema.Tabler, on ...field.Expr) *userProfileDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userProfileDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userProfileDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userProfileDo) RightJoin(table schema.Tabler, on ...field.Expr) *userProfileDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userProfileDo) Group(cols ...field.Expr) *userProfileDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userProfileDo) Having(conds ...gen.Condition) *userProfileDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userProfileDo) Limit(limit int) *userProfileDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userProfileDo) Offset(offset int) *userProfileDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userProfileDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userProfileDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userProfileDo) Unscoped() *userProfileDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userProfileDo) Create(values ...*model.UserProfile) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userProfileDo) CreateInBatches(values []*model.UserProfile, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userProfileDo) Save(values ...*model.UserProfile) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userProfileDo) First() (*model.UserProfile, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserProfile), nil
	}
}

func (u userProfileDo) Take() (*model.UserProfile, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserProfile), nil
	}
}

func (u userProfileDo) Last() (*model.UserProfile, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserProfile), nil
	}
}

func (u userProfileDo) Find() ([]*model.UserProfile, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserProfile), err
}

func (u userProfileDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserProfile, err error) {
	buf := make([]*model.UserProfile, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userProfileDo) FindInBatches(result *[]*model.UserProfile, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userProfileDo) Attrs(attrs ...field.AssignExpr) *userProfileDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userProfileDo) Assign(attrs ...field.AssignExpr) *userProfileDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userProfileDo) Joins(fields ...field.RelationField) *userProfileDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userProfileDo) Preload(fields ...field.RelationField) *userProfileDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userProfileDo) FirstOrInit() (*model.UserProfile, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserProfile), nil
	}
}

func (u userProfileDo) FirstOrCreate() (*model.UserProfile, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserProfile), nil
	}
}

func (u userProfileDo) FindByPage(offset int, limit int) (result []*model.UserProfile, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userProfileDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userProfileDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userProfileDo) Delete(models ...*model.UserProfile) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userProfileDo) withDO(do gen.Dao) *userProfileDo {
	u.DO = *do.(*gen.DO)
	return u
}
