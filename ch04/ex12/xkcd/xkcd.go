package xkcd

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func ReqXKCD(num int) (*XKCD, error) {
	url := "https://xkcd.com/" + strconv.Itoa(num) + "/info.0.json"
	fmt.Println(url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed: %s", resp.Status)
	}
	var xkcd XKCD
	if err := json.NewDecoder(resp.Body).Decode(&xkcd); err != nil {
		return nil, err
	}

	return &xkcd, nil
}

func GetAllXKCDs() {
	// 1 ~ 403, 405 ~ 2033
	db := InitXKCDTable()
	db.Close()
	for i := 1; i <= 2033; i++ {
		if i == 404 {
			continue
		}
		res, err := ReqXKCD(i)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(res)

		if err := PostXKCD(res); err != nil {
			fmt.Println(err)
		}
	}
}

func GetXKCDs() []*XKCD {
	db := InitXKCDTable()
	defer db.Close()
	var xkcds []*XKCD
	db.Find(&xkcds)
	return xkcds
}

func SearchXKCD(query string) ([]*XKCD, error) {
	db := InitXKCDTable()
	defer db.Close()
	var xkcds []*XKCD
	db.Where("transcript LIKE ?", "%"+query+"%").Find(&xkcds)

	if len(xkcds) > 0 {
		return xkcds, nil
	}
	err := fmt.Errorf("query=%s does not match any xkcd.", query)
	return nil, err
}

func PostXKCD(xkcd *XKCD) error {
	db := InitXKCDTable()
	defer db.Close()

	if xkcd.Num > 0 {
		db.Create(&xkcd)
		return nil
	}
	err := errors.New("Values must be int")
	return err
}
