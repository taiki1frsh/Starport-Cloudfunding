package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/taiki-furu/cloudfunding/x/cloudfunding/types"
)

// GetProjectCount get the total number of project
func (k Keeper) GetProjectCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ProjectCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetProjectCount set the total number of project
func (k Keeper) SetProjectCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ProjectCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendProject appends a project in the store with a new id and update the count
func (k Keeper) AppendProject(
	ctx sdk.Context,
	project types.Project,
) uint64 {
	// Create the project
	count := k.GetProjectCount(ctx)

	// Set the ID of the appended value
	project.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	appendedValue := k.cdc.MustMarshal(&project)
	store.Set(GetProjectIDBytes(project.Id), appendedValue)

	// Update project count
	k.SetProjectCount(ctx, count+1)

	return count
}

// SetProject set a specific project in the store
func (k Keeper) SetProject(ctx sdk.Context, project types.Project) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	b := k.cdc.MustMarshal(&project)
	store.Set(GetProjectIDBytes(project.Id), b)
}

// GetProject returns a project from its id
func (k Keeper) GetProject(ctx sdk.Context, id uint64) (val types.Project, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	b := store.Get(GetProjectIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProject removes a project from the store
func (k Keeper) RemoveProject(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	store.Delete(GetProjectIDBytes(id))
}

// GetAllProject returns all project
func (k Keeper) GetAllProject(ctx sdk.Context) (list []types.Project) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Project
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetProjectIDBytes returns the byte representation of the ID
func GetProjectIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetProjectIDFromBytes returns ID in uint64 format from a byte array
func GetProjectIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
