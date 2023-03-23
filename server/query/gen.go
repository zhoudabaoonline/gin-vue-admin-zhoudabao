// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q       = new(Query)
	Strcut3 *strcut3
	Struct1 *struct1
	Struct2 *struct2
	Struct4 *struct4
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Strcut3 = &Q.Strcut3
	Struct1 = &Q.Struct1
	Struct2 = &Q.Struct2
	Struct4 = &Q.Struct4
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:      db,
		Strcut3: newStrcut3(db, opts...),
		Struct1: newStruct1(db, opts...),
		Struct2: newStruct2(db, opts...),
		Struct4: newStruct4(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Strcut3 strcut3
	Struct1 struct1
	Struct2 struct2
	Struct4 struct4
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:      db,
		Strcut3: q.Strcut3.clone(db),
		Struct1: q.Struct1.clone(db),
		Struct2: q.Struct2.clone(db),
		Struct4: q.Struct4.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:      db,
		Strcut3: q.Strcut3.replaceDB(db),
		Struct1: q.Struct1.replaceDB(db),
		Struct2: q.Struct2.replaceDB(db),
		Struct4: q.Struct4.replaceDB(db),
	}
}

type queryCtx struct {
	Strcut3 IStrcut3Do
	Struct1 IStruct1Do
	Struct2 IStruct2Do
	Struct4 IStruct4Do
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Strcut3: q.Strcut3.WithContext(ctx),
		Struct1: q.Struct1.WithContext(ctx),
		Struct2: q.Struct2.WithContext(ctx),
		Struct4: q.Struct4.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
