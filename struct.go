package main

import "time"

type GrafanaBody struct {
	Receiver          string            `json:"receiver"`
	Status            string            `json:"status"`
	OrgId             int               `json:"orgId"`
	Alerts            []*Alert          `json:"alerts"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	TruncatedAlerts   int               `json:"truncatedAlerts"`
	Title             string            `json:"title"`
	State             string            `json:"state"`
	Message           string            `json:"message"`
}

type Alert struct {
	Status       string            `json:"status"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Fingerprint  string            `json:"fingerprint"`
	SilenceURL   string            `json:"silenceURL"`
	DashboardURL string            `json:"dashboardURL"`
	PanelURL     string            `json:"panelURL"`
	ValueString  string            `json:"valueString"`
}

type BarkRequest struct {
	Body      string `json:"body"`
	Title     string `json:"title"`
	Badge     int    `json:"badge"`
	Category  string `json:"category"`
	Sound     string `json:"sound"`
	Icon      string `json:"icon"`
	Group     string `json:"group"`
	Url       string `json:"url"`
	IsArchive bool   `json:"is_archive"`
}

type BarkResponse struct {
	Code    int
	Message string
}

type GotifyMessage struct {
	Id       int64     `json:"id"`
	AppId    int64     `json:"appid"`
	Message  string    `json:"message"`
	Title    string    `json:"title"`
	Priority int64     `json:"priority"`
	Date     time.Time `json:"date"`
}
