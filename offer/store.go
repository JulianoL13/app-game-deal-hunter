package offer

type Store string

const (
	StoreSteam  Store = "steam"
	StoreNuuvem Store = "nuuvem"
	StorePSN    Store = "psn"
	StoreXbox   Store = "xbox"
	StoreEpic   Store = "epic"
	StoreGOG    Store = "gog"
)

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

type Platform string

const (
	PlatformPC          Platform = "pc"
	PlatformPlayStation Platform = "playstation"
	PlatformXbox        Platform = "xbox"
	PlatformNintendo    Platform = "nintendo"
)

func (p Platform) String() string {
	return string(p)
}

func (p Platform) IsValid() bool {
	switch p {
	case PlatformPC, PlatformPlayStation, PlatformXbox, PlatformNintendo:
		return true
	}
	return false
}
