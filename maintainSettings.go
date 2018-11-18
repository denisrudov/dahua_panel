package dahua_panel

type TableData struct {
	AutoRebootDay      int  `json:"AutoRebootDay"`
	AutoRebootEnable   bool `json:"IsAutoRebootEnable"`
	AutoRebootHour     int  `json:"AutoRebootHour"`
	AutoRebootMinute   int  `json:"AutoRebootMinute"`
	AutoShutdownDay    int  `json:"AutoShutdownDay"`
	AutoShutdownHour   int  `json:"AutoShutdownHour"`
	AutoShutdownMinute int  `json:"AutoShutdownMinute"`
	AutoStartUpDay     int  `json:"AutoStartUpDay"`
	AutoStartUpHour    int  `json:"AutoStartUpHour"`
	AutoStartUpMinute  int  `json:"AutoStartUpMinute"`
}

type maintainParams struct {
	Name    string        `json:"name"`
	Table   TableData     `json:"table"`
	Options []interface{} `json:"options"`
}

func NewMaintainParams() *maintainParams {

	return &maintainParams{
		Name: MaintainParamName,
		Table: TableData{
			AutoRebootDay:      -1,
			AutoRebootEnable:   false,
			AutoRebootHour:     2,
			AutoRebootMinute:   0,
			AutoShutdownDay:    -1,
			AutoShutdownHour:   0,
			AutoShutdownMinute: 0,
			AutoStartUpDay:     -1,
			AutoStartUpHour:    0,
			AutoStartUpMinute:  0,
		},
	}
}

func (p *maintainParams) IsAutoRebootEnable(enable bool) {
	p.Table.AutoRebootEnable = enable
}
