package globalModel

import (
	"database/sql/driver"
	"encoding/json"
)

/* 搞定multipic域 */
/* [{"name":"file-1677828241356","key":"4bb0f090dabbafe754d0cdf8c77d9c6f_20230303152400","url":"uploads/file/4bb0f090dabbafe754d0cdf8c77d9c6f_20230303152400"}] */
type MultiPic struct {
	FileName string `json:"fileName"` //文件名
	Key      string `json:"key"`      //文件Key,
	Url      string `json:"url"`      //url
}

func (c MultiPic) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *MultiPic) Scan(src any) error {
	return json.Unmarshal(src.([]byte), c)
}
