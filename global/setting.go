package global

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting (n, p, t string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName(n)
	vp.AddConfigPath(p)
	vp.SetConfigType(t)
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

func (s *Setting) ReadSetting (key string, value any) error {
	err := s.vp.UnmarshalKey(key, value)
	if err != nil {
		return err
	}
	return nil
}