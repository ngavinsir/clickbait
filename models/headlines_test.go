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

func testHeadlines(t *testing.T) {
	t.Parallel()

	query := Headlines()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testHeadlinesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Headline{}
	if err = randomize.Struct(seed, o, headlineDBTypes, true, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
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

	count, err := Headlines().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHeadlinesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Headline{}
	if err = randomize.Struct(seed, o, headlineDBTypes, true, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Headlines().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Headlines().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHeadlinesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Headline{}
	if err = randomize.Struct(seed, o, headlineDBTypes, true, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HeadlineSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Headlines().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHeadlinesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Headline{}
	if err = randomize.Struct(seed, o, headlineDBTypes, true, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := HeadlineExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Headline exists: %s", err)
	}
	if !e {
		t.Errorf("Expected HeadlineExists to return true, but got false.")
	}
}

func testHeadlinesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Headline{}
	if err = randomize.Struct(seed, o, headlineDBTypes, true, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	headlineFound, err := FindHeadline(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if headlineFound == nil {
		t.Error("want a record, got nil")
	}
}

func testHeadlinesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Headline{}
	if err = randomize.Struct(seed, o, headlineDBTypes, true, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Headlines().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testHeadlinesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Headline{}
	if err = randomize.Struct(seed, o, headlineDBTypes, true, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Headlines().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testHeadlinesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	headlineOne := &Headline{}
	headlineTwo := &Headline{}
	if err = randomize.Struct(seed, headlineOne, headlineDBTypes, false, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}
	if err = randomize.Struct(seed, headlineTwo, headlineDBTypes, false, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = headlineOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = headlineTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Headlines().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testHeadlinesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	headlineOne := &Headline{}
	headlineTwo := &Headline{}
	if err = randomize.Struct(seed, headlineOne, headlineDBTypes, false, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}
	if err = randomize.Struct(seed, headlineTwo, headlineDBTypes, false, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = headlineOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = headlineTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Headlines().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func headlineBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Headline) error {
	*o = Headline{}
	return nil
}

func headlineAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Headline) error {
	*o = Headline{}
	return nil
}

func headlineAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Headline) error {
	*o = Headline{}
	return nil
}

func headlineBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Headline) error {
	*o = Headline{}
	return nil
}

func headlineAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Headline) error {
	*o = Headline{}
	return nil
}

func headlineBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Headline) error {
	*o = Headline{}
	return nil
}

func headlineAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Headline) error {
	*o = Headline{}
	return nil
}

func headlineBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Headline) error {
	*o = Headline{}
	return nil
}

func headlineAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Headline) error {
	*o = Headline{}
	return nil
}

func testHeadlinesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Headline{}
	o := &Headline{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, headlineDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Headline object: %s", err)
	}

	AddHeadlineHook(boil.BeforeInsertHook, headlineBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	headlineBeforeInsertHooks = []HeadlineHook{}

	AddHeadlineHook(boil.AfterInsertHook, headlineAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	headlineAfterInsertHooks = []HeadlineHook{}

	AddHeadlineHook(boil.AfterSelectHook, headlineAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	headlineAfterSelectHooks = []HeadlineHook{}

	AddHeadlineHook(boil.BeforeUpdateHook, headlineBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	headlineBeforeUpdateHooks = []HeadlineHook{}

	AddHeadlineHook(boil.AfterUpdateHook, headlineAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	headlineAfterUpdateHooks = []HeadlineHook{}

	AddHeadlineHook(boil.BeforeDeleteHook, headlineBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	headlineBeforeDeleteHooks = []HeadlineHook{}

	AddHeadlineHook(boil.AfterDeleteHook, headlineAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	headlineAfterDeleteHooks = []HeadlineHook{}

	AddHeadlineHook(boil.BeforeUpsertHook, headlineBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	headlineBeforeUpsertHooks = []HeadlineHook{}

	AddHeadlineHook(boil.AfterUpsertHook, headlineAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	headlineAfterUpsertHooks = []HeadlineHook{}
}

func testHeadlinesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Headline{}
	if err = randomize.Struct(seed, o, headlineDBTypes, true, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Headlines().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHeadlinesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Headline{}
	if err = randomize.Struct(seed, o, headlineDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(headlineColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Headlines().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHeadlineToManyLabels(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Headline
	var b, c Label

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, headlineDBTypes, true, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, labelDBTypes, false, labelColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, labelDBTypes, false, labelColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.HeadlineID = a.ID
	c.HeadlineID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Labels().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.HeadlineID == b.HeadlineID {
			bFound = true
		}
		if v.HeadlineID == c.HeadlineID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := HeadlineSlice{&a}
	if err = a.L.LoadLabels(ctx, tx, false, (*[]*Headline)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Labels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Labels = nil
	if err = a.L.LoadLabels(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Labels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testHeadlineToManyAddOpLabels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Headline
	var b, c, d, e Label

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, headlineDBTypes, false, strmangle.SetComplement(headlinePrimaryKeyColumns, headlineColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Label{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, labelDBTypes, false, strmangle.SetComplement(labelPrimaryKeyColumns, labelColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Label{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddLabels(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.HeadlineID {
			t.Error("foreign key was wrong value", a.ID, first.HeadlineID)
		}
		if a.ID != second.HeadlineID {
			t.Error("foreign key was wrong value", a.ID, second.HeadlineID)
		}

		if first.R.Headline != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Headline != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Labels[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Labels[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Labels().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testHeadlinesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Headline{}
	if err = randomize.Struct(seed, o, headlineDBTypes, true, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
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

func testHeadlinesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Headline{}
	if err = randomize.Struct(seed, o, headlineDBTypes, true, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HeadlineSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testHeadlinesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Headline{}
	if err = randomize.Struct(seed, o, headlineDBTypes, true, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Headlines().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	headlineDBTypes = map[string]string{`ID`: `character varying`, `Value`: `text`}
	_               = bytes.MinRead
)

func testHeadlinesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(headlinePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(headlineAllColumns) == len(headlinePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Headline{}
	if err = randomize.Struct(seed, o, headlineDBTypes, true, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Headlines().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, headlineDBTypes, true, headlinePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testHeadlinesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(headlineAllColumns) == len(headlinePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Headline{}
	if err = randomize.Struct(seed, o, headlineDBTypes, true, headlineColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Headlines().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, headlineDBTypes, true, headlinePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(headlineAllColumns, headlinePrimaryKeyColumns) {
		fields = headlineAllColumns
	} else {
		fields = strmangle.SetComplement(
			headlineAllColumns,
			headlinePrimaryKeyColumns,
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

	slice := HeadlineSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testHeadlinesUpsert(t *testing.T) {
	t.Parallel()

	if len(headlineAllColumns) == len(headlinePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Headline{}
	if err = randomize.Struct(seed, &o, headlineDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Headline: %s", err)
	}

	count, err := Headlines().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, headlineDBTypes, false, headlinePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Headline struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Headline: %s", err)
	}

	count, err = Headlines().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
