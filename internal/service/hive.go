package service

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"rig-calculation/internal/config"
	"rig-calculation/internal/model"
	"rig-calculation/pkg/ers"
	"time"
)

type Hive struct {
	client *http.Client
	cfg    config.Hive
	l      zerolog.Logger
}

func NewHive(hiveCfg config.Hive) *Hive {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &Hive{
		client: client,
		cfg:    hiveCfg,
		l:      log.Logger.With().Str("module", "service.hive").Logger(),
	}
}

func (s *Hive) GetFarms() (*model.FarmsResponse, error) {
	req, err := http.NewRequest("GET", s.cfg.ApiUrl+"/farms", nil)
	if err != nil {
		s.l.Error().Err(err).Msg("GetFarms http NewRequest")
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+s.cfg.ApiKey)

	resp, err := s.client.Do(req)
	if err != nil {
		s.l.Error().Err(err).Msg("GetFarms client Do")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.l.Error().Err(err).Int("status_code", resp.StatusCode).Msg("GetFarms client Do")
		return nil, ers.ErrStatusCode
	}

	var farms *model.FarmsResponse
	if err := json.NewDecoder(resp.Body).Decode(&farms); err != nil {
		s.l.Error().Err(err).Msg("GetFarms NewDecoder Decode")
		return nil, err
	}

	return farms, nil
}

func (s *Hive) GetWorkers(farmId int) (*model.WorkersResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/farms/%d/workers", s.cfg.ApiUrl, farmId), nil)
	if err != nil {
		s.l.Error().Err(err).Msg("GetWorkers http NewRequest")
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+s.cfg.ApiKey)

	resp, err := s.client.Do(req)
	if err != nil {
		s.l.Error().Err(err).Msg("GetWorkers client Do")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.l.Error().Err(err).Int("status_code", resp.StatusCode).Msg("GetWorkers client Do")
		return nil, ers.ErrStatusCode
	}

	var workers *model.WorkersResponse
	if err := json.NewDecoder(resp.Body).Decode(&workers); err != nil {
		s.l.Error().Err(err).Msg("GetWorkers NewDecoder Decode")
		return nil, err
	}

	return workers, nil
}

func (s *Hive) GetWorkerStats(farmId int, workerId int) (*model.WorkerStatsResponse, error) {
	currentTime := time.Now()
	oneMonthAgo := currentTime.AddDate(0, -1, 0)
	firstOfMonth := time.Date(oneMonthAgo.Year(), oneMonthAgo.Month(), 1, 0, 0, 0, 0, oneMonthAgo.Location())
	formattedDate := firstOfMonth.Format("2006-01-02")

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/farms/%d/workers/%d/metrics?date=%s&period=1m", s.cfg.ApiUrl, farmId, workerId, formattedDate), nil)
	if err != nil {
		s.l.Error().Err(err).Msg("GetWorkerStats http NewRequest")
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+s.cfg.ApiKey)

	resp, err := s.client.Do(req)
	if err != nil {
		s.l.Error().Err(err).Msg("GetWorkerStats client Do")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.l.Error().Err(err).Int("status_code", resp.StatusCode).Msg("GetWorkerStats client Do")
		return nil, ers.ErrStatusCode
	}

	var worker *model.WorkerStatsResponse
	if err := json.NewDecoder(resp.Body).Decode(&worker); err != nil {
		s.l.Error().Err(err).Msg("GetWorkerStats NewDecoder Decode")
		return nil, err
	}

	return worker, nil
}
