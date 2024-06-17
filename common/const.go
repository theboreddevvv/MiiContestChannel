package common

type ListTag [2]byte

var (
	FirstList       = ListTag{'F', 'D'}
	SpotList        = ListTag{'S', 'L'}
	BargainList     = ListTag{'R', 'L'}
	NewList         = ListTag{'N', 'L'}
	ContestDetail   = ListTag{'C', 'D'}
	ContestInfo     = ListTag{'C', 'I'}
	PopcraftsList   = ListTag{'C', 'L'}
	PopularList     = ListTag{'P', 'L'}
	ExtendedArtisan = ListTag{'R', 'C'}
	RK              = ListTag{'R', 'K'}
	PairMii         = ListTag{'P', 'M'}
	PairArtisan     = ListTag{'P', 'C'}
	EntryList       = ListTag{'E', 'L'}
	ContestMii      = ListTag{'C', 'M'}
	BestList        = ListTag{'B', 'L'}
	ContestTag      = ListTag{'C', 'N'}
	ContestArtisan  = ListTag{'C', 'C'}
	Thumbnail       = ListTag{'T', 'H'}
	MiiPairNumber   = ListTag{'P', 'N'}
)
