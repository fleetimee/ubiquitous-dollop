package repository

import (
	"service-fleetime/cmd/models"
	"service-fleetime/config"
)

func FetchByEmail(email string) (fetcher models.Fetcher, err error) {
	init := config.DB

	query := `
	SELECT p.NRP                                                                                     AS nrp,
		u.KD_KANTOR                                                                               AS kd_kantor,
		s.SOID                                                                                    AS soid,
		s.MasterSOID                                                                              AS master_soid,
		u.UnitKerja                                                                               AS unit_kerja,
		db_HRMIS.dbo.nmjabatanya(p.KDJAB)                                                         AS jabatan,
		(SELECT nrp FROM db_HRMIS.dbo.Pegawai WHERE NRP = (db_HRMIS.dbo.f_nrp_konfirmasi(p.NRP))) AS nrp_atasan_first,
		db_HRMIS.dbo.f_nmatasannya(p.SOID, p.NRP)                                                 AS nama_atasan_first,
		(SELECT email
			FROM db_HRMIS.dbo.Pegawai
			WHERE NRP = (db_HRMIS.dbo.f_nrp_konfirmasi(p.NRP)))                                      AS email_atasan_first,
		(SELECT db_HRMIS.dbo.nmjabatanya(KDJAB)
			FROM db_HRMIS.dbo.Pegawai
			WHERE NRP = (db_HRMIS.dbo.f_nrp_konfirmasi(p.NRP)))                                      AS jabatan_atasan_first,
		(SELECT nrp
			FROM db_HRMIS.dbo.Pegawai
			WHERE NRP = (db_HRMIS.dbo.f_nrp_konfirmasi(db_HRMIS.dbo.f_nrp_konfirmasi(p.NRP))))       AS nrp_atasan_second,
		db_HRMIS.dbo.f_nmatasannya(p.SOID, db_HRMIS.dbo.f_nrp_konfirmasi(p.NRP))                  AS nama_atasan_second,
	      (SELECT email
          FROM db_HRMIS.dbo.Pegawai
          WHERE NRP =
                  (db_HRMIS.dbo.f_nrp_konfirmasi(db_HRMIS.dbo.f_nrp_konfirmasi(p.NRP))))             AS email_atasan_second,
		(SELECT db_HRMIS.dbo.nmjabatanya(KDJAB)
			FROM db_HRMIS.dbo.Pegawai
			WHERE NRP =
				(db_HRMIS.dbo.f_nrp_konfirmasi(db_HRMIS.dbo.f_nrp_konfirmasi(p.NRP))))             AS jabatan_atasan_second
	FROM db_HRMIS.dbo.StrukturOrganisasi s
			JOIN db_HRMIS.dbo.UnitKerja u ON s.KUK = u.KUK
			JOIN db_HRMIS.dbo.Pegawai p ON s.SOID = p.SOID
	WHERE p.Email = ?;
    `

	if err = init.Raw(query, email).Scan(&fetcher).Error; err != nil {
		return fetcher, err
	}

	return fetcher, nil

}
