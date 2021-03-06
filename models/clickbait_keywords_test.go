// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testClickbaitKeywords(t *testing.T) {
	t.Parallel()

	query := ClickbaitKeywords()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testClickbaitKeywordsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClickbaitKeyword{}
	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, true, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClickbaitKeywords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClickbaitKeywordsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClickbaitKeyword{}
	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, true, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ClickbaitKeywords().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClickbaitKeywords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClickbaitKeywordsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClickbaitKeyword{}
	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, true, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClickbaitKeywordSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClickbaitKeywords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClickbaitKeywordsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClickbaitKeyword{}
	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, true, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ClickbaitKeywordExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ClickbaitKeyword exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ClickbaitKeywordExists to return true, but got false.")
	}
}

func testClickbaitKeywordsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClickbaitKeyword{}
	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, true, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	clickbaitKeywordFound, err := FindClickbaitKeyword(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if clickbaitKeywordFound == nil {
		t.Error("want a record, got nil")
	}
}

func testClickbaitKeywordsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClickbaitKeyword{}
	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, true, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ClickbaitKeywords().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testClickbaitKeywordsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClickbaitKeyword{}
	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, true, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ClickbaitKeywords().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testClickbaitKeywordsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	clickbaitKeywordOne := &ClickbaitKeyword{}
	clickbaitKeywordTwo := &ClickbaitKeyword{}
	if err = randomize.Struct(seed, clickbaitKeywordOne, clickbaitKeywordDBTypes, false, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}
	if err = randomize.Struct(seed, clickbaitKeywordTwo, clickbaitKeywordDBTypes, false, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = clickbaitKeywordOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = clickbaitKeywordTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClickbaitKeywords().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testClickbaitKeywordsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	clickbaitKeywordOne := &ClickbaitKeyword{}
	clickbaitKeywordTwo := &ClickbaitKeyword{}
	if err = randomize.Struct(seed, clickbaitKeywordOne, clickbaitKeywordDBTypes, false, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}
	if err = randomize.Struct(seed, clickbaitKeywordTwo, clickbaitKeywordDBTypes, false, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = clickbaitKeywordOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = clickbaitKeywordTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClickbaitKeywords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func clickbaitKeywordBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ClickbaitKeyword) error {
	*o = ClickbaitKeyword{}
	return nil
}

func clickbaitKeywordAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ClickbaitKeyword) error {
	*o = ClickbaitKeyword{}
	return nil
}

func clickbaitKeywordAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ClickbaitKeyword) error {
	*o = ClickbaitKeyword{}
	return nil
}

func clickbaitKeywordBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ClickbaitKeyword) error {
	*o = ClickbaitKeyword{}
	return nil
}

func clickbaitKeywordAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ClickbaitKeyword) error {
	*o = ClickbaitKeyword{}
	return nil
}

func clickbaitKeywordBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ClickbaitKeyword) error {
	*o = ClickbaitKeyword{}
	return nil
}

func clickbaitKeywordAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ClickbaitKeyword) error {
	*o = ClickbaitKeyword{}
	return nil
}

func clickbaitKeywordBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ClickbaitKeyword) error {
	*o = ClickbaitKeyword{}
	return nil
}

func clickbaitKeywordAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ClickbaitKeyword) error {
	*o = ClickbaitKeyword{}
	return nil
}

func testClickbaitKeywordsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ClickbaitKeyword{}
	o := &ClickbaitKeyword{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword object: %s", err)
	}

	AddClickbaitKeywordHook(boil.BeforeInsertHook, clickbaitKeywordBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	clickbaitKeywordBeforeInsertHooks = []ClickbaitKeywordHook{}

	AddClickbaitKeywordHook(boil.AfterInsertHook, clickbaitKeywordAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	clickbaitKeywordAfterInsertHooks = []ClickbaitKeywordHook{}

	AddClickbaitKeywordHook(boil.AfterSelectHook, clickbaitKeywordAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	clickbaitKeywordAfterSelectHooks = []ClickbaitKeywordHook{}

	AddClickbaitKeywordHook(boil.BeforeUpdateHook, clickbaitKeywordBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	clickbaitKeywordBeforeUpdateHooks = []ClickbaitKeywordHook{}

	AddClickbaitKeywordHook(boil.AfterUpdateHook, clickbaitKeywordAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	clickbaitKeywordAfterUpdateHooks = []ClickbaitKeywordHook{}

	AddClickbaitKeywordHook(boil.BeforeDeleteHook, clickbaitKeywordBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	clickbaitKeywordBeforeDeleteHooks = []ClickbaitKeywordHook{}

	AddClickbaitKeywordHook(boil.AfterDeleteHook, clickbaitKeywordAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	clickbaitKeywordAfterDeleteHooks = []ClickbaitKeywordHook{}

	AddClickbaitKeywordHook(boil.BeforeUpsertHook, clickbaitKeywordBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	clickbaitKeywordBeforeUpsertHooks = []ClickbaitKeywordHook{}

	AddClickbaitKeywordHook(boil.AfterUpsertHook, clickbaitKeywordAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	clickbaitKeywordAfterUpsertHooks = []ClickbaitKeywordHook{}
}

func testClickbaitKeywordsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClickbaitKeyword{}
	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, true, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClickbaitKeywords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClickbaitKeywordsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClickbaitKeyword{}
	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(clickbaitKeywordColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ClickbaitKeywords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClickbaitKeywordToOneLabelUsingLabel(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ClickbaitKeyword
	var foreign Label

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, clickbaitKeywordDBTypes, false, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, labelDBTypes, false, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.LabelID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Label().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ClickbaitKeywordSlice{&local}
	if err = local.L.LoadLabel(ctx, tx, false, (*[]*ClickbaitKeyword)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Label == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Label = nil
	if err = local.L.LoadLabel(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Label == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testClickbaitKeywordToOneSetOpLabelUsingLabel(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ClickbaitKeyword
	var b, c Label

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, clickbaitKeywordDBTypes, false, strmangle.SetComplement(clickbaitKeywordPrimaryKeyColumns, clickbaitKeywordColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, labelDBTypes, false, strmangle.SetComplement(labelPrimaryKeyColumns, labelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, labelDBTypes, false, strmangle.SetComplement(labelPrimaryKeyColumns, labelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Label{&b, &c} {
		err = a.SetLabel(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Label != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ClickbaitKeywords[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.LabelID != x.ID {
			t.Error("foreign key was wrong value", a.LabelID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.LabelID))
		reflect.Indirect(reflect.ValueOf(&a.LabelID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.LabelID != x.ID {
			t.Error("foreign key was wrong value", a.LabelID, x.ID)
		}
	}
}

func testClickbaitKeywordsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClickbaitKeyword{}
	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, true, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testClickbaitKeywordsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClickbaitKeyword{}
	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, true, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClickbaitKeywordSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testClickbaitKeywordsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClickbaitKeyword{}
	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, true, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClickbaitKeywords().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	clickbaitKeywordDBTypes = map[string]string{`ID`: `text`, `LabelID`: `text`, `Keyword`: `text`}
	_                       = bytes.MinRead
)

func testClickbaitKeywordsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(clickbaitKeywordPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(clickbaitKeywordAllColumns) == len(clickbaitKeywordPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClickbaitKeyword{}
	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, true, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClickbaitKeywords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, true, clickbaitKeywordPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testClickbaitKeywordsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(clickbaitKeywordAllColumns) == len(clickbaitKeywordPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClickbaitKeyword{}
	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, true, clickbaitKeywordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClickbaitKeywords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, clickbaitKeywordDBTypes, true, clickbaitKeywordPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(clickbaitKeywordAllColumns, clickbaitKeywordPrimaryKeyColumns) {
		fields = clickbaitKeywordAllColumns
	} else {
		fields = strmangle.SetComplement(
			clickbaitKeywordAllColumns,
			clickbaitKeywordPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ClickbaitKeywordSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testClickbaitKeywordsUpsert(t *testing.T) {
	t.Parallel()

	if len(clickbaitKeywordAllColumns) == len(clickbaitKeywordPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ClickbaitKeyword{}
	if err = randomize.Struct(seed, &o, clickbaitKeywordDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClickbaitKeyword: %s", err)
	}

	count, err := ClickbaitKeywords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, clickbaitKeywordDBTypes, false, clickbaitKeywordPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClickbaitKeyword struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClickbaitKeyword: %s", err)
	}

	count, err = ClickbaitKeywords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
