package models

import "github.com/provider-go/district/global"

type District struct {
	Id           int32  `json:"id" gorm:"auto_increment;primary_key;comment:'主键'"`
	DistrictName string `json:"districtName" gorm:"column:district_name;type:varchar(50);not null;default:'';comment:地区名称"`
	Level        int    `json:"level" gorm:"column:level;type:tinyint(1);not null;default:0;comment:地区等级"`
	UpId         int32  `json:"upId" gorm:"column:upid;not null;default:0;comment:上级地区Id"`
	Seq          int    `json:"seq" gorm:"column:seq;type:tinyint(1);not null;default:0;comment:排序"`
}

func ListDistrict(upId int) ([]*District, error) {
	var rows []*District

	if err := global.DB.Table("districts").Where("upid = ?", upId).Order("seq desc").Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
