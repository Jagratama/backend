package seeder

import (
	"fmt"
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type UserSeeder struct{}

func (s UserSeeder) Run(db *gorm.DB) error {
	users := []model.User{
		{ID: 1, ImageID: 1, RoleID: 1, PositionID: 1, Name: "Admin", Email: "admin@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 2, ImageID: 1, RoleID: 4, PositionID: 2, Name: "Pengaju 1", Email: "pengaju1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 3, ImageID: 1, RoleID: 4, PositionID: 2, Name: "Pengaju 2", Email: "pengaju2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 4, ImageID: 1, RoleID: 3, PositionID: 3, Name: "Ketua SBH 1", Email: "ketuasbh1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 5, ImageID: 1, RoleID: 3, PositionID: 3, Name: "Ketua SBH 2", Email: "ketuasbh2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 6, ImageID: 1, RoleID: 3, PositionID: 4, Name: "PEMBINA SBH 1", Email: "pembinasbh1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 7, ImageID: 1, RoleID: 3, PositionID: 4, Name: "PEMBINA SBH 2", Email: "pembinasbh2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 8, ImageID: 1, RoleID: 3, PositionID: 5, Name: "PRESIDEN BEM 1", Email: "presidenbem1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 9, ImageID: 1, RoleID: 3, PositionID: 5, Name: "PRESIDEN BEM 2", Email: "presidenbem2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 10, ImageID: 1, RoleID: 2, PositionID: 6, Name: "KOMISI B BLM 1", Email: "komisibblm1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 11, ImageID: 1, RoleID: 2, PositionID: 6, Name: "KOMISI B BLM 2", Email: "komisibblm2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 12, ImageID: 1, RoleID: 2, PositionID: 7, Name: "PENANGGUNG JAWAB MAHASISWA DAN ALUMNI 1", Email: "penanggungjawabmahasiswaalumni1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 13, ImageID: 1, RoleID: 2, PositionID: 7, Name: "PENANGGUNG JAWAB MAHASISWA DAN ALUMNI 2", Email: "penanggungjawabmahasiswaalumni2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 14, ImageID: 1, RoleID: 2, PositionID: 8, Name: "KA SUB BAG ADM AKADEMIK 1", Email: "kasubbagadmakademik1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 15, ImageID: 1, RoleID: 2, PositionID: 8, Name: "KA SUB BAG ADM AKADEMIK 2", Email: "kasubbagadmakademik2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 16, ImageID: 1, RoleID: 2, PositionID: 9, Name: "KA BAG ADM AKADEMIK UMUM 1", Email: "kabagadmakademikumum1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 17, ImageID: 1, RoleID: 2, PositionID: 9, Name: "KA BAG ADM AKADEMIK UMUM 2", Email: "kabagadmakademikumum2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 18, ImageID: 1, RoleID: 3, PositionID: 10, Name: "WADIR III 1", Email: "wadiriii1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 19, ImageID: 1, RoleID: 3, PositionID: 10, Name: "WADIR III 2", Email: "wadiriii2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 20, ImageID: 1, RoleID: 3, PositionID: 11, Name: "DIREKTUR 1", Email: "direktur1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 21, ImageID: 1, RoleID: 3, PositionID: 11, Name: "DIREKTUR 2", Email: "direktur2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 22, ImageID: 1, RoleID: 3, PositionID: 12, Name: "Ketua UKM 1", Email: "ketuaukm1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 23, ImageID: 1, RoleID: 3, PositionID: 12, Name: "Ketua UKM 2", Email: "ketuaukm2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 24, ImageID: 1, RoleID: 3, PositionID: 13, Name: "PEMBINA UKM 1", Email: "pembinaukm1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 25, ImageID: 1, RoleID: 3, PositionID: 13, Name: "PEMBINA UKM 2", Email: "pembinaukm2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 26, ImageID: 1, RoleID: 3, PositionID: 14, Name: "MENTERI MINBA BEM 1", Email: "menteriminbabem1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 27, ImageID: 1, RoleID: 3, PositionID: 14, Name: "MENTERI MINBA BEM 2", Email: "menteriminbabem2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 28, ImageID: 1, RoleID: 3, PositionID: 15, Name: "Ketua HMPS 1", Email: "ketuahmps1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 29, ImageID: 1, RoleID: 3, PositionID: 15, Name: "Ketua HMPS 2", Email: "ketuahmps2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 30, ImageID: 1, RoleID: 3, PositionID: 16, Name: "PJ MAHASISWA DAN ALUMNI JURUSAN 1", Email: "pjmahasiswaalumnijurusan1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 31, ImageID: 1, RoleID: 3, PositionID: 16, Name: "PJ MAHASISWA DAN ALUMNI JURUSAN 2", Email: "pjmahasiswaalumnijurusan2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 32, ImageID: 1, RoleID: 3, PositionID: 17, Name: "KAPRODI 1", Email: "kaprodi1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 33, ImageID: 1, RoleID: 3, PositionID: 17, Name: "KAPRODI 2", Email: "kaprodi2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 34, ImageID: 1, RoleID: 3, PositionID: 18, Name: "KAJUR 1", Email: "kajur1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 35, ImageID: 1, RoleID: 3, PositionID: 18, Name: "KAJUR 2", Email: "kajur2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 36, ImageID: 1, RoleID: 3, PositionID: 19, Name: "Ketua HMJ 1", Email: "ketuahmj1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 37, ImageID: 1, RoleID: 3, PositionID: 19, Name: "Ketua HMJ 2", Email: "ketuahmj2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 38, ImageID: 1, RoleID: 3, PositionID: 20, Name: "KETUA BLM 1", Email: "ketuablm1@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
		{ID: 39, ImageID: 1, RoleID: 3, PositionID: 20, Name: "KETUA BLM 2", Email: "ketuablm2@gmail.com", Password: "$2a$10$D6kb.AIRHOt1lR66f6SeDearq/cnmI9ElSEZi7ZYbi4GqAZ7mmPvm"},
	}

	for _, user := range users {
		err := db.FirstOrCreate(&model.User{}, user).Error
		if err != nil {
			fmt.Printf("Error seeding user %s: %v\n", user.Name, err)
			return err
		}
	}

	err := db.Exec("SELECT setval('users_id_seq', (SELECT MAX(id) FROM users))").Error
	if err != nil {
		fmt.Printf("Error resetting user ID sequence: %v\n", err)
		return err
	}

	return nil
}
