// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// ClickbaitKeyword is an object representing the database table.
type ClickbaitKeyword struct {
	ID      string `boil:"id" json:"id" toml:"id" yaml:"id"`
	LabelID string `boil:"label_id" json:"label_id" toml:"label_id" yaml:"label_id"`
	Keyword string `boil:"keyword" json:"keyword" toml:"keyword" yaml:"keyword"`

	R *clickbaitKeywordR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L clickbaitKeywordL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ClickbaitKeywordColumns = struct {
	ID      string
	LabelID string
	Keyword string
}{
	ID:      "id",
	LabelID: "label_id",
	Keyword: "keyword",
}

// Generated where

var ClickbaitKeywordWhere = struct {
	ID      whereHelperstring
	LabelID whereHelperstring
	Keyword whereHelperstring
}{
	ID:      whereHelperstring{field: "\"clickbait_keywords\".\"id\""},
	LabelID: whereHelperstring{field: "\"clickbait_keywords\".\"label_id\""},
	Keyword: whereHelperstring{field: "\"clickbait_keywords\".\"keyword\""},
}

// ClickbaitKeywordRels is where relationship names are stored.
var ClickbaitKeywordRels = struct {
	Label string
}{
	Label: "Label",
}

// clickbaitKeywordR is where relationships are stored.
type clickbaitKeywordR struct {
	Label *Label
}

// NewStruct creates a new relationship struct
func (*clickbaitKeywordR) NewStruct() *clickbaitKeywordR {
	return &clickbaitKeywordR{}
}

// clickbaitKeywordL is where Load methods for each relationship are stored.
type clickbaitKeywordL struct{}

var (
	clickbaitKeywordAllColumns            = []string{"id", "label_id", "keyword"}
	clickbaitKeywordColumnsWithoutDefault = []string{"id", "label_id", "keyword"}
	clickbaitKeywordColumnsWithDefault    = []string{}
	clickbaitKeywordPrimaryKeyColumns     = []string{"id"}
)

type (
	// ClickbaitKeywordSlice is an alias for a slice of pointers to ClickbaitKeyword.
	// This should generally be used opposed to []ClickbaitKeyword.
	ClickbaitKeywordSlice []*ClickbaitKeyword
	// ClickbaitKeywordHook is the signature for custom ClickbaitKeyword hook methods
	ClickbaitKeywordHook func(context.Context, boil.ContextExecutor, *ClickbaitKeyword) error

	clickbaitKeywordQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	clickbaitKeywordType                 = reflect.TypeOf(&ClickbaitKeyword{})
	clickbaitKeywordMapping              = queries.MakeStructMapping(clickbaitKeywordType)
	clickbaitKeywordPrimaryKeyMapping, _ = queries.BindMapping(clickbaitKeywordType, clickbaitKeywordMapping, clickbaitKeywordPrimaryKeyColumns)
	clickbaitKeywordInsertCacheMut       sync.RWMutex
	clickbaitKeywordInsertCache          = make(map[string]insertCache)
	clickbaitKeywordUpdateCacheMut       sync.RWMutex
	clickbaitKeywordUpdateCache          = make(map[string]updateCache)
	clickbaitKeywordUpsertCacheMut       sync.RWMutex
	clickbaitKeywordUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var clickbaitKeywordBeforeInsertHooks []ClickbaitKeywordHook
var clickbaitKeywordBeforeUpdateHooks []ClickbaitKeywordHook
var clickbaitKeywordBeforeDeleteHooks []ClickbaitKeywordHook
var clickbaitKeywordBeforeUpsertHooks []ClickbaitKeywordHook

var clickbaitKeywordAfterInsertHooks []ClickbaitKeywordHook
var clickbaitKeywordAfterSelectHooks []ClickbaitKeywordHook
var clickbaitKeywordAfterUpdateHooks []ClickbaitKeywordHook
var clickbaitKeywordAfterDeleteHooks []ClickbaitKeywordHook
var clickbaitKeywordAfterUpsertHooks []ClickbaitKeywordHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ClickbaitKeyword) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clickbaitKeywordBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ClickbaitKeyword) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clickbaitKeywordBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ClickbaitKeyword) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clickbaitKeywordBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ClickbaitKeyword) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clickbaitKeywordBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ClickbaitKeyword) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clickbaitKeywordAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ClickbaitKeyword) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clickbaitKeywordAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ClickbaitKeyword) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clickbaitKeywordAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ClickbaitKeyword) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clickbaitKeywordAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ClickbaitKeyword) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clickbaitKeywordAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddClickbaitKeywordHook registers your hook function for all future operations.
func AddClickbaitKeywordHook(hookPoint boil.HookPoint, clickbaitKeywordHook ClickbaitKeywordHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		clickbaitKeywordBeforeInsertHooks = append(clickbaitKeywordBeforeInsertHooks, clickbaitKeywordHook)
	case boil.BeforeUpdateHook:
		clickbaitKeywordBeforeUpdateHooks = append(clickbaitKeywordBeforeUpdateHooks, clickbaitKeywordHook)
	case boil.BeforeDeleteHook:
		clickbaitKeywordBeforeDeleteHooks = append(clickbaitKeywordBeforeDeleteHooks, clickbaitKeywordHook)
	case boil.BeforeUpsertHook:
		clickbaitKeywordBeforeUpsertHooks = append(clickbaitKeywordBeforeUpsertHooks, clickbaitKeywordHook)
	case boil.AfterInsertHook:
		clickbaitKeywordAfterInsertHooks = append(clickbaitKeywordAfterInsertHooks, clickbaitKeywordHook)
	case boil.AfterSelectHook:
		clickbaitKeywordAfterSelectHooks = append(clickbaitKeywordAfterSelectHooks, clickbaitKeywordHook)
	case boil.AfterUpdateHook:
		clickbaitKeywordAfterUpdateHooks = append(clickbaitKeywordAfterUpdateHooks, clickbaitKeywordHook)
	case boil.AfterDeleteHook:
		clickbaitKeywordAfterDeleteHooks = append(clickbaitKeywordAfterDeleteHooks, clickbaitKeywordHook)
	case boil.AfterUpsertHook:
		clickbaitKeywordAfterUpsertHooks = append(clickbaitKeywordAfterUpsertHooks, clickbaitKeywordHook)
	}
}

// One returns a single clickbaitKeyword record from the query.
func (q clickbaitKeywordQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ClickbaitKeyword, error) {
	o := &ClickbaitKeyword{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for clickbait_keywords")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ClickbaitKeyword records from the query.
func (q clickbaitKeywordQuery) All(ctx context.Context, exec boil.ContextExecutor) (ClickbaitKeywordSlice, error) {
	var o []*ClickbaitKeyword

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ClickbaitKeyword slice")
	}

	if len(clickbaitKeywordAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ClickbaitKeyword records in the query.
func (q clickbaitKeywordQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count clickbait_keywords rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q clickbaitKeywordQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if clickbait_keywords exists")
	}

	return count > 0, nil
}

// Label pointed to by the foreign key.
func (o *ClickbaitKeyword) Label(mods ...qm.QueryMod) labelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.LabelID),
	}

	queryMods = append(queryMods, mods...)

	query := Labels(queryMods...)
	queries.SetFrom(query.Query, "\"labels\"")

	return query
}

// LoadLabel allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (clickbaitKeywordL) LoadLabel(ctx context.Context, e boil.ContextExecutor, singular bool, maybeClickbaitKeyword interface{}, mods queries.Applicator) error {
	var slice []*ClickbaitKeyword
	var object *ClickbaitKeyword

	if singular {
		object = maybeClickbaitKeyword.(*ClickbaitKeyword)
	} else {
		slice = *maybeClickbaitKeyword.(*[]*ClickbaitKeyword)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &clickbaitKeywordR{}
		}
		args = append(args, object.LabelID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &clickbaitKeywordR{}
			}

			for _, a := range args {
				if a == obj.LabelID {
					continue Outer
				}
			}

			args = append(args, obj.LabelID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`labels`), qm.WhereIn(`labels.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Label")
	}

	var resultSlice []*Label
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Label")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for labels")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for labels")
	}

	if len(clickbaitKeywordAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Label = foreign
		if foreign.R == nil {
			foreign.R = &labelR{}
		}
		foreign.R.ClickbaitKeywords = append(foreign.R.ClickbaitKeywords, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.LabelID == foreign.ID {
				local.R.Label = foreign
				if foreign.R == nil {
					foreign.R = &labelR{}
				}
				foreign.R.ClickbaitKeywords = append(foreign.R.ClickbaitKeywords, local)
				break
			}
		}
	}

	return nil
}

// SetLabel of the clickbaitKeyword to the related item.
// Sets o.R.Label to related.
// Adds o to related.R.ClickbaitKeywords.
func (o *ClickbaitKeyword) SetLabel(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Label) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"clickbait_keywords\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"label_id"}),
		strmangle.WhereClause("\"", "\"", 2, clickbaitKeywordPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.LabelID = related.ID
	if o.R == nil {
		o.R = &clickbaitKeywordR{
			Label: related,
		}
	} else {
		o.R.Label = related
	}

	if related.R == nil {
		related.R = &labelR{
			ClickbaitKeywords: ClickbaitKeywordSlice{o},
		}
	} else {
		related.R.ClickbaitKeywords = append(related.R.ClickbaitKeywords, o)
	}

	return nil
}

// ClickbaitKeywords retrieves all the records using an executor.
func ClickbaitKeywords(mods ...qm.QueryMod) clickbaitKeywordQuery {
	mods = append(mods, qm.From("\"clickbait_keywords\""))
	return clickbaitKeywordQuery{NewQuery(mods...)}
}

// FindClickbaitKeyword retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindClickbaitKeyword(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*ClickbaitKeyword, error) {
	clickbaitKeywordObj := &ClickbaitKeyword{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"clickbait_keywords\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, clickbaitKeywordObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from clickbait_keywords")
	}

	return clickbaitKeywordObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ClickbaitKeyword) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no clickbait_keywords provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(clickbaitKeywordColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	clickbaitKeywordInsertCacheMut.RLock()
	cache, cached := clickbaitKeywordInsertCache[key]
	clickbaitKeywordInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			clickbaitKeywordAllColumns,
			clickbaitKeywordColumnsWithDefault,
			clickbaitKeywordColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(clickbaitKeywordType, clickbaitKeywordMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(clickbaitKeywordType, clickbaitKeywordMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"clickbait_keywords\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"clickbait_keywords\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into clickbait_keywords")
	}

	if !cached {
		clickbaitKeywordInsertCacheMut.Lock()
		clickbaitKeywordInsertCache[key] = cache
		clickbaitKeywordInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ClickbaitKeyword.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ClickbaitKeyword) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	clickbaitKeywordUpdateCacheMut.RLock()
	cache, cached := clickbaitKeywordUpdateCache[key]
	clickbaitKeywordUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			clickbaitKeywordAllColumns,
			clickbaitKeywordPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update clickbait_keywords, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"clickbait_keywords\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, clickbaitKeywordPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(clickbaitKeywordType, clickbaitKeywordMapping, append(wl, clickbaitKeywordPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update clickbait_keywords row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for clickbait_keywords")
	}

	if !cached {
		clickbaitKeywordUpdateCacheMut.Lock()
		clickbaitKeywordUpdateCache[key] = cache
		clickbaitKeywordUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q clickbaitKeywordQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for clickbait_keywords")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for clickbait_keywords")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ClickbaitKeywordSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clickbaitKeywordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"clickbait_keywords\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, clickbaitKeywordPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in clickbaitKeyword slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all clickbaitKeyword")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ClickbaitKeyword) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no clickbait_keywords provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(clickbaitKeywordColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	clickbaitKeywordUpsertCacheMut.RLock()
	cache, cached := clickbaitKeywordUpsertCache[key]
	clickbaitKeywordUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			clickbaitKeywordAllColumns,
			clickbaitKeywordColumnsWithDefault,
			clickbaitKeywordColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			clickbaitKeywordAllColumns,
			clickbaitKeywordPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert clickbait_keywords, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(clickbaitKeywordPrimaryKeyColumns))
			copy(conflict, clickbaitKeywordPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"clickbait_keywords\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(clickbaitKeywordType, clickbaitKeywordMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(clickbaitKeywordType, clickbaitKeywordMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert clickbait_keywords")
	}

	if !cached {
		clickbaitKeywordUpsertCacheMut.Lock()
		clickbaitKeywordUpsertCache[key] = cache
		clickbaitKeywordUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ClickbaitKeyword record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ClickbaitKeyword) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ClickbaitKeyword provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), clickbaitKeywordPrimaryKeyMapping)
	sql := "DELETE FROM \"clickbait_keywords\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from clickbait_keywords")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for clickbait_keywords")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q clickbaitKeywordQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no clickbaitKeywordQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from clickbait_keywords")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for clickbait_keywords")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ClickbaitKeywordSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(clickbaitKeywordBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clickbaitKeywordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"clickbait_keywords\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, clickbaitKeywordPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from clickbaitKeyword slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for clickbait_keywords")
	}

	if len(clickbaitKeywordAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ClickbaitKeyword) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindClickbaitKeyword(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ClickbaitKeywordSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ClickbaitKeywordSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clickbaitKeywordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"clickbait_keywords\".* FROM \"clickbait_keywords\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, clickbaitKeywordPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ClickbaitKeywordSlice")
	}

	*o = slice

	return nil
}

// ClickbaitKeywordExists checks if the ClickbaitKeyword row exists.
func ClickbaitKeywordExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"clickbait_keywords\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if clickbait_keywords exists")
	}

	return exists, nil
}
