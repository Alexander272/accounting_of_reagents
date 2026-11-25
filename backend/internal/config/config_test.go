package config

import "testing"

func TestLimiterEnvNames(t *testing.T) {
	// Старое общее имя больше не должно влиять на оба лимитера.
	t.Setenv("RPS", "999")
	t.Setenv("API_RPS", "42")
	t.Setenv("STATIC_RPS", "77")

	conf, err := Init("../../configs/config.yaml")
	if err != nil {
		t.Fatalf("не удалось прочитать конфиг: %v", err)
	}

	if conf.ApiLimiter.RPS != 42 {
		t.Errorf("API_RPS: хочу 42, получил %d", conf.ApiLimiter.RPS)
	}
	if conf.StaticLimiter.RPS != 77 {
		t.Errorf("STATIC_RPS: хочу 77, получил %d", conf.StaticLimiter.RPS)
	}
}
