package db

import (
	"config"
	"fmt"
	"testing"
)

func Test_my_sk_v_car_dealer_redis(t *testing.T) {

	db := MyDBOpen(config.DB_TYPE, config.SK_DB_URL)
	rd := MyRedisOpen(config.REDIS_HOST, config.REDIS_PORT)

	defer func() {
		MyDBClose(db)
		MyRedisQuit(rd)
	}()

	my_sk_v_car_dealer_redis(db, rd)

	fmt.Println("my_sk_v_car_dealer_redis test end.")
}

func Test_my_sk_v_car_dealer_model_redis(t *testing.T) {

	db := MyDBOpen(config.DB_TYPE, config.SK_DB_URL)
	rd := MyRedisOpen(config.REDIS_HOST, config.REDIS_PORT)

	defer func() {
		MyDBClose(db)
		MyRedisQuit(rd)
	}()

	my_sk_v_car_dealer_model_redis(db, rd)

	fmt.Println("my_sk_v_car_dealer_model_redis test end.")
}

func Test_my_sk_v_car_dealer_report_redis(t *testing.T) {
	db := MyDBOpen(config.DB_TYPE, config.SK_DB_URL)
	rd := MyRedisOpen(config.REDIS_HOST, config.REDIS_PORT)

	defer func() {
		MyDBClose(db)
		MyRedisQuit(rd)
	}()

	my_sk_v_car_dealer_report_redis(db, rd)

	fmt.Println("my_sk_v_car_dealer_report_redis test end.")
}

func Test_my_sk_v_car_dealer_wp_redis(t *testing.T) {
	db := MyDBOpen(config.DB_TYPE, config.SK_DB_URL)
	rd := MyRedisOpen(config.REDIS_HOST, config.REDIS_PORT)

	defer func() {
		MyDBClose(db)
		MyRedisQuit(rd)
	}()

	my_sk_v_car_dealer_wp_redis(db, rd)

	fmt.Println("my_sk_v_car_dealer_wp_redis test end.")
}

func Test_my_sk_v_dealer_carline_potentialcustomer_redis(t *testing.T) {
	db := MyDBOpen(config.DB_TYPE, config.SK_DB_URL)
	rd := MyRedisOpen(config.REDIS_HOST, config.REDIS_PORT)

	defer func() {
		MyDBClose(db)
		MyRedisQuit(rd)
	}()

	my_sk_v_dealer_carline_potentialcustomer_redis(db, rd)

	fmt.Println("my_sk_v_dealer_carline_potentialcustomer_redis test end.")
}

func Test_my_sk_v_dealer_fence_redis(t *testing.T) {

	db := MyDBOpen(config.DB_TYPE, config.SK_DB_URL)
	rd := MyRedisOpen(config.REDIS_HOST, config.REDIS_PORT)

	defer func() {
		MyDBClose(db)
		MyRedisQuit(rd)
	}()

	my_sk_v_dealer_fence_redis(db, rd)

	fmt.Println("my_sk_v_dealer_fence_redis test end.")
}

func Test_my_sk_t_loc(t *testing.T) {

	db := MyDBOpen(config.DB_TYPE, config.SK_DB_URL)
	rd := MyRedisOpen(config.REDIS_HOST, config.REDIS_PORT)

	defer func() {
		MyDBClose(db)
		MyRedisQuit(rd)
	}()

	my_sk_t_loc(db, rd)

	fmt.Println("my_sk_t_loc test end.")
}
