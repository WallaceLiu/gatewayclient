package db

import (
	"database/sql"
	"fmt"
	"github.com/gosexy/redis"
	"log"
)

const (
	REDIS_KEY_VW_V_CAR_DEALER                       = "vw_v_car_dealer"
	REDIS_KEY_VW_V_CAR_DEALER_MODEL                 = "vw_v_car_dealer_model"
	REDIS_KEY_VW_V_CAR_DEALER_REPORT                = "vw_v_car_dealer_report"
	REDIS_KEY_VW_V_CAR_DEALER_WP                    = "svw_v_car_dealer_wp"
	REDIS_KEY_VW_V_DEALER_CARLINE_POTENTIALCUSTOMER = "sk_v_dealer_carline_potentialcustomer"
	REDIS_KEY_VW_V_DEALER_FENCE                     = "sk_v_dealer_fence"
)

func MyVwRedis(db *sql.DB, rd *redis.Client) {
	my_vw_v_car_dealer_redis(db, rd)
	my_vw_v_car_dealer_model_redis(db, rd)
	my_vw_v_car_dealer_report_redis(db, rd)
	my_vw_v_car_dealer_wp_redis(db, rd)
	my_vw_v_dealer_carline_potentialcustomer_redis(db, rd)
	my_vw_v_dealer_fence_redis(db, rd)
}

func my_vw_v_car_dealer_redis(db *sql.DB, rd *redis.Client) {

	var (
		id                 string
		t_plan_id          string
		t_carline_model_id string
		t_dealer_id        string
		t_org_id           string
		obdSN              string
		vin                string
		plateNumber        string
		status             int
		dealerstatus       int
		ondutytime         string
		offdutytime        string
		lng_baidu          float64
		lat_baidu          float64
	)

	sql := `select v_car_dealer.id,
		ifnull(v_car_dealer.t_plan_id,"") as t_plan_id,
		ifnull(v_car_dealer.t_carline_model_id,"") as t_carline_model_id,
		ifnull(v_car_dealer.t_dealer_id,"") as t_dealer_id,
		ifnull(v_car_dealer.t_org_id,"") as t_org_id,
		ifnull(v_car_dealer.obdSN,"") as obdSN,
		ifnull(v_car_dealer.vin,"") as vin,
		ifnull(v_car_dealer.plateNumber,"") as plateNumber,
		ifnull(v_car_dealer.status,"") as status,
		ifnull(v_car_dealer.dealerstatus,"") as dealerstatus,
		ifnull(v_car_dealer.ondutytime,"") as ondutytime,
		ifnull(v_car_dealer.offdutytime,"") as offdutytime,
		ifnull(v_car_dealer.lng_baidu,"") as lng_baidu,
		ifnull(v_car_dealer.lat_baidu,"") as lat_baidu from v_car_dealer;`

	rows, err := db.Query(sql)

	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		err := rows.Scan(&id,
			&t_plan_id,
			&t_carline_model_id,
			&t_dealer_id,
			&t_org_id,
			&obdSN,
			&vin,
			&plateNumber,
			&status,
			&dealerstatus,
			&ondutytime,
			&offdutytime,
			&lng_baidu,
			&lat_baidu)

		if err != nil {
			fmt.Println(err)
		}

		key := fmt.Sprintf("%s:%s", REDIS_KEY_VW_V_CAR_DEALER, id)

		res, err := rd.HMSet(key,
			"t_plan_id", t_plan_id,
			"t_carline_model_id", t_carline_model_id,
			"t_dealer_id", t_dealer_id,
			"t_org_id", t_org_id,
			"obdSN", obdSN,
			"vin", vin,
			"plateNumber", plateNumber,
			"status", status,
			"ondutytime", ondutytime,
			"offdutytime", offdutytime,
			"lng_baidu", lng_baidu,
			"lat_baidu", lat_baidu)

		if len(res) > 0 {
		}

		log.Println(key)
	}

}

func my_vw_v_car_dealer_model_redis(db *sql.DB, rd *redis.Client) {

	var (
		t_car_id            string
		vin                 string
		t_org_rboid         string
		t_org_smallregionid string
		t_sys_dic_china_pid string
		t_sys_dic_china_id  string
		t_dealer_id         string
		dealerstatus        int
		plateNumber         string
		carStatus           string
	)

	sql := `select v_car_dealer_model.t_car_id,
		  ifnull(v_car_dealer_model.vin,"") as vin,
		  ifnull(v_car_dealer_model.t_org_rboid,"") as t_org_rboid,
		  ifnull(v_car_dealer_model.t_org_smallregionid ,"") as t_org_smallregionid,
		  ifnull(v_car_dealer_model.t_sys_dic_china_pid,"") as t_sys_dic_china_pid,
		  ifnull(v_car_dealer_model.t_sys_dic_china_id,"") as t_sys_dic_china_id,
		  ifnull(v_car_dealer_model.t_dealer_id,"") as t_dealer_id,
		  ifnull(v_car_dealer_model.dealerstatus,"") as dealerstatus,
		  ifnull(v_car_dealer_model.t_car_id,"") as t_car_id,
		  ifnull(v_car_dealer_model.plateNumber,"") as plateNumber,
		  ifnull(v_car_dealer_model.carStatus,"") as carStatus
		  from v_car_dealer_model;`

	rows, err := db.Query(sql)

	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		err := rows.Scan(&t_car_id,
			&vin,
			&t_org_rboid,
			&t_org_smallregionid,
			&t_sys_dic_china_pid,
			&t_sys_dic_china_id,
			&t_dealer_id,
			&dealerstatus,
			&t_car_id,
			&plateNumber,
			&carStatus)

		if err != nil {
			fmt.Println(err)
		}

		key := fmt.Sprintf("%s:%s", REDIS_KEY_VW_V_CAR_DEALER_MODEL, t_car_id)

		res, err := rd.HMSet(key,
			"vin", vin,
			"t_org_rboid", t_org_rboid,
			"t_org_smallregionid", t_org_smallregionid,
			"t_sys_dic_china_pid", t_sys_dic_china_pid,
			"t_sys_dic_china_id", t_sys_dic_china_id,
			"t_dealer_id", t_dealer_id,
			"dealerstatus", dealerstatus,
			"plateNumber", plateNumber,
			"carStatus", carStatus)

		if len(res) > 0 {
		}

		log.Println(key)
	}
}

func my_vw_v_car_dealer_report_redis(db *sql.DB, rd *redis.Client) {
	var (
		id                 string
		t_plan_id          string
		t_carline_model_id string
		t_dealer_id        string
		obdSN              string
		vin                string
		status             int
		dealerStatus       int
		onDutyTime         string
		offDutyTime        string
		reportBeginTime    string
		reportStatus       int
	)

	sql := `select v_car_dealer_report.id,
		  ifnull(v_car_dealer_report.t_plan_id,"") as t_plan_id,
		  ifnull(v_car_dealer_report.t_carline_model_id,"") as t_carline_model_id,
		  ifnull(v_car_dealer_report.t_dealer_id ,"") as t_dealer_id,
		  ifnull(v_car_dealer_report.obdSN,"") as obdSN,
		  ifnull(v_car_dealer_report.vin,"") as vin,
		  ifnull(v_car_dealer_report.status,"") as status,
		  ifnull(v_car_dealer_report.dealerstatus,"") as dealerstatus,
		  ifnull(v_car_dealer_report.onDutyTime,"") as onDutyTime,
		  ifnull(v_car_dealer_report.offDutyTime,"") as offDutyTime,
		  ifnull(v_car_dealer_report.reportBeginTime,"") as reportBeginTime,
		  ifnull(v_car_dealer_report.reportStatus,"") as reportStatus
		  from v_car_dealer_report;`

	rows, err := db.Query(sql)

	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		err := rows.Scan(&id,
			&t_plan_id,
			&t_carline_model_id,
			&t_dealer_id,
			&obdSN,
			&vin,
			&status,
			&dealerStatus,
			&onDutyTime,
			&offDutyTime,
			&reportBeginTime,
			&reportStatus)

		if err != nil {
			log.Println(err)
			return
		}

		key := fmt.Sprintf("%s:%s", REDIS_KEY_VW_V_CAR_DEALER_REPORT, id)

		res, err := rd.HMSet(key,
			"id", id,
			"t_plan_id", t_plan_id,
			"t_carline_model_id", t_carline_model_id,
			"t_dealer_id", t_dealer_id,
			"obdSN", obdSN,
			"vin", vin,
			"status", status,
			"dealerStatus", dealerStatus,
			"onDutyTime", onDutyTime,
			"offDutyTime", offDutyTime,
			"reportBeginTime", reportBeginTime,
			"reportStatus", reportStatus)

		if len(res) > 0 {
		}

		log.Println(key)
	}
}

func my_vw_v_car_dealer_wp_redis(db *sql.DB, rd *redis.Client) {
	var (
		id                 string
		t_plan_id          string
		t_carline_model_id string
		t_dealer_id        string
		obdSN              string
		vin                string
		status             int
		dealerStatus       string
		onDutyTime         string
		offDutyTime        string
		reportBeginTime    string
		reportStatus       string
	)

	sql := `select v_car_dealer_wp.id,
		  ifnull(v_car_dealer_wp.t_plan_id,"") as t_plan_id,
		  ifnull(v_car_dealer_wp.t_carline_model_id,"") as t_carline_model_id,
		  ifnull(v_car_dealer_wp.t_dealer_id ,"") as t_dealer_id,
		  ifnull(v_car_dealer_wp.obdSN,"") as obdSN,
		  ifnull(v_car_dealer_wp.vin,"") as vin,
		  ifnull(v_car_dealer_wp.status,"") as status,
		  ifnull(v_car_dealer_wp.dealerstatus,"") as dealerstatus,
		  ifnull(v_car_dealer_wp.onDutyTime,"") as onDutyTime,
		  ifnull(v_car_dealer_wp.offDutyTime,"") as offDutyTime,
		  ifnull(v_car_dealer_wp.reportBeginTime,"") as reportBeginTime,
		  ifnull(v_car_dealer_wp.reportStatus,"") as reportStatus
		  from v_car_dealer_wp;`

	rows, err := db.Query(sql)

	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		err := rows.Scan(&id,
			&t_plan_id,
			&t_carline_model_id,
			&t_dealer_id,
			&obdSN,
			&vin,
			&status,
			&dealerStatus,
			&onDutyTime,
			&offDutyTime,
			&reportBeginTime,
			&reportStatus)

		if err != nil {
			log.Println(err)
			return
		}

		key := fmt.Sprintf("%s:%s", REDIS_KEY_VW_V_CAR_DEALER_WP, id)

		res, err := rd.HMSet(key,
			"id", id,
			"t_plan_id", t_plan_id,
			"t_carline_model_id", t_carline_model_id,
			"t_dealer_id", t_dealer_id,
			"obdSN", obdSN,
			"vin", vin,
			"status", status,
			"dealerStatus", dealerStatus,
			"onDutyTime", onDutyTime,
			"offDutyTime", offDutyTime,
			"reportBeginTime", reportBeginTime,
			"reportStatus", reportStatus)

		if len(res) > 0 {
		}

		log.Println(key)
	}
}

func my_vw_v_dealer_carline_potentialcustomer_redis(db *sql.DB, rd *redis.Client) {
	var (
		dealerId   string
		carlineId  string
		recordTime string
		number     string
	)

	sql := `select v_dealer_carline_potentialcustomer.t_dealer_id,
		  ifnull(v_dealer_carline_potentialcustomer.t_carline_id,"") as t_carline_id,
		  ifnull(v_dealer_carline_potentialcustomer.recordtime,"") as recordtime,
		  ifnull(v_dealer_carline_potentialcustomer.number ,"") as number
		  from v_dealer_carline_potentialcustomer;`

	rows, err := db.Query(sql)

	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		err := rows.Scan(&dealerId,
			&carlineId,
			&recordTime,
			&number)

		if err != nil {
			log.Println(err)
			return
		}

		key := fmt.Sprintf("%s:%s",
			REDIS_KEY_VW_V_DEALER_CARLINE_POTENTIALCUSTOMER, dealerId)

		res, err := rd.HMSet(key,
			"t_dealer_id", dealerId,
			"t_carline_id", carlineId,
			"recordTime", recordTime,
			"number", number)
		if len(res) > 0 {
		}

		log.Println(key)
	}
}

func my_vw_v_dealer_fence_redis(db *sql.DB, rd *redis.Client) {
	var (
		id         string
		vin        string
		dealerName string
		dealerCode string
		lng_baidu  float32
		lat_baidu  float32
		fenceName  int
		innerLen   int
		outerLen   int
	)

	sql := `select v_dealer_fence.id,
		  ifnull(v_dealer_fence.vin,"") as vin,
		  ifnull(v_dealer_fence.dealername,"") as dealername,
		  ifnull(v_dealer_fence.dealercode ,"") as dealercode,
		  ifnull(v_dealer_fence.lng_baidu,"") as lng_baidu,
		  ifnull(v_dealer_fence.lat_baidu,"") as lat_baidu,
		  ifnull(v_dealer_fence.fencename,"") as fencename,
		  ifnull(v_dealer_fence.innerlen,"") as innerlen,
		  ifnull(v_dealer_fence.outerlen,"") as outerlen
		  from v_dealer_fence;`

	rows, err := db.Query(sql)

	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		err := rows.Scan(&id,
			&vin,
			&dealerName,
			&dealerCode,
			&lng_baidu,
			&lat_baidu,
			&fenceName,
			&innerLen,
			&outerLen)

		if err != nil {
			log.Println(err)
			return
		}

		key := fmt.Sprintf("%s:%s",
			REDIS_KEY_VW_V_DEALER_FENCE, id)

		res, err := rd.HMSet(key,
			"id", id,
			"vin", vin,
			"dealerName", dealerName,
			"dealerCode", dealerCode,
			"lng_baidu", lng_baidu,
			"lat_baidu", lat_baidu,
			"fenceName", fenceName,
			"innerLen", innerLen,
			"outerLen", outerLen)

		if len(res) > 0 {
		}

		log.Println(key)
	}
}
