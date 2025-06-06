// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// TournamentsUser is an object representing the database table.
type TournamentsUser struct {
	ID           int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID       int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	TournamentID int64     `boil:"tournament_id" json:"tournament_id" toml:"tournament_id" yaml:"tournament_id"`
	TeamID       int64     `boil:"team_id" json:"team_id" toml:"team_id" yaml:"team_id"`
	JoinedAt     time.Time `boil:"joined_at" json:"joined_at" toml:"joined_at" yaml:"joined_at"`

	R *tournamentsUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tournamentsUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TournamentsUserColumns = struct {
	ID           string
	UserID       string
	TournamentID string
	TeamID       string
	JoinedAt     string
}{
	ID:           "id",
	UserID:       "user_id",
	TournamentID: "tournament_id",
	TeamID:       "team_id",
	JoinedAt:     "joined_at",
}

var TournamentsUserTableColumns = struct {
	ID           string
	UserID       string
	TournamentID string
	TeamID       string
	JoinedAt     string
}{
	ID:           "tournaments_users.id",
	UserID:       "tournaments_users.user_id",
	TournamentID: "tournaments_users.tournament_id",
	TeamID:       "tournaments_users.team_id",
	JoinedAt:     "tournaments_users.joined_at",
}

// Generated where

var TournamentsUserWhere = struct {
	ID           whereHelperint64
	UserID       whereHelperint64
	TournamentID whereHelperint64
	TeamID       whereHelperint64
	JoinedAt     whereHelpertime_Time
}{
	ID:           whereHelperint64{field: "\"tournaments_users\".\"id\""},
	UserID:       whereHelperint64{field: "\"tournaments_users\".\"user_id\""},
	TournamentID: whereHelperint64{field: "\"tournaments_users\".\"tournament_id\""},
	TeamID:       whereHelperint64{field: "\"tournaments_users\".\"team_id\""},
	JoinedAt:     whereHelpertime_Time{field: "\"tournaments_users\".\"joined_at\""},
}

// TournamentsUserRels is where relationship names are stored.
var TournamentsUserRels = struct {
	Tournament string
	User       string
}{
	Tournament: "Tournament",
	User:       "User",
}

// tournamentsUserR is where relationships are stored.
type tournamentsUserR struct {
	Tournament *Tournament `boil:"Tournament" json:"Tournament" toml:"Tournament" yaml:"Tournament"`
	User       *User       `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*tournamentsUserR) NewStruct() *tournamentsUserR {
	return &tournamentsUserR{}
}

func (r *tournamentsUserR) GetTournament() *Tournament {
	if r == nil {
		return nil
	}
	return r.Tournament
}

func (r *tournamentsUserR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// tournamentsUserL is where Load methods for each relationship are stored.
type tournamentsUserL struct{}

var (
	tournamentsUserAllColumns            = []string{"id", "user_id", "tournament_id", "team_id", "joined_at"}
	tournamentsUserColumnsWithoutDefault = []string{"user_id", "tournament_id", "team_id"}
	tournamentsUserColumnsWithDefault    = []string{"id", "joined_at"}
	tournamentsUserPrimaryKeyColumns     = []string{"id"}
	tournamentsUserGeneratedColumns      = []string{}
)

type (
	// TournamentsUserSlice is an alias for a slice of pointers to TournamentsUser.
	// This should almost always be used instead of []TournamentsUser.
	TournamentsUserSlice []*TournamentsUser
	// TournamentsUserHook is the signature for custom TournamentsUser hook methods
	TournamentsUserHook func(context.Context, boil.ContextExecutor, *TournamentsUser) error

	tournamentsUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tournamentsUserType                 = reflect.TypeOf(&TournamentsUser{})
	tournamentsUserMapping              = queries.MakeStructMapping(tournamentsUserType)
	tournamentsUserPrimaryKeyMapping, _ = queries.BindMapping(tournamentsUserType, tournamentsUserMapping, tournamentsUserPrimaryKeyColumns)
	tournamentsUserInsertCacheMut       sync.RWMutex
	tournamentsUserInsertCache          = make(map[string]insertCache)
	tournamentsUserUpdateCacheMut       sync.RWMutex
	tournamentsUserUpdateCache          = make(map[string]updateCache)
	tournamentsUserUpsertCacheMut       sync.RWMutex
	tournamentsUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tournamentsUserAfterSelectMu sync.Mutex
var tournamentsUserAfterSelectHooks []TournamentsUserHook

var tournamentsUserBeforeInsertMu sync.Mutex
var tournamentsUserBeforeInsertHooks []TournamentsUserHook
var tournamentsUserAfterInsertMu sync.Mutex
var tournamentsUserAfterInsertHooks []TournamentsUserHook

var tournamentsUserBeforeUpdateMu sync.Mutex
var tournamentsUserBeforeUpdateHooks []TournamentsUserHook
var tournamentsUserAfterUpdateMu sync.Mutex
var tournamentsUserAfterUpdateHooks []TournamentsUserHook

var tournamentsUserBeforeDeleteMu sync.Mutex
var tournamentsUserBeforeDeleteHooks []TournamentsUserHook
var tournamentsUserAfterDeleteMu sync.Mutex
var tournamentsUserAfterDeleteHooks []TournamentsUserHook

var tournamentsUserBeforeUpsertMu sync.Mutex
var tournamentsUserBeforeUpsertHooks []TournamentsUserHook
var tournamentsUserAfterUpsertMu sync.Mutex
var tournamentsUserAfterUpsertHooks []TournamentsUserHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TournamentsUser) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tournamentsUserAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TournamentsUser) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tournamentsUserBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TournamentsUser) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tournamentsUserAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TournamentsUser) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tournamentsUserBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TournamentsUser) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tournamentsUserAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TournamentsUser) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tournamentsUserBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TournamentsUser) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tournamentsUserAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TournamentsUser) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tournamentsUserBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TournamentsUser) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tournamentsUserAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTournamentsUserHook registers your hook function for all future operations.
func AddTournamentsUserHook(hookPoint boil.HookPoint, tournamentsUserHook TournamentsUserHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		tournamentsUserAfterSelectMu.Lock()
		tournamentsUserAfterSelectHooks = append(tournamentsUserAfterSelectHooks, tournamentsUserHook)
		tournamentsUserAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		tournamentsUserBeforeInsertMu.Lock()
		tournamentsUserBeforeInsertHooks = append(tournamentsUserBeforeInsertHooks, tournamentsUserHook)
		tournamentsUserBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		tournamentsUserAfterInsertMu.Lock()
		tournamentsUserAfterInsertHooks = append(tournamentsUserAfterInsertHooks, tournamentsUserHook)
		tournamentsUserAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		tournamentsUserBeforeUpdateMu.Lock()
		tournamentsUserBeforeUpdateHooks = append(tournamentsUserBeforeUpdateHooks, tournamentsUserHook)
		tournamentsUserBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		tournamentsUserAfterUpdateMu.Lock()
		tournamentsUserAfterUpdateHooks = append(tournamentsUserAfterUpdateHooks, tournamentsUserHook)
		tournamentsUserAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		tournamentsUserBeforeDeleteMu.Lock()
		tournamentsUserBeforeDeleteHooks = append(tournamentsUserBeforeDeleteHooks, tournamentsUserHook)
		tournamentsUserBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		tournamentsUserAfterDeleteMu.Lock()
		tournamentsUserAfterDeleteHooks = append(tournamentsUserAfterDeleteHooks, tournamentsUserHook)
		tournamentsUserAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		tournamentsUserBeforeUpsertMu.Lock()
		tournamentsUserBeforeUpsertHooks = append(tournamentsUserBeforeUpsertHooks, tournamentsUserHook)
		tournamentsUserBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		tournamentsUserAfterUpsertMu.Lock()
		tournamentsUserAfterUpsertHooks = append(tournamentsUserAfterUpsertHooks, tournamentsUserHook)
		tournamentsUserAfterUpsertMu.Unlock()
	}
}

// One returns a single tournamentsUser record from the query.
func (q tournamentsUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TournamentsUser, error) {
	o := &TournamentsUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for tournaments_users")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TournamentsUser records from the query.
func (q tournamentsUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (TournamentsUserSlice, error) {
	var o []*TournamentsUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TournamentsUser slice")
	}

	if len(tournamentsUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TournamentsUser records in the query.
func (q tournamentsUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count tournaments_users rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tournamentsUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if tournaments_users exists")
	}

	return count > 0, nil
}

// Tournament pointed to by the foreign key.
func (o *TournamentsUser) Tournament(mods ...qm.QueryMod) tournamentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TournamentID),
	}

	queryMods = append(queryMods, mods...)

	return Tournaments(queryMods...)
}

// User pointed to by the foreign key.
func (o *TournamentsUser) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadTournament allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (tournamentsUserL) LoadTournament(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTournamentsUser interface{}, mods queries.Applicator) error {
	var slice []*TournamentsUser
	var object *TournamentsUser

	if singular {
		var ok bool
		object, ok = maybeTournamentsUser.(*TournamentsUser)
		if !ok {
			object = new(TournamentsUser)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTournamentsUser)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTournamentsUser))
			}
		}
	} else {
		s, ok := maybeTournamentsUser.(*[]*TournamentsUser)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTournamentsUser)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTournamentsUser))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &tournamentsUserR{}
		}
		args[object.TournamentID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tournamentsUserR{}
			}

			args[obj.TournamentID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`tournaments`),
		qm.WhereIn(`tournaments.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Tournament")
	}

	var resultSlice []*Tournament
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Tournament")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for tournaments")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for tournaments")
	}

	if len(tournamentAfterSelectHooks) != 0 {
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
		object.R.Tournament = foreign
		if foreign.R == nil {
			foreign.R = &tournamentR{}
		}
		foreign.R.TournamentsUsers = append(foreign.R.TournamentsUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TournamentID == foreign.ID {
				local.R.Tournament = foreign
				if foreign.R == nil {
					foreign.R = &tournamentR{}
				}
				foreign.R.TournamentsUsers = append(foreign.R.TournamentsUsers, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (tournamentsUserL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTournamentsUser interface{}, mods queries.Applicator) error {
	var slice []*TournamentsUser
	var object *TournamentsUser

	if singular {
		var ok bool
		object, ok = maybeTournamentsUser.(*TournamentsUser)
		if !ok {
			object = new(TournamentsUser)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTournamentsUser)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTournamentsUser))
			}
		}
	} else {
		s, ok := maybeTournamentsUser.(*[]*TournamentsUser)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTournamentsUser)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTournamentsUser))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &tournamentsUserR{}
		}
		args[object.UserID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tournamentsUserR{}
			}

			args[obj.UserID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.TournamentsUsers = append(foreign.R.TournamentsUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.TournamentsUsers = append(foreign.R.TournamentsUsers, local)
				break
			}
		}
	}

	return nil
}

// SetTournament of the tournamentsUser to the related item.
// Sets o.R.Tournament to related.
// Adds o to related.R.TournamentsUsers.
func (o *TournamentsUser) SetTournament(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Tournament) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"tournaments_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"tournament_id"}),
		strmangle.WhereClause("\"", "\"", 2, tournamentsUserPrimaryKeyColumns),
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

	o.TournamentID = related.ID
	if o.R == nil {
		o.R = &tournamentsUserR{
			Tournament: related,
		}
	} else {
		o.R.Tournament = related
	}

	if related.R == nil {
		related.R = &tournamentR{
			TournamentsUsers: TournamentsUserSlice{o},
		}
	} else {
		related.R.TournamentsUsers = append(related.R.TournamentsUsers, o)
	}

	return nil
}

// SetUser of the tournamentsUser to the related item.
// Sets o.R.User to related.
// Adds o to related.R.TournamentsUsers.
func (o *TournamentsUser) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"tournaments_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, tournamentsUserPrimaryKeyColumns),
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

	o.UserID = related.ID
	if o.R == nil {
		o.R = &tournamentsUserR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			TournamentsUsers: TournamentsUserSlice{o},
		}
	} else {
		related.R.TournamentsUsers = append(related.R.TournamentsUsers, o)
	}

	return nil
}

// TournamentsUsers retrieves all the records using an executor.
func TournamentsUsers(mods ...qm.QueryMod) tournamentsUserQuery {
	mods = append(mods, qm.From("\"tournaments_users\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"tournaments_users\".*"})
	}

	return tournamentsUserQuery{q}
}

// FindTournamentsUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTournamentsUser(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*TournamentsUser, error) {
	tournamentsUserObj := &TournamentsUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"tournaments_users\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, tournamentsUserObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from tournaments_users")
	}

	if err = tournamentsUserObj.doAfterSelectHooks(ctx, exec); err != nil {
		return tournamentsUserObj, err
	}

	return tournamentsUserObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TournamentsUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tournaments_users provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tournamentsUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tournamentsUserInsertCacheMut.RLock()
	cache, cached := tournamentsUserInsertCache[key]
	tournamentsUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tournamentsUserAllColumns,
			tournamentsUserColumnsWithDefault,
			tournamentsUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tournamentsUserType, tournamentsUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tournamentsUserType, tournamentsUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"tournaments_users\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"tournaments_users\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into tournaments_users")
	}

	if !cached {
		tournamentsUserInsertCacheMut.Lock()
		tournamentsUserInsertCache[key] = cache
		tournamentsUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TournamentsUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TournamentsUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tournamentsUserUpdateCacheMut.RLock()
	cache, cached := tournamentsUserUpdateCache[key]
	tournamentsUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tournamentsUserAllColumns,
			tournamentsUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update tournaments_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"tournaments_users\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, tournamentsUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tournamentsUserType, tournamentsUserMapping, append(wl, tournamentsUserPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update tournaments_users row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for tournaments_users")
	}

	if !cached {
		tournamentsUserUpdateCacheMut.Lock()
		tournamentsUserUpdateCache[key] = cache
		tournamentsUserUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q tournamentsUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for tournaments_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for tournaments_users")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TournamentsUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tournamentsUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"tournaments_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, tournamentsUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in tournamentsUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all tournamentsUser")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TournamentsUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no tournaments_users provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tournamentsUserColumnsWithDefault, o)

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

	tournamentsUserUpsertCacheMut.RLock()
	cache, cached := tournamentsUserUpsertCache[key]
	tournamentsUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			tournamentsUserAllColumns,
			tournamentsUserColumnsWithDefault,
			tournamentsUserColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			tournamentsUserAllColumns,
			tournamentsUserPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert tournaments_users, could not build update column list")
		}

		ret := strmangle.SetComplement(tournamentsUserAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(tournamentsUserPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert tournaments_users, could not build conflict column list")
			}

			conflict = make([]string, len(tournamentsUserPrimaryKeyColumns))
			copy(conflict, tournamentsUserPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"tournaments_users\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(tournamentsUserType, tournamentsUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tournamentsUserType, tournamentsUserMapping, ret)
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
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert tournaments_users")
	}

	if !cached {
		tournamentsUserUpsertCacheMut.Lock()
		tournamentsUserUpsertCache[key] = cache
		tournamentsUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TournamentsUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TournamentsUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TournamentsUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tournamentsUserPrimaryKeyMapping)
	sql := "DELETE FROM \"tournaments_users\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from tournaments_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for tournaments_users")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tournamentsUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no tournamentsUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tournaments_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tournaments_users")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TournamentsUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(tournamentsUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tournamentsUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"tournaments_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tournamentsUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tournamentsUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tournaments_users")
	}

	if len(tournamentsUserAfterDeleteHooks) != 0 {
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
func (o *TournamentsUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTournamentsUser(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TournamentsUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TournamentsUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tournamentsUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"tournaments_users\".* FROM \"tournaments_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tournamentsUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TournamentsUserSlice")
	}

	*o = slice

	return nil
}

// TournamentsUserExists checks if the TournamentsUser row exists.
func TournamentsUserExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"tournaments_users\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if tournaments_users exists")
	}

	return exists, nil
}

// Exists checks if the TournamentsUser row exists.
func (o *TournamentsUser) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TournamentsUserExists(ctx, exec, o.ID)
}
