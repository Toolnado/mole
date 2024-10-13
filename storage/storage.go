package storage

type StorageOptions struct {
	Transform func(string) string
}

type Storage struct {
	Opts StorageOptions
}

func NewStorage(opts StorageOptions) Storage {
	return Storage{
		Opts: opts,
	}
}
