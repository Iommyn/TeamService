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

// GameMap is an object representing the database table.
type GameMap struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	GameID    int64     `boil:"game_id" json:"game_id" toml:"game_id" yaml:"game_id"`
	MapName   string    `boil:"map_name" json:"map_name" toml:"map_name" yaml:"map_name"`
	UpdateAt  time.Time `boil:"update_at" json:"update_at" toml:"update_at" yaml:"update_at"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *gameMapR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L gameMapL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GameMapColumns = struct {
	ID        string
	GameID    string
	MapName   string
	UpdateAt  string
	CreatedAt string
}{
	ID:        "id",
	GameID:    "game_id",
	MapName:   "map_name",
	UpdateAt:  "update_at",
	CreatedAt: "created_at",
}

var GameMapTableColumns = struct {
	ID        string
	GameID    string
	MapName   string
	UpdateAt  string
	CreatedAt string
}{
	ID:        "game_map.id",
	GameID:    "game_map.game_id",
	MapName:   "game_map.map_name",
	UpdateAt:  "game_map.update_at",
	CreatedAt: "game_map.created_at",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var GameMapWhere = struct {
	ID        whereHelperint64
	GameID    whereHelperint64
	MapName   whereHelperstring
	UpdateAt  whereHelpertime_Time
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "\"game_map\".\"id\""},
	GameID:    whereHelperint64{field: "\"game_map\".\"game_id\""},
	MapName:   whereHelperstring{field: "\"game_map\".\"map_name\""},
	UpdateAt:  whereHelpertime_Time{field: "\"game_map\".\"update_at\""},
	CreatedAt: whereHelpertime_Time{field: "\"game_map\".\"created_at\""},
}

// GameMapRels is where relationship names are stored.
var GameMapRels = struct {
	Game       string
	MapMatches string
}{
	Game:       "Game",
	MapMatches: "MapMatches",
}

// gameMapR is where relationships are stored.
type gameMapR struct {
	Game       *Game      `boil:"Game" json:"Game" toml:"Game" yaml:"Game"`
	MapMatches MatchSlice `boil:"MapMatches" json:"MapMatches" toml:"MapMatches" yaml:"MapMatches"`
}

// NewStruct creates a new relationship struct
func (*gameMapR) NewStruct() *gameMapR {
	return &gameMapR{}
}

func (r *gameMapR) GetGame() *Game {
	if r == nil {
		return nil
	}
	return r.Game
}

func (r *gameMapR) GetMapMatches() MatchSlice {
	if r == nil {
		return nil
	}
	return r.MapMatches
}

// gameMapL is where Load methods for each relationship are stored.
type gameMapL struct{}

var (
	gameMapAllColumns            = []string{"id", "game_id", "map_name", "update_at", "created_at"}
	gameMapColumnsWithoutDefault = []string{"game_id", "map_name"}
	gameMapColumnsWithDefault    = []string{"id", "update_at", "created_at"}
	gameMapPrimaryKeyColumns     = []string{"id"}
	gameMapGeneratedColumns      = []string{}
)

type (
	// GameMapSlice is an alias for a slice of pointers to GameMap.
	// This should almost always be used instead of []GameMap.
	GameMapSlice []*GameMap
	// GameMapHook is the signature for custom GameMap hook methods
	GameMapHook func(context.Context, boil.ContextExecutor, *GameMap) error

	gameMapQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	gameMapType                 = reflect.TypeOf(&GameMap{})
	gameMapMapping              = queries.MakeStructMapping(gameMapType)
	gameMapPrimaryKeyMapping, _ = queries.BindMapping(gameMapType, gameMapMapping, gameMapPrimaryKeyColumns)
	gameMapInsertCacheMut       sync.RWMutex
	gameMapInsertCache          = make(map[string]insertCache)
	gameMapUpdateCacheMut       sync.RWMutex
	gameMapUpdateCache          = make(map[string]updateCache)
	gameMapUpsertCacheMut       sync.RWMutex
	gameMapUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var gameMapAfterSelectMu sync.Mutex
var gameMapAfterSelectHooks []GameMapHook

var gameMapBeforeInsertMu sync.Mutex
var gameMapBeforeInsertHooks []GameMapHook
var gameMapAfterInsertMu sync.Mutex
var gameMapAfterInsertHooks []GameMapHook

var gameMapBeforeUpdateMu sync.Mutex
var gameMapBeforeUpdateHooks []GameMapHook
var gameMapAfterUpdateMu sync.Mutex
var gameMapAfterUpdateHooks []GameMapHook

var gameMapBeforeDeleteMu sync.Mutex
var gameMapBeforeDeleteHooks []GameMapHook
var gameMapAfterDeleteMu sync.Mutex
var gameMapAfterDeleteHooks []GameMapHook

var gameMapBeforeUpsertMu sync.Mutex
var gameMapBeforeUpsertHooks []GameMapHook
var gameMapAfterUpsertMu sync.Mutex
var gameMapAfterUpsertHooks []GameMapHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *GameMap) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMapAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *GameMap) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMapBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *GameMap) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMapAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *GameMap) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMapBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *GameMap) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMapAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *GameMap) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMapBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *GameMap) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMapAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *GameMap) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMapBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *GameMap) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMapAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGameMapHook registers your hook function for all future operations.
func AddGameMapHook(hookPoint boil.HookPoint, gameMapHook GameMapHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		gameMapAfterSelectMu.Lock()
		gameMapAfterSelectHooks = append(gameMapAfterSelectHooks, gameMapHook)
		gameMapAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		gameMapBeforeInsertMu.Lock()
		gameMapBeforeInsertHooks = append(gameMapBeforeInsertHooks, gameMapHook)
		gameMapBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		gameMapAfterInsertMu.Lock()
		gameMapAfterInsertHooks = append(gameMapAfterInsertHooks, gameMapHook)
		gameMapAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		gameMapBeforeUpdateMu.Lock()
		gameMapBeforeUpdateHooks = append(gameMapBeforeUpdateHooks, gameMapHook)
		gameMapBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		gameMapAfterUpdateMu.Lock()
		gameMapAfterUpdateHooks = append(gameMapAfterUpdateHooks, gameMapHook)
		gameMapAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		gameMapBeforeDeleteMu.Lock()
		gameMapBeforeDeleteHooks = append(gameMapBeforeDeleteHooks, gameMapHook)
		gameMapBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		gameMapAfterDeleteMu.Lock()
		gameMapAfterDeleteHooks = append(gameMapAfterDeleteHooks, gameMapHook)
		gameMapAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		gameMapBeforeUpsertMu.Lock()
		gameMapBeforeUpsertHooks = append(gameMapBeforeUpsertHooks, gameMapHook)
		gameMapBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		gameMapAfterUpsertMu.Lock()
		gameMapAfterUpsertHooks = append(gameMapAfterUpsertHooks, gameMapHook)
		gameMapAfterUpsertMu.Unlock()
	}
}

// One returns a single gameMap record from the query.
func (q gameMapQuery) One(ctx context.Context, exec boil.ContextExecutor) (*GameMap, error) {
	o := &GameMap{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for game_map")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all GameMap records from the query.
func (q gameMapQuery) All(ctx context.Context, exec boil.ContextExecutor) (GameMapSlice, error) {
	var o []*GameMap

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to GameMap slice")
	}

	if len(gameMapAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all GameMap records in the query.
func (q gameMapQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count game_map rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q gameMapQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if game_map exists")
	}

	return count > 0, nil
}

// Game pointed to by the foreign key.
func (o *GameMap) Game(mods ...qm.QueryMod) gameQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.GameID),
	}

	queryMods = append(queryMods, mods...)

	return Games(queryMods...)
}

// MapMatches retrieves all the match's Matches with an executor via map_id column.
func (o *GameMap) MapMatches(mods ...qm.QueryMod) matchQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"match\".\"map_id\"=?", o.ID),
	)

	return Matches(queryMods...)
}

// LoadGame allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (gameMapL) LoadGame(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGameMap interface{}, mods queries.Applicator) error {
	var slice []*GameMap
	var object *GameMap

	if singular {
		var ok bool
		object, ok = maybeGameMap.(*GameMap)
		if !ok {
			object = new(GameMap)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeGameMap)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeGameMap))
			}
		}
	} else {
		s, ok := maybeGameMap.(*[]*GameMap)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeGameMap)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeGameMap))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &gameMapR{}
		}
		args[object.GameID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &gameMapR{}
			}

			args[obj.GameID] = struct{}{}

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
		qm.From(`games`),
		qm.WhereIn(`games.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Game")
	}

	var resultSlice []*Game
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Game")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for games")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for games")
	}

	if len(gameAfterSelectHooks) != 0 {
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
		object.R.Game = foreign
		if foreign.R == nil {
			foreign.R = &gameR{}
		}
		foreign.R.GameMaps = append(foreign.R.GameMaps, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GameID == foreign.ID {
				local.R.Game = foreign
				if foreign.R == nil {
					foreign.R = &gameR{}
				}
				foreign.R.GameMaps = append(foreign.R.GameMaps, local)
				break
			}
		}
	}

	return nil
}

// LoadMapMatches allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (gameMapL) LoadMapMatches(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGameMap interface{}, mods queries.Applicator) error {
	var slice []*GameMap
	var object *GameMap

	if singular {
		var ok bool
		object, ok = maybeGameMap.(*GameMap)
		if !ok {
			object = new(GameMap)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeGameMap)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeGameMap))
			}
		}
	} else {
		s, ok := maybeGameMap.(*[]*GameMap)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeGameMap)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeGameMap))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &gameMapR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &gameMapR{}
			}
			args[obj.ID] = struct{}{}
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
		qm.From(`match`),
		qm.WhereIn(`match.map_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load match")
	}

	var resultSlice []*Match
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice match")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on match")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for match")
	}

	if len(matchAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.MapMatches = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &matchR{}
			}
			foreign.R.Map = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.MapID {
				local.R.MapMatches = append(local.R.MapMatches, foreign)
				if foreign.R == nil {
					foreign.R = &matchR{}
				}
				foreign.R.Map = local
				break
			}
		}
	}

	return nil
}

// SetGame of the gameMap to the related item.
// Sets o.R.Game to related.
// Adds o to related.R.GameMaps.
func (o *GameMap) SetGame(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Game) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"game_map\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"game_id"}),
		strmangle.WhereClause("\"", "\"", 2, gameMapPrimaryKeyColumns),
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

	o.GameID = related.ID
	if o.R == nil {
		o.R = &gameMapR{
			Game: related,
		}
	} else {
		o.R.Game = related
	}

	if related.R == nil {
		related.R = &gameR{
			GameMaps: GameMapSlice{o},
		}
	} else {
		related.R.GameMaps = append(related.R.GameMaps, o)
	}

	return nil
}

// AddMapMatches adds the given related objects to the existing relationships
// of the game_map, optionally inserting them as new records.
// Appends related to o.R.MapMatches.
// Sets related.R.Map appropriately.
func (o *GameMap) AddMapMatches(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Match) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.MapID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"match\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"map_id"}),
				strmangle.WhereClause("\"", "\"", 2, matchPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.MapID = o.ID
		}
	}

	if o.R == nil {
		o.R = &gameMapR{
			MapMatches: related,
		}
	} else {
		o.R.MapMatches = append(o.R.MapMatches, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &matchR{
				Map: o,
			}
		} else {
			rel.R.Map = o
		}
	}
	return nil
}

// GameMaps retrieves all the records using an executor.
func GameMaps(mods ...qm.QueryMod) gameMapQuery {
	mods = append(mods, qm.From("\"game_map\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"game_map\".*"})
	}

	return gameMapQuery{q}
}

// FindGameMap retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGameMap(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*GameMap, error) {
	gameMapObj := &GameMap{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"game_map\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, gameMapObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from game_map")
	}

	if err = gameMapObj.doAfterSelectHooks(ctx, exec); err != nil {
		return gameMapObj, err
	}

	return gameMapObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GameMap) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no game_map provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gameMapColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	gameMapInsertCacheMut.RLock()
	cache, cached := gameMapInsertCache[key]
	gameMapInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			gameMapAllColumns,
			gameMapColumnsWithDefault,
			gameMapColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(gameMapType, gameMapMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(gameMapType, gameMapMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"game_map\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"game_map\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into game_map")
	}

	if !cached {
		gameMapInsertCacheMut.Lock()
		gameMapInsertCache[key] = cache
		gameMapInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the GameMap.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GameMap) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	gameMapUpdateCacheMut.RLock()
	cache, cached := gameMapUpdateCache[key]
	gameMapUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			gameMapAllColumns,
			gameMapPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update game_map, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"game_map\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, gameMapPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(gameMapType, gameMapMapping, append(wl, gameMapPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update game_map row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for game_map")
	}

	if !cached {
		gameMapUpdateCacheMut.Lock()
		gameMapUpdateCache[key] = cache
		gameMapUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q gameMapQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for game_map")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for game_map")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GameMapSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gameMapPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"game_map\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, gameMapPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in gameMap slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all gameMap")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GameMap) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no game_map provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gameMapColumnsWithDefault, o)

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

	gameMapUpsertCacheMut.RLock()
	cache, cached := gameMapUpsertCache[key]
	gameMapUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			gameMapAllColumns,
			gameMapColumnsWithDefault,
			gameMapColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			gameMapAllColumns,
			gameMapPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert game_map, could not build update column list")
		}

		ret := strmangle.SetComplement(gameMapAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(gameMapPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert game_map, could not build conflict column list")
			}

			conflict = make([]string, len(gameMapPrimaryKeyColumns))
			copy(conflict, gameMapPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"game_map\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(gameMapType, gameMapMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(gameMapType, gameMapMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert game_map")
	}

	if !cached {
		gameMapUpsertCacheMut.Lock()
		gameMapUpsertCache[key] = cache
		gameMapUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single GameMap record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GameMap) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no GameMap provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), gameMapPrimaryKeyMapping)
	sql := "DELETE FROM \"game_map\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from game_map")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for game_map")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q gameMapQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no gameMapQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from game_map")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for game_map")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GameMapSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(gameMapBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gameMapPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"game_map\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gameMapPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from gameMap slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for game_map")
	}

	if len(gameMapAfterDeleteHooks) != 0 {
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
func (o *GameMap) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGameMap(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GameMapSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GameMapSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gameMapPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"game_map\".* FROM \"game_map\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gameMapPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GameMapSlice")
	}

	*o = slice

	return nil
}

// GameMapExists checks if the GameMap row exists.
func GameMapExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"game_map\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if game_map exists")
	}

	return exists, nil
}

// Exists checks if the GameMap row exists.
func (o *GameMap) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return GameMapExists(ctx, exec, o.ID)
}
