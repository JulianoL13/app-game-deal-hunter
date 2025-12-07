package offer

import "errors"

var (
	ErrInvalidStore = errors.New("invalid store")
)

type Store string

const (
	StoreSteam  Store = "steam"
	StoreNuuvem Store = "nuuvem"
	StorePSN    Store = "psn"
	StoreXbox   Store = "xbox"
	StoreEpic   Store = "epic"
	StoreGOG    Store = "gog"
)

func NewStore(s string) (Store, error) {
	store := Store(s)
	if !store.IsValid() {
		return "", ErrInvalidStore
	}
	return store, nil
}

func (s Store) String() string {
	return string(s)
}

func (s Store) IsValid() bool {
	switch s {
	case StoreSteam, StoreNuuvem, StorePSN, StoreXbox, StoreEpic, StoreGOG:
		return true
	}
	return false
}
