package datasets

// Dataset is an interface for machine learning datasets
type Dataset interface {
	LoadData()
	Preprocess()
}
