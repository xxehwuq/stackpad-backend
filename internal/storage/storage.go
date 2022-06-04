package storage

type Storage struct {
	Notebook NotebookStorage
}

func New() *Storage {
	return &Storage{
		Notebook: newNotebookStorage(),
	}
}
