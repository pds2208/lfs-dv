package types

import "time"

type DV struct {
	Id             int       `db:"id" csv:"id"`
	DvName         string    `db:"dv_name" csv:"dvName"`
	DVCode         string    `db:"dv" csv:"dv"`
	DVType         int       `db:"dv_type" csv:"dvType"`
	ExecutionOrder int       `db:"exec_order" csv:"execOrder"`
	SectionOrder   int       `db:"section_order" csv:"sectionOrder"`
	EndDate        time.Time `db:"end_date" csv:"endDate"`
}
