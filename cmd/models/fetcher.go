package models

type Fetcher struct {
	Nrp                 int    `json:"nrp"`
	KdKantor            int    `json:"kd_kantor"`
	Soid                int    `json:"soid"`
	MasterSoid          int    `json:"master_soid"`
	UnitKerja           string `json:"unit_kerja"`
	Jabatan             string `json:"jabatan"`
	NrpAtasanFirst      int    `json:"nrp_atasan_first"`
	NamaAtasanFirst     string `json:"nama_atasan_first"`
	EmailAtasanFirst    string `json:"email_atasan_first"`
	JabatanAtasanFirst  string `json:"jabatan_atasan_first"`
	NrpAtasanSecond     int    `json:"nrp_atasan_second"`
	NamaAtasanSecond    string `json:"nama_atasan_second"`
	EmailAtasanSecond   string `json:"email_atasan_second"`
	JabatanAtasanSecond string `json:"jabatan_atasan_second"`
}
