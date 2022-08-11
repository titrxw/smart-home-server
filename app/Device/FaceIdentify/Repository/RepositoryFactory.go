package repository

type RepositoryFactory struct {
	FaceModelRepository
}

var FaceIdentifyDeviceRepository = new(RepositoryFactory)
