package model

import "encoding/json"

type WorkersResponse struct {
	Data []WorkerData `json:"data"`
	Tags []WorkerTag  `json:"tags"`
}

type WorkerData struct {
	ID                   int                  `json:"id"`
	Platform             int                  `json:"platform"`
	Name                 string               `json:"name"`
	Description          string               `json:"description"`
	UnitsCount           int                  `json:"units_count"`
	FansCount            int                  `json:"fans_count"`
	Active               bool                 `json:"active"`
	Password             string               `json:"password"`
	TagIDs               []int                `json:"tag_ids"`
	MirrorURL            string               `json:"mirror_url"`
	RepoURLs             []string             `json:"repo_urls"`
	IPAddress            []string             `json:"ip_addresses"`
	RemoteAddress        RemoteAddress        `json:"remote_address"`
	VPN                  bool                 `json:"vpn"`
	HasAMD               bool                 `json:"has_amd"`
	HasNvidia            bool                 `json:"has_nvidia"`
	NeedsUpgrade         bool                 `json:"needs_upgrade"`
	PackagesHash         string               `json:"packages_hash"`
	LANConfig            LANConfig            `json:"lan_config"`
	SystemType           string               `json:"system_type"`
	OSName               string               `json:"os_name"`
	HasOctofan           bool                 `json:"has_octofan"`
	HasCoolbox           bool                 `json:"has_coolbox"`
	HasFanflap           bool                 `json:"has_fanflap"`
	HasPowermeter        bool                 `json:"has_powermeter"`
	HasAsichub           bool                 `json:"has_asichub"`
	HasDonnagerRelay     bool                 `json:"has_donnager_relay"`
	HasYkedaAutofan      bool                 `json:"has_ykeda_autofan"`
	HasWindtankAutofan   bool                 `json:"has_windtank_autofan"`
	HasMknetAutofan      bool                 `json:"has_mknet_autofan"`
	PersonalSettings     PersonalSettings     `json:"personal_settings"`
	Versions             Versions             `json:"versions"`
	Stats                WorkerStats          `json:"stats"`
	FlightSheet          FlightSheet          `json:"flight_sheet"`
	Overclock            Overclock            `json:"overclock"`
	TunedMiners          []string             `json:"tuned_miners"`
	MinersSummary        MinersSummary        `json:"miners_summary"`
	MinersStats          MinersStats          `json:"miners_stats"`
	HardwareInfo         HardwareInfo         `json:"hardware_info"`
	HardwareStats        HardwareStats        `json:"hardware_stats"`
	GPUSummary           GPUSummary           `json:"gpu_summary"`
	GPUInfo              []GPUInfo            `json:"gpu_info"`
	GPUStats             []GPUStats           `json:"gpu_stats"`
	ASICInfo             ASICInfo             `json:"asic_info"`
	ASICStats            ASICStats            `json:"asic_stats"`
	Options              Options              `json:"options"`
	HardwarePowerDraw    int                  `json:"hardware_power_draw"`
	PSUEfficiency        int                  `json:"psu_efficiency"`
	OctofanCorrectPower  bool                 `json:"octofan_correct_power"`
	ASICPowerModes       ASICPowerModes       `json:"asic_power_modes"`
	Autofan              Autofan              `json:"autofan"`
	Octofan              Octofan              `json:"octofan"`
	OctofanStats         OctofanStats         `json:"octofan_stats"`
	Coolbox              Coolbox              `json:"coolbox"`
	CoolboxInfo          CoolboxInfo          `json:"coolbox_info"`
	CoolboxStats         CoolboxStats         `json:"coolbox_stats"`
	Fanflap              Fanflap              `json:"fanflap"`
	FanflapStats         FanflapStats         `json:"fanflap_stats"`
	Powermeter           Powermeter           `json:"powermeter"`
	PowermeterStats      PowermeterStats      `json:"powermeter_stats"`
	DonnagerRelay        DonnagerRelay        `json:"donnager_relay"`
	DonnagerRelayInfo    DonnagerRelayInfo    `json:"donnager_relay_info"`
	DonnagerRelayStats   DonnagerRelayStats   `json:"donnager_relay_stats"`
	YkedaAutofan         YkedaAutofan         `json:"ykeda_autofan"`
	YkedaAutofanStats    YkedaAutofanStats    `json:"ykeda_autofan_stats"`
	WindtankAutofan      WindtankAutofan      `json:"windtank_autofan"`
	WindtankAutofanInfo  WindtankAutofanInfo  `json:"windtank_autofan_info"`
	WindtankAutofanStats WindtankAutofanStats `json:"windtank_autofan_stats"`
	Commands             []Command            `json:"commands"`
	BenchmarkID          int                  `json:"benchmark_id"`
	ASICConfig           ASICConfig           `json:"asic_config"`
	PerfProfile          PerfProfile          `json:"perf_profile"`
	MessagesCounts       MessagesCounts       `json:"messages_counts"`
	PowerOptions         PowerOptions         `json:"power_options"`
}

type RemoteAddress struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
}

type LANConfig struct {
	DHCP    bool   `json:"dhcp"`
	Address string `json:"address"`
	Gateway string `json:"gateway"`
	DNS     string `json:"dns"`
}

type Versions struct {
	Hive         string `json:"hive"`
	Kernel       string `json:"kernel"`
	AMDDriver    string `json:"amd_driver"`
	NvidiaDriver string `json:"nvidia_driver"`
}

type WorkerStats struct {
	Online           bool         `json:"online"`
	BootTime         int64        `json:"boot_time"`
	StatsTime        int64        `json:"stats_time"`
	MinerStartTime   int64        `json:"miner_start_time"`
	GPUsOnline       int          `json:"gpus_online"`
	GPUsOffline      int          `json:"gpus_offline"`
	GPUsOverheated   int          `json:"gpus_overheated"`
	CPUsOffline      int          `json:"cpus_offline"`
	BoardsOnline     int          `json:"boards_online"`
	BoardsOffline    int          `json:"boards_offline"`
	BoardsOverheated int          `json:"boards_overheated"`
	PowerDraw        int          `json:"power_draw"`
	Overheated       bool         `json:"overheated"`
	Overloaded       bool         `json:"overloaded"`
	Invalid          bool         `json:"invalid"`
	LowASR           bool         `json:"low_asr"`
	Problems         []string     `json:"problems"`
	AvgHashrates     AvgHashrates `json:"avg_hashrates"`
}

type AvgHashrates struct {
	Ethash struct {
		M15 int `json:"15m"`
		H1  int `json:"1h"`
	} `json:"ethash"`
}

type FlightSheet struct {
	ID     int               `json:"id"`
	FarmID int               `json:"farm_id"`
	UserID int               `json:"user_id"`
	Name   string            `json:"name"`
	Items  []FlightSheetItem `json:"items"`
}

type FlightSheetItem struct {
	Coin     string `json:"coin"`
	Pool     string `json:"pool"`
	WalID    int    `json:"wal_id"`
	DCoin    string `json:"dcoin"`
	DPool    string `json:"dpool"`
	DWalID   int    `json:"dwal_id"`
	Miner    string `json:"miner"`
	MinerAlt string `json:"miner_alt"`
}

type Overclock struct {
	Algo     string          `json:"algo"`
	AMD      json.RawMessage `json:"amd"`
	Nvidia   json.RawMessage `json:"nvidia"`
	Tweakers json.RawMessage `json:"tweakers"`
}

type AMDOverclock struct {
	CoreClock  string `json:"core_clock"`
	CoreState  string `json:"core_state"`
	CoreVDDC   string `json:"core_vddc"`
	MemClock   string `json:"mem_clock"`
	MemState   string `json:"mem_state"`
	MemMVDD    string `json:"mem_mvdd"`
	MemVDDCI   string `json:"mem_vddci"`
	FanSpeed   string `json:"fan_speed"`
	PowerLimit string `json:"power_limit"`
	TrefTiming string `json:"tref_timing"`
	SOCClock   string `json:"soc_clock"`
	SOCVDDMax  string `json:"soc_vddmax"`
	Aggressive bool   `json:"aggressive"`
}

type NvidiaOverclock struct {
	CoreClock              string `json:"core_clock"`
	LockCoreClock          string `json:"lock_core_clock"`
	MemClock               string `json:"mem_clock"`
	LockMemClock           string `json:"lock_mem_clock"`
	FanSpeed               string `json:"fan_speed"`
	PowerLimit             string `json:"power_limit"`
	LogoOff                bool   `json:"logo_off"`
	OhGodAPill             bool   `json:"ohgodapill"`
	OhGodAPillStartTimeout int    `json:"ohgodapill_start_timeout"`
	OhGodAPillArgs         string `json:"ohgodapill_args"`
	RunningDelay           int    `json:"running_delay"`
	ReducePower            bool   `json:"reduce_power"`
	ForceP0                bool   `json:"force_p0"`
}

type Tweakers struct {
	AMDMemTweak []AMDMemTweak `json:"amdmemtweak"`
}

type AMDMemTweak struct {
	GPUs   []int             `json:"gpus,omitempty"`
	Params AMDMemTweakParams `json:"params"`
}

type AMDMemTweakParams struct {
	CL  int `json:"cl"`
	RAS int `json:"ras"`
}

type MinersSummary struct {
	Hashrates []MinerSummary `json:"hashrates"`
}

type MinerSummary struct {
	Miner  string  `json:"miner"`
	Ver    string  `json:"ver"`
	Algo   string  `json:"algo"`
	Coin   string  `json:"coin"`
	Hash   float64 `json:"hash"`
	DAlgo  string  `json:"dalgo"`
	DCoin  string  `json:"dcoin"`
	DHash  float64 `json:"dhash"`
	Shares Shares  `json:"shares"`
}

type Shares struct {
	Accepted int     `json:"accepted"`
	Rejected int     `json:"rejected"`
	Invalid  int     `json:"invalid"`
	Ratio    float64 `json:"ratio"`
}

type MinersStats struct {
	Hashrates []MinerStats `json:"hashrates"`
}

type MinerStats struct {
	Miner         string    `json:"miner"`
	Algo          string    `json:"algo"`
	Coin          string    `json:"coin"`
	Hashes        []float64 `json:"hashes"`
	DAlgo         string    `json:"dalgo"`
	DCoin         string    `json:"dcoin"`
	DHashes       []float64 `json:"dhashes"`
	Temps         []int     `json:"temps"`
	Fans          []int     `json:"fans"`
	InvalidShares []int     `json:"invalid_shares"`
	BusNumbers    []int     `json:"bus_numbers"`
	DBusNumbers   []int     `json:"dbus_numbers"`
}

type HardwareInfo struct {
	Motherboard   Motherboard    `json:"motherboard"`
	CPU           CPU            `json:"cpu"`
	Disk          Disk           `json:"disk"`
	NetInterfaces []NetInterface `json:"net_interfaces"`
}

type Motherboard struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	BIOS         string `json:"bios"`
}

type CPU struct {
	ID    string `json:"id"`
	Model string `json:"model"`
	Cores int    `json:"cores"`
	AES   bool   `json:"aes"`
}

type Disk struct {
	Model string `json:"model"`
}

type NetInterface struct {
	Interface string `json:"iface"`
	MAC       string `json:"mac"`
}

type HardwareStats struct {
	DF       string    `json:"df"`
	CPUAvg   []float64 `json:"cpuavg"`
	CPUTemp  []int     `json:"cputemp"`
	CPUCores int       `json:"cpu_cores"`
	Memory   Memory    `json:"memory"`
}

type Memory struct {
	Total int `json:"total"`
	Free  int `json:"free"`
}

type GPUSummary struct {
	GPUs    []GPUItem `json:"gpus"`
	MaxTemp int       `json:"max_temp"`
	MaxFan  int       `json:"max_fan"`
}

type GPUItem struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type GPUInfo struct {
	BusID      string         `json:"bus_id"`
	BusNumber  int            `json:"bus_number"`
	Brand      string         `json:"brand"`
	Model      string         `json:"model"`
	ShortName  string         `json:"short_name"`
	Details    GPUInfoDetails `json:"details"`
	PowerLimit PowerLimit     `json:"power_limit"`
	Index      int            `json:"index"`
}

type GPUInfoDetails struct {
	Mem       string `json:"mem"`
	MemGB     int    `json:"mem_gb"`
	MemType   string `json:"mem_type"`
	MemOEM    string `json:"mem_oem"`
	VBIOS     string `json:"vbios"`
	Subvendor string `json:"subvendor"`
	OEM       string `json:"oem"`
}

type PowerLimit struct {
	Min string `json:"min"`
	Def string `json:"def"`
	Max string `json:"max"`
}

type GPUStats struct {
	BusNumber int `json:"bus_number"`
	Temp      int `json:"temp"`
	Fan       int `json:"fan"`
	Power     int `json:"power"`
	CoreClock int `json:"coreclk"`
	MemClock  int `json:"memclk"`
	MemTemp   int `json:"memtemp"`
}

type ASICInfo struct {
	Model        string `json:"model"`
	ShortName    string `json:"short_name"`
	LogicVersion string `json:"logic_version"`
	Firmware     string `json:"firmware"`
	Hiveon       bool   `json:"hiveon"`
	InstallType  string `json:"install_type"`
	ControlBoard string `json:"control_board"`
}

type ASICStats struct {
	Fans      []ASICFan   `json:"fans"`
	FansCount int         `json:"fans_count"`
	Boards    []ASICBoard `json:"boards"`
	ASICBoost bool        `json:"asicboost"`
}

type ASICFan struct {
	Index  int `json:"index"`
	Fan    int `json:"fan"`
	FanRPM int `json:"fan_rpm"`
}

type ASICBoard struct {
	Chain        int     `json:"chain"`
	ACN          int     `json:"acn"`
	Freq         float64 `json:"freq"`
	Status       []int   `json:"status"`
	Temp         int     `json:"temp"`
	BoardTemp    int     `json:"board_temp"`
	HWErrors     int     `json:"hw_errors"`
	Power        float64 `json:"power"`
	ChainVoltage float64 `json:"chain_voltage"`
}

type WatchdogOption struct {
	Miner   string `json:"miner"`
	MinHash int    `json:"minhash"`
	Units   string `json:"units"`
}

type Options struct {
	DisableGUI        bool   `json:"disable_gui"`
	MaintenanceMode   int    `json:"maintenance_mode"`
	PushInterval      int    `json:"push_interval"`
	MinerDelay        int    `json:"miner_delay"`
	DOH               int    `json:"doh"`
	PowerCycle        bool   `json:"power_cycle"`
	ShellinaboxEnable bool   `json:"shellinabox_enable"`
	SSHEnable         bool   `json:"ssh_enable"`
	SSHPasswordEnable bool   `json:"ssh_password_enable"`
	SSHAuthorizedKeys string `json:"ssh_authorized_keys"`
	VNCEnable         bool   `json:"vnc_enable"`
	VNCPassword       string `json:"vnc_password"`
}

type ASICPowerModes struct {
	Mode          int  `json:"mode"`
	ModePower     bool `json:"mode_power"`
	ModeEconomy   bool `json:"mode_economy"`
	DisableFan    bool `json:"disable_fan"`
	DisablePools  bool `json:"disable_pools"`
	DisableSwitch bool `json:"disable_switch"`
}

type Autofan struct {
	Fan           int  `json:"fan"`
	Auto          bool `json:"auto"`
	TargetTemp    int  `json:"target_temp"`
	TargetMemTemp int  `json:"target_mem_temp"`
	MinFan        int  `json:"min_fan"`
	MaxFan        int  `json:"max_fan"`
}

type Octofan struct {
	Casefan       []int `json:"casefan"`
	Thermosensors []int `json:"thermosensors"`
}

type OctofanStats struct {
	Casefan       []int `json:"casefan"`
	Thermosensors []int `json:"thermosensors"`
}

type Coolbox struct {
	Fan           int  `json:"fan"`
	Auto          bool `json:"auto"`
	TargetTemp    int  `json:"target_temp"`
	TargetMemTemp int  `json:"target_mem_temp"`
	WDEnabled     bool `json:"wd_enabled"`
	WDInterval    int  `json:"wd_interval"`
}

type CoolboxInfo struct {
	Version  string `json:"version"`
	Revision string `json:"revision"`
	Model    string `json:"model"`
}

type CoolboxStats struct {
	Casefan       []int `json:"casefan"`
	Thermosensors []int `json:"thermosensors"`
}

type Fanflap struct {
	Fan        int  `json:"fan"`
	Auto       bool `json:"auto"`
	TargetTemp int  `json:"target_temp"`
	MinFan     int  `json:"min_fan"`
	MaxFan     int  `json:"max_fan"`
}

type FanflapStats struct {
	Fan []int `json:"fan"`
}

type Powermeter struct {
	Meters []PowermeterMeter `json:"meters"`
}

type PowermeterMeter struct {
	URL  string `json:"url"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

type PowermeterStats struct {
	Power       [][]float64 `json:"power"`
	PowerTotal  []float64   `json:"power_total"`
	EnergyTotal []float64   `json:"energy_total"`
}

type DonnagerRelay struct {
	Channels []DonnagerRelayChannel `json:"channels"`
}

type DonnagerRelayChannel struct {
	Index    int `json:"index"`
	WorkerID int `json:"worker_id"`
}

type DonnagerRelayInfo struct {
	Firmware string `json:"firmware"`
}

type DonnagerRelayStats struct {
	Channels []DonnagerRelayChannelStats `json:"channels"`
}

type DonnagerRelayChannelStats struct {
	Index   int `json:"index"`
	State   int `json:"state"`
	Current int `json:"current"`
}

type YkedaAutofan struct {
	Fan           int  `json:"fan"`
	Auto          bool `json:"auto"`
	TargetTemp    int  `json:"target_temp"`
	TargetMemTemp int  `json:"target_mem_temp"`
	MinFan        int  `json:"min_fan"`
	MaxFan        int  `json:"max_fan"`
}

type YkedaAutofanStats struct {
	Casefan       []int `json:"casefan"`
	Thermosensors []int `json:"thermosensors"`
}

type WindtankAutofan struct {
	Fan           int  `json:"fan"`
	Auto          bool `json:"auto"`
	TargetTemp    int  `json:"target_temp"`
	TargetMemTemp int  `json:"target_mem_temp"`
	MinFan        int  `json:"min_fan"`
}

type WindtankAutofanInfo struct {
	Model string `json:"model"`
}

type WindtankAutofanStats struct {
	Casefan       []int `json:"casefan"`
	Thermosensors []int `json:"thermosensors"`
}

type Command struct {
	Command string   `json:"command"`
	ID      int      `json:"id"`
	Data    struct{} `json:"data"`
}

type Benchmark struct {
	ID int `json:"benchmark_id"`
}

type ASICConfig struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}

type PerfProfile struct {
	Active bool               `json:"active"`
	Tuned  []PerfProfileTuned `json:"tuned"`
}

type PerfProfileTuned struct {
	Profile string `json:"profile"`
}

type MessagesCounts struct {
	Success int `json:"success"`
	Danger  int `json:"danger"`
	Warning int `json:"warning"`
	Info    int `json:"info"`
	Default int `json:"default"`
	File    int `json:"file"`
}

type PowerOptions struct {
	Mode string `json:"mode"`
}

type WorkerTag struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Color        int    `json:"color"`
	FarmsCount   int    `json:"farms_count"`
	WorkersCount int    `json:"workers_count"`
	IsAuto       bool   `json:"is_auto"`
	FarmID       int    `json:"farm_id"`
	UserID       int    `json:"user_id"`
}
