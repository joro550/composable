package settings

type ComposableSetting struct {
	fileName string
}

func NewSettings() ComposableSetting {
	return ComposableSetting{
		fileName: "settings.json",
	}
}

func (model ComposableSetting) WriteSettings() {
}
