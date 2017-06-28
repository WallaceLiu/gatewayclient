package orbit

import (
	"config"
	"fmt"
	"log"
	"math"
	"testing"
)

func Test_F100_ab01_c0(t *testing.T) {
	log.Println(">>> Function F_a/b=0.1_c=0,Testing...")
	log.Println(fmt.Sprintf(">>> start loc：(%f,%f)", config.LNG, config.LAT))
	lng, lat := config.LNG, config.LAT
	for i := 1; i <= 10; i++ {
		lng, lat = F100_ab01_c0(lng, lat, config.LINEAR_STEP)
		log.Print(fmt.Sprintf("(%f,%f)", lng, lat))
	}
}

func Test_F101_ab02_c0(t *testing.T) {
	log.Println("Function F_a/b=0.2_c=0,Testing...")
	log.Println(fmt.Sprintf(">>> start loc：(%f,%f)", config.LNG, config.LAT))
	lng, lat := config.LNG, config.LAT
	for i := 1; i <= 10; i++ {
		lng, lat = F101_ab02_c0(lng, lat, config.LINEAR_STEP)
		log.Print(fmt.Sprintf("(%f,%f)", lng, lat))
	}
}

func Test_F102_ab025_c0(t *testing.T) {
	log.Println("Function F_a/b=0.25_c=0,Testing...")
	log.Println(fmt.Sprintf(">>> start loc：(%f,%f)", config.LNG, config.LAT))
	lng, lat := config.LNG, config.LAT
	for i := 1; i <= 10; i++ {
		lng, lat = F102_ab025_c0(lng, lat, config.LINEAR_STEP)
		log.Print(fmt.Sprintf("(%f,%f)", lng, lat))
	}
}

func Test_F103_ab03_c0(t *testing.T) {
	log.Println("Function F_a/b=0.3_c=0,Testing...")
	log.Println(fmt.Sprintf(">>> start loc：(%f,%f)", config.LNG, config.LAT))
	lng, lat := config.LNG, config.LAT
	for i := 1; i <= 10; i++ {
		lng, lat = F103_ab03_c0(lng, lat, config.LINEAR_STEP)
		log.Print(fmt.Sprintf("(%f,%f)", lng, lat))
	}
}

func Test_F104_ab04_c0(t *testing.T) {
	log.Println("Function F_a/b=0.4_c=0,Testing...")
	log.Println(fmt.Sprintf(">>> start loc：(%f,%f)", config.LNG, config.LAT))
	lng, lat := config.LNG, config.LAT
	for i := 1; i <= 10; i++ {
		lng, lat = F104_ab04_c0(lng, lat, config.LINEAR_STEP)
		log.Print(fmt.Sprintf("(%f,%f)", lng, lat))
	}
}

func Test_F105_ab05_c0(t *testing.T) {
	log.Println("Function F_a/b=0.5_c=0,Testing...")
	log.Println(fmt.Sprintf(">>> start loc：(%f,%f)", config.LNG, config.LAT))
	lng, lat := config.LNG, config.LAT
	for i := 1; i <= 10; i++ {
		lng, lat = F105_ab05_c0(lng, lat, config.LINEAR_STEP)
		log.Print(fmt.Sprintf("(%f,%f)", lng, lat))
	}
}

func Test_F106_ab06_c0(t *testing.T) {
	log.Println("Function F_a/b=0.6_c=0,Testing...")
	log.Println(fmt.Sprintf(">>> start loc：(%f,%f)", config.LNG, config.LAT))
	lng, lat := config.LNG, config.LAT
	for i := 1; i <= 10; i++ {
		lng, lat = F106_ab06_c0(lng, lat, config.LINEAR_STEP)
		log.Print(fmt.Sprintf("(%f,%f)", lng, lat))
	}
}

func Test_F107_ab07_c0(t *testing.T) {
	log.Println("Function F_a/b=0.7_c=0,Testing...")
	log.Println(fmt.Sprintf(">>> start loc：(%f,%f)", config.LNG, config.LAT))
	lng, lat := config.LNG, config.LAT
	for i := 1; i <= 10; i++ {
		lng, lat = F107_ab07_c0(lng, lat, config.LINEAR_STEP)
		log.Print(fmt.Sprintf("(%f,%f)", lng, lat))
	}
}

func Test_F108_ab08_c0(t *testing.T) {
	log.Println("Function F_a/b=0.8_c=0,Testing...")
	log.Println(fmt.Sprintf(">>> start loc：(%f,%f)", config.LNG, config.LAT))
	lng, lat := config.LNG, config.LAT
	for i := 1; i <= 10; i++ {
		lng, lat = F108_ab08_c0(lng, lat, config.LINEAR_STEP)
		log.Print(fmt.Sprintf("(%f,%f)", lng, lat))
	}
}

func Test_F109_ab09_c0(t *testing.T) {
	log.Println("Function F_a/b=0.9_c=0,Testing...")
	log.Println(fmt.Sprintf(">>> start loc：(%f,%f)", config.LNG, config.LAT))
	lng, lat := config.LNG, config.LAT
	for i := 1; i <= 10; i++ {
		lng, lat = F109_ab09_c0(lng, lat, config.LINEAR_STEP)
		log.Print(fmt.Sprintf("(%f,%f)", lng, lat))
	}
}

func Test_F110_ab1_c0(t *testing.T) {
	log.Println("Function F_a/b=1_c=0,Testing...")
	log.Println(fmt.Sprintf(">>> start loc：(%f,%f)", config.LNG, config.LAT))
	lng, lat := config.LNG, config.LAT
	for i := 1; i <= 10; i++ {
		lng, lat = F110_ab1_c0(lng, lat, config.LINEAR_STEP)
		log.Print(fmt.Sprintf("(%f,%f)", lng, lat))
	}
}

func Test_F110_ab1_c0_control(t *testing.T) {
	log.Println("Function F_a/b=1_c=0,Testing...")
	log.Println(fmt.Sprintf(">>> start loc：(%f,%f)", config.LNG, config.LAT))
	lng, lat := config.LNG, config.LAT
	for i := 1; i <= 1000; i++ {
		lng, lat = F110_ab1_c0(lng, lat, 10)
		log.Print(fmt.Sprintf("(%f,%f)", lng, lat))
	}
}

func Test_F111_ab09_c1(t *testing.T) {
	log.Println("Function F_a/b=0.9_c=0.000001,Testing...")
	log.Println(fmt.Sprintf(">>> start loc：(%f,%f)", config.LNG, config.LAT))
	lng, lat := config.LNG, config.LAT
	for i := 1; i <= 10; i++ {
		lng, lat = F111_ab09_c1(lng, lat, config.LINEAR_STEP)
		log.Print(fmt.Sprintf("(%f,%f)", lng, lat))
	}
}

func Test_F200_sin(t *testing.T) {
	log.Println("Function F_sin,Testing...")
	log.Println(fmt.Sprintf(">>> start loc：(%f,%f)", config.LNG, config.LAT))
	lng, lat := config.LNG, config.LAT
	for i := 1; i <= 10; i++ {
		lng, lat = F200_sin(lng, lat, float64(i)*math.Pi)
		log.Print(fmt.Sprintf("(%f,%f)", lng, lat))
	}
}

func Test_F200_cos(t *testing.T) {
	log.Println("Function F_sin,Testing...")
	log.Println(fmt.Sprintf(">>> start loc：(%f,%f)", config.LNG, config.LAT))
	lng, lat := config.LNG, config.LAT
	for i := 1; i <= 10; i++ {
		lng, lat = F201_cos(lng, lat, float64(i)*math.Pi)
		log.Print(fmt.Sprintf("(%f,%f)", lng, lat))
	}
}
