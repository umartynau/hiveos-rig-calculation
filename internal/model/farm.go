package model

type FarmsResponse struct {
	Data []FarmData `json:"data"`
	Tags []FarmTag  `json:"tags"`
}

type FarmData struct {
	ID                  int                  `json:"id"`
	Name                string               `json:"name"`
	Timezone            string               `json:"timezone"`
	AsicRedLa           float64              `json:"asic_red_la"`
	AsicRedHashrate     []FarmHashrate       `json:"asic_red_hashrate"`
	RepoUrls            []string             `json:"repo_urls"`
	PowerPrice          float64              `json:"power_price"`
	PowerPriceCurrency  string               `json:"power_price_currency"`
	ChargeOnPool        bool                 `json:"charge_on_pool"`
	WorkerNameTemplate  string               `json:"worker_name_template"`
	AutocreateHash      string               `json:"autocreate_hash"`
	Nonfree             bool                 `json:"nonfree"`
	Locked              bool                 `json:"locked"`
	TwofaRequired       bool                 `json:"twofa_required"`
	Trusted             bool                 `json:"trusted"`
	Owner               User                 `json:"owner"`
	Payer               User                 `json:"payer"`
	PersonalSettings    PersonalSettings     `json:"personal_settings"`
	Role                string               `json:"role"`
	WorkersCount        int                  `json:"workers_count"`
	RigsCount           int                  `json:"rigs_count"`
	AsicsCount          int                  `json:"asics_count"`
	DisabledRigsCount   int                  `json:"disabled_rigs_count"`
	DisabledAsicsCount  int                  `json:"disabled_asics_count"`
	Money               FarmMoney            `json:"money"`
	Stats               FarmStats            `json:"stats"`
	Hashrates           []FarmHashrate       `json:"hashrates"`
	HashratesByCoin     []FarmHashrateByCoin `json:"hashrates_by_coin"`
	TagIds              []int                `json:"tag_ids"`
	HardwarePowerDraw   int                  `json:"hardware_power_draw"`
	PsuEfficiency       int                  `json:"psu_efficiency"`
	OctofanCorrectPower bool                 `json:"octofan_correct_power"`
	AutoTags            bool                 `json:"auto_tags"`
	DefaultFs           map[string]int       `json:"default_fs"`
}

type FarmHashrate struct {
	Algo     string  `json:"algo"`
	Hashrate float64 `json:"hashrate"`
}

type User struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
	Me    bool   `json:"me"`
}

type PersonalSettings struct {
	IsFavorite bool   `json:"is_favorite"`
	Note       string `json:"note"`
}

type FarmMoney struct {
	IsPaid            bool         `json:"is_paid"`
	IsFree            bool         `json:"is_free"`
	Balance           float64      `json:"balance"`
	Discount          float64      `json:"discount"`
	MonthlyCost       float64      `json:"monthly_cost"`
	DailyCost         float64      `json:"daily_cost"`
	CostDetails       []CostDetail `json:"cost_details"`
	DaysLeft          int          `json:"days_left"`
	Overdraft         bool         `json:"overdraft"`
	OverdraftDaysLeft int          `json:"overdraft_days_left"`
}

type CostDetail struct {
	Type         int     `json:"type"`
	TypeName     string  `json:"type_name"`
	Amount       float64 `json:"amount"`
	MonthlyPrice float64 `json:"monthly_price"`
	MonthlyCost  float64 `json:"monthly_cost"`
	DailyCost    float64 `json:"daily_cost"`
}

type FarmStats struct {
	WorkersTotal      int     `json:"workers_total"`
	WorkersOnline     int     `json:"workers_online"`
	WorkersOffline    int     `json:"workers_offline"`
	WorkersOverheated int     `json:"workers_overheated"`
	WorkersNoTemp     int     `json:"workers_no_temp"`
	WorkersOverloaded int     `json:"workers_overloaded"`
	WorkersInvalid    int     `json:"workers_invalid"`
	WorkersLowAsr     int     `json:"workers_low_asr"`
	WorkersNoHashrate int     `json:"workers_no_hashrate"`
	RigsTotal         int     `json:"rigs_total"`
	RigsOnline        int     `json:"rigs_online"`
	RigsOffline       int     `json:"rigs_offline"`
	RigsPower         int     `json:"rigs_power"`
	GpusTotal         int     `json:"gpus_total"`
	GpusOnline        int     `json:"gpus_online"`
	GpusOffline       int     `json:"gpus_offline"`
	GpusOverheated    int     `json:"gpus_overheated"`
	GpusNoTemp        int     `json:"gpus_no_temp"`
	AsicsTotal        int     `json:"asics_total"`
	AsicsOnline       int     `json:"asics_online"`
	AsicsPower        int     `json:"asics_power"`
	AsicsOffline      int     `json:"asics_offline"`
	BoardsTotal       int     `json:"boards_total"`
	BoardsOnline      int     `json:"boards_online"`
	BoardsOffline     int     `json:"boards_offline"`
	BoardsOverheated  int     `json:"boards_overheated"`
	BoardsNoTemp      int     `json:"boards_no_temp"`
	CpusOnline        int     `json:"cpus_online"`
	DevicesTotal      int     `json:"devices_total"`
	DevicesOnline     int     `json:"devices_online"`
	DevicesOffline    int     `json:"devices_offline"`
	PowerDraw         int     `json:"power_draw"`
	PowerCost         float64 `json:"power_cost"`
	Asr               float64 `json:"asr"`
}

type FarmHashrateByCoin struct {
	Coin     string  `json:"coin"`
	Algo     string  `json:"algo"`
	Hashrate float64 `json:"hashrate"`
}

type FarmTag struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Color        int    `json:"color"`
	FarmsCount   int    `json:"farms_count"`
	WorkersCount int    `json:"workers_count"`
	TypeId       int    `json:"type_id"`
	UserId       int    `json:"user_id"`
}
