package models

type Employee struct {
	EMAIL       string `json:"email"`
	KD_KANTOR   string `json:"kd_kantor"`
	SOID        int    `json:"soid"`
	MASTERSOID  int    `json:"master_soid"`
	UNITKERJA   string `json:"unit_kerja"`
	JABATAN     string `json:"jabatan"`
	STS_JABATAN string `json:"sts_jabatan"`
}
